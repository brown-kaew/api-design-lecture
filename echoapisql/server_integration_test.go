//go:build integration

package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllUser(t *testing.T) {
	seedUser(t)
	var us []User

	res := request(http.MethodGet, uri("users"), nil)
	err := res.Decode(&us)

	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, res.StatusCode)
	assert.Greater(t, len(us), 0)
}

func TestCreateUser(t *testing.T) {
	body := bytes.NewBufferString(`{
		"name": "Rose",
		"age": 19
	}`)
	var u User

	res := request(http.MethodPost, uri("users"), body)
	err := res.Decode(&u)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.NotEqual(t, 0, u.ID)
	assert.Equal(t, "Rose", u.Name)
	assert.Equal(t, 19, u.Age)
}

func TestGetUserByID(t *testing.T) {
	c := seedUser(t)

	var latest User
	var id string = strconv.Itoa(c.ID)
	res := request(http.MethodGet, uri("users", id), nil)
	err := res.Decode(&latest)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, c.ID, latest.ID)
	assert.NotEmpty(t, latest.Name)
	assert.NotEmpty(t, latest.Age)
}

func TestUpdateUserByID(t *testing.T) {
	c := seedUser(t)
	body := bytes.NewBufferString(`{
		"name": "Nancy",
		"age": 22
	}`)

	var latest User
	var id string = strconv.Itoa(c.ID)
	res := request(http.MethodPut, uri("users", id), body)
	err := res.Decode(&latest)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, c.ID, latest.ID)
	assert.Equal(t, "Nancy", latest.Name)
	assert.Equal(t, 22, latest.Age)
}

func TestDeleteUserByID(t *testing.T) {
	c := seedUser(t)

	var id string = strconv.Itoa(c.ID)
	res := request(http.MethodDelete, uri("users", id), nil)

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func seedUser(t *testing.T) User {
	var c User
	body := bytes.NewBufferString(`{
		"name": "Ewka",
		"age": 19
	}`)
	err := request(http.MethodPost, uri("users"), body).Decode(&c)
	if err != nil {
		t.Fatal("can't create uomer:", err)
	}
	return c
}

func uri(paths ...string) string {
	host := "http://localhost:8080/api"
	if paths == nil {
		return host
	}

	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}

func request(method, url string, body io.Reader) *Response {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Authorization", os.Getenv("AUTH_TOKEN"))
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	return &Response{res, err}
}

type Response struct {
	*http.Response
	err error
}

func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}
	return json.NewDecoder(r.Body).Decode(v)
}
