package pipes

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

func FromHttpBody2Map(body io.ReadCloser) (map[string]interface{}, error) {
	var resBody bytes.Buffer
	_, err := io.Copy(&resBody, body)
	defer body.Close()

	if err != nil {
		return nil, errors.New("failed to read body")
	}

	var mapped_body map[string]interface{}

	if err := json.Unmarshal(resBody.Bytes(), &mapped_body); err != nil {
		return nil, errors.New("failed to unmarshal body")
	}

	return mapped_body, nil
}

// `v` should be a pointer to a struct
func Body2Struct[T any](body io.ReadCloser, v T) error {
	// read body -> bytes' buffer
	var resBody bytes.Buffer
	_, err := io.Copy(&resBody, body)
	defer body.Close()
	if err != nil {
		return errors.New("failed to read body")
	}

	return json.Unmarshal(resBody.Bytes(), &v)
}
