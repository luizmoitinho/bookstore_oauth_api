package cassandra

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetSession(t *testing.T) {
	if err := godotenv.Load("../../../.env"); err != nil {
		log.Fatal(err)
	}
	_, err := GetSession()
	if err != nil {
		t.Fatal(fmt.Sprintf("{error:%v, \n env_vars:%v,%v,%v,%v}"), err.Error(), os.Getenv("CASSANDRA_CLUSTER"),
			os.Getenv("CASSANDRA_USERNAME"),
			os.Getenv("CASSANDRA_CLUSTER"),
			os.Getenv("CASSANDRA_KEY_SPACE"))
	}

}
