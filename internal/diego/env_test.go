package diego

import (
	"testing"

	"github.com/peterldowns/testy/assert"
)

func TestLookupString(t *testing.T) {
	t.Run("is set", func(t *testing.T) {
		t.Setenv("FOO", "bar")
		var s string
		lookupString(&s, "FOO")
		assert.Equal(t, "bar", s)
	})

	t.Run("is not set", func(t *testing.T) {
		var s string = "default"
		lookupString(&s, "FOO")
		assert.Equal(t, "default", s)
	})

	t.Run("is set to empty string", func(t *testing.T) {
		t.Setenv("FOO", "")
		var s string = "default"
		lookupString(&s, "FOO")
		assert.Equal(t, "", s)
	})
}

func TestLookupInt(t *testing.T) {
	t.Run("is set", func(t *testing.T) {
		t.Setenv("FOO", "123")
		var i int
		err := lookupInt(&i, "FOO")
		assert.Nil(t, err)
		assert.Equal(t, 123, i)
	})

	t.Run("is not set", func(t *testing.T) {
		var i int = 42
		err := lookupInt(&i, "FOO")
		assert.Nil(t, err)
		assert.Equal(t, 42, i)
	})

	t.Run("is not an int", func(t *testing.T) {
		t.Setenv("FOO", "bar")
		var i int = 42
		err := lookupInt(&i, "FOO")
		assert.NotNil(t, err)
		assert.Equal(t, 42, i)
	})

	t.Run("is set to an empty string", func(t *testing.T) {
		t.Setenv("FOO", "")
		var i int = 42
		err := lookupInt(&i, "FOO")
		assert.NotNil(t, err)
		assert.Equal(t, 42, i)
	})
}

func TestLookupBool(t *testing.T) {
	t.Run("is set to true", func(t *testing.T) {
		t.Setenv("FOO", "true")
		var b bool
		err := lookupBool(&b, "FOO")
		assert.Nil(t, err)
		assert.True(t, b)
	})

	t.Run("is set to arbitrary string", func(t *testing.T) {
		t.Setenv("FOO", "literally anything else")
		var b bool
		err := lookupBool(&b, "FOO")
		assert.Nil(t, err)
		assert.True(t, b)
	})

	t.Run("is set to false", func(t *testing.T) {
		t.Setenv("FOO", "false")
		var b bool = true
		err := lookupBool(&b, "FOO")
		assert.Nil(t, err)
		assert.False(t, b)
	})

	t.Run("is set to an empty string", func(t *testing.T) {
		t.Setenv("FOO", "")
		var b bool = true
		err := lookupBool(&b, "FOO")
		assert.Nil(t, err)
		assert.False(t, b)
	})

	t.Run("is not set", func(t *testing.T) {
		var b bool = true
		err := lookupBool(&b, "FOO")
		assert.Nil(t, err)
		assert.True(t, b)
	})
}
