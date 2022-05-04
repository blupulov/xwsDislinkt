package startup

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/blupulov/xwsDislinkt/user-service/handler"
	"github.com/blupulov/xwsDislinkt/user-service/service"
	"google.golang.org/grpc"

	usGrpc "github.com/blupulov/xwsDislinkt/common/proto/services/user-service"

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
	//User handler
	userHandler := handler.NewUserHandler(userService)

	//All routes (first is test rout)
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello")
	})

	router.GET("/user", userController.GetAll)
	router.POST("/user", userController.Register)
	router.POST("/user/login", userController.Login)
	router.PUT("/user/:userId/workExpirienceItem", userController.AddExpirience)
	router.PUT("/user/:userId/skillItem", userController.AddSkill)
	router.PUT("/user/:userId/educationItem", userController.AddEducation)

	wg := new(sync.WaitGroup)
	wg.Add(2)

	//Starting servers
	go s.startRestServer(router)
	go s.startGrpcServer(userHandler)
	wg.Wait()
}

func (s *Server) initMongoClient() *mongo.Client {
	uri := fmt.Sprintf("mongodb://%s:%s/", s.config.UserDBHost, s.config.UserDBPort)
	options := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), options)
	log.Println("mongodb client connected")
	if err != nil {
		log.Fatal(err)
	}
	return client
}

//Grpc server
func (s *Server) startGrpcServer(postHandler *handler.UserHandler) {
	listener, err := net.Listen("tcp", ":9051")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	usGrpc.RegisterUserServiceServer(grpcServer, postHandler)
	log.Println("user-service (grpc) runing on port: " + s.config.GrpcPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

//Rest server
func (s *Server) startRestServer(router *httprouter.Router) {
	log.Println("user-service (rest) running on port: " + s.config.RestPort)
	err := http.ListenAndServe(":"+s.config.RestPort, router)
	if err != nil {
		panic(err)
	}
}
