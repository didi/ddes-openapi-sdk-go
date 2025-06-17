package core

import "encoding/json"

type Serializer interface {
	Serialize(interface{}) ([]byte, error)
	Deserialize([]byte, interface{}) error
}

type DefaultSerializer struct{}

func (s *DefaultSerializer) Serialize(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
func (s *DefaultSerializer) Deserialize(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
