package infrastructure

import (
	"context"

	serviceAuthServer "github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/services/authserver"

	"github.com/labstack/gommon/log"
	"github.com/savannahghi/authutils"
	"github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/database"
	"github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/services/fcm"
	"github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/services/feedback"
	"github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/services/library"
	"github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/services/mail"
	"github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/services/messaging"
	"github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/services/onboarding"
	"github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/services/otp"
	"github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/services/sms"
	"github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/services/surveys"
	"github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/services/twilio"
	"github.com/savannahghi/engagementcore/pkg/engagement/infrastructure/services/uploads"
	"github.com/savannahghi/serverutils"
)

// Interactor is an implementation of the infrastructure interface
// It combines each individual service implementation
type Interactor struct {
	database.Repository
	*fcm.ServiceFCMImpl
	*library.ServiceLibraryImpl
	*mail.ServiceMailImpl
	messaging.NotificationService
	onboarding.ProfileService
	*otp.ServiceOTPImpl
	*sms.ServiceSMSImpl
	*surveys.ServiceSurveyImpl
	*feedback.ServiceFeedbackImpl
	*twilio.ServiceTwilioImpl
	*uploads.ServiceUploadImpl
}

// NewInteractor initializes a new infrastructure interactor
func NewInteractor() Interactor {
	ctx := context.Background()

	projectID, err := serverutils.GetEnvVar(serverutils.GoogleCloudProjectIDEnvVarName)
	if err != nil {
		log.Fatal(err)
	}

	db := database.NewDbService()

	onboarding := onboarding.NewRemoteProfileService(onboarding.NewOnboardingClient())

	fcmOne := fcm.NewService(db, onboarding)

	lib := library.NewLibraryService(onboarding)

	mail := mail.NewService(db)

	pubsub, err := messaging.NewPubSubNotificationService(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}

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
		log.Fatal(err)
	}

	sms := sms.NewService(db, pubsub, silCommsAuthService)
	twilio := twilio.NewService(*sms, db)

	uploads := uploads.NewUploadsService()

	otp := otp.NewService(*mail, *sms, twilio)

	surveys := surveys.NewService(db)

	feedback := feedback.NewService(db)

	return Interactor{
		db,
		fcmOne,
		lib,
		mail,
		pubsub,
		onboarding,
		otp,
		sms,
		surveys,
		feedback,
		twilio,
		uploads,
	}
}
