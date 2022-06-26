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
	"github.com/hojin-kr/haru/cmd/trace"
	"google.golang.org/grpc"
)

var (
	port      = flag.Int("port", 50051, "The server port")
	apiKey    = flag.String("key", os.Getenv("APP_KEY"), "API Key for using Google Maps API.")
	inputType = flag.String("inputtype", "textquery", "The type of input. This can be one of either textquery or phonenumber.")
)

var tracer trace.Tracer

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
	Nickname string   `example:"myNickname"`
	ID       int64    `datastore:"-" example:"5373899369873408"`
	Likes    []string `example:"ID,ID,ID"`
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
	tracer.Trace(time.Now().UTC(), in)
	var datastoreClient *datastore.Client
	datastoreClient, err := datastore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatalf("Should Not Datastore New Client" + err.Error())
		tracer.Trace(time.Now().UTC(), in, "Should Not Datastore New Client", err.Error())
	}
	in.RegisterTimestamp = time.Now().Unix()
	// Putting an entity into the datastore under an incomplete key will cause a unique key to be generated for that entity, with a non-zero IntID.
	key, err := datastoreClient.Put(ctx, datastore.IncompleteKey("Account", nil), &Account{RegisterTimestamp: in.GetRegisterTimestamp()})
	if err != nil {
		log.Printf("Should Not Account Put:" + err.Error())
		tracer.Trace(time.Now().UTC(), in, "Should Not Account Put", err.Error())
	}
	in.ID = key.ID
	ret := &pb.AccountReply{ID: in.GetID(), RegisterTimestamp: in.GetRegisterTimestamp()}
	tracer.Trace(time.Now().UTC(), ret)
	return ret, nil
}

func (s *server) GetProfile(ctx context.Context, in *pb.ProfileRequest) (*pb.ProfileReply, error) {
	tracer.Trace(time.Now().UTC(), in)
	var datastoreClient *datastore.Client
	datastoreClient, err := datastore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Printf("Should Not Datastore New Client" + err.Error())
		tracer.Trace(time.Now().UTC(), in, "Should Not Datastore New Client", err.Error())
	}
	key := datastore.IDKey("Profile", in.GetID(), nil)
	var profile Profile
	if err := datastoreClient.Get(ctx, key, &profile); err != nil {
		log.Printf("Should Not Profile get: " + err.Error())
		tracer.Trace(time.Now().UTC(), in, "Should Not Profile get", err.Error())
	}
	in.Nickname = profile.Nickname
	ret := &pb.ProfileReply{ID: in.GetID(), Nickname: in.GetNickname(), Likes: profile.Likes}
	tracer.Trace(time.Now().UTC(), ret)
	return ret, nil
}

func (s *server) UpdateProfile(ctx context.Context, in *pb.ProfileRequest) (*pb.ProfileReply, error) {
	tracer.Trace(time.Now().UTC(), in)
	if len(in.GetLikes()) > 1000 {
		tracer.Trace(time.Now().UTC(), in, "Too Many Likes")
		ret := &pb.ProfileReply{ID: in.GetID(), Nickname: in.GetNickname(), Likes: in.GetLikes()}
		return ret, nil
	}
	var datastoreClient *datastore.Client
	datastoreClient, err := datastore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Printf("Should Not Datastore New Client" + err.Error())
		tracer.Trace(time.Now().UTC(), in, "Should Not Datastore New Client", err.Error())
	}
	var profile Profile
	profile.ID = in.GetID()
	profile.Nickname = in.GetNickname()
	profile.Likes = in.GetLikes()
	key := datastore.IDKey("Profile", in.GetID(), nil)
	_, err = datastoreClient.Put(ctx, key, &profile)
	if err != nil {
		log.Printf("Should Not Profile get: " + err.Error())
		tracer.Trace(time.Now().UTC(), in, "Should Not Profile get", err.Error())
	}
	ret := &pb.ProfileReply{ID: in.GetID(), Nickname: in.GetNickname(), Likes: profile.Likes}
	tracer.Trace(time.Now().UTC(), ret)
	return ret, nil
}

func (s *server) GetPoint(ctx context.Context, in *pb.PointRequest) (*pb.PointReply, error) {
	tracer.Trace(time.Now().UTC(), in)
	var datastoreClient *datastore.Client
	datastoreClient, err := datastore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Printf("Should Not Datastore New Client" + err.Error())
		tracer.Trace(time.Now().UTC(), in, "Should Not Datastore New Client", err.Error())
	}
	key := datastore.IDKey("Rating", in.GetID(), nil)
	var rating Rating
	if err := datastoreClient.Get(ctx, key, &rating); err != nil {
		log.Printf("Should Not Rating get: " + err.Error())
		tracer.Trace(time.Now().UTC(), in, "Should Not Rating get", err.Error())
	}
	in.Point = rating.Point
	ret := &pb.PointReply{ID: in.GetID(), Point: in.GetPoint(), Count: rating.Count}
	tracer.Trace(time.Now().UTC(), ret)
	return ret, nil
}

