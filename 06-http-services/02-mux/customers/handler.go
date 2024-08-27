package customers

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"req-id": r.Context().Value("request-id"),
	}).Warn("[Customers Handler]")
	fmt.Fprintln(w, "All customer data will be served")
}
