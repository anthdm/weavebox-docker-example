package httprouter

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/httprouter"
)

func BenchmarkGet(b *testing.B) {
	router := httprouter.New()
	router.GET("/hello/:name", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {})

	for i := 0; i < b.N; i++ {
		r, err := http.NewRequest("GET", "/hello/foo", nil)
		if err != nil {
			panic(err)
		}
		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)
	}

}
