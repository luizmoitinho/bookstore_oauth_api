package rest_users

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gopkg.in/resty.v1"
)

func TestMain(m *testing.M) {
	log.Println("starting test main to test cases...")
	if err := godotenv.Load("../../../.env"); err != nil {
		log.Fatalf("unxepected error was received: %v", err)
	}
	os.Exit(m.Run())
}

func TestLoginTimeoutFromApi(t *testing.T) {
	restyClient := resty.New()
	httpmock.ActivateNonDefault(restyClient.GetClient())
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "http://bookstore.api.com/user/authenticate",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewJsonResponse(-1, `{}`)
		},
	)

	client := NewClient(restyClient)
	repository := NewRepository(client)
	user, err := repository.Login("email@gmail.com", "123456")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid rest client response when trying to login user", err.Message)
}

func TestLoginInvalidErrorInterface(t *testing.T) {
	restyClient := resty.New()
	httpmock.ActivateNonDefault(restyClient.GetClient())
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "http://bookstore.api.com/users/authenticate",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewJsonResponse(http.StatusNotFound, `{"message": "invalid login credentials", "status": "404", "error": "not_found_error"}`)
		},
	)

	client := NewClient(restyClient)
	repository := NewRepository(client)
	user, err := repository.Login("email@gmail.com", "123456")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when trying to login user", err.Message)
}

func TestLoginInvalidLoginCredentials(t *testing.T) {
	restyClient := resty.New()
	httpmock.ActivateNonDefault(restyClient.GetClient())
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "http://bookstore.api.com/users/authenticate",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewJsonResponse(http.StatusNotFound,
				map[string]any{
					"message": "invalid login credentials",
					"status":  404,
					"error":   "not_found_error",
				},
			)
		},
	)

	client := NewClient(restyClient)
	repository := NewRepository(client)
	user, err := repository.Login("email@gmail.com", "123456")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "invalid login credentials", err.Message)
}

func TestLoginInvalidUserResponse(t *testing.T) {
	restyClient := resty.New()
	httpmock.ActivateNonDefault(restyClient.GetClient())
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "http://bookstore.api.com/users/authenticate",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewJsonResponse(http.StatusOK,
				map[string]any{
					"id":         "25",
					"first_name": "Luiz Carlos",
					"last_name":  "Costa Moitinho",
					"email":      "email@gmail.com",
				},
			)
		},
	)

	client := NewClient(restyClient)
	repository := NewRepository(client)
	user, err := repository.Login("email@gmail.com", "123456")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unsmarshall users login response", err.Message)
}

func TestLoginNoError(t *testing.T) {
	restyClient := resty.New()
	httpmock.ActivateNonDefault(restyClient.GetClient())
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "http://bookstore.api.com/users/authenticate",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewJsonResponse(http.StatusOK,
				map[string]any{
					"id":         1,
					"first_name": "Luiz Carlos",
					"last_name":  "Costa Moitinho",
					"email":      "email@gmail.com",
				},
			)
		},
	)

	client := NewClient(restyClient)
	repository := NewRepository(client)
	user, err := repository.Login("email@gmail.com", "123456")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, user.Id, 1)
	assert.EqualValues(t, user.FirstName, "Luiz Carlos")
	assert.EqualValues(t, user.LastName, "Costa Moitinho")
	assert.EqualValues(t, user.Email, "email@gmail.com")
}
