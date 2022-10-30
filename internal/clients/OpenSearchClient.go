package clients

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/the-jay-team/jaytube-search-recommendation-service/pkg/models"
	"io"
	"net/http"
	"strings"
	"time"
)

type OpenSearchClient struct {
	hostUrl    string
	username   string
	password   string
	httpClient *http.Client
}

func NewOpenSearchClient(hostUrl string, index string, username string, password string) *OpenSearchClient {
	if !strings.HasSuffix(hostUrl, "/") {
		hostUrl += "/"
	}
	hostUrl += index
	client := &OpenSearchClient{hostUrl, username, password, createHttpClient()}
	return client
}

func createHttpClient() *http.Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	httpClient := &http.Client{Transport: transport}
	return httpClient
}

func (client OpenSearchClient) SearchVideoDataContaining(query string) ([]models.VideoData, error) {
	queryStruct := models.OpenSearchQueryStruct{
		Size: 10,
		Query: models.OpenSearchMultiMatchQuery{
			MultiMatch: models.OpenSearchMultiMatch{
				Query:  query,
				Fields: []string{"title", "description", "tags", "creator"},
			}}}
	body := bytes.Buffer{}
	encoderError := json.NewEncoder(&body).Encode(queryStruct)
	if encoderError != nil {
		return []models.VideoData{}, encoderError
	}

	request := client.createRequest(http.MethodGet, "/_search", &body)
	requestResponse, requestError := client.httpClient.Do(request)
	if requestError != nil {
		return []models.VideoData{}, requestError
	}

	responseStruct := models.OpenSearchQueryResult{}
	jsonDecodeError := json.NewDecoder(requestResponse.Body).Decode(&responseStruct)
	if jsonDecodeError != nil {
		return []models.VideoData{}, jsonDecodeError
	}

	var videoDataArray []models.VideoData
	for _, element := range responseStruct.Hits.HitsArray {
		openSearchVideoData := element.Source
		videoDataArray = append(videoDataArray, models.VideoData{
			Id:          element.Id,
			Title:       openSearchVideoData.Title,
			Description: openSearchVideoData.Description,
			UploadDate:  openSearchVideoData.UploadDate,
			Tags:        openSearchVideoData.Tags,
			Creator:     openSearchVideoData.Creator,
			Visibility:  openSearchVideoData.Visibility,
		})
	}

	return videoDataArray, nil
}

func (client OpenSearchClient) QueryRandom() ([]models.VideoData, error) {
	queryStruct := models.OpenSearchQueryStruct{
		Size: 10,
		Query: models.OpenSearchFunctionScoreQuery{
			FunctionScore: models.OpenSearchFunctionScore{
				Functions: []any{
					models.OpenSearchRandomScoreFunction{RandomScore: models.OpenSearchRandomScore{
						Seed: time.Now().Unix()}}}}}}

	body := bytes.Buffer{}
	encoderError := json.NewEncoder(&body).Encode(queryStruct)
	if encoderError != nil {
		return []models.VideoData{}, encoderError
	}

	request := client.createRequest(http.MethodGet, "/_search", &body)
	requestResponse, requestError := client.httpClient.Do(request)
	if requestError != nil {
		return []models.VideoData{}, requestError
	}

	responseStruct := models.OpenSearchQueryResult{}
	jsonDecodeError := json.NewDecoder(requestResponse.Body).Decode(&responseStruct)
	if jsonDecodeError != nil {
		return []models.VideoData{}, jsonDecodeError
	}

	var videoDataArray []models.VideoData
	for _, element := range responseStruct.Hits.HitsArray {
		openSearchVideoData := element.Source
		videoDataArray = append(videoDataArray, models.VideoData{
			Id:          element.Id,
			Title:       openSearchVideoData.Title,
			Description: openSearchVideoData.Description,
			UploadDate:  openSearchVideoData.UploadDate,
			Tags:        openSearchVideoData.Tags,
			Creator:     openSearchVideoData.Creator,
			Visibility:  openSearchVideoData.Visibility,
		})
	}

	return videoDataArray, nil
}

func (client OpenSearchClient) createRequest(methode string, requestPath string, body io.Reader) *http.Request {
	if strings.HasPrefix(requestPath, "/") {
		requestPath = requestPath[1:]
	}
	req, _ := http.NewRequest(methode,
		fmt.Sprintf("%s/%s/", client.hostUrl, requestPath), body)
	req.SetBasicAuth(client.username, client.password)
	req.Header.Add("Content-Type", "application/json")
	return req
}
