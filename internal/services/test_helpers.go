package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kwakubiney/canonical-take-home/internal/domain/model"
	"github.com/kwakubiney/canonical-take-home/internal/postgres"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var dbConnPool *gorm.DB

func NewUUID() string {
	uid, _ := uuid.NewV4()
	return uid.String()
}

func MakeRequest(route string, port string, requestBody interface{}, method string) (*http.Response, error) {
	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, route, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func SeedDB(r ...interface{}) *gorm.DB {
	if dbConnPool == nil {
		db, err := postgres.Init()
		if err != nil {
			log.Fatal(err)
		}
		dbConnPool = db
	}
	tx := dbConnPool.Begin()
	for _, m := range r {
		if err := tx.Create(m).Error; err != nil {
			tx.Rollback()
			log.Fatalf("[data insert failed] %v ", err)
		}
	}
	tx.Commit()
	return dbConnPool
}

func CreateTestUser(t *testing.T) *model.User {
	f := faker.New()

	testUser := model.User{
		Username: NewUUID(),
		Email:    f.Person().Contact().Email,
		Age:      fmt.Sprint(f.RandomNumber(2)),
	}

	return &testUser
}


func CreateTestGame(t *testing.T) *model.Game {
	f := faker.New()

	testGame := model.Game{
		AgeRating:  "12+",
		URL:         f.Internet().URL(),
		Description: f.Lorem().Sentence(200),
		Publisher:   f.Company().Name(),
		Title:       f.App().Name(),
	}

	return &testGame
}

func BootstrapServer(req *http.Request, routeHandlers *gin.Engine) *httptest.ResponseRecorder {
	responseRecorder := httptest.NewRecorder()
	routeHandlers.ServeHTTP(responseRecorder, req)
	return responseRecorder
}

func MakeTestRequest(t *testing.T, route string, requestBody interface{}, method string) *http.Request {
	body, err := json.Marshal(requestBody)
	assert.NoError(t, err)
	req, err := http.NewRequest(method, route, bytes.NewReader(body))

	assert.NoError(t, err)

	return req
}

func DecodeResponse(t *testing.T, response *httptest.ResponseRecorder) map[string]interface{} {
	var responseBody map[string]interface{}
	log.Println(response.Result().Header)
	assert.NoError(t, json.Unmarshal(response.Body.Bytes(), &responseBody))
	return responseBody
}
