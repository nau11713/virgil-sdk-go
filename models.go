package virgilcards

import (
	"time"

	"gopkg.in/virgil.v5/crypto-api"
)

type CardSignature struct {
	SignerCardId string
	Signature    []byte
	ExtraFields  map[string]string
	SignerType   SignerType
	Snapshot     []byte
}

type Card struct {
	ID        string
	Identity  string
	PublicKey cryptoapi.PublicKey
	Version   string
	CreatedAt time.Time
	Signature []*CardSignature
	Snapshot  []byte
}

type RawCardSignature struct {
	SignerCardId string `json:"signer_id"`
	Signature    []byte `json:"signature"`
	ExtraFields  []byte `json:"snapshot,omitempty"`
	SignerType   string `json:"signer_type"`
}

type RawCardMeta struct {
	Signatures map[string][]byte `json:"signs"`
	CreatedAt  string            `json:"created_at"`
	Version    string            `json:"card_version"`
}

type RawCardSnapshot struct {
	Identity       string `json:"identity"`
	PublicKeyBytes []byte `json:"public_key"`
	PreviousCardID string `json:"previous_card_id"`
	Version        string `json:"version"`
	CreatedAt      int64  `json:"created_at"`
}
type RawCard struct {
	Snapshot   []byte              `json:"content_snapshot"`
	Signatures []*RawCardSignature `json:"signatures"`
	Meta       *RawCardMeta        `json:"meta,omitempty"`
}
