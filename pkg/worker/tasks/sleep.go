package tasks

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/RichardKnop/machinery/v1/log"
	"github.com/valyala/fastjson"
)

type Result struct {
	Nums []int `json:"nums"`
}

func Sleep(jobID string, input string) (string, error) {
	log.INFO.Print("sleep task started")

	var jsonParser fastjson.Parser
	jsonFields, err := jsonParser.Parse(input)
	if err != nil {
		return "", err
	}
	if !jsonFields.Exists("seconds") {
		message := "sleep task failed, seconds not exist in input"
		return "", fmt.Errorf(message)
	}
	seconds := jsonFields.GetInt("seconds")

	nums := []int{}
	for i := 0; i < seconds; i++ {
		log.INFO.Print(i)
		nums = append(nums, i)
		time.Sleep(1 * time.Second)
	}

	result := Result{
		Nums: nums,
	}
	resultJson, _ := json.Marshal(result)

	log.INFO.Print("sleep task finished")
	return string(resultJson), nil
}
