package main

import (
	"os"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"cloud.google.com/go/datastore"
	"google.golang.org/grpc"
	pb "github.com/hojin-kr/haru/cmd/proto"
)


var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement UnimplementedServiceServer
type server struct {
	pb.UnimplementedServiceServer
}

// Account account infomation
type Account struct {
	RegisterTimestamp int64 `example:"1639056738"`
	DeviceID          string
	GoogleID          string
	AppleID           string
	LineID            string
	KakaoID           string
	ID                int64 `datastore:"-" example:"5373899369873408"`
}

// GetAccount implements GetAccount
func (s *server) GetAccount(ctx context.Context, in *pb.AccountRequest) (*pb.AccountReply, error) {
	var account Account
	var datastoreClient *datastore.Client
	datastoreClient, err := datastore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatalf("Should Not Datastore New Client"+err.Error())
	}
	account.RegisterTimestamp = time.Now().Unix()
	// Putting an entity into the datastore under an incomplete key will cause a unique key to be generated for that entity, with a non-zero IntID.
	key, err := datastoreClient.Put(ctx, datastore.IncompleteKey("Account", nil), &Account{RegisterTimestamp: account.RegisterTimestamp})
	if err != nil {
		log.Fatalf("Should Not Account Put:"+err.Error())
	}
	account.ID = key.ID
	return &pb.AccountReply{ID: account.ID }, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
