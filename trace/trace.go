package trace

import (
	"net/http"

	"github.com/google/uuid"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmhttp"
)

// TraceID is the header that is passed in outgoing response
const TraceID = "traceid"

// SetTraceID is a http middleware that checks for "Traceparent" header in the
// incoming requests. If present, parse and extract the traceid. Else create a custom traceid
// and sets it as new "Traceparent" header in incoming request.
// Also sets that traceid in outgoing response header.
func SetTraceID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var traceID apm.TraceID
		if values := r.Header[apmhttp.W3CTraceparentHeader]; len(values) == 1 && values[0] != "" {
			if c, err := apmhttp.ParseTraceparentHeader(values[0]); err == nil {
				traceID = c.Trace
			}
		}
		if err := traceID.Validate(); err != nil {
			uuid := uuid.New()
			var spanID apm.SpanID
			var traceOptions apm.TraceOptions
			copy(traceID[:], uuid[:])
			copy(spanID[:], traceID[8:])
			traceContext := apm.TraceContext{
				Trace:   traceID,
				Span:    spanID,
				Options: traceOptions.WithRecorded(true),
			}
			r.Header.Set(apmhttp.W3CTraceparentHeader, apmhttp.FormatTraceparentHeader(traceContext))
		}

		w.Header().Set(TraceID, traceID.String())
		next.ServeHTTP(w, r)
	})
}
