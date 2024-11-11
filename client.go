package keycloak_go_client

import (
	"context"
	"fmt"
	"github.com/zmotso/keycloak-go-client/internal/generated"
	"net/http"
)

type Client struct {
	client generated.ClientWithResponsesInterface

	Users *usersClient
}

type ClientOpts struct {
	url   string
	token string
}

type ClientOptsSetter func(*ClientOpts)

func WithToken(token string) func(*ClientOpts) {
	return func(opts *ClientOpts) {
		opts.token = token
	}
}

func NewClient(url string, opts ...ClientOptsSetter) (*Client, error) {
	defaults := &ClientOpts{
		url: url,
	}

	for _, o := range opts {
		o(defaults)
	}

	tokenProvider := NewBearerTokenAuthProvider(defaults.token)

	c, err := generated.NewClientWithResponses(
		defaults.url,
		generated.WithRequestEditorFn(tokenProvider.Intercept),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create keycloak client: %w", err)
	}

	return &Client{
		client: c,
		Users:  &usersClient{client: c},
	}, nil
}

type BearerTokenAuthProvider struct {
	token string
}

func NewBearerTokenAuthProvider(token string) *BearerTokenAuthProvider {
	return &BearerTokenAuthProvider{token: token}
}

func (s *BearerTokenAuthProvider) Intercept(_ context.Context, req *http.Request) error {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))

	return nil
}

type Response struct {
	Body         []byte
	HTTPResponse *http.Response
}
