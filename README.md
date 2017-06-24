<a href="https://travis-ci.org/go-siris/middleware"><img src="https://img.shields.io/travis/go-siris/middleware.svg?style=flat-square" alt="Build Status"></a>
<a href="https://github.com/go-siris/siris/blob/master/HISTORY.md"><img src="https://img.shields.io/badge/release-v7.3.0-blue.svg?style=flat-square" alt="CHANGELOG/HISTORY"></a>


This repository provides a way to share your [Iris](https://github.com/go-siris/siris)' specific middleware with the rest of us. You can view the @go-siris supported middleware by pressing [here](https://github.com/go-siris/siris/tree/master/middleware).


Installation
------------
The only requirement is the [Go Programming Language](https://golang.org/dl), at least version 1.8

```bash
$ go get github.com/go-siris/middleware/...
```


FAQ
------------
Explore [these questions](https://github.com/go-siris/middleware/issues) or navigate to the [community chat][Chat].


People
------------
The Community.



[Chat]: https://gitter.im/gosiris/siris


# What?

Middleware are just handlers which can be served before or after the main handler, can transfer data between handlers and communicate with third-party libraries, they are just functions.

### How can I install a middleware?

```sh
$ go get -u github.com/go-siris/middleware/$FOLDERNAME
```

**NOTE**: When you install one middleware you will have all of them downloaded & installed, **no need** to re-run the go get foreach middeware.

### How can I register middleware?


**To a single route**
```go
app := siris.New()
app.Get("/mypath",myMiddleware1,myMiddleware2,func(ctx context.Context){}, func(ctx context.Context){},myMiddleware5,myMainHandlerLast)
```

**To a party of routes or subdomain**
```go

myparty := app.Party("/myparty", myMiddleware1,func(ctx context.Context){},myMiddleware3)
{
	//....
}

```

**To all routes**
```go
app.Use(func(ctx context.Context){}, myMiddleware2)
```

**To global, all routes on all subdomains on all parties**
```go
app.UseGlobal(func(ctx context.Context){}, myMiddleware2)
```


# Can I use standard net/http handler with Siris?

**Yes** you can, just pass the Handler inside the `handlerconv.FromStd` in order to be converted into siris.HandlerFunc and register it as you saw before.

## handler which has the form of http.Handler/HandlerFunc

```go
package main

import (
	"github.com/go-siris/siris"
	"github.com/go-siris/siris/context"
	"github.com/go-siris/siris/core/handlerconv"
)

func main() {
	app := siris.New()

	sillyHTTPHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	     println(r.RequestURI)
	})

	sillyConvertedToSiris := handlerconv.FromStd(sillyHTTPHandler)
	// FromStd can take (http.ResponseWriter, *http.Request, next http.Handler) too!
	app.Use(sillyConvertedToSiris)

	app.Run(siris.Addr(":8080"))
}

```
