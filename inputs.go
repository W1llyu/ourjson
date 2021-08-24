package ourjson

func inputConvert(val interface{}) interface{} {
	var r interface{}
	switch val.(type) {
	case *JsonObject:
		r = val.(*JsonObject).toInterface()
	case JsonObject:
		v := val.(JsonObject)
		p := &v
		r = p.toInterface()
	case *JsonArray:
		r = val.(*JsonArray).toInterface()
	case JsonArray:
		v := val.(JsonArray)
		p := &v
		r = p.toInterface()
	case Float:
		r = val.(Float).Value
	case Long:
		r = val.(Long).Value
	default:
		r = val
	}
	return r
}
