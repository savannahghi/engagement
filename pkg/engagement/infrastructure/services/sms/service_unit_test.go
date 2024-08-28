package sms_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/savannahghi/engagementcore/pkg/engagement/application/common/dto"
	"github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/database"
	repositoryMock "github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/database/mock"
	authServerMock "github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/services/authserver/mock"
	"github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/services/messaging"
	pubSubMock "github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/services/messaging/mock"
	"github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/services/sms"
	smsMock "github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/services/sms/mock"
	twilioMock "github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/services/twilio/mock"
	"github.com/savannahghi/enumutils"
	"github.com/savannahghi/silcomms"
)

var fakeInfrastructure repositoryMock.FakeEngagementRepository
var databaseSvc database.Repository = &fakeInfrastructure
var fakePubsub pubSubMock.FakeServiceMessaging
var pubsub messaging.NotificationService = &fakePubsub
var fakeSMS smsMock.FakeServiceSMS
var fakeTwilio twilioMock.FakeServiceTwilio
var fakeAuthserver authServerMock.AuthServerServiceMock

func TestServiceSMSImpl_SendToMany(t *testing.T) {
	e := sms.NewService(databaseSvc, pubsub, fakeAuthserver)
	ctx := context.Background()

	sender := "UNKNOWN"

	type args struct {
		ctx     context.Context
		message string
		to      []string
		from    enumutils.SenderID
	}
	tests := []struct {
		name    string
		args    args
		want    *dto.SendMessageResponse
		wantErr bool
	}{
		{
			name: "valid:successfully send to many using BeWell",
			args: args{
				ctx:     ctx,
				message: "This is a test",
				to:      []string{"+254711223344", "+254700990099"},
				from:    enumutils.SenderIDBewell,
			},
			wantErr: false,
		},
		{
			name: "valid:successfully send to many using Slade360",
			args: args{
				ctx:     ctx,
				message: "This is a test",
				to:      []string{"+254711223344", "+254700990099"},
				from:    enumutils.SenderIDSLADE360,
			},
			wantErr: false,
		},
		{
			name: "invalid: send to many using unknown sender",
			args: args{
				ctx:     ctx,
				message: "This is a test",
				to:      []string{"+254711223344", "+254700990099"},
				from:    enumutils.SenderID(sender),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "valid:successfully send to many using BeWell" {
				fakeSMS.SendFn = func(
					ctx context.Context,
					to, message string,
				) (*silcomms.BulkSMSResponse, error) {
					return &silcomms.BulkSMSResponse{}, nil
				}
			}
			if tt.name == "valid:successfully send to many using Slade360" {
				fakeSMS.SendFn = func(
					ctx context.Context,
					to, message string,
				) (*silcomms.BulkSMSResponse, error) {
					return &silcomms.BulkSMSResponse{}, nil
				}
			}
			if tt.name == "invalid: send to many using unknown sender" {
				fakeSMS.SendFn = func(
					ctx context.Context,
					to, message string,
				) (*silcomms.BulkSMSResponse, error) {
					return &silcomms.BulkSMSResponse{}, nil
				}
			}
			got, err := e.SendToMany(tt.args.ctx, tt.args.to, tt.args.message)
			if tt.wantErr {
				if err == nil {
					t.Errorf("error expected got %v", err)
					return
				}
			}
			if !tt.wantErr {
				if err != nil {
					t.Errorf("error not expected got %v", err)
					return
				}

				if got == nil {
					t.Errorf("nil response returned")
					return
				}
			}
		})
	}
}

func TestServiceSMSImpl_Send(t *testing.T) {
	e := sms.NewService(databaseSvc, pubsub, fakeAuthserver)
	ctx := context.Background()
	type args struct {
		ctx     context.Context
		to      string
		message string
		sender  enumutils.SenderID
	}
	tests := []struct {
		name    string
		args    args
		want    *dto.SendMessageResponse
		wantErr bool
	}{
		{
			name: "valid:successfully send",
			args: args{
				ctx:     ctx,
				message: "This is a test",
				to:      "+254711223344",
				sender:  enumutils.SenderIDSLADE360,
			},
			wantErr: false,
		},
		{
			name: "invalid:fail to send",
			args: args{
				ctx:     ctx,
				message: "",
				to:      "+",
				sender:  enumutils.SenderIDSLADE360,
			},
			wantErr: true,
		},
		{
			name: "invalid: send from an unknown sender",
			args: args{
				ctx:     ctx,
				message: "This is a test",
				to:      "+254711223344",
				sender:  "what-shall-we-do",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := e.Send(tt.args.ctx, tt.args.to, tt.args.message)
			if tt.name == "valid:successfully send" {
				fakeTwilio.SendSMSFn = func(ctx context.Context, to string, msg string) error {
					return nil
				}
			}
			if tt.name == "invalid:fail to send" {
				fakeTwilio.SendSMSFn = func(ctx context.Context, to string, msg string) error {
					return fmt.Errorf("sms error")
				}
			}
			if tt.name == "invalid: send from an unknown sender" {
				fakeTwilio.SendSMSFn = func(ctx context.Context, to string, msg string) error {
					return fmt.Errorf("unknown sender")
				}
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceSMSImpl.SendSMS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err == nil {
					t.Errorf("error expected got %v", err)
					return
				}
			}
			if !tt.wantErr {
				if err != nil {
					t.Errorf("error not expected got %v", err)
					return
				}

				if got == nil {
					t.Errorf("nil response returned")
					return
				}
			}
		})
	}
}
