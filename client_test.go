package keycloak_go_client

import "context"

func ExampleNewClient() {
	cl, _ := NewClient("https://keycloak-eamplec.com", WithToken("token"))
	_, _, _ = cl.Users.GetUsersProfile(context.Background(), "master")
}
