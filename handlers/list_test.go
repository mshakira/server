package handlers_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"server/handlers"
	"server/handlers/test/data"
	"server/store/snow"
	"testing"
)

func TestHandler_ListHandler(t *testing.T) {

	listHandler := &handlers.Handler{}

	// failure case
	content := `{
  "Name": "ServiceNowQuery",
  "Report": [
    {`

	// initialize http list handler
	listHandler.SnowStore, _ = snow.Init(content)

	req, err := http.NewRequest("GET", "/api/v1/list/incidents", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(listHandler.ListHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status == http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want non 200",
			status)
	}

	listHandler.SnowStore, _ = snow.Init("")
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func BenchmarkHandler_ListHandler(b *testing.B) {
	content := map[int]string{
		6: "", 50: data.Json50, 100: data.Json100, 200: data.Json200,
	}

	req, _ := http.NewRequest("GET", "/api/v1/list/incidents", nil)
	listHandler := &handlers.Handler{}
	for key, js := range content {
		jsonBytes, _ := json.Marshal(&js)
		listHandler.SnowStore, _ = snow.Init(string(jsonBytes))
		b.Run(fmt.Sprintf("%d_incs", key), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				rr := httptest.NewRecorder()
				listHandler.ListHandler(rr, req)
			}
		})
	}
}

func TestRequestLogger(t *testing.T) {
	listHandler := &handlers.Handler{}

	// initialize http list handler
	listHandler.SnowStore, _ = snow.Init("")

	ts := httptest.NewServer(listHandler.RequestLogger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})))
	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	expected := "Hello World"

	if equal(expected, body) {
		t.Errorf("expected %s, got %s", expected, body)
	}
}

func equal(s string, b []byte) bool {
	if len(s) != len(b) {
		return false
	}
	for i, x := range b {
		if x != s[i] {
			return false
		}
	}
	return true
}
