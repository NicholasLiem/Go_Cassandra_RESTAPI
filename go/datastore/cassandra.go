package datastore

import (
	"fmt"
	"github.com/gocql/gocql"
	"os"
)

var Session *gocql.Session

func init() {
	var err error
	cassandraHost := os.Getenv("CASSANDRA_HOST")
	cluster := gocql.NewCluster(cassandraHost)
	cluster.Port = 9042
	cluster.Keyspace = "restfulapi"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("Cassandra is initialized")
}
