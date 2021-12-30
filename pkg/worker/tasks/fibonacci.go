package tasks

import (
	"encoding/json"
	"fmt"

	"github.com/RichardKnop/machinery/v1/log"
	"github.com/valyala/fastjson"
)

func Fibonacci(jobID string, input string) (string, error) {
	log.INFO.Print("fibonacci task started")

	var jsonParser fastjson.Parser
	jsonFields, err := jsonParser.Parse(input)
	if err != nil {
		return "", err
	}
	if !jsonFields.Exists("target") {
		message := "fibonacci task failed, target not exist in input"
		return "", fmt.Errorf(message)
	}
	target := jsonFields.GetInt("target")

	result := FibonacciRecursion(target)
	resultJson, _ := json.Marshal(result)

	log.INFO.Print("fibonacci task finished")

	return string(resultJson), nil
}

func FibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}
