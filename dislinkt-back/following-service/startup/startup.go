package startup

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	fsGrpc "github.com/blupulov/xwsDislinkt/common/proto/services/following-service"

	"github.com/blupulov/xwsDislinkt/following-service/controller"
	"github.com/blupulov/xwsDislinkt/following-service/handler"
	"github.com/blupulov/xwsDislinkt/following-service/service"
	"github.com/blupulov/xwsDislinkt/following-service/startup/config"
	"github.com/julienschmidt/httprouter"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
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
	//neo4j driver
	driver, err := s.createDriver()
	if err != nil {
		log.Fatal(err)
	}
	//Follow service implementation
	followServiceImpl := service.NewFollowServiceImpl(driver)
	//Follow service
	followService := service.NewFollowService(followServiceImpl)
	//Follow rest contoller
	followController := controller.NewFollowController(followService)
	//Follow grpc handler
	followHandler := handler.NewFollowingHandler(followService)

	//All routes (first is test rout)
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello")
	})
	router.POST("/follow/:userId", followController.Insert)
	router.POST("/follow/:userId/follow/:targetUserId", followController.Follow)
	router.POST("/follow/:userId/request/:targetUserId", followController.SendRequest)
	//answer := {t, f}
	router.PUT("/follow/:senderId/request/:receiverId/answer/:answer", followController.RequestAnswer)

	router.DELETE("/follow/unFollow/:userId/:targetUserId", followController.UnFollow)
	router.DELETE("/follow/request/:senderId/:receiverId", followController.DeleteRequest)

	router.GET("/follow/:userId/followers", followController.GetAllUserFollowers)
	router.GET("/follow/:userId/following", followController.GetAllFollowingUsers)
	router.GET("/follow/:userId/received/request", followController.GetAllReceivedRequests)
	router.GET("/follow/:userId/sent/request", followController.GetAllSentRequests)

	wg := new(sync.WaitGroup)
	wg.Add(2)

	//Starting servers
	go s.startRestServer(router)
	go s.startGrpcServer(followHandler)
	wg.Wait()
}

//Rest server
func (s *Server) startRestServer(router *httprouter.Router) {
	log.Println("following-service (rest) running on port: " + s.config.RestPort)
	err := http.ListenAndServe(":"+s.config.RestPort, router)
	if err != nil {
		panic(err)
	}
}

//Grpc server
func (s *Server) startGrpcServer(followingHandler *handler.FollowingHandler) {
	listener, err := net.Listen("tcp", ":"+s.config.GrpcPort)
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	fsGrpc.RegisterFollowingServiceServer(grpcServer, followingHandler)
	log.Println("following-service (grpc) running on port: " + s.config.GrpcPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

func (s *Server) createDriver() (neo4j.Driver, error) {
	uri := fmt.Sprintf("neo4j://%s:%s", s.config.FollowDBHost, s.config.FollowDBPort)
	return neo4j.NewDriver(uri, neo4j.BasicAuth(s.config.FollowDBUsername, s.config.FollowDBPassword, ""))
}

func closeDriver(driver neo4j.Driver) error {
	return driver.Close()
}
