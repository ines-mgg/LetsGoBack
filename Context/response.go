package Context

import (
	"net/http"
)

// RespondContinue sends a 100 Continue response with the provided message.
// This response is typically used to indicate that the initial part of a request has been received
// and the client can continue sending the rest of the request.
// It is often used in situations where the client needs to send a large request body,
// and the server wants to confirm that it is ready to receive it.
// The message can be any data type, and it will be serialized to JSON format before being sent.
func (c *Context) RespondContinue(msg any) {
	c.json(http.StatusContinue, msg)
}

// RespondSwitchingProtocols sends a 101 Switching Protocols response with the provided message.
// This response is used when the server agrees to switch protocols as requested by the client,
// typically in response to an Upgrade header in the request.
// It indicates that the server is switching to a different protocol, such as from HTTP/1.1 to WebSocket.
// The message can be any data type, and it will be serialized to JSON format before being sent.
func (c *Context) RespondSwitchingProtocols(msg any) {
	c.json(http.StatusSwitchingProtocols, msg)
}

// RespondProcessing sends a 102 Processing response with the provided message.
// This response is used to indicate that the server has received the request
// and is processing it, but no response is available yet.
// It is typically used in situations where the server needs to perform a long-running operation
// before sending a final response to the client.
func (c *Context) RespondProcessing(msg any) {
	c.json(http.StatusProcessing, msg)
}

// RespondEarlyHints sends a 103 Early Hints response with the provided message.
// This response is used to provide hints to the client about resources that might be needed
// for the request, allowing the client to start preloading resources before the final response is sent.
// It is typically used to improve performance by reducing latency in loading resources.
// The message can be any data type, and it will be serialized to JSON format before being sent.
// The 103 Early Hints response is part of the HTTP/2 and HTTP/3 protocols and is not widely supported in all browsers.
func (c *Context) RespondEarlyHints(msg any) {
	c.json(http.StatusEarlyHints, msg)
}

// RespondOK sends a 200 OK response with the provided message.
// This response indicates that the request has succeeded and the server is returning the requested data.
// It is the standard response for successful HTTP requests.
// The message can be any data type, and it will be serialized to JSON format before being sent.
func (c *Context) RespondOK(msg any) {
	c.json(http.StatusOK, msg)
}

// RespondCreated sends a 201 Created response with the provided message.
// This response indicates that a new resource has been successfully created as a result of the request.
// It is typically used in response to POST requests that create new resources.
// The message can be any data type, and it will be serialized to JSON format before being sent.
func (c *Context) RespondCreated(msg any) {
	c.json(http.StatusCreated, msg)
}

// RespondAccepted sends a 202 Accepted response with the provided message.
// This response indicates that the request has been accepted for processing,
// but the processing has not been completed yet.
// It is typically used for asynchronous operations where the server will process the request later.
func (c *Context) RespondAccepted(msg any) {
	c.json(http.StatusAccepted, msg)
}

// RespondNonAuthoritativeInfo sends a 203 Non-Authoritative Information response with the provided message.
// This response indicates that the server successfully processed the request,
// but is returning information that may be from a third-party source.
// It is typically used when the server is acting as a proxy or gateway and the information is not authoritative.
func (c *Context) RespondNonAuthoritativeInfo(msg any) {
	c.json(http.StatusNonAuthoritativeInfo, msg)
}

// RespondNoContent sends a 204 No Content response with the provided message.
// This response indicates that the server has successfully processed the request,
// but there is no content to return in the response body.
// It is typically used for successful DELETE requests or when the server has no additional information to provide.
func (c *Context) RespondNoContent(msg any) {
	c.json(http.StatusNoContent, msg)
}

// RespondResetContent sends a 205 Reset Content response with the provided message.
// This response indicates that the server has successfully processed the request,
// but the client should reset the document view or clear the form.
func (c *Context) RespondResetContent(msg any) {
	c.json(http.StatusResetContent, msg)
}

// RespondPartialContent sends a 206 Partial Content response with the provided message.
// This response indicates that the server is returning only a portion of the requested resource,
// typically in response to a Range header in the request.
func (c *Context) RespondPartialContent(msg any) {
	c.json(http.StatusPartialContent, msg)
}

// RespondMultiStatus sends a 207 Multi-Status response with the provided message.
// This response is used to convey multiple status codes for different parts of a request,
// typically in response to a WebDAV request.
func (c *Context) RespondMultiStatus(msg any) {
	c.json(http.StatusMultiStatus, msg)
}

