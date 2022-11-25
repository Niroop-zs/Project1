package TypeTest

func typeof(a interface{}) interface{} {
	switch a.(type) {
	case int:
		return a.(int) * 10
	case float64:
		return a.(float64) * 10.5
	case string:
		return "hello" + a.(string)
	case bool:
		return !a.(bool)
	default:
		return "unknown"
	}
}
