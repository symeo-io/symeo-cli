package values

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"gopkg.in/yaml.v2"
	"os"
)

func FetchFromApi(apiUrl string, apiKey string) (map[string]any, error) {
	httpClient := resty.New()

	var result map[string]map[string]any
	response, err := httpClient.
		R().
		SetResult(&result).
		SetHeader("X-API-KEY", apiKey).
		Get(apiUrl)

	if err != nil {
		return map[string]any{}, fmt.Errorf("FetchFromApi: Unable to complete api request [err=%s]", err)
	}

	if response.IsError() {
		return map[string]any{}, fmt.Errorf("FetchFromApi: Unsuccessful response: [response=%s]", response)
	}

	return result["values"], nil
}

func FetchFromFile(path string) (map[string]any, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return map[string]any{}, fmt.Errorf("FetchFromFile: Unable read local file [err=%s]", err)
	}

	var values map[string]any
	err = yaml.Unmarshal(data, &values)

	if err != nil {
		return map[string]any{}, fmt.Errorf("FetchFromFile: Unable parse yaml file [err=%s]", err)
	}

	return values, nil
}
