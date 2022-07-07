package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"cloud.google.com/go/datastore"
	FindPlaceFromText "github.com/hojin-kr/haru/cmd/places/findplacefromtext"
	NearbySearch "github.com/hojin-kr/haru/cmd/places/nearbysearch"
	PlaceDetails "github.com/hojin-kr/haru/cmd/places/placedetails"
	pb "github.com/hojin-kr/haru/cmd/proto"
	"github.com/hojin-kr/haru/cmd/trace"
	"google.golang.org/grpc"
)

var (
	port      = flag.Int("port", 50051, "The server port")
	apiKey    = flag.String("key", os.Getenv("APP_KEY"), "API Key for using Google Maps API.")
	inputType = flag.String("inputtype", "textquery", "The type of input. This can be one of either textquery or phonenumber.")
	tracer    trace.Tracer
)

// server is used to implement UnimplementedServiceServer
type server struct {
	pb.UnimplementedVersion1Server
}

// Account account infomation
type Account struct {
	RegisterTimestamp int64  `json:"register_timestamp,omitempty"`
	DeviceID          string `json:"device_id,omitempty"`
	GoogleID          string `json:"google_id,omitempty"`
	AppleID           string `json:"apple_id,omitempty"`
	LineID            string `json:"line_id,omitempty"`
	KakaoID           string `json:"kakao_id,omitempty"`
	ID                int64  `json:"id,omitempty"`
}

// Profile profile inforamtion
type Profile struct {
	Nickname string   `json:"nickname,omitempty"`
	ID       int64    `json:"id,omitempty"`
	Likes    []string `json:"likes,omitempty"`
}

// Wallet wallet
type Wallet struct {
	Coin string `json:"coin,omitempty"`
	ID   int64  `json:"id,omitempty"`
}

// Rating rating of place
type Rating struct {
	Point int64 `json:"point,omitempty"`
	Count int64 `json:"count,omitempty"`
	ID    int64 `json:"id,omitempty"`
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
	in.Id = key.ID
	ret := &pb.AccountReply{Id: in.GetId(), RegisterTimestamp: in.GetRegisterTimestamp()}
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
	key := datastore.IDKey("Profile", in.GetId(), nil)
	var profile Profile
	if err := datastoreClient.Get(ctx, key, &profile); err != nil {
		log.Printf("Should Not Profile get: " + err.Error())
		tracer.Trace(time.Now().UTC(), in, "Should Not Profile get", err.Error())
	}
	in.Nickname = profile.Nickname
	ret := &pb.ProfileReply{Id: in.GetId(), Nickname: in.GetNickname(), Likes: profile.Likes}
	tracer.Trace(time.Now().UTC(), ret)
	return ret, nil
}

func (s *server) UpdateProfile(ctx context.Context, in *pb.ProfileRequest) (*pb.ProfileReply, error) {
	tracer.Trace(time.Now().UTC(), in)
	if len(in.GetLikes()) > 1000 {
		tracer.Trace(time.Now().UTC(), in, "Too Many Likes")
		ret := &pb.ProfileReply{Id: in.GetId(), Nickname: in.GetNickname(), Likes: in.GetLikes()}
		return ret, nil
	}
	var datastoreClient *datastore.Client
	datastoreClient, err := datastore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Printf("Should Not Datastore New Client" + err.Error())
		tracer.Trace(time.Now().UTC(), in, "Should Not Datastore New Client", err.Error())
	}
	var profile Profile
	profile.ID = in.GetId()
	profile.Nickname = in.GetNickname()
	profile.Likes = in.GetLikes()
	key := datastore.IDKey("Profile", in.GetId(), nil)
	_, err = datastoreClient.Put(ctx, key, &profile)
	if err != nil {
		log.Printf("Should Not Profile get: " + err.Error())
		tracer.Trace(time.Now().UTC(), in, "Should Not Profile get", err.Error())
	}
	ret := &pb.ProfileReply{Id: in.GetId(), Nickname: in.GetNickname(), Likes: profile.Likes}
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
	key := datastore.IDKey("Rating", in.GetId(), nil)
	var rating Rating
	if err := datastoreClient.Get(ctx, key, &rating); err != nil {
		log.Printf("Should Not Rating get: " + err.Error())
		tracer.Trace(time.Now().UTC(), in, "Should Not Rating get", err.Error())
	}
	in.Point = rating.Point
	ret := &pb.PointReply{Id: in.GetId(), Point: in.GetPoint(), Count: rating.Count}
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
	key := datastore.IDKey("Rating", in.GetId(), nil)
	_ = datastoreClient.Get(ctx, key, &rating)
	rating.Count += 1
	rating.Point += in.GetPoint()
	rating.ID = in.GetId()
	_, err = datastoreClient.Put(ctx, key, &rating)
	if err != nil {
		log.Printf("Should Not Incr Rating Point: " + err.Error())
		tracer.Trace(time.Now().UTC(), in, "Should Not Rating get", err.Error())
	}
	in.Point = rating.Point
	ret := &pb.PointReply{Id: in.GetId(), Point: in.GetPoint(), Count: rating.Count}
	tracer.Trace(time.Now().UTC(), ret)
	return ret, nil
}

