package main

import (
	"testing"

	"github.com/peterldowns/testy/assert"
	"github.com/peterldowns/testy/check"
)

func TestLukasVars_Parse(t *testing.T) {
	t.Run("no env or flags", func(t *testing.T) {
		vars := &LUKASVars{}
		err := vars.Parse([]string{})
		check.Nil(t, err)
		assert.False(t, vars.Color)
		assert.False(t, vars.Verbose)
		assert.Equal(t, "", vars.File)
		assert.Equal(t, 0, vars.Workers)
	})

	t.Run("env only", func(t *testing.T) {
		t.Setenv("LUKAS_COLOR", "true")
		t.Setenv("LUKAS_VERBOSE", "true")
		t.Setenv("LUKAS_FILE", "/tmp/foo")
		t.Setenv("LUKAS_WORKERS", "10")

		vars := &LUKASVars{}
		err := vars.Parse([]string{})
		check.Nil(t, err)

		assert.True(t, vars.Color)
		assert.True(t, vars.Verbose)
		assert.Equal(t, "/tmp/foo", vars.File)
		assert.Equal(t, 10, vars.Workers)
	})

	t.Run("env and flags", func(t *testing.T) {
		t.Setenv("LUKAS_COLOR", "false")
		t.Setenv("LUKAS_VERBOSE", "false")
		t.Setenv("LUKAS_FILE", "/tmp/foo")
		t.Setenv("LUKAS_WORKERS", "10")

		args := []string{
			"--color=true",
			"--verbose=true",
			"--file=/tmp/bar",
			"--workers=20",
		}
		vars := &LUKASVars{}
		err := vars.Parse(args)
		check.Nil(t, err)

		assert.True(t, vars.Color)
		assert.True(t, vars.Verbose)
		assert.Equal(t, "/tmp/bar", vars.File)
		assert.Equal(t, 20, vars.Workers)
	})
}
