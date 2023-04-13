package values

import (
	"github.com/go-faker/faker/v4"
	"reflect"
	"testing"
)

func TestInitializeValues(t *testing.T) {
	contract := map[string]any{
		"database": map[string]any{
			"host": map[string]any{
				"type":     "string",
				"optional": true,
			},
			"password": map[string]any{
				"type":     "string",
				"optional": true,
			},
			"responseLimit": map[string]any{
				"type": "float",
			},
		},
		"vcsProvider": map[string]any{
			"paginationLength": map[string]any{
				"type": "integer",
			},
		},
		"auth0": map[string]any{
			"isAdmin": map[string]any{
				"type": "boolean",
			},
		},
	}

	host := faker.URL()
	responseLimit, _ := faker.RandomInt(0)
	paginationLength, _ := faker.RandomInt(0)
	values := map[string]any{
		"database": map[string]any{
			"host":          host,
			"responseLimit": responseLimit,
		},
		"vcsProvider": map[string]any{
			"paginationLength": paginationLength,
		},
	}

	expectedValues := map[string]any{
		"database": map[string]any{
			"host":          host,
			"password":      nil,
			"responseLimit": responseLimit,
		},
		"vcsProvider": map[string]any{
			"paginationLength": paginationLength,
		},
		"auth0": map[string]any{
			"isAdmin": nil,
		},
	}

	initializedValues := InitializeValues(contract, values)

	if !reflect.DeepEqual(initializedValues, expectedValues) {
		t.Errorf("got %q, wanted %q", initializedValues, expectedValues)
	}
}
