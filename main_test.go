package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

var (
	svr      *httptest.Server
	username string = "username"
)

func TestMain(m *testing.M) {
	expected := User{}
	svr = httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		json.NewEncoder(rw).Encode(expected)
	}))
	defer svr.Close()
	os.Exit(m.Run())
}

func TestGetGithubUser(t *testing.T) {
	if svr == nil {
		t.Errorf("svr is nil")
		return
	}
	user, err := getGithubUser(svr.URL, username)
	if err != nil {
		t.Errorf("GET github user error : %v\n", err)
	}
	expected := reflect.TypeOf(User{})
	got := reflect.TypeOf(user)
	if got != expected {
		t.Errorf("GET github user error : types dont match %s != %s\n", expected, got)
	}
}

func TestGetGithubName(t *testing.T) {
	if svr == nil {
		t.Errorf("svr is nil")
		return
	}
	name, err := getGithubName(svr.URL, username)
	if err != nil {
		t.Errorf("GET github user error : %v\n", err)
	}
	expected := reflect.TypeOf("")
	got := reflect.TypeOf(name)
	if got != expected {
		t.Errorf("GET github user error : types dont match %s != %s", expected, got)
	}
}
