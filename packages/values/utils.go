package values

func AnyToValues(el any) Values {
	values := make(Values)

	if el == nil {
		return values
	}

	if castValues, ok := el.(Values); ok {
		return castValues
	}

	for propertyName, valuesProperty := range el.(map[any]any) {
		values[propertyName.(string)] = valuesProperty
	}

	return values
}
