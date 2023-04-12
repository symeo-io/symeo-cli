package contracts

import (
	"reflect"
)

func IsContractProperty(el Contract) bool {
	return el["type"] != nil && reflect.TypeOf(el["type"]).Kind() == reflect.String
}

func IsContractPropertyOptional(el Contract) bool {
	return el["optional"] != nil && el["optional"] == true
}

func AnyToContract(el any) Contract {
	contract := make(Contract)

	for propertyName, contractProperty := range el.(map[any]any) {
		contract[propertyName.(string)] = contractProperty
	}

	return contract
}
