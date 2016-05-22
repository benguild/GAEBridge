// +build appengine

package log

// Provides standard debug logging functionality by wrapping Google App Engine (GAE)'s logging methods.

/*
For example, this can be used to log GORM's database queries:

```
db.LogMode(true)
db.SetLogger(log.NewDebugLogWriter(context.Context(r)))
```
*/

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	l "log"
)

type DebugLogWriter struct {
	context *context.Context
}

func NewDebugLogWriter(ctx *context.Context) *DebugLogWriter {
	return &DebugLogWriter{context: ctx}
}

func (w DebugLogWriter) Print(v ...interface{}) {
	if !appengine.IsDevAppServer() {
		log.Debugf(*w.Context, fmt.Sprint(v))
	} else {
		l.Print(v)
	}
}
