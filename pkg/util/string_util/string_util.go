package string_util

func GetStringFromPointer(ptrString *string, defaultString string) string {
	if ptrString == nil {
		return defaultString
	}
	return *ptrString
}
