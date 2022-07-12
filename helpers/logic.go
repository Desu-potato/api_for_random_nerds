package helpers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type columnResponse struct {
	Stdr float64 `json:"stddev"`
	Data []int   `json:"data"`
}

func MeanEndpoint(Context *gin.Context) {

	apiKey := Initial_data()
	requests := Context.Query("requests")
	length := Context.Query("length")

	requestsInt, err := strconv.Atoi(requests)
	if err != nil {
		log.Print("error: couldn't convert request value to int")
		Context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "couldn't convert request value to int"})
		return
	}
	lengthInt, err := strconv.Atoi(length)
	if err != nil {
		log.Print("error: couldn't convert lenght value to int")
		Context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "couldn't convert lenght value to int"})
		return
	}

	ch := make(chan *http.Response)
	for i := 0; i < requestsInt; i++ {
		res := fmt.Sprintf(`
		{
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
		}`,
			apiKey, lengthInt, i)
		readRequest := strings.NewReader(res)
		go makeRequest(readRequest, ch)
	}
	dataSheet := make([][]int, lengthInt)
	for i := 0; i < requestsInt; i++ {
		rsp := <-ch
		defer rsp.Body.Close()
		bodyBytes, err := io.ReadAll(rsp.Body)
		if err != nil {
			log.Printf("error: parser cannot convert bytes from https response, %v", err)
			Context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "parser cannot convert bytes https from response"})
			return
		}
		data, err := parseJSON(bodyBytes)
		if err != nil {
			log.Printf("error: parser cannot convert bytes from response, %v", err)
			Context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "parser cannot convert bytes from response"})
			return
		}

		dataSheet[i] = data

	}
	var outputArray []columnResponse
	var buffor []int
	for _, data := range dataSheet {
		buffor = append(buffor, data...)
		outputArray = append(outputArray, columnResponse{
			Data: data,
			Stdr: Stddev(data),
		})
	}
	outputArray = append(outputArray, columnResponse{
		Data: buffor,
		Stdr: Stddev(buffor)})

	Context.PureJSON(http.StatusOK, outputArray)
}
