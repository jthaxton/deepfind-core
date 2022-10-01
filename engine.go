package main

import (
	// "bytes"
	"bytes"
	"encoding/json"
	"log"

	// "log"

	// "io/ioutil"
	// "net/http"

	"github.com/gin-gonic/gin"
	"github.com/jthaxton/deepfind-core/services"
)

type Handler struct{}
type StoreResponse struct {
	Document         map[string]string         `json:"document"`
}

type BodyId struct {
	CustomId string        `json:"customId`
}

type DocBody struct {
	CustomId *string                   `json:"customId"`
}
type Body struct {
	Document *DocBody            `json:"document"`
}
func (e *Handler) HandleCheckVideo(ctx *gin.Context) {
	url := ctx.DefaultQuery("youtubeUrl", "")
	if len(url) == 0 {
		ctx.JSON(403,map[string]string{"error":"youtubeUrl not found"})
		return
	}

	// var m Body
	existsInStore, err := services.FetchFromStore(url)

	if err == nil {
		ctx.JSON(200,gin.H{"document": string(existsInStore)})
		return
	} else {
		log.Println(err.Error())
	}
		
	// err := json.Unmarshal(existsInStore, &m)
	// if err != nil {
	// 	ctx.JSON(403,map[string]string{"error": err.Error()})
	// 	return
	// }

	// log.Println(string(existsInStore))
	// log.Println(m)
	// if existsInStore != nil {
	// 	ctx.JSON(200,gin.H{"document": string(existsInStore)})
	// 	return
	// }
	// if err == nil {
	// 	if err != nil {
	// 		ctx.JSON(403,map[string]string{"error": err.Error()})
	// 		return
	// 		} else {
	// 		ctx.JSON(200,map[string]*bytes.Buffer{"document": bytes.NewBuffer(existsInStoreBody)})
	// 		return
	// 	}
	// }
		body, err := json.Marshal(ctx.Request.Body)
		if err != nil {
			ctx.JSON(403,map[string]string{"error":err.Error()})
			return
		}

		_, err = services.AddToJobs(url, body)
		if err != nil {
			ctx.JSON(403,map[string]string{"error": err.Error()})
			return
		}

		addedToStore, err := services.AddToStore(url, body)
		if err != nil {
			ctx.JSON(403,map[string]string{"error": err.Error()})
			return
		}

		existsInStoreBody, err := json.Marshal(addedToStore)
		if err != nil {
			ctx.JSON(403,map[string]string{"error": err.Error()})
			return
		}
	
		existsInStore, err = services.FetchFromStore(url)

		if err == nil {
			ctx.JSON(200,gin.H{"document": string(existsInStore)})
			return
		} else {
			log.Println(err.Error())
		}

	ctx.JSON(200,map[string]*bytes.Buffer{"document": bytes.NewBuffer(existsInStoreBody)})
}
