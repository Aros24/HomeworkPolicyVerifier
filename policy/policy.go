package policy

import (
	"encoding/json"
)

type Policy struct {
	PolicyName     string         `json:"PolicyName"`
	PolicyDocument PolicyDocument `json:"PolicyDocument"`
}

type PolicyDocument struct {
	Version   string      `json:"Version"`
	Statement []Statement `json:"Statement"`
}

type Statement struct {
	Sid       string      `json:"Sid,omitempty"`
	Effect    string      `json:"Effect"`
	Action    SliceString `json:"Action"`
	Resource  SliceString `json:"Resource"`
	Condition *Condition  `json:"Condition,omitempty"`
}

type SliceString []string

type Condition struct {
	Bool map[string]string `json:"Bool"`
}

// Overload if has only one or multiple values (Array)
func (textToMarshall *SliceString) UnmarshalJSON(data []byte) error {
	var singleValue string
	if err := json.Unmarshal(data, &singleValue); err == nil {
		*textToMarshall = []string{singleValue}
		return nil
	}

	var multiValue []string
	if err := json.Unmarshal(data, &multiValue); err == nil {
		*textToMarshall = multiValue
		return nil
	}

	return UnmarshalError("invalid format for sliced string")
}

func UnmarshalPolicy(data []byte) (*Policy, error) {
	var p Policy
	err := json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
