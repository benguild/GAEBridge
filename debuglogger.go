// +build appengine

package gaebridge

// Provides standard debug logging functionality by wrapping Google App Engine (GAE)'s logging methods.

/*
For example, this can be used to log GORM's database queries:

```
db.LogMode(true);
db.SetLogger(log.New(gaebridge.DebugLogWriter{ gaebridge.Context(c) }, "", log.LstdFlags));
```
 */

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

type DebugLogWriter struct {
	Context *context.Context

}

func (w DebugLogWriter) Write(p []byte) (n int, err error) {
	log.Debugf(*w.Context, string(p));
	return len(p), nil;

}