package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

type request struct {
	Jsonrpc string         `json:"jsonrpc"`
	Result  request_result `json:"result"`
	Id      int            `json:"id"`
}

type request_result struct {
	Random request_result_random `json:"random"`
}

type request_result_random struct {
	Data           []int  `json:"data"`
	CompletionTime string `json:"completionTime"`
}

func parseJSON(bodyBytes []byte) (int, []int, error) {
	var res request
	err := json.Unmarshal(bodyBytes, &res)
	if err != nil {
		return 0, nil, err
	}
	return res.Id, res.Result.Random.Data, nil
}

func initial_data() string {
	api_key_random := os.Getenv("API_KEY_RANDOM")
	if api_key_random == "" {
		log.Fatal("error: no api_key env set")
	}
	return api_key_random
}

func makeRequest(request *strings.Reader, ch chan<- *http.Response) {
	resp, err := http.Post("https://api.random.org/json-rpc/2/invoke", "application/json", request)
	if err != nil {
		log.Printf("error in cocurrency: problem with request %v", err)
		return
	}
	ch <- resp
}
