package ourjson

type JsonArray struct {
	s []*Value
}

func (j *JsonArray) Get(index int) (*Value, error) {
	if index >= len(j.s) {
		return nil, IndexOutOfRangeError
	}
	return j.s[index], nil
}

func (j *JsonArray) Values() ([]*Value) {
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
		return false , err
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