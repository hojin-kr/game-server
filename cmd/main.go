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
	ds "github.com/hojin-kr/haru/cmd/ds"
	pb "github.com/hojin-kr/haru/cmd/proto"
	"github.com/hojin-kr/haru/cmd/trace"
	"google.golang.org/grpc"
)

var (
	port       = flag.Int("port", 50051, "The server port")
	project_id = os.Getenv("PROJECT_ID")
	tracer     trace.Tracer
)

// server is used to implement UnimplementedServiceServer
type server struct {
	pb.UnimplementedVersion1Server
}

// Account account infomation

// CreateAccount implements CreateAccount
func (s *server) CreateAccount(ctx context.Context, in *pb.AccountRequest) (*pb.AccountReply, error) {
	tracer.Trace(time.Now().Unix(), in)
	tm := time.Now().Unix()
	// Putting an entity into the datastore under an incomplete key will cause a unique key to be generated for that entity, with a non-zero IntID.
	key := ds.Put(ctx, datastore.IncompleteKey("Account", nil), &pb.AccountRequest{RegisterTimestamp: tm})
	ret := &pb.AccountReply{Id: key.ID, RegisterTimestamp: tm}
	tracer.Trace(time.Now().Unix(), ret)
	return ret, nil
}

func (s *server) GetProfile(ctx context.Context, in *pb.ProfileRequest) (*pb.ProfileReply, error) {
	tracer.Trace(time.Now().UTC(), in)
	key := datastore.IDKey("Profile", in.GetId(), nil)
	ds.Get(ctx, key, in)
	ret := &pb.ProfileReply{Id: in.GetId(), Nickname: in.GetNickname()}
	tracer.Trace(time.Now().UTC(), ret)
	return ret, nil
}

func (s *server) UpdateProfile(ctx context.Context, in *pb.ProfileRequest) (*pb.ProfileReply, error) {
	tracer.Trace(time.Now().UTC(), in)
	if in.GetId() == 0 {
		tracer.Trace(time.Now().UTC(), in, "ID is 0")
		ret := &pb.ProfileReply{Id: in.GetId(), Nickname: in.GetNickname()}
		return ret, nil
	}
	key := datastore.IDKey("Profile", in.GetId(), nil)
	ds.Put(ctx, key, in)
	ret := &pb.ProfileReply{Id: in.GetId(), Nickname: in.GetNickname()}
	tracer.Trace(time.Now().UTC(), ret)
	return ret, nil
}

// CreateRound
func (s *server) CreateRound(ctx context.Context, in *pb.RoundRequest) (*pb.RoundReply, error) {
	tracer.Trace(time.Now().UTC(), in)
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
		Lat:                   in.Round.GetLat(),
		Long:                  in.Round.GetLong(),
		Updated:               time.Now().Unix(),
		PersonFull:            in.Round.GetPersonFull(),
		Person:                in.Round.GetPerson(),
		PlaceImg:              in.Round.GetPlaceImg(),
		TypeHole:              in.Round.GetTypeHole(),
	}
	// Putting an entity into the datastore under an incomplete key will cause a unique key to be generated for that entity, with a non-zero IntID.
	key := ds.Put(ctx, datastore.IncompleteKey("Round", nil), &round)
	round.Id = key.ID
	ret := &pb.RoundReply{Round: &round, Place: in.GetPlace(), Attend: in.GetAttend()}
	tracer.Trace(time.Now().UTC(), ret)
	return ret, nil
}

func (s *server) GetRound(ctx context.Context, in *pb.RoundRequest) (*pb.RoundReply, error) {
	tracer.Trace(time.Now().UTC(), in)
	key := datastore.IDKey("Round", in.Round.GetId(), nil)
	ds.Get(ctx, key, in.Round)
	ret := &pb.RoundReply{Round: in.GetRound(), Place: in.GetPlace(), Attend: in.GetAttend()}
	tracer.Trace(time.Now().UTC(), ret)
	return ret, nil
}

// filterdRounds에서는 Round 목록만 반환하고 GetRound에서는 attend, place 부가 정보 반환
func (s *server) GetFilterdRounds(ctx context.Context, in *pb.FilterdRoundsRequest) (*pb.FilterdRoundsReply, error) {
	tracer.Trace(time.Now().UTC(), in)
	// q := datastore.NewQuery("Round").Filter("A =", 12).Limit(30)
	var rounds []*pb.Round
	q := datastore.NewQuery("Round").Limit(30)
	ds.GetAll(ctx, q, &rounds)
	ret := &pb.FilterdRoundsReply{Rounds: rounds}
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
