package sms

import (
	"context"
	"fmt"

	"github.com/savannahghi/authutils"
	"github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/database"
	"github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/services/messaging"
	"github.com/savannahghi/silcomms"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("github.com/savannahghi/engagementcore/pkg/engagement/services/sms")

// ServiceSMS defines the interactions with sms service
type ServiceSMS interface {
	SendToMany(
		ctx context.Context,
		to []string,
		message string,
	) (*silcomms.BulkSMSResponse, error)
	Send(
		ctx context.Context,
		to, message string,
	) (*silcomms.BulkSMSResponse, error)
}

// AuthServerImpl defines the methods provided by
// the auth server library
type AuthServerImpl interface {
	LoginUser(ctx context.Context, input *authutils.LoginUserPayload) (*authutils.OAUTHResponse, error)
	RefreshToken(ctx context.Context, refreshToken string) (*authutils.OAUTHResponse, error)
}

// ServiceSMSImpl defines a sms service struct
type ServiceSMSImpl struct {
	SILComms silcomms.CommsLib
}

// NewService returns a new service
func NewService(
	repository database.Repository,
	pubsub messaging.NotificationService,
	authService AuthServerImpl,
) *ServiceSMSImpl {
	silCommsLib := silcomms.MustNewSILCommsLib(authService)
	return &ServiceSMSImpl{
		*silCommsLib,
	}
}

// SendToMany is a utility method to send to many recipients at the same time
func (s ServiceSMSImpl) SendToMany(
	ctx context.Context,
	to []string,
	message string,
) (*silcomms.BulkSMSResponse, error) {
	bulkSMSResponse, err := s.SendSMS(ctx, to, message)
	if err != nil {
		return nil, fmt.Errorf("")
	}
	return bulkSMSResponse, nil
}

// Send is a method used to send to a single recipient
func (s ServiceSMSImpl) Send(
	ctx context.Context,
	to, message string,
) (*silcomms.BulkSMSResponse, error) {
	recipients := []string{to}
	smsResponse, err := s.SendSMS(ctx, recipients, message)
	if err != nil {
		return nil, fmt.Errorf("")
	}
	return smsResponse, nil
}

// SendSMS is a method used to send SMS
func (s ServiceSMSImpl) SendSMS(
	ctx context.Context,
	to []string, message string,
) (*silcomms.BulkSMSResponse, error) {
	ctx, span := tracer.Start(ctx, "SendSMS")
	defer span.End()
	smsResponse, err := s.SILComms.SendBulkSMS(ctx, message, to)
	if err != nil {
		return nil, fmt.Errorf("failed to send sms")
	}

	return smsResponse, nil
}
