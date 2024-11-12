package keycloak_go_client

import (
	"context"
	"github.com/zmotso/keycloak-go-client/internal/generated"
)

type (
	UserProfileConfig               = generated.UPConfig
	UserProfileAttribute            = generated.UPAttribute
	UserProfileAttributePermissions = generated.UPAttributePermissions
	UserProfileAttributeRequired    = generated.UPAttributeRequired
	UserProfileAttributeSelector    = generated.UPAttributeSelector
	UserProfileGroup                = generated.UPGroup
)

type usersClient struct {
	client generated.ClientWithResponsesInterface
}

func (c *usersClient) GetUsersProfile(ctx context.Context, realm string) (*UserProfileConfig, *Response, error) {
	res, err := c.client.GetAdminRealmsRealmUsersProfileWithResponse(ctx, realm)
	if res != nil {
		return res.JSON200, &Response{HTTPResponse: res.HTTPResponse, Body: res.Body}, err
	}

	return nil, nil, err
}

func (c *usersClient) UpdateUsersProfile(ctx context.Context, realm string, userProfile UserProfileConfig) (*UserProfileConfig, *Response, error) {
	res, err := c.client.PutAdminRealmsRealmUsersProfileWithResponse(ctx, realm, userProfile)
	if res != nil {
		return res.JSON200, &Response{HTTPResponse: res.HTTPResponse, Body: res.Body}, err
	}

	return nil, nil, err
}
