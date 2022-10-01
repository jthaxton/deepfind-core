package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	// "io"
	"log"
	"net/http"

	// "github.com/gin-gonic/gin"
)

type DocBody struct {
	CustomId *string                   `json:"custom_id"`
}
type Body struct {
	Document DocBody            `json:"document"`
}

type AddToStoreBody struct {
	CustomId 	string             `json:"custom_id"`
	Meta      interface{}        `json:"meta"`
}

type AddToJobsBody struct {
	CustomId       string        `json:"custom_id"`
	Kind           string        `json:"kind"`
	DataJson       map[string]interface{}   `json:"data_json"`
}

type Unmarshalled struct {
	CustomID      *string                   `json:"customId"`
}

type NullableBody struct {
	Document      interface{}     `json:"document"`
}

func FetchFromStore(url string) ([]byte, error) {
	route := fmt.Sprintf("http://store_service:8080/find?customId=%s", url)
	resp, err := http.Get(route)
	// log.Println(resp.Body)
	if err != nil {
		return nil, err
	}
	var nb NullableBody
	json.NewDecoder(resp.Body).Decode(&nb)
	log.Println(nb)
	log.Println(&nb)
	m, err := json.Marshal(nb)
	if err != nil {
		return nil, err
	}
	if nb.Document == nil {
		return nil, fmt.Errorf("document not found")
	}
	log.Println(m)
	log.Println(&m)

	return m, nil
}

func AddToStore(url string, b interface{}) ([]byte, error) {
	send := &AddToStoreBody{
		CustomId: url,
		Meta: b,
	}

	body, err := json.Marshal(send)
	if err != nil {
		return nil, err
	}
	route := fmt.Sprintf("http://store_service:8080/create?customId=%s", url)
	resp, err := http.Post(route, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	res, err := json.Marshal(resp.Body)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func AddToJobs(url string, ctx []byte) (interface{}, error) {
	var dj map[string]interface{}
	err := json.Unmarshal(ctx, &dj)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	send := AddToJobsBody{
		Kind: "youtube_processing",
		CustomId: url,
		DataJson: dj,
	}
	sendbody, err := json.Marshal(send)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	resp, err := http.Post("http://atlas:8080/create", "application/json", bytes.NewBuffer(sendbody))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}
	var target interface{}

	err = json.NewDecoder(resp.Body).Decode(&target)
	if err != nil {
		return nil, err
	}

	return target, nil
}