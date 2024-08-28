package authserver_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/jarcoal/httpmock"
	"github.com/savannahghi/authutils"
	"github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/services/authserver"
	"github.com/savannahghi/serverutils"
)

var config = authutils.Config{
	AuthServerEndpoint: serverutils.MustGetEnvVar("AUTHSERVER_DOMAIN"),
	ClientID:           serverutils.MustGetEnvVar("AUTHSERVER_CLIENT_ID"),
	ClientSecret:       serverutils.MustGetEnvVar("AUTHSERVER_CLIENT_SECRET"),
	GrantType:          serverutils.MustGetEnvVar("AUTHSERVER_GRANT_TYPE"),
	Username:           serverutils.MustGetEnvVar("AUTHSERVER_USERNAME"),
	Password:           serverutils.MustGetEnvVar("AUTHSERVER_PASSWORD"),
}

func Test_client_LoginUser(t *testing.T) {
	type args struct {
		ctx   context.Context
		input *authutils.LoginUserPayload
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success: happy login",
			args: args{
				ctx: context.Background(),
				input: &authutils.LoginUserPayload{
					Email:    gofakeit.Email(),
					Password: gofakeit.BeerAlcohol(),
				},
			},
			wantErr: false,
		},
		{
			name: "fail: unsuccessful login",
			args: args{
				ctx: context.Background(),
				input: &authutils.LoginUserPayload{
					Email:    gofakeit.Email(),
					Password: gofakeit.BeerAlcohol(),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, _ := authserver.NewServiceAuthServer(config)

			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			domain := serverutils.MustGetEnvVar("AUTHSERVER_DOMAIN")
			authURL := fmt.Sprintf("%s/oauth2/token/", domain)

			if tt.name == "success: happy login" {
				httpmock.RegisterResponder(http.MethodPost, authURL, func(r *http.Request) (*http.Response, error) { //nolint:all
					resp := authutils.OAUTHResponse{}

					return httpmock.NewJsonResponse(http.StatusOK, resp)
				})
			}

			if tt.name == "fail: unsuccessful login" {
				httpmock.RegisterResponder(http.MethodPost, authURL, func(r *http.Request) (*http.Response, error) { //nolint:all
					resp := authutils.OAUTHResponse{}

					return httpmock.NewJsonResponse(http.StatusUnauthorized, resp)
				})
			}

			_, err := s.LoginUser(tt.args.ctx, tt.args.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("client.LoginUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
