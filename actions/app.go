package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	forcessl "github.com/gobuffalo/mw-forcessl"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
	"github.com/gobuffalo/x/sessions"
	"github.com/unrolled/secure"

	"livegift_back/actions/middleware/authorization"
	"livegift_back/models"

	"github.com/gobuffalo/buffalo-pop/v2/pop/popmw"
	contenttype "github.com/gobuffalo/mw-contenttype"
	"github.com/rs/cors"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var (
	baseURL = envy.Get("BASE_URL", "http://localhost:3000")
	ENV     = envy.Get("GO_ENV", "development")
	app     *buffalo.App
)

func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:          ENV,
			SessionStore: sessions.Null{},
			PreWares: []buffalo.PreWare{cors.New(cors.Options{
				AllowedOrigins: []string{"*"},
				AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
				AllowedHeaders: []string{
					"Content-Type", "application/json",
					"Authorization",
				},
				AllowCredentials: true,
				Debug:            true,
			}).Handler},
			SessionName: "_livegift_back_session",
		})

		app.Use(forceSSL())
		app.Use(paramlogger.ParameterLogger)
		app.Use(contenttype.Set("application/json"))
		app.Use(popmw.Transaction(models.DB))

		app.Use(authorization.Authorizator)

		SetRoutes(app)
	}

	return app
}

func forceSSL() buffalo.MiddlewareFunc {
	return forcessl.Middleware(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}
