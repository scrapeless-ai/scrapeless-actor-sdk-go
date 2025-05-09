package scrapeless

import (
	"fmt"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/httpserver"
	log "github.com/sirupsen/logrus"
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
	actor.Server.AddHandle("", getData)
	if err := actor.Start(); err != nil {
		log.Fatal(err)
	}
}

func getData(inputStruct []byte) (httpserver.Response, error) {
	fmt.Println(string(inputStruct))
	// proxy
	// browser
	// get data
	return httpserver.Response{
		Code: 200,
		Data: string(inputStruct),
		Msg:  "good",
	}, nil
}
