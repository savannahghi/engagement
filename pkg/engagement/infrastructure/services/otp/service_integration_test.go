package otp_test

import (
	"context"
	"testing"
	"time"

	"github.com/savannahghi/authutils"
	"github.com/savannahghi/engagementcore/pkg/engagement/application/common/dto"
	db "github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/database/firestore"
	serviceAuthServer "github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/services/authserver"
	"github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/services/mail"
	"github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/services/messaging"
	"github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/services/otp"
	"github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/services/sms"
	"github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/services/twilio"
	"github.com/savannahghi/firebasetools"
	"github.com/savannahghi/interserviceclient"
	"github.com/savannahghi/serverutils"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
)

func newTestOtpSService(t *testing.T) (*otp.ServiceOTPImpl, error) {
	ctx := context.Background()
	fr, err := db.NewFirebaseRepository(ctx)
	if err != nil {
		t.Errorf("can't instantiate firebase repository in resolver: %v", err)
		return nil, err
	}

	ps, err := messaging.NewPubSubNotificationService(
		ctx,
		serverutils.MustGetEnvVar(serverutils.GoogleCloudProjectIDEnvVarName),
	)
	mail := mail.NewService(fr)
	silCommsConfig := authutils.Config{
		AuthServerEndpoint: serverutils.MustGetEnvVar("SIL_COMMS_AUTHSERVER_DOMAIN"),
		ClientID:           serverutils.MustGetEnvVar("SIL_COMMS_AUTHSERVER_CLIENT_ID"),
		ClientSecret:       serverutils.MustGetEnvVar("SIL_COMMS_AUTHSERVER_CLIENT_SECRET"),
		GrantType:          serverutils.MustGetEnvVar("SIL_COMMS_AUTHSERVER_GRANT_TYPE"),
		Username:           serverutils.MustGetEnvVar("SIL_COMMS_AUTHSERVER_USERNAME"),
		Password:           serverutils.MustGetEnvVar("SIL_COMMS_AUTHSERVER_PASSWORD"),
	}

	silCommsAuthService, err := serviceAuthServer.NewServiceAuthServer(silCommsConfig)
	if err != nil {
		return nil, err
	}

	sms := sms.NewService(fr, ps, silCommsAuthService)
	twilio := twilio.NewService(sms, fr)

	if err != nil {
		t.Errorf(
			"can't instantiate notification service in resolver: %v",
			err,
		)
		return nil, err
	}
	return otp.NewService(mail, sms, twilio), nil
}

func TestNewRemoteOtpService(t *testing.T) {
	s, err := newTestOtpSService(t)
	if err != nil {
		t.Errorf("unable to initialize test service with error %v", err)
		return
	}
	if s == nil {
		t.Errorf("got back nil remote Otp service")
		return
	}
}

func TestServiceOTPImpl_GenerateAndSendOTP(t *testing.T) {
	ctx := context.Background()
	s, err := newTestOtpSService(t)
	if err != nil {
		t.Errorf("unable to initialize test service with error %v", err)
		return
	}

	msisdn := interserviceclient.TestUserPhoneNumber
	appID := ksuid.New().String()

	type args struct {
		ctx    context.Context
		msisdn string
		appID  *string
	}
	tests := []struct {
		name      string
		args      args
		wantValue bool
		wantErr   bool
	}{
		{
			name: "valid: correct params passed",
			args: args{
				ctx:    ctx,
				msisdn: msisdn,
				appID:  &appID,
			},
			wantErr:   false,
			wantValue: true,
		},
		{
			name: "invalid: missing msisdn",
			args: args{
				ctx:   ctx,
				appID: &appID,
			},
			wantErr:   true,
			wantValue: false,
		},
		{
			name: "invalid: missing app id",
			args: args{
				ctx:    ctx,
				msisdn: msisdn,
			},
			wantErr:   false,
			wantValue: false,
		},
		{
			name: "invalid: missing args",
			args: args{
				ctx: ctx,
			},
			wantErr:   true,
			wantValue: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GenerateAndSendOTP(tt.args.ctx, tt.args.msisdn, tt.args.appID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceOTPImpl.GenerateAndSendOTP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantValue && got == "" {
				t.Errorf("ServiceOTPImpl.GenerateAndSendOTP() = %v,", got)
			}
		})
	}
}

