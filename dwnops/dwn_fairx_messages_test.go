package dwnops_test

import (
	"encoding/json"
	"github.com/fairxio/go/did"
	"github.com/fairxio/protocol-integration-tests/auth"
	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DWN FairX Messages", func() {

	Describe("Get No Messages", func() {

		// Authenticate
		jwt := auth.Authenticate("did:fairx:aW50ZWdyYXRpb250ZXN0QGZhaXJ4Lmlv")

		It("Returns a successful response with no messages", func() {

			fairxMessage := did.Message{
				Descriptor: did.Descriptor{
					Method: "FairXSessionMessages",
				},
			}
			requestObject := did.RequestObject{
				Target:   "did:fairx:aW50ZWdyYXRpb250ZXN0QGZhaXJ4Lmlv",
				Messages: []did.Message{fairxMessage},
			}
			rawRequestObject, _ := json.Marshal(&requestObject)

			client := resty.New()
			resp, err := client.R().
				SetHeader("Content-Type", "application/json").
				SetAuthToken(jwt).
				SetBody(rawRequestObject).
				Post("http://localhost:8002/v1.0.0")

			Expect(err).To(BeNil())
			Expect(resp.StatusCode()).To(BeEquivalentTo(200))

			var responseObject did.ResponseObject
			bodyContent := resp.Body()
			Expect(bodyContent).ToNot(BeNil())

			err = json.Unmarshal(bodyContent, &responseObject)
			Expect(err).To(BeNil())
			Expect(responseObject.Status.Code).To(BeEquivalentTo(200))
			Expect(len(responseObject.Replies)).To(BeEquivalentTo(1))
			Expect(responseObject.Replies[0].Status.Code).To(BeEquivalentTo(200))

		})

	})

	Describe("Post a valid message and retrieve it", func() {

		// Authenticate
		jwt := auth.Authenticate("did:fairx:dGVzdEBmYWlyeC5pbw")

		fairxMessage := did.Message{
			Descriptor: did.Descriptor{
				Method: "FairXSessionEstablish",
			},
		}
		requestObject := did.RequestObject{
			Target:   "did:fairx:dGVzdEBmYWlyeC5pbw",
			Messages: []did.Message{fairxMessage},
		}

		It("Posts a Session Establish message to myself", func() {

			rawRequestObject, _ := json.Marshal(&requestObject)

			// Post to DWN
			client := resty.New()
			resp, err := client.R().
				SetHeader("Content-Type", "application/json").
				SetAuthToken(jwt).
				SetBody(rawRequestObject).
				Post("http://localhost:8002/v1.0.0")

			Expect(err).To(BeNil())
			Expect(resp.StatusCode()).To(BeEquivalentTo(200))

			var responseObject did.ResponseObject
			err = json.Unmarshal(resp.Body(), &responseObject)

			Expect(err).To(BeNil())
			Expect(responseObject.Status.Code).To(BeEquivalentTo(200))
			Expect(len(responseObject.Replies)).To(BeEquivalentTo(1))
			Expect(responseObject.Replies[0].Status).To(BeEquivalentTo(200))

			// Get from DWN
			getMessages := did.Message{Descriptor: did.Descriptor{Method: "FairXSessionMessages"}}
			roMessages := did.RequestObject{
				Target:   "did:fairx:dGVzdEBmYWlyeC5pbw",
				Messages: []did.Message{getMessages},
			}
			rawRoMessages, _ := json.Marshal(&roMessages)

			getMsgsResp, err := client.R().
				SetHeader("Content-Type", "application/json").
				SetAuthToken(jwt).
				SetBody(rawRoMessages).
				Post("http://localhost:8002/v1.0.0")

			Expect(err).To(BeNil())
			Expect(getMsgsResp.StatusCode()).To(BeEquivalentTo(200))

			var getMessagesResponseObject did.ResponseObject
			err = json.Unmarshal(getMsgsResp.Body(), &getMessagesResponseObject)

			Expect(err).To(BeNil())
			Expect(getMessagesResponseObject.Status.Code).To(BeEquivalentTo(200))
			Expect(len(getMessagesResponseObject.Replies)).To(BeEquivalentTo(1))
			Expect(getMessagesResponseObject.Replies[0].Status).To(BeEquivalentTo(200))
		})

	})

})
