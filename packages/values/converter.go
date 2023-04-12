package values

import (
	"fmt"
	"github.com/gobeam/stringy"
	"reflect"
)

func ValuesToEnv(values Values) []string {
	var env []string
	envMap := convertValuesToEnv(values, "")

	for key, value := range envMap {
		if value != "" {
			s := key + "=" + value
			env = append(env, s)
		}
	}

	return env
}

func convertValuesToEnv(values Values, path string) map[string]string {
	env := make(map[string]string)

	for propertyName, valuesProperty := range values {
		if valuesProperty != nil && reflect.TypeOf(valuesProperty).Kind() == reflect.Map {
			subEnv := convertValuesToEnv(AnyToValues(valuesProperty), concatPath(path, propertyName))

			for k, v := range subEnv {
				env[k] = v
			}
		} else {
			key := stringy.New(concatPath(path, propertyName)).SnakeCase().ToUpper()
			env[key] = toString(valuesProperty)
		}
	}

	return env
}

func concatPath(path string, propertyName string) string {
	if path == "" {
		return propertyName
	}

	return path + "_" + propertyName
}

func toString(value any) string {
	if value == nil {
		return ""
	}

	return fmt.Sprintf("%v", value)
}
