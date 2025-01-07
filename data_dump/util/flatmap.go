package util

import (
	"fmt"
	
)

func Flatten(nestedMap map[string]interface{}) map[string]interface{} {
	newMap := make(map[string]interface{})
	for k, v := range nestedMap {
		if reflect.ValueOf(v).Kind() == reflect.Map {
			flattenedChild := Flatten(v.(map[string]interface{}))
			for fk, fv := range flattenedChild {
				newMap[fmt.Sprintf("%s:%s", k, fk)] = fv
			}
		} else {
			newMap[k] = v
		}
	}

	return newMap
}
