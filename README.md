# Let's Go Back

Let's Go Back is a backend framework written in Go, inspired by popular frameworks like Gin and Fiber.

## How Does It work ?

```Bash
.
├── Context/
│   └── context.go
│   └── error.go
│   └── file.go
│   └── json.go
│   └── types.go
│   └── utils.go
├── Middleware/
│   └── apply.go
│   └── cors.go
│   └── error.go
│   └── logger.go
│   └── recover.go
│   └── requestID.go
│   └── types.go
│   └── upload.go
│   └── utils.go
├── Router/
│   └── router.go
│   └── routerGroup.go
│   └── server.go
│   └── types.go
│   └── utils.go
├── Utils/
│   └── jwt.go
├── go.mod
└── go.sum
```

### Context

**Purpose**: encapsulate everything related to a specific HTTP request.

It plays a central role (it’s the "core" of the processing):

- Provides convenient access to the Request, Writer, Params, Method, Path, etc.
- Manages data shared between middlewares.
- Offers helpers.

### Middleware

**Purpose**: define how to chain and execute middlewares.

Here you find:

- Reusable middlewares: LoggerMiddleware, ErrorRecoveryMiddleware, UploadValidatorMiddleware, etc.

### Router

**Purpose**: define the routing system and route resolution.

Here you:

- Declare routes: GET, POST, etc.
- Store the association between method + path → handler.
- Apply middlewares during route resolution.
- Resolve dynamic parameters (e.g., `/users/:id`).

### Utils

## Disclaimer

This project is currently under development and is not yet finished. Features and functionality may change as the project evolves.
