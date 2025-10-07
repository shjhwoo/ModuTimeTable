package util

func SafeStr(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

func StringPtr(s string) *string {
	return &s
}

func SafeInt64(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

func SafeInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}
