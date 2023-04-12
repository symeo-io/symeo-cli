package values

func AnyToValues(el any) Values {
	values := make(Values)

	if el == nil {
		return values
	}

	for propertyName, contractProperty := range el.(map[any]any) {
		values[propertyName.(string)] = contractProperty
	}

	return values
}
