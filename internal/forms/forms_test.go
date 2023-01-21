package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// naming convention for receiver funcs as follows:
func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	// set the required fields to a, b, c (with no values associated)
	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}

	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows invalid when valid")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	if form.Has("a", r) {
		t.Error("shows valid when invalid")
	}

	postedData := url.Values{}

	postedData.Add("a", "b")

	form = New(postedData)

	if !form.Has("a", r) {
		t.Error("shows invalid when valid")
	}
}

func Test_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.MinLength("a", 2, r)
	if form.Valid() {
		t.Error("shows valid when field is non-existent")
	}

	postedData := url.Values{}

	postedData.Add("a", "bc")
	
	form = New(postedData)

	form.MinLength("a", 2, r) 
	if !form.Valid(){
		t.Error("shows invalid when valid")
	}

	form.MinLength("a", 3, r) 
	if form.Valid(){
		t.Error("shows valid when invalid")
	}
}

func Test_IsEmail(t *testing.T) {

	postedData := url.Values{}

	form := New(postedData)
	form.IsEmail("x")
	if form.Valid() {
		t.Error("validator shows valid email for non-existent field")
	}

	postedData.Add("good email", "test@example.com")
	postedData.Add("bad email", "test@example")

	form = New(postedData)

	form.IsEmail("good email")
	if !form.Valid() {
		t.Error("field shows invalid when valid")
	}

	form.IsEmail("bad email")
	if form.Valid() {
		t.Error("field shows valid when invalid")
	}
}
