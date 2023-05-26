package ims

import (
	"context"

	"github.com/Shopify/sarama"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type TokenProvider struct {
	tokenSource oauth2.TokenSource
}

func NewTokenProvider(clientID string, clientSecret string, tokenURL string) sarama.AccessTokenProvider {
	cfg := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     tokenURL,
	}

	return &TokenProvider{
		tokenSource: cfg.TokenSource(context.Background()),
	}
}

func (t *TokenProvider) Token() (*sarama.AccessToken, error) {
	token, err := t.tokenSource.Token()
	if err != nil {
		return nil, err
	}

	return &sarama.AccessToken{Token: token.AccessToken}, nil
}
