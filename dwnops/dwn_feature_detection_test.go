package dwnops_test

import (
	"encoding/json"
	"github.com/fairxio/protocol-integration-tests/auth"
	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type FeatureDetection struct {
	Type       string           `json:"type"`
	Interfaces FeatureInterface `json:"interfaces"`
}

type FeatureInterface struct {
	Collections CollectionsFeatures `json:"collections,omitempty"`
	Actions     ActionsFeatures     `json:"actions,omitempty"`
	Permissions PermissionsFeatures `json:"permissions,omitempty"`
	Messaging   MessagingFeatures   `json:"messaging,omitempty"`
	FairX       FairXFeatures       `json:"fairx,omitempty"`
}

type CollectionsFeatures struct {
	CollectionsQuery  bool `json:"CollectionsQuery"`
	CollectionsWrite  bool `json:"CollectionsWrite"`
	CollectionsCommit bool `json:"CollectionsCommit"`
	CollectionsDelete bool `json:"CollectionsDelete"`
}

type ActionsFeatures struct {
	ThreadsQuery  bool `json:"ThreadsQuery"`
	ThreadsCreate bool `json:"ThreadsCreate"`
	ThreadsReply  bool `json:"ThreadsReply"`
	ThreadsClose  bool `json:"ThreadsClose"`
	ThreadsDelete bool `json:"ThreadsDelete"`
}

type PermissionsFeatures struct {
	PermissionsRequest bool `json:"PermissionsRequest"`
	PermissionsGrant   bool `json:"PermissionsGrant"`
	PermissionsRevoke  bool `json:"PermissionsRevoke"`
}

type MessagingFeatures struct {
	Batching bool `json:"batching"`
}

type FairXFeatures struct {
	SessionEstablish       bool `json:"FairXSessionEstablish"`
	SessionExecuteFunction bool `json:"FairXSessionExecuteFunction"`
	SessionMessages        bool `json:"FairXSessionMessages"`
}

var _ = Describe("DWN Feature Detection", func() {

	Describe("Getting supported DWN Features", func() {

		// Authenticate
		jwt := auth.Authenticate("did:fairx:aW50ZWdyYXRpb250ZXN0QGZhaXJ4Lmlv")

		It("Returns a protocol-compliant response", func() {

			client := resty.New()
			resp, err := client.R().
				SetHeader("Content-Type", "application/json").
				SetAuthToken(jwt).
				Post("http://localhost:8002/v1.0.0")

			Expect(err).To(BeNil())
			Expect(resp.StatusCode()).To(BeEquivalentTo(200))

			var featureDetectionBody FeatureDetection
			bodyContent := resp.Body()
			Expect(bodyContent).ToNot(BeNil())

			err = json.Unmarshal(bodyContent, &featureDetectionBody)
			Expect(err).To(BeNil())
			Expect(featureDetectionBody.Type).To(BeEquivalentTo("FeatureDetection"))
			Expect(featureDetectionBody.Interfaces.FairX.SessionEstablish).To(BeTrue())
			Expect(featureDetectionBody.Interfaces.FairX.SessionExecuteFunction).To(BeTrue())
			Expect(featureDetectionBody.Interfaces.FairX.SessionMessages).To(BeTrue())

		})

	})

})
