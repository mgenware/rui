package handler

import (
	"errors"
	"net/http"
)

// HTMLResponse helps you create a HTTP response in HTML with MasterPageData.
type HTMLResponse struct {
	BaseResponse

	writer      http.ResponseWriter
	isCompleted bool
}

// HTML is a dummy type returned by HTTPResponse to make sure response is completed.
type HTML = int

// NewHTMLResponse creates a new HTMLResponse.
func NewHTMLResponse(r *http.Request, mgr *Manager, wr http.ResponseWriter) *HTMLResponse {
	return &HTMLResponse{
		BaseResponse: newBaseResponse(r, mgr),
		writer:       wr,
	}
}

// MustCompleteWithContent finished the response with the given HTML content.
func (h *HTMLResponse) MustCompleteWithContent(content string, w http.ResponseWriter) {
	h.checkCompletion()
	h.mgr.MustCompleteWithContent([]byte(content), w)
}

// MustComplete finishes the response with the given MasterPageData, and panics if unexpected error happens.
func (h *HTMLResponse) MustComplete(d *MasterPageData) HTML {
	h.checkCompletion()
	h.mgr.MustComplete(h.Request(), h.lang, d, h.writer)
	return HTML(0)
}

// MustFail finishes the response with the given error object.
func (h *HTMLResponse) MustFail(err error) HTML {
	h.MustFailWithError(err, false)
	return HTML(0)
}

// MustFailWithUserError finishes the response with an user error (expected error) message.
func (h *HTMLResponse) MustFailWithUserError(msg string) HTML {
	h.MustFailWithError(errors.New(msg), true)
	return HTML(0)
}

// MustFailWithError finishes the response with the given error and `expected` arguments, and panics if unexpected error happens.
func (h *HTMLResponse) MustFailWithError(err error, expected bool) HTML {
	h.checkCompletion()
	h.mgr.MustError(h.Request(), h.lang, err, expected, h.writer)
	return HTML(0)
}

func (h *HTMLResponse) checkCompletion() {
	if h.isCompleted {
		panic(errors.New("Result has completed"))
	}
	h.isCompleted = true
}
