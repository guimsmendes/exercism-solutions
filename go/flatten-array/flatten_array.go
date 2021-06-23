package flatten

func Flatten(arr interface{}) []interface{} {
	arr2 := []interface{}{}
	for _, v := range arr.([]interface{}) {
		switch v.(type) {
		case int:
			arr2 = append(arr2, v)
		case []interface{}:
			arr2 = append(arr2, Flatten(v)...)
		}
	}
	return arr2
}