// RespondAlreadyReported sends a 208 Already Reported response with the provided message.
// This response indicates that the server has already reported the status of the requested resource,
// typically in response to a WebDAV request.
func (c *Context) RespondAlreadyReported(msg any) {
	c.json(http.StatusAlreadyReported, msg)
}

// RespondIMUsed sends a 226 IM Used response with the provided message.
// This response indicates that the server has fulfilled a request for the resource,
// and the response is a representation of the result of one or more instance manipulations.
func (c *Context) RespondIMUsed(msg any) {
	c.json(http.StatusIMUsed, msg)
}

// RespondMultipleChoices sends a 300 Multiple Choices response with the provided message.
// This response indicates that there are multiple options for the requested resource,
// and the client should choose one of them.
// It is typically used when the server has multiple representations of a resource,
// such as different formats or languages, and the client needs to select one.
func (c *Context) RespondMultipleChoices(msg any) {
	c.json(http.StatusMultipleChoices, msg)
}

// RespondMovedPermanently sends a 301 Moved Permanently response with the provided message.
// This response indicates that the requested resource has been permanently moved to a new URL,
// and the client should use the new URL for future requests.
func (c *Context) RespondMovedPermanently(msg any) {
	c.json(http.StatusMovedPermanently, msg)
}

// RespondFound sends a 302 Found response with the provided message.
// This response indicates that the requested resource has been temporarily moved to a different URL,
// and the client should use the new URL for this request.
func (c *Context) RespondFound(msg any) {
	c.json(http.StatusFound, msg)
}

// RespondSeeOther sends a 303 See Other response with the provided message.
// This response indicates that the server is redirecting the client to a different URL,
// typically in response to a POST request that has been processed successfully.
func (c *Context) RespondSeeOther(msg any) {
	c.json(http.StatusSeeOther, msg)
}

// RespondNotModified sends a 304 Not Modified response with the provided message.
// This response indicates that the requested resource has not been modified since the last request,
// and the client can use the cached version of the resource.
func (c *Context) RespondNotModified(msg any) {
	c.json(http.StatusNotModified, msg)
}

// RespondUseProxy sends a 305 Use Proxy response with the provided message.
// This response indicates that the requested resource must be accessed through a proxy,
// and the client should use the specified proxy to access the resource.
// Note: The 305 Use Proxy status code is deprecated and should not be used in new applications.
// It is included here for completeness, but it is recommended to use other status codes for proxy-related responses.
func (c *Context) RespondUseProxy(msg any) {
	c.json(http.StatusUseProxy, msg)
}

// RespondTemporaryRedirect sends a 307 Temporary Redirect response with the provided message.
// This response indicates that the requested resource is temporarily located at a different URL,
// and the client should use the new URL for this request.
func (c *Context) RespondTemporaryRedirect(msg any) {
	c.json(http.StatusTemporaryRedirect, msg)
}

// RespondPermanentRedirect sends a 308 Permanent Redirect response with the provided message.
// This response indicates that the requested resource has been permanently moved to a new URL,
// and the client should use the new URL for future requests.
// It is similar to the 301 Moved Permanently status code, but it does not allow the HTTP method to change.
func (c *Context) RespondPermanentRedirect(msg any) {
	c.json(http.StatusPermanentRedirect, msg)
}

// ErrorBadRequest sends a 400 Bad Request response with the provided message.
// This response indicates that the server cannot process the request due to a client error,
// such as malformed syntax or invalid request parameters.
func (c *Context) ErrorBadRequest(msg string) {
	c.abortWithStatusJSON(http.StatusBadRequest, msg)
}

// ErrorUnauthorized sends a 401 Unauthorized response with the provided message.
// This response indicates that the request requires user authentication,
// or the provided authentication credentials are invalid or missing.
// It is typically used when the client needs to provide valid credentials to access the requested resource.
// It is often used in conjunction with authentication mechanisms such as Basic Auth or Bearer tokens.
func (c *Context) ErrorUnauthorized(msg string) {
	c.abortWithStatusJSON(http.StatusUnauthorized, msg)
}

// ErrorPaymentRequired sends a 402 Payment Required response with the provided message.
// This response indicates that the request cannot be processed until the client makes a payment.
// It is typically used in scenarios where the requested resource requires payment or subscription,
// such as accessing premium content or services.
// It is not commonly used in practice, as most APIs do not implement payment requirements directly in HTTP responses.
func (c *Context) ErrorPaymentRequired(msg string) {
	c.abortWithStatusJSON(http.StatusPaymentRequired, msg)
}

