# LetsGoBack

**LetsGoBack** is a backend framework written in **[Go](https://go.dev/doc/)**, inspired by popular frameworks like **[Gin](https://github.com/gin-gonic/gin)** and **[Fiber](https://github.com/gofiber/fiber)**.

It's built on top of the **[net/http](https://pkg.go.dev/net/http)** package.

## ⚠️ Disclaimer

This project is currently under development and is not yet finished. Features and documentation may change frequently.

## ⚙️ Installation and Quickstart

LetsGoBack requires **Go `v1.24` or higher to work**. If you need to install or upgrade Go, visit the official [Go download page](https://go.dev/dl/).
After creating your project, install LetsGoBack with the following command:

```bash
go get -u github.com/ines-mgg/LetsGoBack@latest
```

This command fetches the package and adds it to your project's dependencies, allowing you to start building your web applications with it.
Using LetsGoBack is very easy, let's create a very basic web server that responds with "Hello World, I'm LetsGoBack Framework!" on the root path.

```Go
package main

import (
    "log"
    router "github.com/ines-mgg/LetsGoBack/Router"
)

func main() {
    // Initialize a new LetsGoBack router
    log.Println("Starting server...")
    r := router.NewRouter()

    // Define a route for the GET method on the root path '/'
    r.GET("/", func(c *context.Context) {
    c.RespondOK("Hello World, I'm LetsGoBack Framework!")
    })

    // Start the server on port 8080
    log.Fatal(r.Listen(":8080"))
}
```

Now, just run this Go program, and visit `http://localhost:8080` in your browser to see the message.
For a more complete example, go to [LetsGoBackExample](https://github.com/ines-mgg/LetsGoBackExample).

## ⚡️ Features

- **Fast and Lightweight**: Minimal overhead, built directly on top of Go's `net/http`.
- **Intuitive Routing**: Simple route definitions with support for dynamic parameters.
- **Middleware Support**: Easily add, chain, and manage middleware functions.
- **Context Object**: Centralized context for each request, with helpers for JSON, files, and more.
- **Error Handling**: Built-in error recovery and customizable error responses.
- **Request Validation**: Helpers for validating and parsing incoming data.
- **CORS Support**: Easily enable and configure Cross-Origin Resource Sharing.
- **File Uploads**: Helpers and middleware for handling file uploads.

## 🧑‍💻 Examples

**Basic routing**:

```Go
package main

import (
    "log"
    context "github.com/ines-mgg/LetsGoBack/Context"
    router "github.com/ines-mgg/LetsGoBack/Router"
)

func main() {
    r := router.NewRouter()
    r.GET("/", func(c *context.Context) {
        c.RespondOK("Hello World, I'm LetsGoBack Framework!")
    })
    r.GET("/hello", func(c *context.Context) {
        c.RespondOK("Hello, this is a simple route!")
    })
    r.GET("/greet/:name", func(c *context.Context) {
        name := c.Param("name")
        if name == "" {
            c.ErrorBadRequest("Name parameter is required")
        return
    }
        c.RespondOK("Hello, "+name+"!")
    })
    r.GET("/header", func(c *context.Context) {
        customHeader := c.Request.Header.Get("X-Custom-Header")
        if customHeader == "" {
            c.ErrorBadRequest("X-Custom-Header is required")
        return
         }
        c.RespondOK("Custom header value: " + customHeader)
    })
    log.Fatal(r.Listen(":8080"))
}
```

**Custom NotFound handler**:

```Go
package main

import (
    "log"
    context "github.com/ines-mgg/LetsGoBack/Context"
    router "github.com/ines-mgg/LetsGoBack/Router"
)

func main() {
    r := router.NewRouter()
    r.NotFoundHandler = func(c *context.Context) {
        c.ErrorNotFound("Page not found")
    }
    log.Fatal(r.Listen(":8080"))
}
```

**Custom MethodNotAllowed handler**:

```Go
package main

import (
    "log"
    context "github.com/ines-mgg/LetsGoBack/Context"
    router "github.com/ines-mgg/LetsGoBack/Router"
)

func main() {
    r := router.NewRouter()
    r.MethodNotAllowedHandler = func(c *context.Context) {
        c.ErrorMethodNotAllowed("Method not allowed")
    }
    log.Fatal(r.Listen(":8080"))
}
```

**Group routes**:

```Go
package main

import (
    "log"
    context "github.com/ines-mgg/LetsGoBack/Context"
    router "github.com/ines-mgg/LetsGoBack/Router"
)

func main() {
    r := router.NewRouter()
    auth := r.Group("/auth")
    auth.POST("/register", func(c *context.Context) {
        // Add logic here...
    })
    auth.POST("/login", func(c *context.Context) {
        // Add logic here...
    })
    log.Fatal(r.Listen(":8080"))
}
```

**Serving Static Files**:

```Go
package main

import (
    "log"
    context "github.com/ines-mgg/LetsGoBack/Context"
    router "github.com/ines-mgg/LetsGoBack/Router"
)

func main() {
    r := router.NewRouter()
    // Serve static files from the "uploads" directory
    r.ServeStatic("/uploads", "./uploads")
    log.Fatal(r.Listen(":8080"))
}
```

**JSON Responses**:

```Go
package main

import (
    "log"
    context "github.com/ines-mgg/LetsGoBack/Context"
    router "github.com/ines-mgg/LetsGoBack/Router"
)

func main() {
    r := router.NewRouter()
    r.GET("/json", func(c *context.Context) {
        data := map[string]interface{}{
            "message": "Hello, JSON!",
            "success": true,
        }
        c.RespondOK(data)
    })
    r.GET("/json/custom-status", func(c *context.Context) {
        c.ErrorNotFound("Resource not found")
    })
    log.Fatal(r.Listen(":8080"))
}
```

**Middleware custom**:

```Go
package main

import (
    "log"
    context "github.com/ines-mgg/LetsGoBack/Context"
    router "github.com/ines-mgg/LetsGoBack/Router"
)

func main() {
    r := router.NewRouter()
    r.Use(func(next context.HandlerFunc) context.HandlerFunc {
        return func(c *context.Context) {
        // Custom logic before the handler
        log.Println("Custom middleware before handler")
        // Call the next handler in the chain
        next(c)
        // Custom logic after the handler
        log.Println("Custom middleware after handler")
        }
    })
    log.Fatal(r.Listen(":8080"))
}
```

**Middleware logger**:

```Go
package main

import (
    "log"
    router "github.com/ines-mgg/LetsGoBack/Router"
    middleware "github.com/ines-mgg/LetsGoBack/Middleware"
)

func main() {
    r := router.NewRouter()
    r.Use(middleware.RequestIDMiddleware())
    r.Use(middleware.LoggerMiddleware("2006-01-02 15:04:05"))
    log.Fatal(r.Listen(":8080"))
}
```

**Middleware CORS**:

```Go
package main

import (
    "log"
    router "github.com/ines-mgg/LetsGoBack/Router"
    middleware "github.com/ines-mgg/LetsGoBack/Middleware"
)

func main() {
    r := router.NewRouter()
    corsHeaders := map[string]any{
        "Access-Control-Allow-Methods":     "GET, POST, PUT, PATCH, DELETE",
        "Access-Control-Allow-Headers":     "Content-Type, Authorization",
        "Access-Control-Allow-Credentials": "true",
        "Access-Control-Allow-Origin":      "*", 
    }
    r.Use(middleware.CORSMiddleware(corsHeaders))
    log.Fatal(r.Listen(":8080"))
}
```

**Middleware Recovery**:

```Go
package main

import (
    "log"
    router "github.com/ines-mgg/LetsGoBack/Router"
    middleware "github.com/ines-mgg/LetsGoBack/Middleware"
)

func main() {
    r := router.NewRouter()
    r.Use(middleware.RecoverMiddleware())
    log.Fatal(r.Listen(":8080"))
}
```

**Middleware jwtAuth**:

```Go
package main

import (
    "log"
    router "github.com/ines-mgg/LetsGoBack/Router"
    middleware "github.com/ines-mgg/LetsGoBack/Middleware"
)

func main() {
    r := router.NewRouter()
    r.Use(middleware.JWTAuthMiddleware("jwtClaims"))
    log.Fatal(r.Listen(":8080"))
}
```

## Contributing

Help is always appreciated ! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details on submitting patches and the contribution workflow.
