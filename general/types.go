package general

import (
	"football_api/ent"
	"football_api/profile"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartyemailpassword"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartyemailpassword/tpepmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
	"os"
)

type Server struct {
	Router *chi.Mux
	Client *ent.Client
}

func (s *Server) SetSuperTokens() {
	apiBasePath := "/auth"
	websiteBasePath := "/auth"

	err := supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			// https://try.supertokens.com is for demo purposes. Replace this with the address of your core instance (sign up on supertokens.com), or self host a core.
			ConnectionURI: "https://try.supertokens.com",
			// APIKey: <API_KEY(if configured)>,
		},
		AppInfo: supertokens.AppInfo{
			AppName:         os.Getenv("APP_NAME"),
			APIDomain:       os.Getenv("API_DOMAIN"),
			WebsiteDomain:   os.Getenv("WEB_DOMAIN"),
			APIBasePath:     &apiBasePath,
			WebsiteBasePath: &websiteBasePath,
		},
		RecipeList: []supertokens.Recipe{
			thirdpartyemailpassword.Init(&tpepmodels.TypeInput{ /*TODO: See next step*/ }),
			session.Init(nil), // initializes session features
		},
	})

	s.Router.Use(supertokens.Middleware)

	if err != nil {
		panic(err.Error())
	}
}

func (s *Server) SetupMiddlewares() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   append([]string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}, supertokens.GetAllCORSHeaders()...),
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
}

func (s *Server) SetupHandlers() {
	// Profile
	pController := profile.Controller{
		Client: s.Client,
	}
	pGroup := chi.NewRouter()
	// Wrap the API handler in session.VerifySession
	//sessionRequired := false
	//r.Post("/likecomment", session.VerifySession(&sessmodels.VerifySessionOptions{
	//	SessionRequired: &sessionRequired,
	//}, likeCommentAPI))
	pGroup.Post("/", session.VerifySession(nil, pController.Create))
	pGroup.Put("/", session.VerifySession(nil, pController.Update))
	pGroup.Get("/", session.VerifySession(nil, pController.CurrentProfile))
	s.Router.Mount("/profile", pGroup)
}
