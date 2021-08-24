//

package ourjson

import "encoding/json"

func New() *JsonObject {
	j, _ := ParseObject("{}")
	return j
}

func ParseObject(jsonStr string) (*JsonObject, error) {
	value := new(Value)

	err := json.Unmarshal([]byte(jsonStr), &value.data)
	if err != nil {
		return nil, err
	}

	return value.JsonObject(), nil
}

func FromObject(json *JsonObject) *JsonObject {
	j := New()
	return j.Replace(json)
}

func NewArray() *JsonArray {
	ja, _ := ParseArray("[]")
	return ja
}

func FromArray(json *JsonArray) *JsonArray {
	j := NewArray()
	return j.Replace(json)
}

func ParseArray(jsonStr string) (*JsonArray, error) {
	value := new(Value)

	err := json.Unmarshal([]byte(jsonStr), &value.data)
	if err != nil {
		return nil, err
	}
	return value.JsonArray(), nil
}
