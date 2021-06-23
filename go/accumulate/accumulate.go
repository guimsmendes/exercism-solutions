package accumulate

func Accumulate(arr []string, function func(string) string) []interface{} {
	var arr2 []interface{}
	for _, v := range arr {
		arr2 = append(arr2, function(v))
	}
	return arr2
}
