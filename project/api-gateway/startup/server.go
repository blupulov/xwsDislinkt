package startup

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"

	"github.com/blupulov/xwsDislinkt/api-gateway/customHandlers"
	"github.com/blupulov/xwsDislinkt/api-gateway/startup/config"
	fsGw "github.com/blupulov/xwsDislinkt/common/proto/services/following-service"
	psGw "github.com/blupulov/xwsDislinkt/common/proto/services/post-service"
	usGw "github.com/blupulov/xwsDislinkt/common/proto/services/user-service"
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

	psEndpoint := fmt.Sprintf("%s:%s", s.config.PostServiceHost, s.config.PostServicePort)
	s.log.Println("post-service grpc endpoint port: " + s.config.PostServicePort)
	err := psGw.RegisterPostServiceHandlerFromEndpoint(context.TODO(), s.mux, psEndpoint, options)
	if err != nil {
		s.log.Println("POST-SERVICE GATEWAY PROBLEM")
		panic(err)
	}

	usEndpoint := fmt.Sprintf("%s:%s", s.config.UserServiceHost, s.config.UserServicePort)
	s.log.Println("user-service grpc endpoint port: " + s.config.UserServicePort)
	err = usGw.RegisterUserServiceHandlerFromEndpoint(context.TODO(), s.mux, usEndpoint, options)
	if err != nil {
		s.log.Println("USER-SERVICE GATEWAY PROBLEM")
		panic(err)
	}

	fsEndpoint := fmt.Sprintf("%s:%s", s.config.FollowingServiceHost, s.config.FollowingServicePort)
	s.log.Println("following-service grpc endpoint port: " + s.config.FollowingServicePort)
	err = fsGw.RegisterFollowingServiceHandlerFromEndpoint(context.TODO(), s.mux, fsEndpoint, options)
	if err != nil {
		s.log.Println("FOLLOWING-SERVICE GATEWAY PROBLEM")
		panic(err)
	}
}

func (s *Server) initCustomHandlers() {
	psEndpoint := fmt.Sprintf("%s:%s", s.config.PostServiceHost, s.config.PostServicePort)
	usEndpoint := fmt.Sprintf("%s:%s", s.config.UserServiceHost, s.config.UserServicePort)
	fsEndpoint := fmt.Sprintf("%s:%s", s.config.FollowingServiceHost, s.config.FollowingServicePort)

	customHandler := customHandlers.NewCustomHandler(usEndpoint, psEndpoint, fsEndpoint)
	customHandler.Init(s.mux)
}

func (s *Server) Start() {
	s.log.Println("api-gateway runing on port: " + s.config.Port)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "PUT", "DELETE", "GET", "PATCH"},
	})

	handler := c.Handler(s.mux)
	err := http.ListenAndServe(":"+s.config.Port, handler)
	if err != nil {
		panic(err)
	}
}