func TestServiceOTPImpl_SendOTPToEmail(t *testing.T) {
	ctx := firebasetools.GetAuthenticatedContext(t)

	s, err := newTestOtpSService(t)
	if err != nil {
		t.Errorf("unable to initialize test service with error %v", err)
		return
	}

	msisdn := interserviceclient.TestUserPhoneNumber
	appID := ksuid.New().String()
	email := "test@bewell.co.ke"

	type args struct {
		ctx    context.Context
		msisdn *string
		email  *string
		appID  *string
	}
	tests := []struct {
		name      string
		args      args
		wantValue bool
		wantErr   bool
		panics    bool
	}{
		{
			name: "valid: correct params passed",
			args: args{
				ctx:    ctx,
				msisdn: &msisdn,
				email:  &email,
				appID:  &appID,
			},
			wantErr:   false,
			wantValue: true,
		},
		{
			name: "invalid: missing msisdn",
			args: args{
				ctx:   ctx,
				appID: &appID,
				email: &email,
			},
			panics: true,
		},
		{
			name: "invalid: missing email",
			args: args{
				ctx:    ctx,
				appID:  &appID,
				msisdn: &msisdn,
			},
			wantErr:   false,
			wantValue: false,
		},
		{
			name: "invalid: missing app id",
			args: args{
				ctx:    ctx,
				msisdn: &msisdn,
				email:  &email,
			},
			wantErr:   false,
			wantValue: false,
		},
		{
			name: "invalid: missing args",
			args: args{
				ctx: ctx,
			},
			panics: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.panics {
				got, err := s.SendOTPToEmail(tt.args.ctx, tt.args.msisdn, tt.args.email, tt.args.appID)
				if (err != nil) != tt.wantErr {
					t.Errorf("ServiceOTPImpl.SendOTPToEmail() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if tt.wantValue && got == "" {
					t.Errorf("ServiceOTPImpl.GenerateAndSendOTP() = %v,", got)
				}
			}
			if tt.panics {
				fcSendOTPToEmail := func() { _, _ = s.SendOTPToEmail(tt.args.ctx, tt.args.msisdn, tt.args.email, tt.args.appID) }
				assert.Panics(t, fcSendOTPToEmail)
			}
		})
	}
}

