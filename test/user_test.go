package test

import (
	"bytes"
	"fmt"
	"gopkg.in/go-playground/assert.v1"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func TestUserAll(t *testing.T) {
	//gin.SetMode(gin.TestMode)
	fmt.Println("开始测试 Router 中的相关方法")
	//t.Run("测试获取Router所有用户", testIndexGetRouter)
	t.Run("测试用户登录", testUserLogin)

}
func testUserSave(t *testing.T) {
	username := "lisi"
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user/"+username, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户"+username+"已经保存", w.Body.String())
}

func testUserSaveQuery(t *testing.T) {
	username := "lisi"
	age := 18
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user?name="+username+"&age="+strconv.Itoa(age), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户:"+username+",年龄:"+strconv.Itoa(age)+"已经保存", w.Body.String())
}

func testUserSaveWithNotAge(t *testing.T) {
	username := "lisi"
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user?name="+username, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户:"+username+",年龄:20已经保存", w.Body.String())
}

func testUserPostForm(t *testing.T) {
	value := url.Values{}
	value.Add("email", "youngxhui@gmail.com")
	value.Add("password", "1234")
	value.Add("password-again", "1234")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/user/register", bytes.NewBufferString(value.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusMovedPermanently, w.Code)
}

func testUserPostFormEmailErrorAndPasswordError(t *testing.T) {
	value := url.Values{}
	value.Add("email", "youngxhui")
	value.Add("password", "1234")
	value.Add("password-again", "qwer")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/user/register", bytes.NewBufferString(value.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func testUserLogin(t *testing.T) {
	email := "admin@localhost.com"
	value := url.Values{}
	value.Add("email", email)
	value.Add("password", "1234")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/user/login", bytes.NewBufferString(value.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, strings.Contains(w.Body.String(), email), true)
}
