package cassandra

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load("../../../.env"); err != nil {
		log.Fatalf("unxepected error was received: %v", err)
	}
	InitCluster()
}

func TestGetSession(t *testing.T) {
	_, err := GetSession()
	if err != nil {
		t.Errorf(fmt.Sprintf("{error:%v, \n env_vars:%v,%v,%v,%v}", err.Error(), os.Getenv("CASSANDRA_CLUSTER"),
			os.Getenv("CASSANDRA_USERNAME"),
			os.Getenv("CASSANDRA_PASSWORD"),
			os.Getenv("CASSANDRA_KEY_SPACE")))
	}
}
