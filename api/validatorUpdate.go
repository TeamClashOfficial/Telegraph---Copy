package api

import (
	"io"
	"net/http"
	"telegraph/binance"
	"telegraph/tools"
)

func Update(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(io.LimitReader(r.Body, 1048576))
	tools.Check(err)
	go binance.UpdateFromBytes(bytes)
	_, _ = w.Write([]byte("Update started"))
}
