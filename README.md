# GAEBridge
A set of handy tools for using [Google App Engine](https://cloud.google.com/appengine/docs/go/) with [Golang](https://golang.org/) (Go).

* `context.go` provides context management.
* `debuglogger.go` presents a standard logging interface, wrapping Google App Engine's provided methods.

## Example

Provide a standard debug logging interface to [GORM](https://github.com/jinzhu/gorm) on the Google App Engine context using the Google App Engine logging interface:

```
db.LogMode(true);
db.SetLogger(NewDebugLogger(Context(c.Request)));
```

**Important!** Be sure to clean up the singleton instance of the context when finished:

```
defer Context.CleanUp(c.Request); // You can call this code from custom middleware, for example.
```

Due to the common names of the packages, you may wish to alias them:

```
import (
  GAEBridgeContext "github.com/benguild/GAEBridge/context"
  GAEBridgeLog "github.com/benguild/GAEBridge/log"
)
```

## Author

Ben Guild, email@benguild.com

## License

`GAEBridge` is available under the MIT license. See the LICENSE file for more info.
