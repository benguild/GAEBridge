// +build appengine

package context

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
	value := contexts[r]
	if value != nil {
		defer mutex.RUnlock()
		return value
	}
	mutex.RUnlock()

	// Currently doesn't exist, so make a new context!
	newContext := appengine.NewContext(r)

	mutex.Lock()
	defer mutex.Unlock()

	if contexts[r] != nil {
		contexts[r] = &newContext
		return &newContext
	} else {
		// *****REMOVE VERBOSE TEXT BELOW *****

		//This case is very very unusual but necessary for maximum robustness.
		//It may seem redundant since we did a similar check at the top.
		//We had 2 choices -
		// a) remove the RLocks and just have 1 Full lock to cover this entire method.
		// We would then not need to cater for this case at all since we already did an existence check at the top.
		// b) The other is this current option.

		// Option a is not best because Full Locks slow down code more than ReadLocks. They are only necessary if you are writing to the map.
		// Since the majority of time we'd be checking for existence, rather than setting the newContext, it is better to use a ReadLock.
		// Now since we decided to use ReadLocks, then in between the mutex.RUnlock() and mutex.Lock() leaves a case where 2 concurrent goroutines
		// could believe that a context hasn't yet been created. They both want to create the new context. The Full Lock only allows that to happen
		// one at a time, despite both goroutines wanting to create it. That means the first goroutine will create a new context and the second one
		// will again create a new context.

		return contexts[r]
	}
}

func CleanUp(r *http.Request) { // Must be called directly prior to end of request cycle to prevent a memory leak.
	mutex.Lock()
	delete(contexts, r)
	mutex.Unlock()
}
