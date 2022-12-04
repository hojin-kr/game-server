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
	key := ds.Put(ctx, datastore.IncompleteKey("Account", nil), &pb.AccountRequest{Account: &pb.Account{RegisterTimestamp: tm}})
	ret := &pb.AccountReply{Account: &pb.Account{Id: key.ID, RegisterTimestamp: tm}}
	tracer.Trace(time.Now().Unix(), ret)
	return ret, nil
}

func (s *server) GetProfile(ctx context.Context, in *pb.ProfileRequest) (*pb.ProfileReply, error) {
	tracer.Trace(time.Now().UTC(), in)
	key := datastore.IDKey("Profile", in.Profile.GetAccountId(), nil)
	ds.Get(ctx, key, in.Profile)
	ret := &pb.ProfileReply{Profile: in.GetProfile()}
	tracer.Trace(time.Now().UTC(), ret)
	return ret, nil
}

func (s *server) UpdateProfile(ctx context.Context, in *pb.ProfileRequest) (*pb.ProfileReply, error) {
	tracer.Trace(time.Now().UTC(), in)
	if in.Profile.GetAccountId() == 0 {
		tracer.Trace(time.Now().UTC(), in, "ID is 0")
		ret := &pb.ProfileReply{Profile: in.GetProfile()}
		return ret, nil
	}
	ds.Put(ctx, datastore.IDKey("Profile", in.Profile.GetAccountId(), nil), in.Profile)
	ret := &pb.ProfileReply{Profile: in.GetProfile()}
	tracer.Trace(time.Now().UTC(), ret)
	return ret, nil
}

func (s *server) CreateGame(ctx context.Context, in *pb.GameRequest) (*pb.GameReply, error) {
	tracer.Trace(time.Now().UTC(), in)
	// Game 생성
	var game = in.Game
	key := ds.Put(ctx, datastore.IncompleteKey("Game", nil), &game)
	game.Id = key.ID
	_ = ds.Put(ctx, datastore.IDKey("Game", key.ID, nil), &game)
	ret := &pb.GameReply{Game: game}
	tracer.Trace(time.Now().UTC(), ret)
	return ret, nil
}

func (s *server) GetGame(ctx context.Context, in *pb.GameRequest) (*pb.GameReply, error) {
	tracer.Trace(time.Now().UTC(), in)
	ds.Get(ctx, datastore.IDKey("Game", in.Game.GetId(), nil), in.Game)
	ret := &pb.GameReply{Game: in.GetGame()}
	tracer.Trace(time.Now().UTC(), ret)
	return ret, nil
}

// filterdGames에서는 Game 목록만 반환하고 GetGame에서는 attend, place 부가 정보 반환
func (s *server) GetFilterdGames(ctx context.Context, in *pb.FilterdGamesRequest) (*pb.FilterdGamesReply, error) {
	tracer.Trace(time.Now().UTC(), in)
	// q := datastore.NewQuery("Game").Filter("A =", 12).Limit(30)
	var rounds []*pb.Game
	q := datastore.NewQuery("Game").Limit(30)
	ds.GetAll(ctx, q, &rounds)
	ret := &pb.FilterdGamesReply{Games: rounds}
	tracer.Trace(time.Now().UTC(), ret)
	return ret, nil
}

// func (s *server) JoinGame(ctx context.Context, in *pb.JoinRequest) (*pb.JoinReply, error) {
// 	tracer.Trace(time.Now().UTC(), in)
// 	if in.Join.GetId() == 0 {
// 		tracer.Trace(time.Now().UTC(), in, "ID is 0")
// 		ret := &pb.JoinReply{Join: &pb.Join{Id: in.Join.GetId()}}
// 		return ret, nil
// 	}

// 	// tracer.Trace(time.Now().UTC(), ret)
// 	return in, nil
// }

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
