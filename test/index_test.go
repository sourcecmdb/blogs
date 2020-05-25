package test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	_ "strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sourcecmdb/blogs/ginhello/initRouter"

	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func init() {
	gin.SetMode(gin.TestMode)
	router = initRouter.SetupRouter()
}

// TestMain =函数可以在测试函数执行之前做一些其他的操作
func TestMain(m *testing.M) {
	fmt.Println("测试开始")
	//通过m.Run()来执行测试函数
	m.Run()
}

func TestAll(t *testing.T) {
	//gin.SetMode(gin.TestMode)
	fmt.Println("开始测试 Router 中的相关方法")
	//t.Run("测试获取Router所有用户", testIndexGetRouter)
	//t.Run("测试获取Router所有用户", testUserSave)
	// t.Run("测试获取Router所有用户", testUserSaveQuery)
	//t.Run("测试获取Router所有用户", testIndexHtml)
	//t.Run("测试获取Router所有用户", testUserPostForm)
	t.Run("测试获取Router所有用户", testUserPostFormEmailErrorAndPasswordError)
}

func testIndexGetRouter(t *testing.T) {
	//router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "GIN hello world", w.Body.String())

}

func testUserSave(t *testing.T) {
	username := "lisi"
	//router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user/"+username, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户"+username+"已经保存", w.Body.String())
}

func testUserSaveQuery(t *testing.T) {
	username := "lisi"
	//age := 18
	//router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	//req, _ := http.NewRequest(http.MethodGet, "/user?name="+username+"&age="+strconv.Itoa(age), nil)
	req, _ := http.NewRequest(http.MethodGet, "/user?name="+username, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	//assert.Equal(t, "用户:"+username+",年龄:"+strconv.Itoa(age)+"已经保存", w.Body.String())
	assert.Equal(t, "用户:"+username+",年龄:20已经保存", w.Body.String())
}

func testIndexHtml(t *testing.T) {
	//router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "hello gingetmethod")
}

func testUserPostForm(t *testing.T) {
	value := url.Values{}
	value.Add("email", "youngxhui@gmail.com")
	value.Add("password", "123")
	value.Add("password-again", "123")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/user/register", bytes.NewBufferString(value.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
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
	assert.Equal(t, http.StatusOK, w.Code)
}
