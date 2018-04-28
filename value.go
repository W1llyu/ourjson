package ourjson

import (
	"encoding/json"
	"strconv"
)

const (
	JSONOBJECTTYPE = "JsonObject"
	JSONARRAYTYPE  = "JsonArray"
	STRINGTYPE     = "string"
	BOOLEANTYPE    = "bool"
)

type Value struct {
	data interface {}
}

func (v *Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.data)
}

func (v *Value) JsonObject() *JsonObject {
	if _, ok := v.data.(map[string]interface {}); !ok {
		panic(ValueTransformTypeError{JSONOBJECTTYPE})
	}
	mapValue := make(map[string]*Value)
	for key, val := range v.data.(map[string]interface {}) {
		mapValue[key] = &Value{val}
	}
	return &JsonObject{
		//Value: Value{v.data},
		m: mapValue,
	}
}

func (v *Value) JsonArray() *JsonArray {
	if _, ok := v.data.([]interface {}); !ok {
		panic(ValueTransformTypeError{JSONARRAYTYPE})
	}
	var slice []*Value
	for _, val := range v.data.([]interface {}) {
		slice = append(slice, &Value{val})
	}
	return &JsonArray{
		s: slice,
	}
}

func (v *Value) String() (string, error) {
	if v.data == nil {
		return "", nil
	}
	if _, ok := v.data.(string); !ok {
		return "", ValueTransformTypeError{STRINGTYPE}
	}
	return v.data.(string), nil
}

func (v *Value) Int() (int, error) {
	f, err := v.Float64()
	if err != nil {
		return 0, err
	}
	return int(f), nil
}

func (v *Value) NullInt() (*Integer, error) {
	i := new(Integer)
	nf, err := v.NullFloat()
	if err != nil {
		return i, err
	}
	i.Valid = nf.Valid
	i.Value = int(nf.Value)
	return i, nil
}

func (v *Value) Int64() (int64, error) {
	f, err := v.Float64()
	if err != nil {
		return 0, err
	}
	return int64(f), nil
}

func (v *Value) NullLong() (*Long, error) {
	l := new(Long)
	nf, err := v.NullFloat()
	if err != nil {
		return l, err
	}
	l.Valid = nf.Valid
	l.Value = int64(nf.Value)
	return l, nil
}

func (v *Value) Float64() (float64, error) {
	switch v.data.(type) {
	default:
		return 0, ValueNotNumberError
	case float64:
		return v.data.(float64), nil
	case string:
		return strconv.ParseFloat(v.data.(string), 64)
	}
}

func (v *Value) NullFloat() (*Float, error) {
	nf := new(Float)
	if v.data == nil {
		return nf, nil
	}
	f, err := v.Float64()
	if err != nil {
		return nf, err
	}
	nf.Valid = true
	nf.Value = f
	return nf, nil
}

func (v *Value) Boolean() (bool, error) {
	if _, ok := v.data.(bool); !ok {
		return false, ValueTransformTypeError{BOOLEANTYPE}
	}
	return v.data.(bool), nil
}

func (v *Value) NullBoolean() (*Boolean, error) {
	nb := new(Boolean)
	if v.data == nil {
		return nb, nil
	}
	b, err := v.Boolean()
	if err != nil {
		return nb, err
	}
	nb.Valid = true
	nb.Value = b
	return nb, nil
}