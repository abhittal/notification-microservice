package app

import (
	"code.jtg.tools/ayush.singhal/notifications-microservice/app/controllers/auth"
	"code.jtg.tools/ayush.singhal/notifications-microservice/app/controllers/authorization"
	"code.jtg.tools/ayush.singhal/notifications-microservice/app/controllers/channels"
	"code.jtg.tools/ayush.singhal/notifications-microservice/app/controllers/logs"
	"code.jtg.tools/ayush.singhal/notifications-microservice/app/controllers/notifications"
	"code.jtg.tools/ayush.singhal/notifications-microservice/app/controllers/recipients"
	"code.jtg.tools/ayush.singhal/notifications-microservice/app/controllers/users"
	"code.jtg.tools/ayush.singhal/notifications-microservice/app/middlewares"
	"code.jtg.tools/ayush.singhal/notifications-microservice/configuration"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// InitServer is used to initialize server routes
func InitServer() error {
	router := gin.Default()
	// setting the cors headers

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "ResponseType", "accept", "origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowMethods:     []string{"POST", "OPTIONS", "GET", "PUT", "DELETE"},
	}))

	v1 := router.Group("/api/v1")

	healthCheck := v1.Group("/health-check")

	// healthCheck contains the /health-check Health Check Endpoint
	healthCheck.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	firstSignUp := v1.Group("/signup", middlewares.CheckIfFirst())
	authorization.SignUpRoute(firstSignUp)

	authToken := v1.Group("/auth")
	auth.RefreshAccessTokenRoute(authToken)
	auth.ValidateEmailRoute(authToken)
	auth.CheckIfFirstRoute(authToken)
	users.CreatePasswordRoute(authToken)

	loginGroup := v1.Group("/login", middlewares.CheckIfLogged())
	authorization.SignInRoute(loginGroup)

	ownInfoGroup := v1.Group("/profile", middlewares.AuthorizeJWT())
	users.ChangeOwnPasswordRoute(ownInfoGroup)
	users.GetUserProfileRoute(ownInfoGroup)

	userGroup := v1.Group("/users", middlewares.AuthorizeJWT(), middlewares.CheckIfSystemAdmin())
	users.GetAllUsersRoute(userGroup)
	users.AddUserRoute(userGroup)
	users.DeleteUserRoute(userGroup)
	users.GetUserRoute(userGroup)
	users.UpdateUserRoute(userGroup)
	users.ResetPasswordRoute(userGroup)

	recipientGroup := v1.Group("/recipients", middlewares.AuthorizeJWT())
	recipientSystemAdminGroup := recipientGroup.Group("", middlewares.CheckIfSystemAdmin())
	recipients.AddUpdateRecipientRoute(recipientSystemAdminGroup)
	recipients.GetRecipientRoute(recipientGroup)
	recipients.GetAllRecipientRoute(recipientGroup)

	channelGroup := v1.Group("/channels", middlewares.AuthorizeJWT())
	channelSystemAdminGroup := channelGroup.Group("", middlewares.CheckIfSystemAdmin())
	channels.AddChannelRoute(channelSystemAdminGroup)
	channels.GetAllChannelsRoute(channelGroup)
	channels.UpdateChannelRoute(channelSystemAdminGroup)
	channels.DeleteChannelRoute(channelSystemAdminGroup)
	channels.GetChannelRoute(channelGroup)

	logGroup := v1.Group("/logs", middlewares.AuthorizeJWT(), middlewares.CheckIfSystemAdmin())
	logs.GetLogsRoute(logGroup)

	apiKeyGroup := v1.Group("/api-key", middlewares.AuthorizeJWT(), middlewares.CheckIfSystemAdmin())
	notifications.GetAPILastRoute(apiKeyGroup)
	notifications.GetAPIKeyRoute(apiKeyGroup)

	notificationGroup := v1.Group("/notifications", middlewares.AuthorizeJWT())
	notifications.GetAllNotificationsRoute(notificationGroup)

	sendNotificationGroup := v1.Group("/send-notification")
	notificationOpenAPI := sendNotificationGroup.Group("", middlewares.APIKeyAuth())
	notifications.PostSendNotificationsRoute(notificationOpenAPI)

	err := router.Run(":" + configuration.GetResp().Server.Port)
	return errors.Wrap(err, "Unable to run server")
}
