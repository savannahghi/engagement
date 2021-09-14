package domain

import (
	"fmt"

	"github.com/savannahghi/feedlib"
)

// Feed manages and serializes the nudges, actions and feed items that a
// specific user should see.
//
// A feed is stored and serialized on a per-user basis. If a feed item is sent
// to a group of users, it should be "expanded" before the user's feed gets
// stored.
type Feed struct {
	// a string composed by concatenating the UID, a "|" and a flavour
	ID string `json:"id" firestore:"-"`

	// A higher sequence number means that it came later
	SequenceNumber int `json:"sequenceNumber" firestore:"sequenceNumber"`

	// user identifier - who does this feed belong to?
	// this is also the unique identifier for a feed
	UID string `json:"uid" firestore:"uid"`

	// whether this is a consumer or pro feed
	Flavour feedlib.Flavour `json:"flavour" firestore:"flavour"`

	// what are the global actions available to this user?
	Actions []feedlib.Action `json:"actions" firestore:"actions"`

	// what does this user's feed contain?
	Items []feedlib.Item `json:"items" firestore:"items"`

	// what prompts or nudges should this user see?
	Nudges []feedlib.Nudge `json:"nudges" firestore:"nudges"`

	// indicates whether the user is Anonymous or not
	IsAnonymous *bool `json:"isAnonymous" firestore:"isAnonymous"`

	FeatureImage string `json:"feature_image"`
}

// GetID return the feed ID
func (fe Feed) GetID() string {
	return fmt.Sprintf("%s|%s", fe.UID, fe.Flavour.String())
}

// IsEntity marks a feed as an Apollo federation GraphQL entity
func (fe Feed) IsEntity() {}

// ValidateAndUnmarshal checks that the input data is valid as per the
// relevant JSON schema and unmarshals it if it is
func (fe *Feed) ValidateAndUnmarshal(b []byte) error {
	err := feedlib.ValidateAndUnmarshal(feedlib.FeedSchemaFile, b, fe)
	if err != nil {
		return fmt.Errorf("invalid feed JSON: %w", err)
	}
	return nil
}

// ValidateAndMarshal validates against the JSON schema then marshals to JSON
func (fe *Feed) ValidateAndMarshal() ([]byte, error) {
	return feedlib.ValidateAndMarshal(feedlib.FeedSchemaFile, fe)
}

// EMailMessage holds data required to send emails
type EMailMessage struct {
	Subject string   `json:"subject,omitempty"`
	Text    string   `json:"text,omitempty"`
	To      []string `json:"to,omitempty"`
}
