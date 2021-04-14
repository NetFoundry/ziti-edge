// +build apitests

/*
	Copyright NetFoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package tests

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"github.com/openziti/edge/eid"
	"math/big"
	"net/http"
	"strings"
	"testing"
	"time"
)

func Test_enrollment(t *testing.T) {
	ctx := NewTestContext(t)
	defer ctx.Teardown()
	ctx.StartServer()
	ctx.RequireAdminLogin()

	t.Run("ca auto enrollment", func(t *testing.T) {

		t.Run("setup CA", func(t *testing.T) {
			testCa := newTestCa()
			ctx.testContextChanged(t)

			testCaId := ctx.AdminSession.requireCreateEntity(testCa)
			ctx.Req.NotEmpty(testCaId)

			caContainer := ctx.AdminSession.requireQuery("cas/" + testCaId)
			ctx.Req.NotEmpty(caContainer)

			token := caContainer.Path("data.verificationToken").Data().(string)
			ctx.Req.NotEmpty(token)

			verifyCert, _, err := generateCert(testCa.publicCert, testCa.privateKey, token)
			ctx.Req.NoError(err)

			verificationBlock := &pem.Block{
				Type:  "CERTIFICATE",
				Bytes: verifyCert.Raw,
			}
			verifyPem := pem.EncodeToMemory(verificationBlock)

			resp, err := ctx.AdminSession.newAuthenticatedRequest().SetHeader("content-type", "text/plain").SetBody(verifyPem).Post("cas/" + testCaId + "/verify")
			ctx.Req.NoError(err)
			standardJsonResponseTests(resp, http.StatusOK, t)

			t.Run("can enroll without a name and a 0 length body", func(t *testing.T) {
				ctx.testContextChanged(t)
				cert, key, err := generateCert(testCa.publicCert, testCa.privateKey, "test-can-enroll-"+eid.New())
				ctx.Req.NoError(err)

				restClient, _, transport := ctx.NewClientComponents()
				transport.TLSClientConfig.Certificates = []tls.Certificate{
					{
						Certificate: [][]byte{cert.Raw},
						PrivateKey:  key,
					},
				}

				resp, err := restClient.R().
					SetHeader("content-type", "application/json").
					Post("enroll?method=ca")

				ctx.Req.NoError(err)

				ctx.Req.Equal(http.StatusOK, resp.StatusCode())

				contentTypeHeaders := resp.Header().Values("content-type")
				ctx.Req.Equal(1, len(contentTypeHeaders), "expected only 1 content type header")

				contentType := strings.Split(contentTypeHeaders[0], ";")[0]
				ctx.Req.Equal("application/json", contentType)
			})

			t.Run("can enroll without a name and empty JSON object", func(t *testing.T) {
				ctx.testContextChanged(t)
				cert, key, err := generateCert(testCa.publicCert, testCa.privateKey, "test-can-enroll-"+eid.New())
				ctx.Req.NoError(err)

				restClient, _, transport := ctx.NewClientComponents()
				transport.TLSClientConfig.Certificates = []tls.Certificate{
					{
						Certificate: [][]byte{cert.Raw},
						PrivateKey:  key,
					},
				}

				resp, err := restClient.R().
					SetHeader("content-type", "application/json").
					SetBody("{}").
					Post("enroll?method=ca")

				ctx.Req.NoError(err)

				ctx.Req.Equal(http.StatusOK, resp.StatusCode())

				contentTypeHeaders := resp.Header().Values("content-type")
				ctx.Req.Equal(1, len(contentTypeHeaders), "expected only 1 content type header")

				contentType := strings.Split(contentTypeHeaders[0], ";")[0]
				ctx.Req.Equal("application/json", contentType)
			})

			t.Run("can enroll with a name", func(t *testing.T) {
				t.Run("can enroll without a name and empty JSON object", func(t *testing.T) {
					ctx.testContextChanged(t)
					cert, key, err := generateCert(testCa.publicCert, testCa.privateKey, "test-can-enroll-"+eid.New())
					ctx.Req.NoError(err)

					restClient, _, transport := ctx.NewClientComponents()
					transport.TLSClientConfig.Certificates = []tls.Certificate{
						{
							Certificate: [][]byte{cert.Raw},
							PrivateKey:  key,
						},
					}

					resp, err := restClient.R().
						SetHeader("content-type", "application/json").
						SetBody(`{"name": "` + eid.New() + `"}`).
						Post("enroll?method=ca")

					ctx.Req.NoError(err)

					ctx.Req.Equal(http.StatusOK, resp.StatusCode())
				})
			})
		})

	})
}

func generateCert(caCert *x509.Certificate, caKey crypto.Signer, commonName string) (*x509.Certificate, crypto.Signer, error) {
	id, _ := rand.Int(rand.Reader, big.NewInt(100000000000000000))
	verificationCert := &x509.Certificate{
		SerialNumber: id,
		Subject: pkix.Name{
			CommonName:   commonName,
			Organization: []string{"Ziti CLI Generated Validation Cert"},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(time.Minute * 10),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
	}

	verificationKey, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)

	if err != nil {
		return nil, nil, fmt.Errorf("could not generate private key for verification certificate (%v)", err)
	}

	signedCertBytes, err := x509.CreateCertificate(rand.Reader, verificationCert, caCert, verificationKey.Public(), caKey)

	if err != nil {
		return nil, nil, fmt.Errorf("could not sign verification certificate with CA (%v)", err)
	}

	verificationCert, _ = x509.ParseCertificate(signedCertBytes)

	return verificationCert, verificationKey, nil
}