// 

package ourjson

type JsonObject struct {
	m map[string]*Value
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

// Get a child node of JsonObject from this parent node
// This will raise an error when the value is not existed or not a JsonObject
// Considering having a recover function with it
func (j *JsonObject) GetJsonObject(key string) *JsonObject {
	val, err := j.Get(key)
	if err != nil {
		panic(err)
	}
	return val.JsonObject()
}

// Get a child node of JsonArray from this parent node
// This will raise an error when the value if not existed or not a JsonArray
// Considering having a recover function with it
func (j *JsonObject) GetJsonArray(key string) *JsonArray {
	val, err := j.Get(key)
	if err != nil {
		panic(err)
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
		return false , err
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

func (j *JsonObject) Put(key string, val interface {}) {
	j.m[key] = &Value{val}
}