// ErrorForbidden sends a 403 Forbidden response with the provided message.
// This response indicates that the server understands the request,
// but refuses to authorize it. The client does not have permission to access the requested resource,
// even if the request is authenticated.
func (c *Context) ErrorForbidden(msg string) {
	c.abortWithStatusJSON(http.StatusForbidden, msg)
}

// ErrorNotFound sends a 404 Not Found response with the provided message.
// This response indicates that the requested resource could not be found on the server.
// It is typically used when the client requests a resource that does not exist,
// such as a non-existent URL or an item that has been deleted.
// It is a common response for APIs when the requested endpoint or resource is not available.
// It is important to provide a clear message in the response body to help the client understand why the resource was not found.
func (c *Context) ErrorNotFound(msg string) {
	c.abortWithStatusJSON(http.StatusNotFound, msg)
}

// ErrorMethodNotAllowed sends a 405 Method Not Allowed response with the provided message.
// This response indicates that the request method is not allowed for the requested resource.
// It is typically used when the client tries to use an HTTP method (such as GET, POST, PUT, DELETE)
// that is not supported by the server for the specified resource.
func (c *Context) ErrorMethodNotAllowed(msg string) {
	c.abortWithStatusJSON(http.StatusMethodNotAllowed, msg)
}

// ErrorNotAcceptable sends a 406 Not Acceptable response with the provided message.
// This response indicates that the server cannot produce a response that matches the criteria specified by the client,
// such as the requested content type or language.
func (c *Context) ErrorNotAcceptable(msg string) {
	c.abortWithStatusJSON(http.StatusNotAcceptable, msg)
}

// ErrorProxyAuthRequired sends a 407 Proxy Authentication Required response with the provided message.
// This response indicates that the client must first authenticate itself with the proxy server
// before the request can be fulfilled. It is typically used in scenarios where the client needs to
// provide authentication credentials to access a resource through a proxy server.
func (c *Context) ErrorProxyAuthRequired(msg string) {
	c.abortWithStatusJSON(http.StatusProxyAuthRequired, msg)
}

// ErrorRequestTimeout sends a 408 Request Timeout response with the provided message.
// This response indicates that the server did not receive a complete request from the client within the
// server's timeout period. It is typically used when the client takes too long to send the request,
// and the server closes the connection to free up resources.
func (c *Context) ErrorRequestTimeout(msg string) {
	c.abortWithStatusJSON(http.StatusRequestTimeout, msg)
}

// ErrorConflict sends a 409 Conflict response with the provided message.
// This response indicates that the request could not be completed due to a conflict with the current state of the resource.
// It is typically used when the client tries to create or update a resource that conflicts with an existing resource,
// such as attempting to create a resource with a duplicate identifier or trying to update a resource that has been modified by another client.
func (c *Context) ErrorConflict(msg string) {
	c.abortWithStatusJSON(http.StatusConflict, msg)
}

// ErrorGone sends a 410 Gone response with the provided message.
// This response indicates that the requested resource is no longer available on the server
// and has been permanently removed. It is typically used when a resource has been deleted
// and the server wants to inform the client that it should not attempt to access it again.
func (c *Context) ErrorGone(msg string) {
	c.abortWithStatusJSON(http.StatusGone, msg)
}

// ErrorLengthRequired sends a 411 Length Required response with the provided message.
// This response indicates that the server requires a Content-Length header in the request,
// but it was not provided by the client. It is typically used when the server expects a request body
// and needs to know the length of the body to process the request correctly.
func (c *Context) ErrorLengthRequired(msg string) {
	c.abortWithStatusJSON(http.StatusLengthRequired, msg)
}

// ErrorPreconditionFailed sends a 412 Precondition Failed response with the provided message.
// This response indicates that one or more preconditions specified in the request headers
// were not met by the server. It is typically used when the client includes conditions in the request,
// such as If-Match or If-None-Match, and those conditions are not satisfied by the current state of the resource.
func (c *Context) ErrorPreconditionFailed(msg string) {
	c.abortWithStatusJSON(http.StatusPreconditionFailed, msg)
}

// ErrorRequestEntityTooLarge sends a 413 Request Entity Too Large response with the provided message.
// This response indicates that the request body is larger than the server is willing or able to process.
func (c *Context) ErrorRequestEntityTooLarge(msg string) {
	c.abortWithStatusJSON(http.StatusRequestEntityTooLarge, msg)
}

