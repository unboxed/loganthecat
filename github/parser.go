package github

import "encoding/json"

type Payload struct {
	Ref string `json:"ref"`
}

func ParseCommit(data []byte) (Payload, error) {
	payload := &Payload{}
	err := json.Unmarshal(data, payload)

	return *payload, err
}
