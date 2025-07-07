package env_test

import (
	"testing"

	"github.com/lukasschwab/diego/pkg/env"
	"github.com/peterldowns/testy/assert"
	"github.com/peterldowns/testy/check"
)

func TestLookupString(t *testing.T) {
	t.Run("is set", func(t *testing.T) {
		t.Setenv("FOO", "bar")
		var s string
		env.LookupString(&s, "FOO")
		assert.Equal(t, "bar", s)
	})

	t.Run("is not set", func(t *testing.T) {
		t.Parallel()
		var s string = "default"
		env.LookupString(&s, "FOO")
		assert.Equal(t, "default", s)
	})
}

func TestLookupInt(t *testing.T) {
	t.Run("is set", func(t *testing.T) {
		t.Setenv("FOO", "123")
		var i int
		err := env.LookupInt(&i, "FOO")
		check.Nil(t, err)
		assert.Equal(t, 123, i)
	})

	t.Run("is not set", func(t *testing.T) {
		t.Parallel()
		var i int = 42
		err := env.LookupInt(&i, "FOO")
		check.Nil(t, err)
		assert.Equal(t, 42, i)
	})

	t.Run("is not an int", func(t *testing.T) {
		t.Setenv("FOO", "bar")
		var i int = 42
		err := env.LookupInt(&i, "FOO")
		check.NotNil(t, err)
		assert.Equal(t, 42, i)
	})
}

func TestLookupBool(t *testing.T) {
	t.Run("is set to a non-empty value", func(t *testing.T) {
		t.Setenv("FOO", "true")
		var b bool
		err := env.LookupBool(&b, "FOO")
		check.Nil(t, err)
		assert.True(t, b)
	})

	t.Run("is set to an empty value", func(t *testing.T) {
		t.Setenv("FOO", "")
		var b bool = true
		err := env.LookupBool(&b, "FOO")
		check.Nil(t, err)
		assert.False(t, b)
	})

	t.Run("is not set", func(t *testing.T) {
		t.Parallel()
		var b bool = true
		err := env.LookupBool(&b, "FOO")
		check.Nil(t, err)
		assert.True(t, b)
	})
}
