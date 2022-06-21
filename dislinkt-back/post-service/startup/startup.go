package startup

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	psGrpc "github.com/blupulov/xwsDislinkt/common/proto/services/post-service"
	"github.com/blupulov/xwsDislinkt/post-service/controller"
	"github.com/blupulov/xwsDislinkt/post-service/handler"
	"github.com/blupulov/xwsDislinkt/post-service/service"
	"github.com/blupulov/xwsDislinkt/post-service/startup/config"
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
	//Post service implementation (implementation of post interface)
	postServiceImpl := service.NewPostServiceImpl(mgClient)
	//Post service
	postService := service.NewPostService(postServiceImpl)
	//Post controller
	postController := controller.NewPostController(postService)
	//Post handler
	postHandler := handler.NewPostHandler(postService)

	//All routes (first is test rout)
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello")
	})

	router.GET("/post", postController.GetAll)
	router.POST("/post", postController.Insert)
	router.PUT("/post/:postId/like/:userId", postController.Like)
	router.PUT("/post/:postId/comment", postController.AddComment)
	router.DELETE("/post/:postId/comment/:commentId", postController.DeleteCommentById)
	router.GET("/post/:postId", postController.GetById)
	router.DELETE("/post/:postId", postController.DeleteById)
	router.PUT("/post/:postId/dislike/:userId", postController.Dislike)
	router.GET("/post/:postId/:ownerId", postController.GetAllByOwnerId)

	wg := new(sync.WaitGroup)
	wg.Add(2)

	//Starting servers
	go s.startRestServer(router)
	go s.startGrpcServer(postHandler)
	wg.Wait()
}

//Rest server
func (s *Server) startRestServer(router *httprouter.Router) {
	log.Println("post-service (rest) running on port: " + s.config.RestPort)
	err := http.ListenAndServe(":"+s.config.RestPort, router)
	if err != nil {
		panic(err)
	}
}

//Mongo client (connection with db)
func (s *Server) initMongoClient() *mongo.Client {
	uri := fmt.Sprintf("mongodb://%s:%s/", s.config.PostDBHost, s.config.PostDBPort)
	options := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("mongodb client connected")
	return client
}

//Grpc server
func (s *Server) startGrpcServer(postHandler *handler.PostHandler) {
	listener, err := net.Listen("tcp", ":"+s.config.GrpcPort)
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	psGrpc.RegisterPostServiceServer(grpcServer, postHandler)
	log.Println("post-service (grpc) runing on port: " + s.config.GrpcPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
