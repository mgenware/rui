package api

import "rui/app/handler"

// Router is the root router for all APIs.
var Router = handler.NewJSONRouter()

func init() {
	Router.Post("/form_api", formAPI)
	Router.Post("/json_api", jsonAPI)
}
