package mock

import (
	"context"
	"net/http"

	"github.com/savannahghi/authutils"
)

// AuthServerServiceMock mocks onboarding implementations
type AuthServerServiceMock struct {
	MockLoginUserFn                   func(ctx context.Context, input *authutils.LoginUserPayload) (*authutils.OAUTHResponse, error)
	MockValidateUserFn                func(ctx context.Context, authTokens *authutils.OAUTHResponse) (*authutils.MeResponse, error)
	MockRefreshTokenFn                func(ctx context.Context, refreshToken string) (*authutils.OAUTHResponse, error)
	MockHasValidSlade360BearerTokenFn func(ctx context.Context, r *http.Request) (bool, map[string]string, *authutils.TokenIntrospectionResponse)
}

// NewAuthServerServiceMock initializes our client mocks
func NewAuthServerServiceMock() *AuthServerServiceMock {
	response := &authutils.OAUTHResponse{
		AccessToken:  "access",
		RefreshToken: "refresh",
	}

	return &AuthServerServiceMock{
		MockLoginUserFn: func(ctx context.Context, input *authutils.LoginUserPayload) (*authutils.OAUTHResponse, error) { //nolint: revive
			return response, nil
		},
		MockValidateUserFn: func(ctx context.Context, authTokens *authutils.OAUTHResponse) (*authutils.MeResponse, error) { //nolint:revive
			return &authutils.MeResponse{}, nil
		},
		MockRefreshTokenFn: func(ctx context.Context, refreshToken string) (*authutils.OAUTHResponse, error) { //nolint:all
			return response, nil
		},
		MockHasValidSlade360BearerTokenFn: func(_ context.Context, _ *http.Request) (bool, map[string]string, *authutils.TokenIntrospectionResponse) {
			return true, nil, &authutils.TokenIntrospectionResponse{}
		},
	}
}

// LoginUser mocks the implementation of proxying login requests for users not found on consumer backend to onboarding
func (oc AuthServerServiceMock) LoginUser(ctx context.Context, input *authutils.LoginUserPayload) (*authutils.OAUTHResponse, error) {
	return oc.MockLoginUserFn(ctx, input)
}

// MockValidateUserFn mocks the implementation of validating user exists on authserver
func (oc AuthServerServiceMock) ValidateUser(ctx context.Context, authTokens *authutils.OAUTHResponse) (*authutils.MeResponse, error) {
	return oc.MockValidateUserFn(ctx, authTokens)
}

// RefreshToken mocks the implementation of refreshing a token
func (oc AuthServerServiceMock) RefreshToken(ctx context.Context, refreshToken string) (*authutils.OAUTHResponse, error) {
	return oc.MockRefreshTokenFn(ctx, refreshToken)
}

// HasValidSlade360BearerToken mocks implementation of validating a bearer token
func (oc AuthServerServiceMock) HasValidSlade360BearerToken(ctx context.Context, r *http.Request) (bool, map[string]string, *authutils.TokenIntrospectionResponse) {
	return oc.MockHasValidSlade360BearerTokenFn(ctx, r)
}
