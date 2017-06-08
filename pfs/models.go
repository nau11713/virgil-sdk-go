package pfs

import "gopkg.in/virgil.v4"

// Credentials represent user's cards needed to establish a PFS session
type Credentials struct {
	IdentityCard *virgil.Card
	LTC          *virgil.Card
	OTC          *virgil.Card
}

type Recipient struct {
	LTC  *virgil.Card
	OTCs []*virgil.Card
}

type CredentialsResponse struct {
	IdentityCard *virgil.CardResponse `json:"identity_card"`
	LTC          *virgil.CardResponse `json:"long_time_card"`
	OTC          *virgil.CardResponse `json:"one_time_card"`
}

type CredentialsRequest struct {
	Identities []string `json:"identities"`
}

type CreateRecipientRequest struct {
	LTC  *virgil.SignableRequest   `json:"long_time_card"`
	OTCS []*virgil.SignableRequest `json:"one_time_cards"`
}

type CreateRecipientResponse struct {
	LTC  *virgil.CardResponse   `json:"long_time_card"`
	OTCS []*virgil.CardResponse `json:"one_time_cards"`
}

type Message struct {
	ID         string `json:"id,omitempty"`
	SessionId  []byte `json:"session_id,omitempty"`
	Eph        []byte `json:"eph,omitempty"`
	Signature  []byte `json:"sign,omitempty"`
	ICID       string `json:"ic_id"`
	LTCID      string `json:"ltc_id,omitempty"`
	OTCID      string `json:"otc_id,omitempty"`
	Salt       []byte `json:"salt"`
	Ciphertext []byte `json:"ciphertext"`
}