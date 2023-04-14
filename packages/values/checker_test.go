package values

import (
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

var contract = map[string]any{
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
	"currentUser": map[string]any{
		"email": map[string]any{
			"type":  "string",
			"regex": "^([a-zA-Z0-9_\\-\\.]+)@([a-zA-Z0-9_\\-\\.]+).([a-zA-Z]{2,5})$",
		},
	},
}

func TestCheckContractCompatibilityForMissingProperty(t *testing.T) {
	host := faker.URL()
	responseLimit := 100
	paginationLength := 100
	email := faker.Email()
	values := map[string]any{
		"badName": map[string]any{
			"host":          host,
			"responseLimit": responseLimit,
		},
		"vcsProvider": map[string]any{
			"paginationLength": paginationLength,
		},
		"auth0": map[string]any{
			"isAdmin": true,
		},
		"currentUser": map[string]any{
			"email": email,
		},
	}

	errors := CheckContractCompatibility(contract, values)

	assert.Equal(t, 1, len(errors))
	assert.Contains(t, errors, "The property \"database\" of your configuration contract is missing in your configuration values.")
}

func TestCheckContractCompatibilityForMissingNonOptionalValue(t *testing.T) {
	host := faker.URL()
	password := faker.Password()
	paginationLength := 100
	email := faker.Email()
	values := map[string]any{
		"database": map[string]any{
			"host":     host,
			"password": password,
		},
		"vcsProvider": map[string]any{
			"paginationLength": paginationLength,
		},
		"auth0": map[string]any{
			"isAdmin": true,
		},
		"currentUser": map[string]any{
			"email": email,
		},
	}

	errors := CheckContractCompatibility(contract, values)

	assert.Equal(t, 1, len(errors))
	assert.Contains(t, errors, "The property \"database.responseLimit\" of your configuration contract is missing in your configuration values.")
}

func TestCheckContractCompatibilityForDifferentTypeFromContract(t *testing.T) {
	host := faker.URL()
	password := faker.Password()
	responseLimit := "wrongType"
	paginationLength := 100
	email := faker.Email()
	values := map[string]any{
		"database": map[string]any{
			"host":          host,
			"password":      password,
			"responseLimit": responseLimit,
		},
		"vcsProvider": map[string]any{
			"paginationLength": paginationLength,
		},
		"auth0": map[string]any{
			"isAdmin": true,
		},
		"currentUser": map[string]any{
			"email": email,
		},
	}

	errors := CheckContractCompatibility(contract, values)

	assert.Equal(t, 1, len(errors))
	assert.Contains(t, errors, "The property \"database.responseLimit\" has type \"string\" while configuration contract defined \"database.responseLimit\" as \"float\".")
}

func TestCheckContractCompatibilityForInValidValuesForRegex(t *testing.T) {
	host := faker.URL()
	responseLimit := 100
	paginationLength := 100
	email := "nonMatchingString"
	values := map[string]any{
		"database": map[string]any{
			"host":          host,
			"responseLimit": responseLimit,
		},
		"vcsProvider": map[string]any{
			"paginationLength": paginationLength,
		},
		"auth0": map[string]any{
			"isAdmin": true,
		},
		"currentUser": map[string]any{
			"email": email,
		},
	}

	errors := CheckContractCompatibility(contract, values)

	assert.Equal(t, 1, len(errors))
	assert.Contains(t, errors, "The property \"currentUser.email\" with value \"nonMatchingString\" does not match regex \"^([a-zA-Z0-9_\\-\\.]+)@([a-zA-Z0-9_\\-\\.]+).([a-zA-Z]{2,5})$\" defined in contract.")
}

func TestCheckContractCompatibilityForValidValues(t *testing.T) {
	host := faker.URL()
	password := faker.Password()
	responseLimit := 100
	paginationLength := 100
	email := faker.Email()
	values := map[string]any{
		"database": map[string]any{
			"host":          host,
			"password":      password,
			"responseLimit": responseLimit,
		},
		"vcsProvider": map[string]any{
			"paginationLength": paginationLength,
		},
		"auth0": map[string]any{
			"isAdmin": true,
		},
		"currentUser": map[string]any{
			"email": email,
		},
	}

	errors := CheckContractCompatibility(contract, values)

	assert.Equal(t, 0, len(errors))
}

func TestCheckContractCompatibilityForValidValuesWithoutOptionals(t *testing.T) {
	host := faker.URL()
	responseLimit := 100
	paginationLength := 100
	email := faker.Email()
	values := map[string]any{
		"database": map[string]any{
			"host":          host,
			"responseLimit": responseLimit,
		},
		"vcsProvider": map[string]any{
			"paginationLength": paginationLength,
		},
		"auth0": map[string]any{
			"isAdmin": true,
		},
		"currentUser": map[string]any{
			"email": email,
		},
	}

	errors := CheckContractCompatibility(contract, values)

	assert.Equal(t, 0, len(errors))
}
