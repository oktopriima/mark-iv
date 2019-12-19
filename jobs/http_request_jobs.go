package jobs

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/oktopriima/mark-v/configurations"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

type responseBody struct {
	Data struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Birthday  time.Time `json:"birthday"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"data"`
}

type HttpRequestJobContract interface {
	GetHttpRequest()
	PostHttpRequest()
	BulkPostHttpRequest()
}

type httpRequestJobContract struct {
	cfg configurations.Config
	db  *gorm.DB
}

func NewHttpRequestJobs(cfg configurations.Config,
	db *gorm.DB) HttpRequestJobContract {
	return &httpRequestJobContract{cfg, db}
}

func (j *httpRequestJobContract) GetHttpRequest() {
	var baseUrl string
	baseUrl = "http://localhost:9000"

	var client = &http.Client{}

	request, err := http.NewRequest("GET", baseUrl+"/user/1", nil)
	if err != nil {
		panic(err)
	}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var resp responseBody
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}

func (j *httpRequestJobContract) PostHttpRequest() {

	var responseBody struct {
		Message string `json:"message"`
	}

	var param = url.Values{}
	param.Set("name", "Ginanjar")
	param.Set("email", "ginanjar@gmai.com")
	param.Set("birthday", "1992-10-05T00:00:00Z")
	param.Set("password", "secret")

	bodyBuffer := bytes.NewBufferString(param.Encode())

	var baseUrl string
	baseUrl = "http://localhost:9000"

	var client = &http.Client{}

	request, err := http.NewRequest("POST", baseUrl+"/user", bodyBuffer)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", "Bearer TOKEN SECRET")

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		err = json.NewDecoder(response.Body).Decode(&responseBody)
		if err != nil {
			panic(err)
		}

		fmt.Println(responseBody)
	}
	return
}

func (j *httpRequestJobContract) BulkPostHttpRequest() {
	var responseBody struct {
		Message string `json:"message"`
	}

	file, err := os.Open("practice.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	param := url.Values{}
	for _, line := range lines {
		if err == io.EOF {
			break
		}

		if line[0] == "Name" {
			continue
		}

		param.Set("name", line[0])
		param.Set("email", line[1])
		param.Set("birthday", line[3])
		param.Set("password", line[2])

		bodyBuffer := bytes.NewBufferString(param.Encode())

		fmt.Println(param.Encode())
		var baseUrl string
		baseUrl = "http://localhost:9000"

		var client = &http.Client{}

		request, err := http.NewRequest("POST", baseUrl+"/user", bodyBuffer)
		if err != nil {
			panic(err)
		}
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		response, err := client.Do(request)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()

		err = json.NewDecoder(response.Body).Decode(&responseBody)
		if err != nil {
			panic(err)
		}

		fmt.Println(responseBody)
	}

}
