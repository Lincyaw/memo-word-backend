package service

import (
	"fmt"
	"github.com/mini-ecs/back-end/internal/dao/pool"
	"github.com/mini-ecs/back-end/internal/model"
	"io/ioutil"
	"net/http"
	"time"
)

func QueryWord(word string) string {
	db := pool.GetDB()
	w := model.Word{}
	res := db.First(&w, "label = ?", word)
	if res.Error != nil {
		req, err := http.NewRequest("GET",
			"https://od-api.oxforddictionaries.com/api/v2/entries/en-gb/"+word+"?strictMatch=false",
			nil)
		if err != nil {
			fmt.Println(err)
		}
		// 比如说设置个token
		req.Header.Set("app_id", "232bb462")
		req.Header.Set("app_key", "3be4b5a8c79e2ca68c89c62b77713569")
		resp, err := (&http.Client{}).Do(req)
		if err != nil {
			return ""
		}
		defer resp.Body.Close()
		respByte, _ := ioutil.ReadAll(resp.Body)
		db.Save(&model.Word{
			Label: word,
			Data:  string(respByte),
		})
		return string(respByte)
	}
	return w.Data
}
func RememberWord(word string) {
	db := pool.GetDB()
  fmt.Println("remember ", word)

	t := time.Now()
	d, _ := time.ParseDuration("72h")
	newTime := t.Add(d)

	res := db.Model(&model.Word{}).Where("label = ?", word).Update("updated_at", newTime)
  if res.Error != nil {
    fmt.Println(res.Error)
  }
}

func ForgetWord(word string)  {
  db := pool.GetDB()
  fmt.Println("remember ", word)

  t := time.Now()
  d, _ := time.ParseDuration("-24h")
  newTime := t.Add(d)

  res := db.Model(&model.Word{}).Where("label = ?", word).Update("updated_at", newTime)
  if res.Error != nil {
    fmt.Println(res.Error)
  }
}
func GetWordByTime(num int) []model.Word {
	db := pool.GetDB()
	var words []model.Word
	res := db.Order("updated_at").Select("label", "updated_at").Limit(num).Find(&words)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
	return words
}
