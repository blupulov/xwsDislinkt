package startup

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/blupulov/xwsDislinkt/user-service/service"

	"github.com/blupulov/xwsDislinkt/user-service/controller"
	"github.com/blupulov/xwsDislinkt/user-service/startup/config"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (s *Server) Start() {
	//Router
	router := httprouter.New()
	//Mongo client
	mgClient := s.initMongoClient()
	//User service implementation (implementation of user interface)
	userServiceImpl := service.NewUserServiceImpl(mgClient)
	//User service
	userService := service.NewUserService(userServiceImpl)
	//User controller
	userController := controller.NewUserController(userService)

	//All routes (first is test rout)
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello")
	})

	router.GET("/user", userController.GetAll)
	router.POST("/user", userController.Register)

	//Starting server
	log.Println("user-service running on port " + s.config.Port)
	err := http.ListenAndServe(":"+s.config.Port, router)
	if err != nil {
		panic(err)
	}
}

func (s *Server) initMongoClient() *mongo.Client {
	uri := fmt.Sprintf("mongodb://%s:%s/", s.config.UserDBHost, s.config.UserDBPort)
	options := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
