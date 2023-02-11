package cassandra

import (
	"os"

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

	cluster.Keyspace = os.Getenv("CASSANDRA_KEY_SPACE")
	cluster.Consistency = gocql.Quorum
}

func GetSession() (*gocql.Session, error) {
	InitCluster()
	return cluster.CreateSession()
}
