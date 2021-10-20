package config

import (
	"testing"
)

func TestEnvConfigViaApplication(t *testing.T) {
	New()
	want := ExpectedConfig()
	got := *Get()
	if got != want {
		t.Errorf("Config should be equal to: %s but there was such a values : %s", got, want)
	}
}

func ExpectedConfig() Config {
	want := Config{Port: "3000"}
	return want
}
