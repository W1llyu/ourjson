//

package ourjson

import "encoding/json"

func ParseObject(jsonStr string) (*JsonObject, error) {
	value := new(Value)

	err := json.Unmarshal([]byte(jsonStr), &value.data)
	if err != nil {
		return nil, err
	}

	return value.JsonObject(), nil
}

func ParseArray(jsonStr string) (*JsonArray, error) {
	value := new(Value)

	err := json.Unmarshal([]byte(jsonStr), &value.data)
	if err != nil {
		return nil, err
	}
	return value.JsonArray(), nil
}