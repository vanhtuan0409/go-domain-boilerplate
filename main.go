package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	nsq "github.com/nsqio/go-nsq"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
	"github.com/vanhtuan0409/go-domain-boilerplate/application/accesscontrol"
	"github.com/vanhtuan0409/go-domain-boilerplate/application/goal"
	"github.com/vanhtuan0409/go-domain-boilerplate/application/member"
	domaingoal "github.com/vanhtuan0409/go-domain-boilerplate/domain/goal"
	"github.com/vanhtuan0409/go-domain-boilerplate/infrastructure/eventbus"
	"github.com/vanhtuan0409/go-domain-boilerplate/infrastructure/logger"
	"github.com/vanhtuan0409/go-domain-boilerplate/infrastructure/tokenprovider"
	eventhandler "github.com/vanhtuan0409/go-domain-boilerplate/interface/eventbus"
	"github.com/vanhtuan0409/go-domain-boilerplate/interface/http/handler"
	"github.com/vanhtuan0409/go-domain-boilerplate/interface/http/middleware"
	"github.com/vanhtuan0409/go-domain-boilerplate/interface/repository"
)

var (
	tokenContextKey = "authInfo"
	nsqServer       = "127.0.0.1:4150"
)

func main() {
	// Set logger
	logwriter := logger.NewLogrusStdLogger()
	logger.SetLogger(logwriter)

	// Init nsq dispatcher
	nsqConfig := nsq.NewConfig()
	w, _ := nsq.NewProducer(nsqServer, nsqConfig)
	defer w.Stop()

	// Init repo
	goalRepo := repository.NewInMemGoalRepo()
	memberRepo := repository.NewInMemMemberRepo()

	// Init application service
	accessControlService := accesscontrol.NewAccessControl()
	dispatcher := eventbus.NewEventDispatcher(w)

	// Init event handler
	InitEventHandler()

	// Init controller
	goalUsecase := goal.NewGoalUsecase(
		goalRepo, memberRepo,
		accessControlService, dispatcher,
	)
	memberUsecase := member.NewMemberUsecase(memberRepo)
	controller := handler.NewController(goalUsecase, memberUsecase)

	// Init http router
	routes := InitRoute(controller)
	server := &http.Server{
		Addr:         ":8080",
		Handler:      routes,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server deployed at :8080")
	log.Fatal(server.ListenAndServe())
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

func InitEventHandler() {
	config := nsq.NewConfig()

	addTaskHandler, _ := nsq.NewConsumer(domaingoal.EventAddTaskToGoalType, "ch1", config)
	addTaskHandler.AddHandler(eventbus.MakeEventHandlerFunc(eventhandler.HandleAddTaskToGoal))
	addTaskHandler.ConnectToNSQD(nsqServer)

	checkinHandler, _ := nsq.NewConsumer(domaingoal.EventCheckInTaskType, "ch1", config)
	checkinHandler.AddHandler(eventbus.MakeEventHandlerFunc(eventhandler.HandleCheckInTask))
	checkinHandler.ConnectToNSQD(nsqServer)
}
