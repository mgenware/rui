package api

import (
	"net/http"
	"rui/app/defs"
	"rui/app/handler"

	"rui/app"
)

func jsonAPI(w http.ResponseWriter, r *http.Request) handler.JSON {
	dict := defs.BodyContext(r.Context())
	return app.JSONResponse(w, r).MustComplete(dict)
}
