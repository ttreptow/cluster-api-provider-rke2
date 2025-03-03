/*
Copyright 2022 SUSE.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package secret

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
	"sigs.k8s.io/cluster-api/util/certs"

	bootstrapv1 "github.com/rancher-sandbox/cluster-api-provider-rke2/bootstrap/api/v1alpha1"
	"github.com/rancher-sandbox/cluster-api-provider-rke2/pkg/consts"
)

const (
	// DefaultCertificatesDir is the default location (file path) where the provider will put the certificates, this location will then
	// be automatically used by RKE2 to use the pre-defined certificates instead of generating them.
	DefaultCertificatesDir = "/var/lib/rancher/rke2/server/tls"

	// Kubeconfig is the secret name suffix storing the Cluster Kubeconfig.
	Kubeconfig = Purpose("kubeconfig")

	// KubeconfigDataName is the data entry name for the Kubeconfig file content.
	KubeconfigDataName string = "value"

	// EtcdCA is the secret name suffix for the Etcd CA.
	EtcdCA Purpose = "etcd"

	// ClusterCA is the secret name suffix for APIServer CA.
	ClusterCA = Purpose("ca")

	// ClientClusterCA is the secret name suffix for APIServer CA.
	ClientClusterCA = Purpose("cca")

	// TLSKeyDataName is the key used to store a TLS private key in the secret's data field.
	TLSKeyDataName = "tls.key"

	// TLSCrtDataName is the key used to store a TLS certificate in the secret's data field.
	TLSCrtDataName = "tls.crt"

	// APIServerEtcdClient is the secret name of user-supplied secret containing the apiserver-etcd-client key/cert.
	APIServerEtcdClient Purpose = "apiserver-etcd-client"

	// ServiceAccount is the secret name suffix for the Service Account keys.
	ServiceAccount Purpose = "sa"

	// TenYears is the duration of one year.
	TenYears = time.Hour * 24 * 365 * 10
)

// Purpose is the name to append to the secret generated for a cluster.
type Purpose string

// Certificate represents a single certificate CA.
type Certificate struct {
	Generated         bool
	External          bool
	Purpose           Purpose
	KeyPair           *certs.KeyPair
	CertFile, KeyFile string
}

// Certificates are the certificates necessary to bootstrap a cluster.
type Certificates []*Certificate

// NewCertificatesForInitialControlPlane returns a list of certificates configured for a control plane node.
func NewCertificatesForInitialControlPlane() Certificates {
	certificatesDir := DefaultCertificatesDir

	certificates := Certificates{
		&Certificate{
			Purpose:  ClusterCA,
			CertFile: filepath.Join(certificatesDir, "server-ca.crt"),
			KeyFile:  filepath.Join(certificatesDir, "server-ca.key"),
		},
		&Certificate{
			Purpose:  ClientClusterCA,
			CertFile: filepath.Join(certificatesDir, "client-ca.crt"),
			KeyFile:  filepath.Join(certificatesDir, "client-ca.key"),
		},
	}

	return certificates
}

// GetByPurpose returns a certificate by the given name.
// This could be removed if we use a map instead of a slice to hold certificates, however other code becomes more complex.
func (c Certificates) GetByPurpose(purpose Purpose) *Certificate {
	for _, certificate := range c {
		if certificate.Purpose == purpose {
			return certificate
		}
	}

	return nil
}

// Lookup looks up each certificate from secrets and populates the certificate with the secret data.
func (c Certificates) Lookup(ctx context.Context, ctrlclient client.Client, clusterName client.ObjectKey) error {
	// Look up each certificate as a secret and populate the certificate/key
	for _, certificate := range c {
		s := &corev1.Secret{}
		key := client.ObjectKey{
			Name:      Name(clusterName.Name, certificate.Purpose),
			Namespace: clusterName.Namespace,
		}

		if err := ctrlclient.Get(ctx, key, s); err != nil {
			if apierrors.IsNotFound(err) {
				if certificate.External {
					return errors.WithMessage(err, "external certificate not found")
				}

				continue
			}

			return errors.WithStack(err)
		}
		// If a user has a badly formatted secret it will prevent the cluster from working.
		kp, err := secretToKeyPair(s)
		if err != nil {
			return err
		}

		certificate.KeyPair = kp
	}

	return nil
}

// Generate will generate any certificates that do not have KeyPair data.
func (c *Certificate) Generate() error {
	// Do not generate the APIServerEtcdClient key pair. It is user supplied
	if c.Purpose == APIServerEtcdClient {
		return nil
	}

	generator := generateCACert
	if c.Purpose == ServiceAccount {
		generator = generateServiceAccountKeys
	}

	kp, err := generator()
	if err != nil {
		return err
	}

	c.KeyPair = kp
	c.Generated = true

	return nil
}

// Generate will generate any certificates that do not have KeyPair data.
func (c Certificates) Generate() error {
	for _, certificate := range c {
		if certificate.KeyPair == nil {
			err := certificate.Generate()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// SaveGenerated will save any certificates that have been generated as Kubernetes secrets.
func (c Certificates) SaveGenerated(ctx context.Context, ctrlclient client.Client, clusterName client.ObjectKey, owner metav1.OwnerReference) error {
	for _, certificate := range c {
		if !certificate.Generated {
			continue
		}

		s := certificate.AsSecret(clusterName, owner)
		if err := ctrlclient.Create(ctx, s); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

// LookupOrGenerate is a convenience function that wraps cluster bootstrap certificate behavior.
func (c Certificates) LookupOrGenerate(
	ctx context.Context,
	ctrlclient client.Client,
	clusterName client.ObjectKey,
	owner metav1.OwnerReference,
) error {
	// Find the certificates that exist
	if err := c.Lookup(ctx, ctrlclient, clusterName); err != nil {
		return err
	}

	// Generate the certificates that don't exist
	if err := c.Generate(); err != nil {
		return err
	}

	// Save any certificates that have been generated
	if err := c.SaveGenerated(ctx, ctrlclient, clusterName, owner); err != nil {
		return err
	}

	return nil
}

// AsSecret converts a single certificate into a Kubernetes secret.
func (c *Certificate) AsSecret(clusterName client.ObjectKey, owner metav1.OwnerReference) *corev1.Secret {
	s := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: clusterName.Namespace,
			Name:      Name(clusterName.Name, c.Purpose),
			Labels: map[string]string{
				clusterv1.ClusterNameLabel: clusterName.Name,
			},
		},
		Data: map[string][]byte{
			TLSKeyDataName: c.KeyPair.Key,
			TLSCrtDataName: c.KeyPair.Cert,
		},
		Type: clusterv1.ClusterSecretType,
	}

	if c.Generated {
		s.OwnerReferences = []metav1.OwnerReference{owner}
	}

	return s
}

// AsFiles converts the certificate to a slice of Files that may have 0, 1 or 2 Files.
func (c *Certificate) AsFiles() []bootstrapv1.File {
	out := make([]bootstrapv1.File, 0)
	if len(c.KeyPair.Cert) > 0 {
		out = append(out, bootstrapv1.File{
			Path:        c.CertFile,
			Owner:       consts.DefaultFileOwner,
			Permissions: "0640",
			Content:     string(c.KeyPair.Cert),
		})
	}

	if len(c.KeyPair.Key) > 0 {
		out = append(out, bootstrapv1.File{
			Path:        c.KeyFile,
			Owner:       consts.DefaultFileOwner,
			Permissions: "0600",
			Content:     string(c.KeyPair.Key),
		})
	}

	return out
}

// Name returns the name of the secret for a cluster.
func Name(cluster string, suffix Purpose) string {
	return fmt.Sprintf("%s-%s", cluster, suffix)
}

// AsFiles converts a slice of certificates into bootstrap files.
func (c Certificates) AsFiles() []bootstrapv1.File {
	clusterCA := c.GetByPurpose(ClusterCA)
	clientClusterCA := c.GetByPurpose(ClientClusterCA)

	etcdCA := c.GetByPurpose(EtcdCA)

	certFiles := make([]bootstrapv1.File, 0)
	if clusterCA != nil {
		certFiles = append(certFiles, clusterCA.AsFiles()...)
	}

	if clientClusterCA != nil {
		certFiles = append(certFiles, clientClusterCA.AsFiles()...)
	}

	if etcdCA != nil {
		certFiles = append(certFiles, etcdCA.AsFiles()...)
	}

	// these will only exist if external etcd was defined and supplied by the user
	apiserverEtcdClientCert := c.GetByPurpose(APIServerEtcdClient)
	if apiserverEtcdClientCert != nil {
		certFiles = append(certFiles, apiserverEtcdClientCert.AsFiles()...)
	}

	return certFiles
}

// secretToKeyPair gets a Certificate Keypair from a data entry in a secret.
func secretToKeyPair(s *corev1.Secret) (*certs.KeyPair, error) {
	c, exists := s.Data[TLSCrtDataName]
	if !exists {
		return nil, errors.Errorf("missing data for key %s", TLSCrtDataName)
	}

	// In some cases (external etcd) it's ok if the etcd.key does not exist.
	key, exists := s.Data[TLSKeyDataName]
	if !exists {
		key = []byte("")
	}

	return &certs.KeyPair{
		Cert: c,
		Key:  key,
	}, nil
}

func generateCACert() (*certs.KeyPair, error) {
	x509Cert, privKey, err := newCertificateAuthority()
	if err != nil {
		return nil, err
	}

	return &certs.KeyPair{
		Cert: certs.EncodeCertPEM(x509Cert),
		Key:  certs.EncodePrivateKeyPEM(privKey),
	}, nil
}

// newCertificateAuthority creates new certificate and private key for the certificate authority.
func newCertificateAuthority() (*x509.Certificate, *rsa.PrivateKey, error) {
	key, err := certs.NewPrivateKey()
	if err != nil {
		return nil, nil, err
	}

	c, err := newSelfSignedCACert(key)
	if err != nil {
		return nil, nil, err
	}

	return c, key, nil
}

// newSelfSignedCACert creates a CA certificate.
func newSelfSignedCACert(key *rsa.PrivateKey) (*x509.Certificate, error) {
	cfg := certs.Config{
		CommonName: "kubernetes",
	}

	now := time.Now().UTC()

	tmpl := x509.Certificate{
		SerialNumber: new(big.Int).SetInt64(0),
		Subject: pkix.Name{
			CommonName:   cfg.CommonName,
			Organization: cfg.Organization,
		},
		NotBefore:             now.Add(time.Minute * -5),
		NotAfter:              now.Add(TenYears), // 10 years
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		MaxPathLenZero:        true,
		BasicConstraintsValid: true,
		MaxPathLen:            0,
		IsCA:                  true,
	}

	b, err := x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, key.Public(), key)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create self signed CA certificate: %+v", tmpl)
	}

	c, err := x509.ParseCertificate(b)

	return c, errors.WithStack(err)
}

func generateServiceAccountKeys() (*certs.KeyPair, error) {
	saCreds, err := certs.NewPrivateKey()
	if err != nil {
		return nil, err
	}

	saPub, err := certs.EncodePublicKeyPEM(&saCreds.PublicKey)
	if err != nil {
		return nil, err
	}

	return &certs.KeyPair{
		Cert: saPub,
		Key:  certs.EncodePrivateKeyPEM(saCreds),
	}, nil
}

// GetFromNamespacedName retrieves the specified Secret (if any) from the given
// cluster name and namespace.
func GetFromNamespacedName(ctx context.Context, c client.Reader, clusterName client.ObjectKey, purpose Purpose) (*corev1.Secret, error) {
	secret := &corev1.Secret{}
	secretKey := client.ObjectKey{
		Namespace: clusterName.Namespace,
		Name:      Name(clusterName.Name, purpose),
	}

	if err := c.Get(ctx, secretKey, secret); err != nil {
		return nil, err
	}

	return secret, nil
}
