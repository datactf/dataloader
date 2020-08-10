package loader

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

var mdbURI string

func LoadMongoDB(dbname, collname string, data []interface{}) error {

	client, err := mongo.NewClient(options.Client().ApplyURI(mdbURI))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if err := client.Connect(ctx); err != nil {
		return err
	}
	defer client.Disconnect(ctx)

	fmt.Printf("Loading data into `%s.%s`\n", dbname, collname)

	coll := client.Database(dbname).Collection(collname)
	opts := options.InsertMany().SetOrdered(false)
	iResult, err := coll.InsertMany(ctx, data, opts)

	if err != nil {
		return err
	}
	fmt.Printf("inserted: %d\n %+v", len(iResult.InsertedIDs), iResult.InsertedIDs)
	return nil
}

func buildMongoDBConnectionString() string {

	viper.AddConfigPath("/tmp")
	viper.SetConfigType("yaml")
	viper.SetConfigName("envs")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	port := viper.GetInt("port")
	replSetName := viper.Get("replicasetname")
	hostname, err := os.Hostname()

	if err != nil {
		panic("cannot read hostname from os")
	}

	fmt.Printf("the hostname of this machine is: %s\n", hostname)
	mdbURI := fmt.Sprintf("mongodb://%s:%d/admin?replSet=%s", hostname, port, replSetName)

	fmt.Println(mdbURI)
	return mdbURI
}

func init() {
	mdbURI = buildMongoDBConnectionString()
}
