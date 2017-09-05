package main

// import "fmt"
import "github.com/gin-gonic/gin"

// Data:
// - time
// - value percentile
// - numbers used
type Result struct {
	Time int `json:"time" binding:"exists"`
	Value int `json:"value" binding:"exists"`
	Used int `json:"used" binding:"exists"`
}

type Results struct {
	Times []int `json:"times" binding:"required"`
	Values []int `json:"values" binding:"required"`
	Used []int `json:"used" binding:"required"`
}

var results []Result

func dataKeeper(data chan Result, clear chan bool) {
	for {
		select {
		case result := <-data:
			results = append(results, result)
		case <-clear:
			results = nil
		}
	}
}

func gatherData() Results {
	data := Results{}
	for _, v := range results {
		data.Times = append(data.Times, v.Time)
		data.Values = append(data.Values, v.Value)
		data.Used = append(data.Used, v.Used)
	}
	return data
}

func main() {
	save := make(chan Result)
	clear := make(chan bool)
	go dataKeeper(save, clear)

	r := gin.Default()

	r.Static("/game", "./web/game")
	r.Static("/results", "./web/results")
	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/game")
	})

	r.GET("/data", func(c *gin.Context) {
		results := gatherData()
		c.JSON(200, results)
	})

	r.POST("/result", func(c *gin.Context) {
		var result Result
		if c.BindJSON(&result) == nil {
			save <- result
			c.JSON(200, result)
		}
	})

	r.POST("/clear", func(c *gin.Context) {
		p := c.PostForm("p")
		if p == "4ba776a95704205f428291eb3e0fccd6" {
			clear <- true
			c.JSON(200, gin.H{"status": "ok",})
		}
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
