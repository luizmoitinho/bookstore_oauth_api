package cassandra

import (
	"log"
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
	session := GetSession()
	if session == nil {
		t.Errorf("session not be a nil")
	}
}
