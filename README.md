# GAEBridge
A set of handy tools for using Google App Engine with Golang (Go).

* "context" provides context management.
* "debuglogger" presents a standard logging interface, wrapping Google App Engine's provided methods.

## Example

Provide a standard debug logging interface to [GORM](https://github.com/jinzhu/gorm) on the Google App Engine context using the Google App Engine logging interface:

```
db.LogMode(true);
db.SetLogger(log.New(gaebridge.DebugLogWriter{ gaebridge.Context(r) }, "", log.LstdFlags));
```

## Author

Ben Guild, email@benguild.com

## License

`GAEBridge` is available under the MIT license. See the LICENSE file for more info.
