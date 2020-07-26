package examples

import (
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thamizhv/elasticapm-customtrace/trace"
	"go.elastic.co/apm/module/apmgorilla"
)

func GorillaRouter() {
	mux := mux.NewRouter()
	mux.Use(trace.SetTraceID, apmgorilla.Middleware())

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	log.Fatal(http.ListenAndServe(net.JoinHostPort("", "8081"), mux))
}
