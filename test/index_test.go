package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/sourcecmdb/blogs/ginhello"
	"github.com/stretchr/testify/assert"
)

func testIndexGetRouter(t *testing.T) {
	router := ginhello.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "GIN hello world", w.Body.String())

}

// TestMain =函数可以在测试函数执行之前做一些其他的操作
func TestMain(m *testing.M) {
	fmt.Println("测试开始")
	//通过m.Run()来执行测试函数
	m.Run()
}

func TestAll(t *testing.T) {
	gin.SetMode(gin.TestMode)
	fmt.Println("开始测试 Router 中的相关方法")
	t.Run("测试获取Router所有用户", testIndexGetRouter)
}
