package startup

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/blupulov/xwsDislinkt/api-gateway/startup/config"
	psGw "github.com/blupulov/xwsDislinkt/common/proto/services/post-service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	config *config.Config
	mux    *runtime.ServeMux
	log    *log.Logger
}

func NewServer(config *config.Config, logger *log.Logger) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
		log:    logger,
	}
	server.initHandlers()
	server.initCustomHandlers()
	return server
}

func (s *Server) initHandlers() {
	options := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	//endpoints

	psEndpoint := fmt.Sprintf("%s:%s", s.config.PostServiceHost, s.config.PostServicePort)
	s.log.Println("post-service grpc endpoint port: " + s.config.PostServicePort)
	err := psGw.RegisterPostServiceHandlerFromEndpoint(context.TODO(), s.mux, psEndpoint, options)
	if err != nil {
		s.log.Println("POST-SERVICE GATEWAY PROBLEM")
		panic(err)
	}
}

func (s *Server) initCustomHandlers() {

}

func (s *Server) Start() {
	s.log.Println("api-gateway runing on port: " + s.config.Port)
	err := http.ListenAndServe(":"+s.config.Port, s.mux)
	if err != nil {
		panic(err)
	}
}
