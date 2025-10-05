package util

func ParseQueryArrayToInt64List(queryArray []string) []*int64 {
	var result []*int64
	for _, str := range queryArray {
		val := ParseInt64(str)
		result = append(result, &val)
	}

	return result
}
