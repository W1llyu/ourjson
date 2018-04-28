package ourjson

import "strconv"

// ---------- Boolean ----------
type Boolean struct {
	Valid bool
	Value bool
}

func (b *Boolean) UnmarshalJSON(data []byte) error {
	boolValue, err := strconv.ParseBool(string(data))
	b.Valid = err == nil
	if err == nil {
		b.Value = boolValue
	}
	return err
}

// ---------- Integer ----------
type Integer struct {
	Valid bool
	Value int
}

func (b *Integer) UnmarshalJSON(data []byte) error {
	intValue, err := strconv.Atoi(string(data))
	b.Valid = err == nil
	if err == nil {
		b.Value = intValue
	}
	return nil
}

// ---------- Long ----------
type Long struct {
	Valid bool
	Value int64
}

func (b *Long) UnmarshalJSON(data []byte) error {
	longValue, err := strconv.ParseInt(string(data), 10, 64)
	b.Valid = err == nil
	if err == nil {
		b.Value = longValue
	}
	return nil
}

// ---------- Float ----------
type Float struct {
	Valid bool
	Value float64
}

func (b *Float) UnmarshalJSON(data []byte) error {
	floatValue, err := strconv.ParseFloat(string(data), 32)
	b.Valid = err == nil
	if err == nil {
		b.Value = floatValue
	}
	return nil
}
