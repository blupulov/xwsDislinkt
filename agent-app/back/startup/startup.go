package startup

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/blupulov/xwsDislinkt/agent-app/back/controllers"
	"github.com/blupulov/xwsDislinkt/agent-app/back/service"
	"github.com/blupulov/xwsDislinkt/agent-app/back/startup/config"
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
	//Service implementations
	companyServiceImpl := service.NewCompanyServiceImpl(mgClient)
	userServiceImpl := service.NewUserServiceImpl(mgClient)
	//Services
	companyService := service.NewCompanyService(companyServiceImpl)
	userService := service.NewUserService(userServiceImpl)
	//Controller
	controller := controllers.NewController(companyService, userService)

	//All routes (first is test rout)
	registerRoutes(router, controller)

	wg := new(sync.WaitGroup)
	wg.Add(1)

	//Starting servers
	go s.startRestServer(router)
	wg.Wait()
}

func registerRoutes(router *httprouter.Router, controller *controllers.Controller) {
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello")
	})

	//Companies
	router.GET("/agentApp/company", controller.GetAllCompanies)
	//Users

	//Combined
}

func (s *Server) initMongoClient() *mongo.Client {
	uri := fmt.Sprintf("mongodb://%s:%s/", s.config.AgentAppDBHost, s.config.AgentAppDBPort)
	options := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), options)
	log.Println("mongodb client connected")
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (s *Server) startRestServer(router *httprouter.Router) {
	log.Println("company-service (rest) running on port: " + s.config.RestPort)
	err := http.ListenAndServe(":"+s.config.RestPort, router)
	if err != nil {
		panic(err)
	}
}

/*
func (s *Server) startGrpcServer(companyHandler *handler.CompanyHandler) {
	listener, err := net.Listen("tcp", ":"+s.config.GrpcPort)
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	csGrpc.RegisterCompanyServiceServer(grpcServer, companyHandler)
	log.Println("company-service (grpc) running on port: " + s.config.GrpcPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
*/
