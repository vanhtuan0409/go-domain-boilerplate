package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
	"github.com/vanhtuan0409/go-domain-boilerplate/application/accesscontrol"
	"github.com/vanhtuan0409/go-domain-boilerplate/application/goal"
	"github.com/vanhtuan0409/go-domain-boilerplate/application/member"
	"github.com/vanhtuan0409/go-domain-boilerplate/infrastructure/eventbus"
	"github.com/vanhtuan0409/go-domain-boilerplate/infrastructure/logger"
	"github.com/vanhtuan0409/go-domain-boilerplate/infrastructure/tokenprovider"
	"github.com/vanhtuan0409/go-domain-boilerplate/interface/http/handler"
	"github.com/vanhtuan0409/go-domain-boilerplate/interface/http/middleware"
	"github.com/vanhtuan0409/go-domain-boilerplate/interface/repository"
)

var (
	tokenContextKey = "authInfo"
)

func main() {
	// Set logger
	logwriter := logger.NewLogrusStdLogger()
	logger.SetLogger(logwriter)

	// Init repo
	goalRepo := repository.NewInMemGoalRepo()
	memberRepo := repository.NewInMemMemberRepo()

	// Init application service
	accessControlService := accesscontrol.NewAccessControl()
	dispatcher := eventbus.NewEventDispatcher()

	// Init controller
	goalUsecase := goal.NewGoalUsecase(
		goalRepo, memberRepo,
		accessControlService, dispatcher,
	)
	memberUsecase := member.NewMemberUsecase(memberRepo)
	controller := handler.NewController(goalUsecase, memberUsecase)

	// Init router
	routes := InitRoute(controller)
	server := &http.Server{
		Addr:         ":8080",
		Handler:      routes,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server deployed at :8080")
	server.ListenAndServe()
}

func InitRoute(ctrl *handler.Controller) *mux.Router {
	tokenProvider := tokenprovider.NewTokenProvider()

	loggerMdw := middleware.NewLoggerMiddleware()
	tokenMdw := middleware.NewTokenMiddleware(tokenContextKey, tokenProvider)
	corsMdw := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{
			"Accept", "Content-Type", "Content-Length",
			"Accept-Encoding", "X-CSRF-Token", "Authorization",
		},
	})
	commonMdw := negroni.New(loggerMdw, corsMdw)
	protectMdw := negroni.New(loggerMdw, corsMdw, tokenMdw)

	r := mux.NewRouter()
	r.Path("/members").Methods("GET").Handler(
		middleware.AdaptHandleFunc(commonMdw, tokenContextKey, ctrl.ListAllMember),
	)
	r.Path("/members/{memberID}/goals").Methods("GET").Handler(
		middleware.AdaptHandleFunc(commonMdw, tokenContextKey, ctrl.ListMemberGoal),
	)
	r.Path("/goals/{goalID}/checkin").Methods("POST").Handler(
		middleware.AdaptHandleFunc(protectMdw, tokenContextKey, ctrl.CheckInTask),
	)

	return r
}
