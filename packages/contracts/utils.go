package contracts

import (
	"reflect"
)

func IsContractProperty(el map[string]any) bool {
	return el["type"] != nil && reflect.TypeOf(el["type"]).Kind() == reflect.String
}

func IsContractPropertyOptional(el map[string]any) bool {
	return el["optional"] != nil && el["optional"] == true
}

func HasContractPropertyRegex(el map[string]any) bool {
	return el["regex"] != nil && el["regex"] != ""
}

func AnyToContract(el any) map[string]any {
	contract := make(map[string]any)

	if el == nil {
		return contract
	}

	if castContract, ok := el.(map[string]any); ok {
		return castContract
	}

	for propertyName, contractProperty := range el.(map[any]any) {
		contract[propertyName.(string)] = contractProperty
	}

	return contract
}
