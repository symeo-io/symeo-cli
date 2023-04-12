package contracts

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

func LoadContractFile(path string) (Contract, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return Contract{}, fmt.Errorf("FetchFromFile: Unable read local file [err=%s]", err)
	}

	var contract Contract
	err = yaml.Unmarshal(data, &contract)

	if err != nil {
		return Contract{}, fmt.Errorf("FetchFromFile: Unable parse yaml file [err=%s]", err)
	}

	return contract, nil
}
