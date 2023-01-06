package main

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"math/rand"
	"shenaishiren/model"
	"shenaishiren/server"
	"shenaishiren/utils"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())

	utils.InitDb()
	defer func() {
		utils.Db.Close()
	}()

	utils.InitCache()

	initServer := server.InitServer()

	initServer.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"data": nil,
			"msg":  c.FullPath() + " 路由未找到!",
		})
	})
	//找不到的方法的情况
	initServer.NoMethod(func(c *gin.Context) {
		c.JSON(500, gin.H{
			"data": nil,
			"msg":  c.HandlerName() + " 方法未找到!",
		})
	})

	//接口
	initServer.GET("/sasr", func(c *gin.Context) {
		x, found := utils.GlobalCache.Get("sasr")
		if !found {

			bibleAct := make([]*model.Bible, 5)

			var count int64
			utils.Db.Model(bibleAct).Count(&count)

			rand := rand.Int63n(count - 1)

			err := utils.Db.Where("ID >= ?", rand).Limit(5).Find(&bibleAct).Error
			if err != nil {
				c.JSON(500, gin.H{
					"data": nil,
					"msg":  "查询失败",
				})
				return
			}

			calendar := utils.GetCalendar()
			now, _ := calendar.GetTimeNow()

			m := make(map[string]interface{})
			m["shengji"] = bibleAct
			m["rili"] = now
			utils.GlobalCache.Set("sasr", &m, cache.DefaultExpiration)
			x = m
		}

		c.JSON(200, gin.H{
			"data": x,
			"msg":  "ok",
		})
		return
	})

	server.Run()
}
