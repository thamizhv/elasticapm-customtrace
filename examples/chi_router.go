package examples

import (
	"log"
	"net"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/thamizhv/elasticapm-customtrace/trace"
	"go.elastic.co/apm/module/apmchi"
)

func ChiRouter() {
	mux := chi.NewRouter()
	mux.Use(trace.SetTraceID, apmchi.Middleware())

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	log.Fatal(http.ListenAndServe(net.JoinHostPort("", "8081"), mux))
}
