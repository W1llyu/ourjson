package ourjson

import (
	"strconv"
	"strings"
)

type JsonArray struct {
	s []*Value
}

func (j *JsonArray) compareTo(c *JsonArray) bool {
	if c == nil {
		return false
	}
	for i, v := range j.Values() {
		data := v.Data()
		switch data.(type) {
		case map[string]interface{}:
			v.JsonObject().compareTo(c.GetJsonObject(i))
		case []interface{}:
			v.JsonArray().compareTo(c.GetJsonArray(i))
		default:
			jv, err := c.Get(i)
			if err != nil || jv.data != v.data {
				return false
			}
		}
	}
	return true
}
func (j *JsonArray) String() string {
	str := "["
	for _, v := range j.s {
		data := v.Data()
		switch data.(type) {
		case string:
			str += `"` + v.data.(string) + `"`
		case int:
			str += strconv.Itoa(v.data.(int))
		case int64:
			str += strconv.FormatInt(v.data.(int64), 10)
		case float64:
			str += strconv.FormatFloat(v.data.(float64), 'f', -1, 64)
		case bool:
			{
				if v.data.(bool) {
					str += "true"
				} else {
					str += "false"
				}
			}
		case map[string]interface{}:
			str += v.JsonObject().String()
		case []interface{}:
			str += v.JsonArray().String()
		case nil:
			str += "null"
		}
		str += ","
	}
	str = strings.TrimRight(str, ",")
	str += "]"
	return str
}

func (j *JsonArray) Replace(fj *JsonArray) *JsonArray {
	array := make([]*Value, len(fj.Values()))
	for i, v := range fj.Values() {
		data := v.Data()
		switch data.(type) {
		case map[string]interface{}:
			nJson, _ := ParseObject("{}")
			nJson.Replace(v.JsonObject())
			array[i] = &Value{nJson.toInterface(), nJson}
		case []interface{}:
			nArray, _ := ParseArray("[]")
			nArray.Replace(v.JsonArray())
			array[i] = &Value{nArray.toInterface(), nArray}
		default:
			array[i] = &Value{data, nil}
		}
	}
	j.s = array
	return j
}

func (j *JsonArray) Clear() *JsonArray {
	j.s = make([]*Value, 0)
	return j
}

func (j *JsonArray) Remove(index int) *JsonArray {
	j.s = append(j.s[:index], j.s[index+1:]...)
	return j
}

func (j *JsonArray) Has(val interface{}) int {
	for i, v := range j.s {
		if v == val {
			return i
		}
	}
	return -1
}

func (j *JsonArray) Get(index int) (*Value, error) {
	if index >= len(j.s) {
		return nil, IndexOutOfRangeError
	}
	return j.s[index], nil
}

func (j *JsonArray) Values() []*Value {
	return j.s
}

func (j *JsonArray) GetJsonObject(index int) *JsonObject {
	val, err := j.Get(index)
	if err != nil {
		panic(err)
	}
	return val.JsonObject()
}

func (j *JsonArray) GetJsonArray(index int) *JsonArray {
	val, err := j.Get(index)
	if err != nil {
		panic(err)
	}
	return val.JsonArray()
}

func (j *JsonArray) GetString(index int) (string, error) {
	val, err := j.Get(index)
	if err != nil {
		return "", err
	}
	return val.String()
}

func (j *JsonArray) GetInt(index int) (int, error) {
	val, err := j.Get(index)
	if err != nil {
		return 0, err
	}
	return val.Int()
}

func (j *JsonArray) GetNullInt(index int) (*Integer, error) {
	val, err := j.Get(index)
	if err != nil {
		return nil, err
	}
	return val.NullInt()
}

func (j *JsonArray) GetInt64(index int) (int64, error) {
	val, err := j.Get(index)
	if err != nil {
		return 0, err
	}
	return val.Int64()
}

func (j *JsonArray) GetNullLong(index int) (*Long, error) {
	val, err := j.Get(index)
	if err != nil {
		return nil, err
	}
	return val.NullLong()
}

func (j *JsonArray) GetFloat64(index int) (float64, error) {
	val, err := j.Get(index)
	if err != nil {
		return 0, err
	}
	return val.Float64()
}

func (j *JsonArray) GetNullFloat(index int) (*Float, error) {
	val, err := j.Get(index)
	if err != nil {
		return nil, err
	}
	return val.NullFloat()
}

func (j *JsonArray) GetBoolean(index int) (bool, error) {
	val, err := j.Get(index)
	if err != nil {
		return false, err
	}
	return val.Boolean()
}

func (j *JsonArray) GetNullBoolean(index int) (*Boolean, error) {
	val, err := j.Get(index)
	if err != nil {
		return nil, err
	}
	return val.NullBoolean()
}

func (j *JsonArray) toInterface() []interface{} {
	m := make([]interface{}, len(j.s))
	for i, v := range j.s {
		m[i] = v.data
	}
	return m
}

func (j *JsonArray) Put(val interface{}) *JsonArray {
	j.s = append(j.s, &Value{inputConvert(val), nil})
	return j
}
