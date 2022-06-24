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
	FindPlaceFromText "github.com/hojin-kr/haru/cmd/places/findplacefromtext"
	PlaceDetails "github.com/hojin-kr/haru/cmd/places/placedetails"
	pb "github.com/hojin-kr/haru/cmd/proto"
	"google.golang.org/grpc"
)

var (
	port      = flag.Int("port", 50051, "The server port")
	apiKey    = flag.String("key", "", "API Key for using Google Maps API.")
	inputType = flag.String("inputtype", "textquery", "The type of input. This can be one of either textquery or phonenumber.")
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

// Rating rating of place
type Rating struct {
	Point int64 `example:"100"`
	Count int64 `example:"100"`
	ID    int64 `datastore:"-" example:"PlaceID"`
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
		log.Printf("Should Not Account Put:" + err.Error())
	}
	in.ID = key.ID
	return &pb.AccountReply{ID: in.GetID(), RegisterTimestamp: in.GetRegisterTimestamp()}, nil
}

func (s *server) GetProfile(ctx context.Context, in *pb.ProfileRequest) (*pb.ProfileReply, error) {
	var datastoreClient *datastore.Client
	datastoreClient, err := datastore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Printf("Should Not Datastore New Client" + err.Error())
	}
	key := datastore.IDKey("Profile", in.GetID(), nil)
	var profile Profile
	if err := datastoreClient.Get(ctx, key, &profile); err != nil {
		log.Printf("Should Not Profile get: " + err.Error())
	}
	in.Nickname = profile.Nickname
	return &pb.ProfileReply{ID: in.GetID(), Nickname: in.GetNickname()}, nil
}

func (s *server) UpdateProfile(ctx context.Context, in *pb.ProfileRequest) (*pb.ProfileReply, error) {
	var datastoreClient *datastore.Client
	datastoreClient, err := datastore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Printf("Should Not Datastore New Client" + err.Error())
	}
	var profile Profile
	profile.ID = in.GetID()
	profile.Nickname = in.GetNickname()
	key := datastore.IDKey("Profile", in.GetID(), nil)
	_, err = datastoreClient.Put(ctx, key, &profile)
	if err != nil {
		log.Printf("Should Not Profile get: " + err.Error())
	}
	return &pb.ProfileReply{ID: in.GetID(), Nickname: in.GetNickname()}, nil
}

func (s *server) GetPoint(ctx context.Context, in *pb.PointRequest) (*pb.PointReply, error) {
	var datastoreClient *datastore.Client
	datastoreClient, err := datastore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Printf("Should Not Datastore New Client" + err.Error())
	}
	key := datastore.IDKey("Rating", in.GetID(), nil)
	var rating Rating
	if err := datastoreClient.Get(ctx, key, &rating); err != nil {
		log.Printf("Should Not Rating get: " + err.Error())
	}
	in.Point = rating.Point
	return &pb.PointReply{ID: in.GetID(), Point: in.GetPoint(), Count: rating.Count}, nil
}

func (s *server) IncrPoint(ctx context.Context, in *pb.PointRequest) (*pb.PointReply, error) {
	var datastoreClient *datastore.Client
	datastoreClient, err := datastore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Printf("Should Not Datastore New Client" + err.Error())
	}
	var rating Rating
	key := datastore.IDKey("Rating", in.GetID(), nil)
	_ = datastoreClient.Get(ctx, key, &rating)
	rating.Count += 1
	rating.Point += in.GetPoint()
	rating.ID = in.GetID()
	_, err = datastoreClient.Put(ctx, key, &rating)
	if err != nil {
		log.Printf("Should Not Incr Rating Point: " + err.Error())
	}
	in.Point = rating.Point
	return &pb.PointReply{ID: in.GetID(), Point: in.GetPoint(), Count: rating.Count}, nil
}

func (s *server) GetPlace(ctx context.Context, in *pb.PlaceRequest) (*pb.PlaceReply, error) {
	// todo cache
	PlaceID := FindPlaceFromText.Find(*apiKey, in.GetInput(), *inputType)
	// todo cache
	PlaceDetails := PlaceDetails.Find(*apiKey, PlaceID, in.GetLanguage())
	return &pb.PlaceReply{
		Name:                 PlaceDetails.Name,
		FormattedAddress:     PlaceDetails.FormattedAddress,
		FormattedPhoneNumber: PlaceDetails.FormattedPhoneNumber,
		Icon:                 PlaceDetails.Icon,
		PlaceID:              PlaceDetails.PlaceID,
		Website:              PlaceDetails.Website,
		URL:                  PlaceDetails.URL,
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
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
