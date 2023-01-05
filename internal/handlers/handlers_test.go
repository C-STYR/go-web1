package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type PostData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []PostData
	expectedStatusCode int
}{
	{"home", "/", "GET", []PostData{}, http.StatusOK},
	{"about", "/about", "GET", []PostData{}, http.StatusOK},
	{"generals", "/generals-quarters", "GET", []PostData{}, http.StatusOK},
	{"majors", "/majors-suite", "GET", []PostData{}, http.StatusOK},
	{"sa", "/search-availability", "GET", []PostData{}, http.StatusOK},
	{"contact", "/contact", "GET", []PostData{}, http.StatusOK},
	{"make-res", "/make-reservation", "GET", []PostData{}, http.StatusOK},
	// {"res-sum", "/reservation-summary", "GET", []PostData{}, http.StatusOK}, //returns 200 AFTER the homepage redirect, it seems
	{"post-search-availability", "/search-availability", "POST", []PostData{
		{key: "start", value: "2020-10-01"},
		{key: "end", value: "2020-10-02"},
	}, http.StatusOK},
	{"post-search-availability-json", "/search-availability-json", "POST", []PostData{
		{key: "start", value: "2020-10-01"},
		{key: "end", value: "2020-10-02"},
	}, http.StatusOK},
	{"post-make-reservation", "/make-reservation", "POST", []PostData{
		{key: "first_name", value: "john"},
		{key: "last_name", value: "doe"},
		{key: "email", value: "doe@doe.com"},
		{key: "phone", value: "999"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()

	/*
		we need to create a server and a client for testing purposes;
		httptest.NewTLSServer is comes out of the box with everything:
	*/

	ts := httptest.NewTLSServer(routes) // test server
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			// test client is a method on the test server
			res, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if res.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected code %d but got %d", e.name, e.expectedStatusCode, res.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, param := range e.params {
				values.Add(param.key, param.value)
			}
			res, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if res.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected code %d but got %d", e.name, e.expectedStatusCode, res.StatusCode)
			}
		}
	}
}
