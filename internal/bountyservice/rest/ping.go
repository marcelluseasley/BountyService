package rest 

import (
	"net/http"
)



func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("online"))
}
