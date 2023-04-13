package values

import (
	"fmt"
	"reflect"
	"symeo/cli/packages/contracts"
)

func CheckContractCompatibility(contract map[string]any, values map[string]any) []string {
	return executeCompatibilityCheck(contract, values, "")
}

func executeCompatibilityCheck(contract map[string]any, values map[string]any, parentPath string) []string {
	var errors []string

	for propertyName, contractProperty := range contract {
		valueProperty := values[propertyName]

		if !contracts.IsContractProperty(contracts.AnyToContract(contractProperty)) && isUndefined(valueProperty) {
			errors = append(errors, buildMissingPropertyError(propertyName, parentPath))
			continue
		}

		if !contracts.IsContractProperty(contracts.AnyToContract(contractProperty)) && isDefined(valueProperty) {
			errors = append(errors, executeCompatibilityCheck(contracts.AnyToContract(contractProperty), AnyToValues(valueProperty), buildParentPath(propertyName, parentPath))...)
			continue
		}

		if isUndefined(valueProperty) && !contracts.IsContractPropertyOptional(contracts.AnyToContract(contractProperty)) {
			errors = append(errors, buildMissingPropertyError(propertyName, parentPath))
			continue
		}

		if isDefined(valueProperty) && !contractPropertyAndValueHaveSameType(contracts.AnyToContract(contractProperty), valueProperty) {
			errors = append(errors, buildWrongTypeError(propertyName, parentPath, contracts.AnyToContract(contractProperty), valueProperty))
			continue
		}
	}

	return errors
}

func contractPropertyAndValueHaveSameType(contractProperty map[string]any, value any) bool {
	propertyType := contractProperty["type"]

	switch propertyType {
	case "string":
		return reflect.TypeOf(value).Kind() == reflect.String
	case "boolean":
		return reflect.TypeOf(value).Kind() == reflect.Bool
	case "integer":
		return isInteger(value)
	case "float":
		return reflect.TypeOf(value).Kind() == reflect.Int || reflect.TypeOf(value).Kind() == reflect.Float32 || reflect.TypeOf(value).Kind() == reflect.Float64
	}

	return false
}

func buildMissingPropertyError(propertyName string, parentPath string) string {
	displayedPropertyName := buildParentPath(propertyName, parentPath)

	return fmt.Sprintf("The property \"%s\" of your configuration contract is missing in your configuration values.", displayedPropertyName)
}

func buildWrongTypeError(propertyName string, parentPath string, contractProperty map[string]any, value any) string {
	displayedPropertyName := buildParentPath(propertyName, parentPath)

	return fmt.Sprintf("The property \"%s\" has type \"%s\" while configuration contract defined \"%s\" as \"%s\".", displayedPropertyName, displayValueType(value), displayedPropertyName, contractProperty["type"])
}

func buildParentPath(propertyName string, parentPath string) string {
	if parentPath == "" {
		return propertyName
	}

	return parentPath + "." + propertyName
}

func isDefined(value any) bool {
	return value != nil && value != ""
}

func isUndefined(value any) bool {
	return !isDefined(value)
}

func displayValueType(value any) string {
	valueType := reflect.TypeOf(value).Kind()

	switch valueType {
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "boolean"
	case reflect.Int:
		return "integer"
	case reflect.Float32:
	case reflect.Float64:
		if isInteger(value) {
			return "integer"
		}
		return "float"
	}

	return valueType.String()
}

func isInteger(value any) bool {
	valueType := reflect.TypeOf(value).Kind()

	if valueType == reflect.Int {
		return true
	}

	if valueType == reflect.Float32 || valueType == reflect.Float64 {
		floatValue := value.(float64)
		return floatValue == float64(int(floatValue))
	}

	return false
}
