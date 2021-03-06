package homep

import (
	"net/http"
	"time"

	"rui/app"
	"rui/app/handler"
)

var indexView = app.TemplateManager.MustParseLocalizedView("home.html")

// HomeGET is the HTTP handler for root URL.
func HomeGET(w http.ResponseWriter, r *http.Request) handler.HTML {
	resp := app.HTMLResponse(w, r)
	pageData := &HomePageData{Time: time.Now().String()}
	pageHTML := indexView.MustExecuteToString(resp.Lang(), pageData)

	d := app.MasterPageData(resp.Dictionary().Home, pageHTML)
	return resp.MustComplete(d)
}
