package ds

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/datastore"
)

var (
	project_id = os.Getenv("PROJECT_ID")
)

func getClient(ctx context.Context) *datastore.Client {
	var client *datastore.Client
	client, err := datastore.NewClient(ctx, project_id)
	if err != nil {
		log.Printf("get ds client" + err.Error())
	}
	return client
}

func Get(ctx context.Context, key *datastore.Key, dst interface{}) (err error) {
	client := getClient(ctx)
	if err := client.Get(ctx, key, dst); err != nil {
		log.Printf("get ds" + err.Error())
	}
	return err
}

func Put(ctx context.Context, key *datastore.Key, src interface{}) (_key *datastore.Key) {
	client := getClient(ctx)
	key, err := client.Put(ctx, key, src)
	if err != nil {
		log.Printf("put ds" + err.Error())
	}
	return key
}

func GetAll(ctx context.Context, query *datastore.Query, dst interface{}) (keys []*datastore.Key, err error) {
	client := getClient(ctx)
	keys, err = client.GetAll(ctx, query, dst)
	if err != nil {
		log.Printf("query ds" + err.Error())
	}
	return keys, err
}
