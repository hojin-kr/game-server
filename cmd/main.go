package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"cloud.google.com/go/datastore"
	pb "github.com/hojin-kr/haru/cmd/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement UnimplementedServiceServer
type server struct {
	pb.UnimplementedVersion1Server
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

// Profile profile inforamtion
type Profile struct {
	Nickname string `example:"myNickname"`
	ID       int64  `datastore:"-" example:"5373899369873408"`
}

// Wallet wallet
type Wallet struct {
	Coin string
	ID   int64 `datastore:"-" example:"5373899369873408"`
}

// Attack attack boss infomation
type Attack struct {
	Point int64 `example:"100"`
	ID    int64 `datastore:"-" example:"5373899369873408"`
}

// CreateAccount implements CreateAccount
func (s *server) CreateAccount(ctx context.Context, in *pb.AccountRequest) (*pb.AccountReply, error) {
	var datastoreClient *datastore.Client
	datastoreClient, err := datastore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatalf("Should Not Datastore New Client" + err.Error())
	}
	in.RegisterTimestamp = time.Now().Unix()
	// Putting an entity into the datastore under an incomplete key will cause a unique key to be generated for that entity, with a non-zero IntID.
	key, err := datastoreClient.Put(ctx, datastore.IncompleteKey("Account", nil), &Account{RegisterTimestamp: in.GetRegisterTimestamp()})
	if err != nil {
		log.Fatalf("Should Not Account Put:" + err.Error())
	}
	in.ID = key.ID
	return &pb.AccountReply{ID: in.GetID(), RegisterTimestamp: in.GetRegisterTimestamp()}, nil
}

func (s *server) GetProfile(ctx context.Context, in *pb.ProfileRequest) (*pb.ProfileReply, error) {
	var datastoreClient *datastore.Client
	datastoreClient, err := datastore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatalf("Should Not Datastore New Client" + err.Error())
	}
	key := datastore.IDKey("Profile", in.GetID(), nil)
	var profile Profile
	if err := datastoreClient.Get(ctx, key, &profile); err != nil {
		log.Fatalf("Should Not Profile get: " + err.Error())
	}
	in.Nickname = profile.Nickname
	return &pb.ProfileReply{ID: in.GetID(), Nickname: in.GetNickname()}, nil
}

func (s *server) UpdateProfile(ctx context.Context, in *pb.ProfileRequest) (*pb.ProfileReply, error) {
	var datastoreClient *datastore.Client
	datastoreClient, err := datastore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatalf("Should Not Datastore New Client" + err.Error())
	}
	var profile Profile
	profile.ID = in.GetID()
	profile.Nickname = in.GetNickname()
	key := datastore.IDKey("Profile", in.GetID(), nil)
	_, err = datastoreClient.Put(ctx, key, &profile)
	if err != nil {
		log.Fatalf("Should Not Profile get: " + err.Error())
	}
	return &pb.ProfileReply{ID: in.GetID(), Nickname: in.GetNickname()}, nil
}

func (s *server) GetPoint(ctx context.Context, in *pb.PointRequest) (*pb.PointReply, error) {
	var datastoreClient *datastore.Client
	datastoreClient, err := datastore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatalf("Should Not Datastore New Client" + err.Error())
	}
	key := datastore.IDKey("Attack", in.GetID(), nil)
	var attack Attack
	if err := datastoreClient.Get(ctx, key, &attack); err != nil {
		log.Fatalf("Should Not Attack get: " + err.Error())
	}
	in.Point = attack.Point
	return &pb.PointReply{ID: in.GetID(), Point: in.GetPoint()}, nil
}

func (s *server) IncrPoint(ctx context.Context, in *pb.PointRequest) (*pb.PointReply, error) {
	var datastoreClient *datastore.Client
	datastoreClient, err := datastore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatalf("Should Not Datastore New Client" + err.Error())
	}
	var attack Attack
	key := datastore.IDKey("Attack", in.GetID(), nil)
	_ = datastoreClient.Get(ctx, key, &attack)
	attack.Point += in.GetPoint()
	attack.ID = in.GetID()
	_, err = datastoreClient.Put(ctx, key, &attack)
	if err != nil {
		log.Fatalf("Should Not Incr Boss Pint: " + err.Error())
	}
	in.Point = attack.Point
	return &pb.PointReply{ID: in.GetID(), Point: in.GetPoint()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterVersion1Server(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
