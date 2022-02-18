package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouterPattern(t *testing.T) {
	router := httprouter.New()

	router.GET("/product/:id/item/:itemId", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		text := "Product " + params.ByName("id") + " Item " + params.ByName("itemId")
		_, err := fmt.Fprint(writer, text)
		if err != nil {
			panic(err)
		}
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/product/1/item/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1 Item 1", string(body))
}
