// +build appengine

package log

// Provides standard debug logging functionality by wrapping Google App Engine (GAE)'s logging methods.

/*
For example, this can be used to log GORM's database queries:

```
db.LogMode(true)
db.SetLogger(log.NewDebugLogger(context.Context(r)))
```
*/

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	l "log"
)

type DebugLogger struct {
	context *context.Context
}

func NewDebugLogger(ctx *context.Context) *DebugLogger {
	return &DebugLogger{context: ctx}
}

func (w DebugLogger) Print(v ...interface{}) {
	if !appengine.IsDevAppServer() {
		log.Debugf(*w.context, fmt.Sprint(v))
	} else {
		l.Print(v)
	}
}
