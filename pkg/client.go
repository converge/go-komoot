package pkg

import "golang.org/x/oauth2"

type OAuthConfig struct {
	clientId    string
	redirectURL string
	authURL     string
}

func NewClient(config OAuthConfig) *oauth2.Config {
	return &oauth2.Config{
		ClientID: config.clientId,
		Endpoint: oauth2.Endpoint{
			AuthURL: config.authURL,
		},
		RedirectURL: config.redirectURL,
	}
}
