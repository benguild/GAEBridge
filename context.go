// +build appengine

package gaebridge

// Only creates and returns a single Google App Engine (GAE) context, once per request.

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"net/http"
	"sync"
)

var (
	mutex    sync.RWMutex
	contexts = make(map[*http.Request]*context.Context)
)

func Context(r *http.Request) *context.Context {

	mutex.RLock()
	if contexts[r] == nil { //Currently doesn't exist so make a new context
		mutex.RUnlock()
		mutex.Lock()
		newContext := appengine.NewContext(r)
		contexts[r] = &newContext
		mutex.Unlock()
		return &newContext
	} else {
		value := contexts[r]
		mutex.RUnlock()
		return value
	}
}

// Must be called directly prior to end of request cycle.
// This is to prevent a memory leak.
func CleanUp(r *http.Request) {
	mutex.Lock()
	delete(contexts, r)
	mutex.Unlock()
}
