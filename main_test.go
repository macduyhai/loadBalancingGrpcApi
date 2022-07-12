package main_bk

import (
	_ "encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	_ "io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func RunServer(id string) {
	//
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		time.Sleep(10 * time.Millisecond)
		log.Println("message request")
		ctx.JSON(http.StatusOK, gin.H{
			"MyID": id,
		})
	})

	router.Run()

}

func main() {
	var IdServer = os.Getenv("ID")
	log.Printf("ID Server:%s", IdServer)
	RunServer(IdServer)
}
