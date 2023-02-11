package cassandra

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetSession(t *testing.T) {
	if err := godotenv.Load("../../../.env"); err != nil {
		log.Fatal(err)
	}
	_, err := GetSession()
	if err != nil {
		t.Fatal(err)
	}
}
