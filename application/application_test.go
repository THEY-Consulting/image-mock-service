package application

import (
	"github.com/they-consulting/image-mock-service/config"
	"testing"
)

func TestConfigFromApplication(t *testing.T) {
	config.New()
	want := config.Config{Port: "3000"}
	got := Get()

	if *got.Config != want {
		t.Errorf("Application Config is %v; want %v", got.Config, want)
	}
}
