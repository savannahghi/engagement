package authserver

import (
	"context"
	"net/http"
	"time"

	"github.com/savannahghi/authutils"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("gitlab.slade360emr.com/go/consumer-backend/pkg/mycarehub/infrastructure/services/authserver")

// ServiceAuthServer holds the method defined in authutils library
type ServiceAuthServer interface {
	LoginUser(ctx context.Context, input *authutils.LoginUserPayload) (*authutils.OAUTHResponse, error)
	RefreshToken(ctx context.Context, refreshToken string) (*authutils.OAUTHResponse, error)
	ValidateUser(ctx context.Context, authTokens *authutils.OAUTHResponse) (*authutils.MeResponse, error)
	HasValidSlade360BearerToken(ctx context.Context, r *http.Request) (bool, map[string]string, *authutils.TokenIntrospectionResponse)
}

// client is the library's client used to make requests
type client struct {
	authClient ServiceAuthServer
	httpClient *http.Client
}

// NewServiceAuthServer is the constructor which initializes health crm's authentication mechanism
func NewServiceAuthServer(config authutils.Config) (*client, error) { //nolint:revive
	slade360AuthClient, err := authutils.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &client{
		authClient: slade360AuthClient,
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}, nil
}

// Login logs in a user using the authserver
func (s *client) LoginUser(ctx context.Context, input *authutils.LoginUserPayload) (*authutils.OAUTHResponse, error) {
	ctx, span := tracer.Start(ctx, "LoginUser")
	defer span.End()

	return s.authClient.LoginUser(ctx, input)
}

// Validates whether a user exists on authserver
func (s *client) ValidateUser(ctx context.Context, oauthCredentials *authutils.OAUTHResponse) (*authutils.MeResponse, error) {
	ctx, span := tracer.Start(ctx, "ValidateUser")
	defer span.End()

	return s.authClient.ValidateUser(ctx, oauthCredentials)
}

// RefreshToken sends a request to the authServer to refresh a token
// An access token is returned
func (s *client) RefreshToken(ctx context.Context, refreshToken string) (*authutils.OAUTHResponse, error) {
	ctx, span := tracer.Start(ctx, "RefreshToken")
	defer span.End()

	return s.authClient.RefreshToken(ctx, refreshToken)
}

// HasValidSlade360BearerToken validates whether the authToken in
// the request is valid
func (s *client) HasValidSlade360BearerToken(ctx context.Context, r *http.Request) (bool, map[string]string, *authutils.TokenIntrospectionResponse) {
	ctx, span := tracer.Start(ctx, "RefreshToken")
	defer span.End()

	return s.authClient.HasValidSlade360BearerToken(ctx, r)
}
