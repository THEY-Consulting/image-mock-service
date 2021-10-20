package application

import (
	"github.com/they-consulting/gogogo/config"
	"testing"
)

func TestConfigFromApplication(t *testing.T) {
	config.New()
	database := config.Database{
		Port:         "5432",
		User:         "postgres",
		DatabaseName: "postgres",
		Password:     "password",
		Host:         "postgres",
	}
	want := config.Config{Database: database, Port: "3000", DefaultRadiusInMatchesSearch: "10000", GeocoderApiKey: "GStJaAjouMu021G0qLjuGe4oOwOFfjVB"}
	got := Get()

	if *got.Config != want {
		t.Errorf("Application Config is %v; want %v", got.Config, want)
	}
}