func TestServiceOTPImpl_SaveOTPToFirestore(t *testing.T) {
	s, err := newTestOtpSService(t)
	if err != nil {
		t.Errorf("unable to initialize test service with error %v", err)
		return
	}

	otp := dto.OTP{
		MSISDN:            interserviceclient.TestUserPhoneNumber,
		Message:           "test message",
		AuthorizationCode: "test code",
		Timestamp:         time.Now(),
		IsValid:           true,
		Email:             "test@bewell.co.ke",
	}
	type args struct {
		otp dto.OTP
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "default case",
			args: args{
				otp: otp,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.SaveOTPToFirestore(tt.args.otp); (err != nil) != tt.wantErr {
				t.Errorf("ServiceOTPImpl.SaveOTPToFirestore() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestImplOTP_VerifyOtp(t *testing.T) {
	ctx := firebasetools.GetAuthenticatedContext(t)

	s, err := newTestOtpSService(t)
	if err != nil {
		t.Errorf("unable to initialize test service with error %v", err)
		return
	}

	msisdn := interserviceclient.TestUserPhoneNumber
	appID := ksuid.New().String()
	retryStep := 1

	verificationCode, err := s.GenerateRetryOTP(ctx, &msisdn, retryStep, &appID)
	if err != nil {
		t.Errorf("failed to generate retry OTP to phone: %v", err)
	}

	type args struct {
		ctx              context.Context
		msisdn           *string
		verificationCode *string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
		panics  bool
	}{
		{
			name: "valid: correct params passed",
			args: args{
				ctx:              ctx,
				msisdn:           &msisdn,
				verificationCode: &verificationCode,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "invalid: missing  phone number",
			args: args{
				ctx:              ctx,
				verificationCode: &verificationCode,
			},
			panics: true,
		},
		{
			name: "invalid: missing verification code",
			args: args{
				ctx:    ctx,
				msisdn: &msisdn,
			},
			panics: true,
		},
	}
	for _, tt := range tests {
		if !tt.panics {
			t.Run(tt.name, func(t *testing.T) {
				got, err := s.VerifyOtp(tt.args.ctx, tt.args.msisdn, tt.args.verificationCode)
				if (err != nil) != tt.wantErr {
					t.Errorf("ImplOTP.VerifyOtp() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("ImplOTP.VerifyOtp() = %v, want %v", got, tt.want)
				}
			})
		}
		if tt.panics {
			fcVerifyOtp := func() { _, _ = s.VerifyOtp(tt.args.ctx, tt.args.msisdn, tt.args.verificationCode) }
			assert.Panics(t, fcVerifyOtp)
		}
	}
}
func TestServiceOTPImpl_VerifyEmailOtp(t *testing.T) {
	s, err := newTestOtpSService(t)
	if err != nil {
		t.Errorf("unable to initialize test service with error %v", err)
		return
	}
	ctx := firebasetools.GetAuthenticatedContext(t)
	email := "test@bewell.co.ke"

	msisdn := interserviceclient.TestUserPhoneNumber
	appID := ksuid.New().String()

	verificationCode, err := s.SendOTPToEmail(ctx, &msisdn, &email, &appID)
	if err != nil {
		t.Errorf("failed to send OTP to email: %v", err)
	}

	otp := dto.OTP{
		MSISDN:            interserviceclient.TestUserPhoneNumber,
		Message:           "test message",
		AuthorizationCode: verificationCode,
		Timestamp:         time.Now(),
		IsValid:           true,
		Email:             "test@bewell.co.ke",
	}

	err = s.SaveOTPToFirestore(otp)
	if err != nil {
		t.Errorf("failed to save otp to firestore: %v", err)
	}

	invalidVerificationCode := "invalid"

	type args struct {
		ctx              context.Context
		email            *string
		verificationCode *string
	}
	tests := []struct {
		name string

		args    args
		want    bool
		wantErr bool
		panics  bool
	}{
		{
			name: "valid: correct params passed",
			args: args{
				ctx:              ctx,
				email:            &email,
				verificationCode: &verificationCode,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "invalid: invalid verification code",
			args: args{
				ctx:              ctx,
				email:            &email,
				verificationCode: &invalidVerificationCode,
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "invalid: missing  email",
			args: args{
				ctx:              ctx,
				verificationCode: &verificationCode,
			},
			panics: true,
		},
		{
			name: "invalid: missing verification code",
			args: args{
				ctx:   ctx,
				email: &email,
			},
			panics: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.panics {
				got, err := s.VerifyEmailOtp(tt.args.ctx, tt.args.email, tt.args.verificationCode)
				if (err != nil) != tt.wantErr {
					t.Errorf("ServiceOTPImpl.VerifyEmailOtp() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("ServiceOTPImpl.VerifyEmailOtp() = %v, want %v", got, tt.want)
				}
			}
			if tt.panics {
				fcVerifyEmailOtp := func() { _, _ = s.VerifyEmailOtp(tt.args.ctx, tt.args.email, tt.args.verificationCode) }
				assert.Panics(t, fcVerifyEmailOtp)
			}
		})
	}
}

func TestServiceOTPImpl_GenerateRetryOTP(t *testing.T) {
	ctx := firebasetools.GetAuthenticatedContext(t)

	s, err := newTestOtpSService(t)
	if err != nil {
		t.Errorf("unable to initialize test service with error %v", err)
		return
	}

	msisdn := interserviceclient.TestUserPhoneNumber
	appID := ksuid.New().String()

	type args struct {
		ctx       context.Context
		msisdn    *string
		retryStep int
		appID     *string
	}
	tests := []struct {
		name      string
		args      args
		wantValue bool
		wantErr   bool
		panics    bool
	}{
		{
			name: "valid: correct params passed",
			args: args{
				ctx:       ctx,
				msisdn:    &msisdn,
				retryStep: 1,
				appID:     &appID,
			},
			wantErr:   false,
			wantValue: true,
		},
		{
			name: "invalid: missing msisdn",
			args: args{
				ctx:       ctx,
				retryStep: 1,
				appID:     &appID,
			},
			panics: true,
		},
		{
			name: "invalid: missing retryStep",
			args: args{
				ctx:    ctx,
				msisdn: &msisdn,
				appID:  &appID,
			},
			wantErr:   true,
			wantValue: false,
		},
		{
			name: "invalid: missing appid",
			args: args{
				ctx:       ctx,
				msisdn:    &msisdn,
				retryStep: 1,
			},
			wantErr:   false,
			wantValue: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.panics {
				got, err := s.GenerateRetryOTP(tt.args.ctx, tt.args.msisdn, tt.args.retryStep, tt.args.appID)
				if (err != nil) != tt.wantErr {
					t.Errorf("ServiceOTPImpl.GenerateRetryOTP() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if tt.wantValue && got == "" {
					t.Errorf("ServiceOTPImpl.GenerateRetryOTP() expected value got: %v", got)
				}
			}
			if tt.panics {
				fcGenerateRetryOTP := func() { _, _ = s.GenerateRetryOTP(tt.args.ctx, tt.args.msisdn, tt.args.retryStep, tt.args.appID) }
				assert.Panics(t, fcGenerateRetryOTP)
			}
		})
	}
}

func TestServiceOTPImpl_EmailVerificationOtp(t *testing.T) {
	ctx := firebasetools.GetAuthenticatedContext(t)

	s, err := newTestOtpSService(t)
	if err != nil {
		t.Errorf("unable to initialize test service with error %v", err)
		return
	}

	email := "test@bewell.co.ke"

	type args struct {
		ctx   context.Context
		email *string
	}
	tests := []struct {
		name      string
		args      args
		wantValue bool
		wantErr   bool
		panics    bool
	}{
		{
			name: "valid: correct params passed",
			args: args{
				ctx:   ctx,
				email: &email,
			},
			wantValue: true,
			wantErr:   false,
		},
		{
			name: "invalid: missing email",
			args: args{
				ctx: ctx,
			},
			panics: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.panics {
				got, err := s.EmailVerificationOtp(tt.args.ctx, tt.args.email)
				if (err != nil) != tt.wantErr {
					t.Errorf("ServiceOTPImpl.EmailVerificationOtp() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if tt.wantValue && got == "" {
					t.Errorf("ServiceOTPImpl.EmailVerificationOtp(), expected value got none:%v", got)
				}
			}
			if tt.panics {
				fcEmailVerificationOtp := func() { _, _ = s.EmailVerificationOtp(tt.args.ctx, tt.args.email) }
				assert.Panics(t, fcEmailVerificationOtp)
			}
		})
	}
}

func TestServiceOTPImpl_GenerateOTP(t *testing.T) {
	ctx := firebasetools.GetAuthenticatedContext(t)

	s, err := newTestOtpSService(t)
	if err != nil {
		t.Errorf("unable to initialize test service with error %v", err)
		return
	}

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name      string
		args      args
		wantValue bool
		wantErr   bool
	}{
		{
			name: "default case",
			args: args{
				ctx: ctx,
			},
			wantValue: true,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GenerateOTP(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceOTPImpl.GenerateOTP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantValue && got == "" {
				t.Errorf("ServiceOTPImpl.EmailVerificationOtp(), expected value got none:%v", got)
			}
		})
	}
}

func TestServiceOTPImpl_SendTemporaryPIN(t *testing.T) {

	ctx := firebasetools.GetAuthenticatedContext(t)

	s, err := newTestOtpSService(t)
	if err != nil {
		t.Errorf("unable to initialize test service with error %v", err)
		return
	}
	input := dto.TemporaryPIN{
		PhoneNumber: interserviceclient.TestUserPhoneNumber,
		FirstName:   "test",
		PIN:         "test",
		Channel:     1,
	}

	type args struct {
		ctx   context.Context
		input dto.TemporaryPIN
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid: correct params passed",
			args: args{
				ctx:   ctx,
				input: input,
			},
			wantErr: false,
		},
		{
			name: "invalid: missing input",
			args: args{
				ctx: ctx,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.SendTemporaryPIN(tt.args.ctx, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("ServiceOTPImpl.SendTemporaryPIN() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
