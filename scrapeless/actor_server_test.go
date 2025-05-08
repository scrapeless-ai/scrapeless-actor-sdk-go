package scrapeless

import (
	"fmt"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/httpserver"
	log "github.com/sirupsen/logrus"
	"net/http"
	"testing"
)

type Input struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var (
	actor *Actor
)

func TestWithServer(t *testing.T) {
	actor = New(WithServer(httpserver.DebugMode))
	input := &Input{}
	actor.Server.AddHandle(http.MethodGet, "", input, getData)
	if err := actor.Start(); err != nil {
		log.Fatal(err)
	}
}

func getData(inputStruct any) (any, error) {
	fmt.Println(inputStruct)
	// proxy
	// browser
	// get data
	return "test", nil
}
