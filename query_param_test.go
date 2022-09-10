package golang_web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writter http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writter, "Hello")
	} else {
		fmt.Fprintf(writter, "Hello %s", name)
	}
}

func TestQueryParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=arif", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func MultipleParamQuery(writter http.ResponseWriter, request *http.Request) {
	first_name := request.URL.Query().Get("first_name")
	last_name := request.URL.Query().Get("last_name")

	fmt.Fprintf(writter, "Hello %s %s", first_name, last_name)
}

func TestQueryMultipleParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=arif&last_name=hidayat", nil)
	recorder := httptest.NewRecorder()

	MultipleParamQuery(recorder, request)

	response := recorder.Result()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func MultipleValueQuery(writter http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]
	fmt.Fprintf(writter, strings.Join(names, " "))
}

func TestQueryMultipleValue(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=arif&name=hidayat", nil)
	recorder := httptest.NewRecorder()

	MultipleValueQuery(recorder, request)

	response := recorder.Result()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
