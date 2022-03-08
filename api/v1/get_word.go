package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mini-ecs/back-end/internal/service"
	"github.com/mini-ecs/back-end/pkg/log"
)

var logger = log.GetGlobalLogger()

func GetWord(c *gin.Context) {
	word := c.Param("word")
	data := service.QueryWord(word)
	c.JSON(200, data)
}

func Remember(c *gin.Context) {
  word := c.Param("word")
  service.RememberWord(word)
  c.JSON(200, "")
}


func Forget(c *gin.Context) {
  word := c.Param("word")
  service.ForgetWord(word)
  c.JSON(200, "")
}


func GetWordByTime(c *gin.Context){
  data := service.GetWordByTime(50)
  c.JSON(200, data)
}
