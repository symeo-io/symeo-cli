package contracts

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

func LoadContractFile(path string) (map[string]any, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return map[string]any{}, fmt.Errorf("FetchFromFile: Unable read local file [err=%s]", err)
	}

	var contract map[string]any
	err = yaml.Unmarshal(data, &contract)

	if err != nil {
		return map[string]any{}, fmt.Errorf("FetchFromFile: Unable parse yaml file [err=%s]", err)
	}

	return contract, nil
}
