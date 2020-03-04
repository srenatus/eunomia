package iamutil

import (
	"encoding/json"
	"io/ioutil"
)

// Parse reads the content of an IAM policy file and returns
// an in-memory representation of it
func Parse(iamfile string) (pd PolicyDocument, err error) {
	pf, err := ioutil.ReadFile(iamfile)
	if err != nil {
		return pd, err
	}
	err = json.Unmarshal(pf, &pd)
	if err != nil {
		return pd, err
	}
	return pd, nil
}

// DumpJSON reads the content of an IAM policy file and returns
// a textual JSON rendering of it
func DumpJSON(iamfile string) (content string, err error) {
	pd, err := Parse(iamfile)
	if err != nil {
		return content, err
	}
	bc, err := json.MarshalIndent(pd, "", " ")
	if err != nil {
		return content, err
	}
	content = string(bc)
	return content, nil
}
