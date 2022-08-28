package did_test

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/fairxio/go/did"
	"github.com/fairxio/protocol-integration-tests/auth"
	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"math/rand"
	"net/http"
)

var _ = Describe("DID Document Operations", func() {

	Describe("Creating a DID Document for the first time", func() {

		// generate random id
		randomNum := rand.Int()
		randomId := fmt.Sprintf("%v@fairx.io", randomNum)
		randomDid := fmt.Sprintf("did:fairx:%s", randomId)

		// Authenticate
		jwt := auth.Authenticate("did:fairx:aW50ZWdyYXRpb250ZXN0QGZhaXJ4Lmlv")

		// Create the DID Document
		didDocument := did.DIDDocument{
			ID:          randomDid,
			AlsoKnownAs: []string{randomId, "test@fairx.io"},
			Controller:  []string{randomDid},
		}

		// Marshall to json and send to DID service
		didDocumentJson, _ := json.Marshal(&didDocument)

		It("Accepts the document and stores", func() {

			client := resty.New()
			resp, err := client.R().
				SetHeader("Content-Type", "application/json").
				SetAuthToken(jwt).
				SetBody(didDocumentJson).
				Post("http://localhost:8001/v1.0.0")

			Expect(err).To(BeNil())
			Expect(resp).ToNot(BeNil())
			Expect(resp.StatusCode()).To(BeEquivalentTo(http.StatusOK))

		})

		It("Allows us to retrieve the document", func() {

			client := resty.New()
			resp, err := client.R().
				Get(fmt.Sprintf("http://localhost:8001/v1.0.0/%s", base64.RawURLEncoding.EncodeToString([]byte(randomDid))))

			Expect(err).To(BeNil())
			Expect(resp).ToNot(BeNil())
			Expect(resp.StatusCode()).To(BeEquivalentTo(http.StatusOK))

			var responseDidDocument did.DIDDocument
			err = json.Unmarshal(resp.Body(), &responseDidDocument)
			Expect(err).To(BeNil())
			Expect(responseDidDocument.ID).To(BeEquivalentTo(randomDid))

		})

		It("Allows us to update the document with additional attributes", func() {

			client := resty.New()
			resp, err := client.R().
				Get(fmt.Sprintf("http://localhost:8001/v1.0.0/%s", base64.RawURLEncoding.EncodeToString([]byte(randomDid))))

			Expect(err).To(BeNil())
			Expect(resp).ToNot(BeNil())
			Expect(resp.StatusCode()).To(BeEquivalentTo(http.StatusOK))

			var responseDidDocument did.DIDDocument
			err = json.Unmarshal(resp.Body(), &responseDidDocument)
			Expect(err).To(BeNil())
			Expect(responseDidDocument.ID).To(BeEquivalentTo(randomDid))

			responseDidDocument.Service = []did.Service{
				did.Service{
					ID: fmt.Sprintf("%s%s", responseDidDocument.ID, "#fairx"),
					Endpoints: []did.ServiceEndpoint{
						did.ServiceEndpoint{
							URI: "https://fairx.io/v1.0.0",
						},
					},
				},
			}

			docJson, _ := json.Marshal(&responseDidDocument)
			updateResp, err := client.R().
				SetHeader("Content-Type", "application/json").
				SetAuthToken(jwt).
				SetBody(docJson).
				Post("http://localhost:8001/v1.0.0")

			Expect(updateResp).ToNot(BeNil())
			Expect(updateResp.StatusCode()).To(BeEquivalentTo(http.StatusOK))

		})

	})

})
