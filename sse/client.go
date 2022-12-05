package main

// This example creates a client that uses Server-Sent Events (SSE) to receive real-time updates from a server.
// In the Go language, we will use the net/http package to create the SSE client. We will create an http.Client and create the SSE request. We will also register a response handler to handle the incoming events.
// First, we need to create the http.Client and set its request headers for SSE. We can do this by using the DefaultTransport and headers:
