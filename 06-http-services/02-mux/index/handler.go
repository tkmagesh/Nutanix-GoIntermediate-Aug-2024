package index

import (
	"fmt"
	"net/http"

	log "context-app/log"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Context(), "[Index Handler]", log.LogFields{"user": "magesh"})
	fmt.Fprintln(w, "Hello World!")
}
