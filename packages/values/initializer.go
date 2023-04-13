package values

import (
	"symeo/cli/packages/contracts"
)

func InitializeValues(contract map[string]any, values map[string]any) map[string]any {
	initializedValues := make(map[string]any)
	for propertyName, contractProperty := range contract {
		valuesProperty := values[propertyName]

		if !contracts.IsContractProperty(contracts.AnyToContract(contractProperty)) {
			initializedValues[propertyName] = InitializeValues(contracts.AnyToContract(contractProperty), AnyToValues(valuesProperty))
		} else {
			initializedValues[propertyName] = valuesProperty
		}
	}

	return initializedValues
}
