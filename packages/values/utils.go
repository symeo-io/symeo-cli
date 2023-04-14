package values

func AnyToValues(el any) map[string]any {
	values := make(map[string]any)

	if el == nil {
		return values
	}

	if castValues, ok := el.(map[string]any); ok {
		return castValues
	}

	for propertyName, valuesProperty := range el.(map[any]any) {
		values[propertyName.(string)] = valuesProperty
	}

	return values
}
