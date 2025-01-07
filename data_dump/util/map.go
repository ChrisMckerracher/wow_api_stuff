package util

func Clone(start map[string]interface{}) map[string]interface{} {
	newMap := make(map[string]interface{})

	for k, v := range start {
		newMap[k] = v
	}

	return newMap
}
