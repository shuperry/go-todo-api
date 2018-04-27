package utils

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"fmt"
)

func SendRequest(method string, uri string, param map[string]interface{}, router *gin.Engine) []byte {
	jsonByte,_ := json.Marshal(param)

	req := httptest.NewRequest(method, uri, bytes.NewReader(jsonByte))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result := w.Result()
	defer result.Body.Close()

	body,_ := ioutil.ReadAll(result.Body)

	fmt.Println("send requestbody =", body)

	return body
}
