# GAEBridge
A set of handy tools for using [Google App Engine](https://cloud.google.com/appengine/docs/go/) with [Golang](https://golang.org/) (Go).

* ~~`context` provides context management.~~
  * This file has been omitted in v0.2.0 because of superior Golang design patterns, and therefore it's unnecessary.
* `log` presents a standard logging interface, wrapping Google App Engine's provided methods.

## Example

Provide a standard debug logging interface to [GORM](https://github.com/jinzhu/gorm) on the Google App Engine context using the Google App Engine logging interface:

```
db.LogMode(true);
db.SetLogger(NewDebugLogger(appengine.NewContext(request)));
```

Due to the common names of the packages, you may wish to alias them:

```
import (
  GAEBridgeLog "github.com/benguild/GAEBridge/log"
)
```

## Author

Ben Guild, hello@benguild.com

## License

`GAEBridge` is available under the MIT license. See the LICENSE file for more info.
