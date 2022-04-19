package startup

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/blupulov/xwsDislinkt/post-service/controller"
	"github.com/blupulov/xwsDislinkt/post-service/service"
	"github.com/blupulov/xwsDislinkt/post-service/startup/config"
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
	//Post service implementation (implementation of post interface)
	postServiceImpl := service.NewPostServiceImpl(mgClient)
	//Post service
	postService := service.NewPostService(postServiceImpl)
	//Post controller
	postController := controller.NewPostController(postService)

	//All routes (first is test rout)
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello")
	})

	router.GET("/post", postController.GetAll)
	router.POST("/post", postController.Insert)

	//Starting server
	log.Println("post-service running on port " + s.config.Port)
	err := http.ListenAndServe(":"+s.config.Port, router)
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
