package test

import (
	"github.com/astaxie/beego"
	_ "go_web/lesson6/blog-app-mongo/routers"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestCase struct {
	id         string
	StatusCode int
}

func TestPostLoadOK(t *testing.T) {

	for caseNum, item := range createCases() {
		r := httptest.NewRequest("GET", "/post?id="+item.id, nil)
		w := httptest.NewRecorder()

		beego.BeeApp.Handlers.ServeHTTP(w, r)

		if w.Code != item.StatusCode {
			t.Errorf("[%d] wrong status code: got %d, expected %d", caseNum, w.Code, item.StatusCode)
		}
	}
}

func createCases() []TestCase {
	return []TestCase{
		{
			id:         "1",
			StatusCode: http.StatusOK,
		},
		{
			id:         "10",
			StatusCode: http.StatusNotFound,
		},
	}
}
