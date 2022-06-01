package startup

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	csGrpc "github.com/blupulov/xwsDislinkt/common/proto/services/company-service"
	"github.com/blupulov/xwsDislinkt/company-service/handler"
	"github.com/blupulov/xwsDislinkt/company-service/service"
	"github.com/blupulov/xwsDislinkt/company-service/startup/config"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
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
	//Company service implementation
	companyServiceImpl := service.NewCompanyServiceImpl(mgClient)
	//Company service
	companyService := service.NewCompanyService(companyServiceImpl)
	//Company handler
	companyHandler := handler.NewCompanyHandler(companyService)

	//All routes (first is test rout)
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello")
	})

	wg := new(sync.WaitGroup)
	wg.Add(2)

	//Starting servers
	go s.startRestServer(router)
	go s.startGrpcServer(companyHandler)
	wg.Wait()
}

func (s *Server) initMongoClient() *mongo.Client {
	uri := fmt.Sprintf("mongodb://%s:%s/", s.config.CompanyDBHost, s.config.CompanyDBPort)
	options := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), options)
	log.Println("mongodb client connected")
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (s *Server) startRestServer(router *httprouter.Router) {
	log.Println("user-service (rest) running on port: " + s.config.RestPort)
	err := http.ListenAndServe(":"+s.config.RestPort, router)
	if err != nil {
		panic(err)
	}
}

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
