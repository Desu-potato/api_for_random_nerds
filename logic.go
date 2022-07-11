package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func meanEndpoint(Context *gin.Context) {
	api_key := initial_data()
	requests := Context.Query("requests")
	length := Context.Query("length")
	requestsInt, err := strconv.Atoi(requests)
	if err != nil {
		log.Print("error: couldn't convert request value to int")
		Context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "couldn't convert request value to int"})
		return
	}
	lengthInt, err := strconv.Atoi(length)
	if err != nil {
		log.Print("error: couldn't convert lenght value to int")
		Context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "couldn't convert lenght value to int"})
		return
	}

	ch := make(chan *http.Response)
	for i := 0; i < requestsInt; i++ {
		res := fmt.Sprintf(`{
			"jsonrpc": "2.0",
			"method": "generateIntegers",
			"params": {
				"apiKey": "%s",
				"n": %v,
				"min": 1,
				"max": 9,
			"replacement": false
			},
			"id": %v
		}`, api_key, lengthInt, i)
		readRequest := strings.NewReader(res)
		go makeRequest(readRequest, ch)
	}
	dataSheet := make(map[string][]int, lengthInt)
	for i := 0; i < requestsInt; i++ {
		rsp := <-ch
		bodyBytes, err := io.ReadAll(rsp.Body)
		id, data, err := parseJSON(bodyBytes)
		if err != nil {
			log.Printf("error: parser cannot convert bytes from response, %v", err)
			Context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "parser cannot convert bytes from response"})
			return
		}

		dataSheet[fmt.Sprint(id)] = data
		defer rsp.Body.Close()
	}
	fmt.Print(dataSheet)
	Context.PureJSON(http.StatusOK, gin.H{})
	return
}
