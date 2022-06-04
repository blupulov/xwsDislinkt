package grpcClients

import (
	"log"

	company "github.com/blupulov/xwsDislinkt/common/proto/services/company-service"
	following "github.com/blupulov/xwsDislinkt/common/proto/services/following-service"
	posts "github.com/blupulov/xwsDislinkt/common/proto/services/post-service"
	users "github.com/blupulov/xwsDislinkt/common/proto/services/user-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserClient(addr string) users.UserServiceClient {
	conn, err := getConnection(addr)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to User service: %v", err)
	}

	return users.NewUserServiceClient(conn)
}

func NewPostsClient(addr string) posts.PostServiceClient {
	conn, err := getConnection(addr)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Post service: %v", err)
	}

	return posts.NewPostServiceClient(conn)
}

func NewFollowingClient(addr string) following.FollowingServiceClient {
	conn, err := getConnection(addr)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Following service: %v", err)
	}

	return following.NewFollowingServiceClient(conn)
}

func NewCompanyClient(addr string) company.CompanyServiceClient {
	conn, err := getConnection(addr)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Company service: %v", err)
	}

	return company.NewCompanyServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
