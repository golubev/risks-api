package risk

import "github.com/google/uuid"

func (s State) IsValid() bool {
	switch s {
	case Open, Closed, Accepted, Investigating:
		return true
	}

	return false
}

type State string

const (
	Open          State = "open"
	Closed        State = "closed"
	Accepted      State = "accepted"
	Investigating State = "investigating"
)

// swagger:model
type RiskBody struct {
	State       State  `json:"state" example:"investigating" binding:"required" extensions:"x-order=1"`
	Title       string `json:"title" example:"CVE-2022-29217" extensions:"x-order=2"`
	Description string `json:"description" example:"python-jose through 3.3.0 has algorithm confusion with OpenSSH ECDSA keys and other key formats." extensions:"x-order=3"`
}

// swagger:model
type Risk struct {
	ID uuid.UUID `json:"id" example:"add736b0-516b-401c-a4ee-bfa00812bb52" extensions:"x-order=0"`
	RiskBody
}

var RisksStorage = []Risk{}
