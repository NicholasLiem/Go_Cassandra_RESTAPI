package datastore

import (
	"fmt"
	"github.com/gocql/gocql"
)

var Session *gocql.Session

func init() {
	var err error
	cluster := gocql.NewCluster("cassandra1", "cassandra2", "cassandra3")
	cluster.Port = 9042
	cluster.Keyspace = "restfulapi"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("Cassandra is initialized")
}
