package main

import (
	// "bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Engine struct{}
type StoreResponse struct {
	Document         map[string]string         `json:"document"`
}
func (e *Engine) HandleCheckVideo(ctx *gin.Context) {
	url := ctx.DefaultQuery("customId", "")
	if len(url) > 0 {
		e.HandleSendReqToDataService(url, ctx)
	}
}

func (e *Engine) HandleSendReqToDataService(url string, ctx *gin.Context) {
	fmt.Println(url)   
	fmt.Println("GOT HERE")   

	postUrl := fmt.Sprintf("http://store_service:8080/find?customId='bbbbbbbbbbbbb.com'")
	resp, err := http.Get(postUrl)
if err != nil {
   fmt.Println(err.Error())
	 fmt.Println("FAILED TO GET")
	 ctx.JSON(http.StatusOK, gin.H{"res": err.Error()})
	 return
}

defer resp.Body.Close()
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
			fmt.Println("FAILED TO READ BODY")
      fmt.Println(err.Error())
			ctx.JSON(http.StatusOK, gin.H{"res": err.Error()})
			return
   }
	response := StoreResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("FAILED TO UNMARSHAL")
		fmt.Println(err.Error())
		ctx.JSON(http.StatusOK, gin.H{"res": err.Error()})

		return
	}
	fmt.Println(response)
	ctx.JSON(http.StatusOK, gin.H{"res": response})
  //  sb := string(body)
}