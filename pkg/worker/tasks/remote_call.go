package tasks

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/RichardKnop/machinery/v1/log"
	"github.com/valyala/fastjson"
)

func RemoteHTTPCall(jobID string, input string) (string, error) {
	log.INFO.Print("remoteHTTPCall task started")

	var jsonParser fastjson.Parser
	jsonFields, err := jsonParser.Parse(input)
	if err != nil {
		return "", err
	}
	if !jsonFields.Exists("url") {
		message := "remoteHTTPCall task failed, url not exist in input"
		return "", fmt.Errorf(message)
	}
	if !jsonFields.Exists("body") {
		message := "remoteHTTPCall task failed, body not exist in input"
		return "", fmt.Errorf(message)
	}

	url := jsonFields.GetStringBytes("url")
	postBody := jsonFields.Get("body").MarshalTo(nil)

	// Send POST request
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(string(url), "application/json", responseBody)
	if err != nil {
		log.ERROR.Println("remoteHTTPCall failed")
		return "", err
	}
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.ERROR.Println("remoteHTTPCall failed, read response failed")
		return "", err
	}
	result := string(body)

	log.INFO.Print("remoteHTTPCall task finished")
	return result, nil
}
