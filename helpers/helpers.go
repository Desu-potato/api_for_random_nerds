package helpers

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
)

type response struct {
	Jsonrpc string          `json:"jsonrpc"`
	Result  response_result `json:"result"`
	Id      int             `json:"id"`
}

type response_result struct {
	Random response_result_random `json:"random"`
}

type response_result_random struct {
	Data           []int  `json:"data"`
	CompletionTime string `json:"completionTime"`
}

func parseJSON(bodyBytes []byte) ([]int, error) {
	var res response
	err := json.Unmarshal(bodyBytes, &res)
	if err != nil {
		return nil, err
	}
	return res.Result.Random.Data, nil
}

func Initial_data() string {
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

func Stddev(arr []int) float64 {

	lenght := float64(len(arr))
	var floatBufforArray []float64
	var buffor int
	for _, record := range arr {

		buffor = buffor + record

	}
	mean := float64(buffor) / lenght
	for _, record := range arr {

		result := math.Pow(float64(record)-mean, 2)
		floatBufforArray = append(floatBufforArray, result)

	}
	var floatBuffor float64
	for _, record := range floatBufforArray {
		floatBuffor = floatBuffor + record
	}

	floatBuffor = floatBuffor / lenght
	return math.Sqrt(floatBuffor)
}