// ErrorRequestURITooLong sends a 414 Request-URI Too Long response with the provided message.
// This response indicates that the URI provided in the request is too long for the server to process.
func (c *Context) ErrorRequestURITooLong(msg string) {
	c.abortWithStatusJSON(http.StatusRequestURITooLong, msg)
}

// ErrorUnsupportedMediaType sends a 415 Unsupported Media Type response with the provided message.
// This response indicates that the server refuses to accept the request because the payload format is in an unsupported format.
func (c *Context) ErrorUnsupportedMediaType(msg string) {
	c.abortWithStatusJSON(http.StatusUnsupportedMediaType, msg)
}

// ErrorRequestedRangeNotSatisfiable sends a 416 Requested Range Not Satisfiable response with the provided message.
// This response indicates that the server cannot fulfill the request for a specific range of bytes in the resource.
func (c *Context) ErrorRequestedRangeNotSatisfiable(msg string) {
	c.abortWithStatusJSON(http.StatusRequestedRangeNotSatisfiable, msg)
}

// ErrorExpectationFailed sends a 417 Expectation Failed response with the provided message.
// This response indicates that the server cannot meet the requirements specified in the Expect request header.
func (c *Context) ErrorExpectationFailed(msg string) {
	c.abortWithStatusJSON(http.StatusExpectationFailed, msg)
}

// ErrorTeapot sends a 418 I'm a teapot response with the provided message.
// This response is part of the Hyper Text Coffee Pot Control Protocol (HTCPCP),
// which is an April Fools' joke specification. It indicates that the server is a teapot
// and cannot brew coffee. It is not a standard HTTP response and is not intended for use in production applications.
func (c *Context) ErrorTeapot(msg string) {
	c.abortWithStatusJSON(http.StatusTeapot, msg)
}

// ErrorMisdirectedRequest sends a 421 Misdirected Request response with the provided message.
// This response indicates that the request was directed at a server that is not able to produce a response.
func (c *Context) ErrorMisdirectedRequest(msg string) {
	c.abortWithStatusJSON(http.StatusMisdirectedRequest, msg)
}

// ErrorUnprocessableEntity sends a 422 Unprocessable Entity response with the provided message.
// This response indicates that the server understands the content type of the request entity,
// but was unable to process the contained instructions. It is typically used when the request is well-formed,
// but the server is unable to process the request due to semantic errors in the request body.
func (c *Context) ErrorUnprocessableEntity(msg string) {
	c.abortWithStatusJSON(http.StatusUnprocessableEntity, msg)
}

// ErrorLocked sends a 423 Locked response with the provided message.
// This response indicates that the requested resource is currently locked and cannot be accessed.
func (c *Context) ErrorLocked(msg string) {
	c.abortWithStatusJSON(http.StatusLocked, msg)
}

// ErrorFailedDependency sends a 424 Failed Dependency response with the provided message.
// This response indicates that the request failed due to a failure of a previous request,
// such as when a request to create a resource depends on another resource that could not be created.
func (c *Context) ErrorFailedDependency(msg string) {
	c.abortWithStatusJSON(http.StatusFailedDependency, msg)
}

// ErrorTooEarly sends a 425 Too Early response with the provided message.
// This response indicates that the server is unwilling to risk processing a request
// that might be replayed, such as when the request is too early in a sequence of requests.
func (c *Context) ErrorTooEarly(msg string) {
	c.abortWithStatusJSON(http.StatusTooEarly, msg)
}

// ErrorUpgradeRequired sends a 426 Upgrade Required response with the provided message.
// This response indicates that the client should switch to a different protocol,
// such as HTTP/2 or WebSocket, to access the requested resource.
func (c *Context) ErrorUpgradeRequired(msg string) {
	c.abortWithStatusJSON(http.StatusUpgradeRequired, msg)
}

// ErrorPreconditionRequired sends a 428 Precondition Required response with the provided message.
// This response indicates that the server requires the request to be conditional,
// meaning that the client must include certain headers to ensure that the request is safe to process.
func (c *Context) ErrorPreconditionRequired(msg string) {
	c.abortWithStatusJSON(http.StatusPreconditionRequired, msg)
}

// ErrorTooManyRequests sends a 429 Too Many Requests response with the provided message.
// This response indicates that the client has sent too many requests in a given amount of time,
// and the server is limiting the rate of requests to prevent abuse or overload.
func (c *Context) ErrorTooManyRequests(msg string) {
	c.abortWithStatusJSON(http.StatusTooManyRequests, msg)
}

