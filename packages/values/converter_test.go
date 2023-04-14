package values

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValuesToEnv(t *testing.T) {
	host := faker.URL()
	responseLimit := 100
	paginationLength := 100
	values := map[string]any{
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

	env := ValuesToEnv(values)

	assert.Equal(t, 3, len(env))
	assert.Contains(t, env, fmt.Sprintf("DATABASE_HOST=%s", host))
	assert.Contains(t, env, fmt.Sprintf("DATABASE_RESPONSE_LIMIT=%d", responseLimit))
	assert.Contains(t, env, fmt.Sprintf("VCS_PROVIDER_PAGINATION_LENGTH=%d", paginationLength))
}
