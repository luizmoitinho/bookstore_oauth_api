package cassandra

import (
	"log"
	"os"
	"strconv"

	"github.com/gocql/gocql"
)

var (
	cluster *gocql.ClusterConfig
)

func InitCluster() {
	//Connect to Cassandra Cluster:
	cluster = gocql.NewCluster(os.Getenv("CASSANDRA_CLUSTER"))
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: os.Getenv("CASSANDRA_USERNAME"),
		Password: os.Getenv("CASSANDRA_PASSWORD"),
	}

	if port, err := strconv.ParseInt(os.Getenv("CASSANDRA_PORT"), 10, 64); err != nil {
		cluster.Port = 9042
		log.Printf("cannot convert CASSANDRA_PORT at .env: %v", err)
	} else {
		cluster.Port = int(port)
	}

	cluster.Keyspace = os.Getenv("CASSANDRA_KEY_SPACE")
	cluster.Consistency = gocql.Quorum
}

func GetSession() (*gocql.Session, error) {
	return cluster.CreateSession()
}