// ErrorRequestHeaderFieldsTooLarge sends a 431 Request Header Fields Too Large response with the provided message.
// This response indicates that the server is unwilling to process the request
// because the request headers are too large. It is typically used when the client sends
// excessively large headers, such as cookies or custom headers, that exceed the server's limits.
func (c *Context) ErrorRequestHeaderFieldsTooLarge(msg string) {
	c.abortWithStatusJSON(http.StatusRequestHeaderFieldsTooLarge, msg)
}

// ErrorUnavailableForLegalReasons sends a 451 Unavailable For Legal Reasons response with the provided message.
// This response indicates that the requested resource is unavailable due to legal reasons,
// such as government censorship or legal restrictions. It is typically used when the server is unable to provide access to a resource
// because it is prohibited by law or regulation.
func (c *Context) ErrorUnavailableForLegalReasons(msg string) {
	c.abortWithStatusJSON(http.StatusUnavailableForLegalReasons, msg)
}

// ErrorInternalServerError sends a 500 Internal Server Error response with the provided message.
// This response indicates that the server encountered an unexpected condition
// that prevented it from fulfilling the request. It is typically used when there is a server-side error,
// such as a bug in the application code or a failure in the server's infrastructure.
func (c *Context) ErrorInternalServerError(msg string) {
	c.abortWithStatusJSON(http.StatusInternalServerError, msg)
}

// ErrorNotImplemented sends a 501 Not Implemented response with the provided message.
// This response indicates that the server does not support the functionality required to fulfill the request.
func (c *Context) ErrorNotImplemented(msg string) {
	c.abortWithStatusJSON(http.StatusNotImplemented, msg)
}

// ErrorBadGateway sends a 502 Bad Gateway response with the provided message.
// This response indicates that the server, while acting as a gateway or proxy,
// received an invalid response from the upstream server it was trying to communicate with.
func (c *Context) ErrorBadGateway(msg string) {
	c.abortWithStatusJSON(http.StatusBadGateway, msg)
}

// ErrorServiceUnavailable sends a 503 Service Unavailable response with the provided message.
// This response indicates that the server is currently unable to handle the request
// due to temporary overloading or maintenance of the server.
func (c *Context) ErrorServiceUnavailable(msg string) {
	c.abortWithStatusJSON(http.StatusServiceUnavailable, msg)
}

// ErrorGatewayTimeout sends a 504 Gateway Timeout response with the provided message.
// This response indicates that the server, while acting as a gateway or proxy,
// did not receive a timely response from the upstream server it was trying to communicate with.
func (c *Context) ErrorGatewayTimeout(msg string) {
	c.abortWithStatusJSON(http.StatusGatewayTimeout, msg)
}

// ErrorHTTPVersionNotSupported sends a 505 HTTP Version Not Supported response with the provided message.
// This response indicates that the server does not support the HTTP protocol version used in the request.
func (c *Context) ErrorHTTPVersionNotSupported(msg string) {
	c.abortWithStatusJSON(http.StatusHTTPVersionNotSupported, msg)
}

// ErrorVariantAlsoNegotiates sends a 506 Variant Also Negotiates response with the provided message.
// This response indicates that the server has an internal configuration error
// where it is unable to negotiate a variant of the requested resource.
func (c *Context) ErrorVariantAlsoNegotiates(msg string) {
	c.abortWithStatusJSON(http.StatusVariantAlsoNegotiates, msg)
}

// ErrorInsufficientStorage sends a 507 Insufficient Storage response with the provided message.
// This response indicates that the server is unable to store the representation needed to complete the request.
func (c *Context) ErrorInsufficientStorage(msg string) {
	c.abortWithStatusJSON(http.StatusInsufficientStorage, msg)
}

// ErrorLoopDetected sends a 508 Loop Detected response with the provided message.
// This response indicates that the server has detected an infinite loop while processing the request.
func (c *Context) ErrorLoopDetected(msg string) {
	c.abortWithStatusJSON(http.StatusLoopDetected, msg)
}

// ErrorNotExtended sends a 510 Not Extended response with the provided message.
// This response indicates that the server requires further extensions to fulfill the request.
func (c *Context) ErrorNotExtended(msg string) {
	c.abortWithStatusJSON(http.StatusNotExtended, msg)
}

// ErrorNetworkAuthenticationRequired sends a 511 Network Authentication Required response with the provided message.
// This response indicates that the client needs to authenticate to gain network access.
func (c *Context) ErrorNetworkAuthenticationRequired(msg string) {
	c.abortWithStatusJSON(http.StatusNetworkAuthenticationRequired, msg)
}
