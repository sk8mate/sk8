package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_env_struct_should_match_specific_strings(t *testing.T) {
	assert.Equal(t, "development", Env.Development)
	assert.Equal(t, "test", Env.Test)
	assert.Equal(t, "production", Env.Production)
}
func Test_env_should_fallback_to_development_if_not_provided(t *testing.T) {
	t.Setenv("SK8_ENV", "")

	env := GetEnv()

	assert.Equal(t, Env.Development, env)
}

func Test_env_should_resolve_development_for_development_string(t *testing.T) {
	t.Setenv("SK8_ENV", "development")

	env := GetEnv()

	assert.Equal(t, Env.Development, env)
}

func Test_env_should_resolve_test_for_test_string(t *testing.T) {
	t.Setenv("SK8_ENV", "test")

	env := GetEnv()

	assert.Equal(t, Env.Test, env)
}

func Test_env_should_resolve_production_for_production_string(t *testing.T) {
	t.Setenv("SK8_ENV", "production")

	env := GetEnv()

	assert.Equal(t, Env.Production, env)
}

func Test_should_panic_for_undefined_env(t *testing.T) {
	assert.PanicsWithValue(t, "Invalid SK8_ENV value.", func() {
		t.Setenv("SK8_ENV", "shitenv")
		GetEnv()
	})
}
