//

package ourjson

import (
	"strconv"
	"strings"
)

type JsonObject struct {
	m map[string]*Value
}

func (j *JsonObject) compareTo(c *JsonObject) bool {
	if c == nil {
		return false
	}
	for k, v := range j.Values() {
		data := v.Data()
		switch data.(type) {
		case map[string]interface{}:
			v.JsonObject().compareTo(c.GetJsonObject(k))
		case []interface{}:
			v.JsonArray().compareTo(c.GetJsonArray(k))
		default:
			jv, err := c.Get(k)
			if err != nil || jv.data != v.data {
				return false
			}
		}
	}
	return true
}

func (j *JsonObject) String() string {
	str := "{"
	for k, v := range j.m {
		str += `"` + k + `":`
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
		// case JsonObject:
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
	str += "}"
	return str
}

func (j *JsonObject) Replace(fj *JsonObject) *JsonObject {
	json := make(map[string]*Value)
	for k, v := range fj.Values() {
		data := v.Data()
		switch data.(type) {
		case map[string]interface{}:
			nJson, _ := ParseObject("{}")
			nJson.Replace(v.JsonObject())
			json[k] = &Value{nJson.toInterface(), nJson}
		case []interface{}:
			nArray, _ := ParseArray("[]")
			nArray.Replace(v.JsonArray())
			json[k] = &Value{nArray.toInterface(), nArray}
		default:
			json[k] = &Value{data, nil}
		}
	}
	j.m = json
	return j
}

func (j *JsonObject) Clear() *JsonObject {
	j.m = make(map[string]*Value)
	return j
}

func (j *JsonObject) Remove(key string) *JsonObject {
	delete(j.m, key)
	return j
}

// Check if the key is existed
func (j *JsonObject) HasKey(key string) bool {
	if j.m == nil {
		return false
	}
	_, ok := j.m[key]
	return ok
}

func (j *JsonObject) Get(key string) (*Value, error) {
	if !j.HasKey(key) {
		return nil, KeyNotFoundError{key}
	}
	return j.m[key], nil
}

func (j *JsonObject) Values() map[string]*Value {
	return j.m
}

// Get a child node of JsonObject from this parent node
// This will raise an error when the value is not existed or not a JsonObject
// Considering having a recover function with it
func (j *JsonObject) GetJsonObject(key string) *JsonObject {
	val, err := j.Get(key)
	if err != nil {
		// panic(err)
		return nil
	}
	return val.JsonObject()
}

// Get a child node of JsonArray from this parent node
// This will raise an error when the value if not existed or not a JsonArray
// Considering having a recover function with it
func (j *JsonObject) GetJsonArray(key string) *JsonArray {
	val, err := j.Get(key)
	if err != nil {
		// panic(err)
		return nil
	}
	return val.JsonArray()
}

func (j *JsonObject) GetString(key string) (string, error) {
	val, err := j.Get(key)
	if err != nil {
		return "", err
	}
	return val.String()
}

func (j *JsonObject) GetInt(key string) (int, error) {
	val, err := j.Get(key)
	if err != nil {
		return 0, err
	}
	return val.Int()
}

// Get an integer which could be null
// example:
//    i, _ := jsonObject.GetNullInt("age")
//    if !i.Valid {
//        fmt.Println("null age")
//    } else {
//        fmt.Printf("the age is %d", i.Value)
// 	  }
//
func (j *JsonObject) GetNullInt(key string) (*Integer, error) {
	val, err := j.Get(key)
	if err != nil {
		return nil, err
	}
	return val.NullInt()
}

func (j *JsonObject) GetInt64(key string) (int64, error) {
	val, err := j.Get(key)
	if err != nil {
		return 0, err
	}
	return val.Int64()
}

func (j *JsonObject) GetNullLong(key string) (*Long, error) {
	val, err := j.Get(key)
	if err != nil {
		return nil, err
	}
	return val.NullLong()
}

func (j *JsonObject) GetFloat64(key string) (float64, error) {
	val, err := j.Get(key)
	if err != nil {
		return 0, err
	}
	return val.Float64()
}

func (j *JsonObject) GetNullFloat(key string) (*Float, error) {
	val, err := j.Get(key)
	if err != nil {
		return nil, err
	}
	return val.NullFloat()
}

func (j *JsonObject) GetBoolean(key string) (bool, error) {
	val, err := j.Get(key)
	if err != nil {
		return false, err
	}
	return val.Boolean()
}

// Get a boolean which could be null
// example:
//    isHit, _ := jsonObject.GetNullBoolean("is_hit")
//    if !i.Valid {
//        fmt.Println("unknown")
//        return
//    }
//	  if i.Value {
//        fmt.Println("hit!")
//    } else {
// 	      fmt.Println("now hit)
//    }
//
func (j *JsonObject) GetNullBoolean(key string) (*Boolean, error) {
	val, err := j.Get(key)
	if err != nil {
		return nil, err
	}
	return val.NullBoolean()
}

func (j *JsonObject) toInterface() map[string]interface{} {
	m := make(map[string]interface{})
	for k, v := range j.m {
		m[k] = v.data
	}
	return m
}

func (j *JsonObject) Put(key string, val interface{}) *JsonObject {
	j.m[key] = &Value{inputConvert(val), nil}
	return j
}
