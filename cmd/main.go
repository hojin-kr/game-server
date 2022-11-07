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
	"github.com/hojin-kr/haru/cmd/trace"
	"google.golang.org/grpc"
)

var (
	port       = flag.Int("port", 50051, "The server port")
	project_id = flag.String("PROJECT_ID", "golf-367911", "")
	tracer     trace.Tracer
)

// server is used to implement UnimplementedServiceServer
type server struct {
	pb.UnimplementedVersion1Server
}

// Account account infomation

// CreateAccount implements CreateAccount
func (s *server) CreateAccount(ctx context.Context, in *pb.AccountRequest) (*pb.AccountReply, error) {
	tracer.Trace(time.Now().UTC(), in)
	var datastoreClient *datastore.Client
	datastoreClient, err := datastore.NewClient(ctx, *project_id)
	if err != nil {
		log.Fatalf("Should Not Datastore New Client" + err.Error())
		tracer.Trace(time.Now().UTC(), in, "Should Not Datastore New Client", err.Error())
	}
	in.RegisterTimestamp = time.Now().Unix()
	// Putting an entity into the datastore under an incomplete key will cause a unique key to be generated for that entity, with a non-zero IntID.
	key, err := datastoreClient.Put(ctx, datastore.IncompleteKey("Account", nil), &pb.AccountRequest{RegisterTimestamp: in.GetRegisterTimestamp()})
	if err != nil {
		log.Printf("Should Not Account Put:" + err.Error())
		tracer.Trace(time.Now().UTC(), in, "Should Not Account Put", err.Error())
	}
	in.Id = key.ID
	ret := &pb.AccountReply{Id: in.GetId(), RegisterTimestamp: in.GetRegisterTimestamp()}
	tracer.Trace(time.Now().UTC(), ret)
	return ret, nil
}

func (s *server) GetProfile(ctx context.Context, in *pb.ProfileRequest) (*pb.ProfileReply, error) {
	tracer.Trace(time.Now().UTC(), in)
	var datastoreClient *datastore.Client
	datastoreClient, err := datastore.NewClient(ctx, *project_id)
	if err != nil {
		log.Printf("Should Not Datastore New Client" + err.Error())
		tracer.Trace(time.Now().UTC(), in, "Should Not Datastore New Client", err.Error())
	}
	key := datastore.IDKey("Profile", in.GetId(), nil)
	if err := datastoreClient.Get(ctx, key, in); err != nil {
		log.Printf("Should Not Profile get: " + err.Error())
		tracer.Trace(time.Now().UTC(), in, "Should Not Profile get", err.Error())
	}
	ret := &pb.ProfileReply{Id: in.GetId(), Nickname: in.GetNickname()}
	tracer.Trace(time.Now().UTC(), ret)
	return ret, nil
}

func (s *server) UpdateProfile(ctx context.Context, in *pb.ProfileRequest) (*pb.ProfileReply, error) {
	tracer.Trace(time.Now().UTC(), in)
	var datastoreClient *datastore.Client
	datastoreClient, err := datastore.NewClient(ctx, *project_id)
	if err != nil {
		log.Printf("Should Not Datastore New Client" + err.Error())
		tracer.Trace(time.Now().UTC(), in, "Should Not Datastore New Client", err.Error())
	}
	if in.GetId() == 0 {
		tracer.Trace(time.Now().UTC(), in, "ID is 0")
		ret := &pb.ProfileReply{Id: in.GetId(), Nickname: in.GetNickname()}
		return ret, nil
	}
	key := datastore.IDKey("Profile", in.GetId(), nil)
	_, err = datastoreClient.Put(ctx, key, in)
	if err != nil {
		log.Printf("Should Not Profile get: " + err.Error())
		tracer.Trace(time.Now().UTC(), in, "Should Not Profile get", err.Error())
	}
	ret := &pb.ProfileReply{Id: in.GetId(), Nickname: in.GetNickname()}
	tracer.Trace(time.Now().UTC(), ret)
	return ret, nil
}

// CreateRound
func (s *server) CreateRound(ctx context.Context, in *pb.RoundRequest) (*pb.RoundReply, error) {
	tracer.Trace(time.Now().UTC(), in)
	var datastoreClient *datastore.Client
	datastoreClient, err := datastore.NewClient(ctx, *project_id)
	if err != nil {
		log.Fatalf("Should Not Datastore New Client" + err.Error())
		tracer.Trace(time.Now().UTC(), in, "Should Not Datastore New Client", err.Error())
	}
	// Round 생성
	var round = pb.Round{
		Host:                  in.Round.GetHost(),
		Time:                  in.Round.GetTime(),
		Price:                 in.Round.GetPrice(),
		TypePlay:              in.Round.GetTypePlay(),
		TypeAge:               in.Round.GetTypeAge(),
		TypeSex:               in.Round.GetTypeSex(),
		TypeScoreOfGross:      in.Round.GetTypeScoreOfGross(),
		TypeExperienceOfYears: in.Round.GetTypeExperienceOfYears(),
		PlaceId:               in.Round.GetPlaceId(),
		PlaceName:             in.Round.GetPlaceName(),
		PlaceAddress:          in.Round.GetPlaceAddress(),
		ShortAddress:          in.Round.GetShortAddress(),
		Lat:                   in.Place.GetLat(),
		Long:                  in.Place.GetLong(),
		Updated:               time.Now().Unix(),
	}
	// Putting an entity into the datastore under an incomplete key will cause a unique key to be generated for that entity, with a non-zero IntID.
	key, err := datastoreClient.Put(ctx, datastore.IncompleteKey("Round", nil), &round)
	if err != nil {
		log.Printf("Should Not Account Put:" + err.Error())
		tracer.Trace(time.Now().UTC(), in, "Should Not Account Put", err.Error())
	}
	round.Id = key.ID
	ret := &pb.RoundReply{Round: &round, Place: in.GetPlace(), Attend: in.GetAttend()}
	tracer.Trace(time.Now().UTC(), ret)
	return ret, nil
}

func main() {
	flag.Parse()
	tracer = trace.New(os.Stdout)
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		log.Printf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterVersion1Server(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}
}
