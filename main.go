package main

import (
	"github.com/NYTimes/gizmo/server"
	nsq "github.com/nsqio/go-nsq"
	"github.com/vanhtuan0409/go-domain-boilerplate/application/accesscontrol"
	"github.com/vanhtuan0409/go-domain-boilerplate/application/goal"
	"github.com/vanhtuan0409/go-domain-boilerplate/application/member"
	"github.com/vanhtuan0409/go-domain-boilerplate/infrastructure/eventbus"
	"github.com/vanhtuan0409/go-domain-boilerplate/infrastructure/logger"
	httpendpoints "github.com/vanhtuan0409/go-domain-boilerplate/interface/http"
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

	// Init usecase
	goalUsecase := goal.NewGoalUsecase(
		goalRepo, memberRepo,
		accessControlService, dispatcher,
	)
	memberUsecase := member.NewMemberUsecase(memberRepo)

	goalEndPoints := httpendpoints.NewGoalEndPoints(goalUsecase)
	memberEndPoints := httpendpoints.NewMemberEndPoints(memberUsecase)

	sconfig := &server.Config{
		HTTPPort: 8000,
	}
	server.Init("godomain", sconfig)
	server.Register(goalEndPoints)
	server.Register(memberEndPoints)
	server.Log.Fatal(server.Run())
}
