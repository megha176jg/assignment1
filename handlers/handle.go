package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"video_api/dummy"
	"video_api/models"

	"github.com/gin-gonic/gin"
)

func ViewCountById(c *gin.Context) {
	id := c.Param("id")
	vid, err := GetVideoById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Video not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, vid)
}

func ViewVideo(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	vid, err := GetVideoById(id)

	fmt.Printf("recieving the address of required video %p\n", vid)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Video not found."})
		return
	}

	vid.View += 1
	// fmt.Println("incrementing the count to %p", (*vid).View)
	c.IndentedJSON(http.StatusOK, *vid)
	// fmt.Println("count is %p", (*vid).View)

}

func GetVideoById(id string) (*models.Video, error) {
	videos := dummy.GetVideos()
	fmt.Println(videos)
	videoVal := videos
	for i := range videos {
		if videoVal[i].ID == id {
			fmt.Printf("sending the address of required video %p\n", &videoVal[i])
			return &videoVal[i], nil
		}
	}
	return models.NewModel(), errors.New("video not found")
}
