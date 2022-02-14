package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"params_demo/api/controllers"
	"params_demo/api/dtos"

	"github.com/gin-gonic/gin"
)

func main() {

	setupRouter()

	// 驗證username
	validate(dtos.UsernameInput{Username: "ahsing"}, "/api/username")

	// 驗證age
	// validate(dtos.AgeInput{Age: 1}, "/api/age")

	// 驗證age
	validate(dtos.GenderInput{Gender: "female"}, "/api/gender")
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.POST("/api/username", controllers.ValidateUsername)
	router.POST("/api/age", controllers.ValidateAge)
	router.POST("/api/gender", controllers.ValidateGender)

	return router
}

func validate(params interface{}, path string) {
	router := setupRouter()
	b, _ := json.Marshal(params)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", path, bytes.NewReader(b))
	if err != nil {
		return
	}

	router.ServeHTTP(w, req)

	log.Printf("路徑 %s", path)
	log.Printf("參數 %s", string(b))
	log.Printf("回傳 %s", w.Body)
	log.Println("----- ----- -----")
}
