package keycloak_go_client

import (
	"context"
	"fmt"
	"github.com/zmotso/keycloak-go-client/generated"
	"net/http"
)

type Client struct {
	client generated.ClientWithResponsesInterface

	Users *usersClient
}

type ClientOpts struct {
	url    string
	client *http.Client
	token  string
}

type ClientOptsSetter func(*ClientOpts)

func WithToken(token string) func(*ClientOpts) {
	return func(opts *ClientOpts) {
		opts.token = token
	}
}

func WithHTTPClient(httpClient *http.Client) func(*ClientOpts) {
	return func(opts *ClientOpts) {
		opts.client = httpClient
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

	generatedClientOpts := []generated.ClientOption{
		generated.WithRequestEditorFn(tokenProvider.Intercept),
	}

	if defaults.client != nil {
		generatedClientOpts = append(generatedClientOpts, generated.WithHTTPClient(defaults.client))
	}

	c, err := generated.NewClientWithResponses(
		defaults.url,
		generatedClientOpts...,
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
