package gaebridge

// Only creates and returns a single Google App Engine (GAE) context, once per request.

import (
	"google.golang.org/appengine"
	"golang.org/x/net/context"
	"net/http"
)

var (
	ctx *context.Context

)

func Context(r *http.Request) *context.Context {
	if (ctx==nil) {
		newContext := appengine.NewContext(r);
		ctx = &newContext;

	}

	return ctx;

}