func (s *server) IncrPoint(ctx context.Context, in *pb.PointRequest) (*pb.PointReply, error) {
	tracer.Trace(time.Now().UTC(), in)
	var datastoreClient *datastore.Client
	datastoreClient, err := datastore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Printf("Should Not Datastore New Client" + err.Error())
		tracer.Trace(time.Now().UTC(), in, "Should Not Datastore New Client", err.Error())
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
		tracer.Trace(time.Now().UTC(), in, "Should Not Rating get", err.Error())
	}
	in.Point = rating.Point
	ret := &pb.PointReply{ID: in.GetID(), Point: in.GetPoint(), Count: rating.Count}
	tracer.Trace(time.Now().UTC(), ret)
	return ret, nil
}

func (s *server) GetPlace(ctx context.Context, in *pb.PlaceRequest) (*pb.PlaceReply, error) {
	tracer.Trace(time.Now().UTC(), in)
	// todo cache
	PlaceID := FindPlaceFromText.Find(*apiKey, in.GetInput(), *inputType)
	// todo cache
	if PlaceID == "0" {
		return &pb.PlaceReply{PlaceID: "0"}, nil
	}
	PlaceDetails := PlaceDetails.Find(*apiKey, PlaceID, in.GetLanguage())
	var PhotoReferences []string
	for i := 0; i < len(PlaceDetails.Photos); i++ {
		PhotoReferences = append(PhotoReferences, PlaceDetails.Photos[i].PhotoReference)
	}
	var Reviews []string
	for i := 0; i < len(PlaceDetails.Reviews); i++ {
		Reviews = append(Reviews, PlaceDetails.Reviews[i].Text)
	}
	var WeekdayText []string
	if PlaceDetails.OpeningHours != nil {
		WeekdayText = PlaceDetails.OpeningHours.WeekdayText
	}
	ret := &pb.PlaceReply{
		Name:                 PlaceDetails.Name,
		FormattedAddress:     PlaceDetails.FormattedAddress,
		FormattedPhoneNumber: PlaceDetails.FormattedPhoneNumber,
		Icon:                 PlaceDetails.Icon,
		PlaceID:              PlaceDetails.PlaceID,
		Website:              PlaceDetails.Website,
		URL:                  PlaceDetails.URL,
		Types:                PlaceDetails.Types,
		Rating:               PlaceDetails.Rating,
		UserRatingsTotal:     int64(PlaceDetails.UserRatingsTotal),
		WeekdayText:          WeekdayText,
		PhotoReferences:      PhotoReferences,
		BusinessStatus:       PlaceDetails.BusinessStatus,
		Reviews:              Reviews,
		LocationLat:          fmt.Sprintf("%f", PlaceDetails.Geometry.Location.Lat),
		LocationLng:          fmt.Sprintf("%f", PlaceDetails.Geometry.Location.Lng),
	}
	tracer.Trace(time.Now().UTC(), ret)
	return ret, nil
}

type PlaceProifle struct {
	PlaceID      string
	Populartimes []int64
	Likes        int64
}

// UpdatePlace updatePlace additional info
func (s *server) UpdatePlaceProfile(ctx context.Context, in *pb.PlaceRequest) (*pb.PlaceProfileReply, error) {
	tracer.Trace(time.Now().UTC(), in)
	var datastoreClient *datastore.Client
	datastoreClient, err := datastore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Printf("Should Not Datastore New Client" + err.Error())
		tracer.Trace(time.Now().UTC(), in, "Should Not Datastore New Client", err.Error())
	}
	var placeProfile PlaceProifle
	_ = datastoreClient.Get(ctx, datastore.NameKey("Place", in.GetPlaceID(), nil), &placeProfile)
	if placeProfile.PlaceID == "" {
		placeProfile.PlaceID = in.GetPlaceID()
	}
	if in.IsLike {
		placeProfile.Likes += 1
	}
	// todo 시간대별 방문수 누적
	_, err = datastoreClient.Put(ctx, datastore.NameKey("Place", in.GetPlaceID(), nil), &placeProfile)
	if err != nil {
		log.Printf("Should Not Update Place: " + err.Error())
		tracer.Trace(time.Now().UTC(), in, "Should Not Profile get", err.Error())
	}
	ret := &pb.PlaceProfileReply{PlaceID: placeProfile.PlaceID, Likes: placeProfile.Likes}
	tracer.Trace(time.Now().UTC(), ret)
	return ret, nil
}

func main() {
	flag.Parse()
	tracer = trace.New(os.Stdout)
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
