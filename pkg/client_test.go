package pkg

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	config := OAuthConfig{
		clientId:    "test",
		redirectURL: "http://localhost:9000",
		authURL:     "example.com/authorize",
	}
	client := NewClient(config)

	if client.ClientID != config.clientId {
		t.Errorf("wanted: %s, got: %s", client.ClientID, config.clientId)
	}

	if client.RedirectURL != config.redirectURL {
		t.Errorf("wanted: %s, got: %s", client.ClientSecret, config.redirectURL)
	}

	if client.Endpoint.AuthURL != config.authURL {
		t.Errorf("wanted: %s, got: %s", client.Endpoint.AuthURL, config.authURL)
	}
}
