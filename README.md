# WebSocket Playground

A simple WebSocket playground built with Go and JavaScript to learn and experiment with WebSockets.

## Features

- Basic WebSocket echo server in Go
- Simple web interface to test WebSocket connections
- Real-time message sending and receiving
- Connection status monitoring
- Automatic reconnection

## Prerequisites

- Go 1.16 or later
- A modern web browser

## Getting Started

1. Clone the repository
2. Install dependencies:
   ```
   go mod tidy
   ```
3. Run the server:
   ```
   go run main.go
   ```
4. Open your browser and navigate to `http://localhost:8080`
5. Start sending messages through the WebSocket connection

## How It Works

- The server runs on port 8080
- WebSocket endpoint is available at `/ws`
- The web interface connects automatically when loaded
- Messages are echoed back from the server

## Project Structure

- `main.go` - The Go WebSocket server
- `static/` - Contains the web interface
  - `index.html` - The main web page with WebSocket client code

## Dependencies

- [gorilla/websocket](https://github.com/gorilla/websocket) - A fast, well-tested WebSocket implementation for Go

## License

MIT
# sockets-playground
