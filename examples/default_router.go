package examples

import (
	"log"
	"net"
	"net/http"

	"github.com/thamizhv/elasticapm-customtrace/trace"
	"go.elastic.co/apm/module/apmhttp"
)

func DefaultRouter() {
	mux := http.DefaultServeMux
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	log.Fatal(http.ListenAndServe(net.JoinHostPort("", "8081"), trace.SetTraceID(apmhttp.Wrap(mux))))
}
