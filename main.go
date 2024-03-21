package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
)

type value struct {
	Value string `json:"value"`
	Goal  string `json:"goal"`
	Unit  string `json:"unit"`
}
type meters struct {
	Gas      value `json:"gas"`
	Electric value `json:"electric"`
	Water    value `json:"water"`
}

func randValue(min, max float64, precision int) string {
	var rand2 = min + rand.Float64()*(max-min)
	return strconv.FormatFloat(rand2, 'f', precision, 32)
}

func getMetersValues(c *gin.Context) {
	var meters = []meters{
		{
			Gas: value{
				Value: randValue(100, 1000, 2),
				Goal:  randValue(400, 1000, 2),
				Unit:  "m3",
			},
			Electric: value{
				Value: randValue(100, 600, 2),
				Goal:  randValue(200, 600, 2),
				Unit:  "kWh",
			},
			Water: value{
				Value: randValue(1, 30, 3),
				Goal:  randValue(1, 30, 3),
				Unit:  "m3",
			},
		},
	}
	c.IndentedJSON(http.StatusOK, meters)
}

func main() {

	router := gin.Default()
	router.GET("/api/meters/", getMetersValues)

	router.Run("localhost:8080")
}