func (s *server) GetPlaceByInput(ctx context.Context, in *pb.PlaceRequest) (*pb.PlaceReply, error) {
	tracer.Trace(time.Now().UTC(), in)
	// todo cache
	PlaceID := FindPlaceFromText.Find(*apiKey, in.GetInput(), *inputType)
	// todo cache
	if PlaceID == "0" {
		return &pb.PlaceReply{PlaceId: "0"}, nil
	}
	var placeReply pb.PlaceReply
	marsharled, _ := json.Marshal(PlaceDetails.Find(*apiKey, PlaceID, in.GetLanguage()))
	_ = json.Unmarshal(marsharled, &placeReply)
	tracer.Trace(time.Now().UTC(), &placeReply)
	return &placeReply, nil
}

func (s *server) GetPlaceByID(ctx context.Context, in *pb.PlaceRequest) (*pb.PlaceReply, error) {
	tracer.Trace(time.Now().UTC(), in)
	var placeReply pb.PlaceReply
	place := PlaceDetails.Find(*apiKey, in.GetPlaceId(), in.GetLanguage())
	marsharled, _ := json.Marshal(place)
	tracer.Trace(time.Now().UTC(), string(marsharled))
	_ = json.Unmarshal(marsharled, &placeReply)
	tracer.Trace(time.Now().UTC(), &placeReply)
	return &placeReply, nil
}

type PlaceProifle struct {
	PlaceID      string         `json:"place_id,omitempty"`
	Populartimes []Populartimes `json:"popular_times,omitempty"`
	Likes        int64          `json:"likes,omitempty"`
}

type Populartimes struct {
	Weekday int64 `json:"weekday,omitempty"`
	Time    int64 `json:"time,omitempty"`
	Count   int64 `json:"count,omitempty"`
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
	_ = datastoreClient.Get(ctx, datastore.NameKey("Place", in.GetPlaceId(), nil), &placeProfile)
	if placeProfile.PlaceID == "" {
		placeProfile.PlaceID = in.GetPlaceId()
	}
	if in.IsLike {
		placeProfile.Likes += 1
	}
	if in.IsVisit {
		isHit := false
		WeekdayNow := int64(time.Now().UTC().Weekday())
		TimeNow := int64(time.Now().UTC().Hour())
		for k, v := range placeProfile.Populartimes {
			if v.Weekday == WeekdayNow && v.Time == TimeNow {
				placeProfile.Populartimes[k].Count++
				isHit = true
			}
		}
		if !isHit {
			placeProfile.Populartimes = append(placeProfile.Populartimes, Populartimes{Weekday: WeekdayNow, Time: TimeNow, Count: 1})
		}
	}
	_, err = datastoreClient.Put(ctx, datastore.NameKey("Place", in.GetPlaceId(), nil), &placeProfile)
	if err != nil {
		log.Printf("Should Not Update Place: " + err.Error())
		tracer.Trace(time.Now().UTC(), in, "Should Not Profile get", err.Error())
	}
	var PlaceProfileReply pb.PlaceProfileReply
	marsharled, _ := json.Marshal(placeProfile)
	_ = json.Unmarshal(marsharled, &PlaceProfileReply)
	tracer.Trace(time.Now().UTC(), &PlaceProfileReply)
	return &PlaceProfileReply, nil
}

func (s *server) GetNearbySearch(ctx context.Context, in *pb.PlaceRequest) (*pb.PlaceReplyList, error) {
	tracer.Trace(time.Now().UTC(), in)
	// todo next page token
	places := NearbySearch.Find(*apiKey, in.GetLocation(), uint(in.GetRadius()), in.GetKeyword(), in.GetLanguage(), in.GetPageToken(), in.GetOpenNow()).Results
	marsharled, _ := json.Marshal(places)
	var placeList = make([]*pb.PlaceReply, len(places))
	_ = json.Unmarshal(marsharled, &placeList)
	tracer.Trace(time.Now().UTC(), placeList)
	return &pb.PlaceReplyList{PlaceReplyList: placeList}, nil
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
