# Go HTTP Logger

Adapted from [Gin](https://github.com/gin-gonic/gin) for `net/http`.

```go
http.ListenAndServe("localhost:3456", logger.Logger(http.DefaultServeMux))
```

License: [WTFPL](http://www.wtfpl.net/)
