# Node.js ↔ Go IPC Example (Unix Domain Sockets)

This project demonstrates a minimal IPC (Inter‑Process
Communication) channel between **Node.js** and **Golang** using **Unix
Domain Sockets**.

## Overview

Both processes communicate by **message passing** over a local socket
file (`/tmp/ipc.sock`).\
Messages are simple newline‑terminated strings.

### Why Unix Domain Sockets?

-   Faster than TCP (0.3--1 µs latency)
-   No network overhead
-   Bidirectional
-   Perfect for processes on the same machine

### Why Message Passing?

-   No shared memory, no race conditions
-   Easy to serialize (text, JSON, binary)
-   Idiomatic for both Node.js and Go
-   Works exactly like Docker, systemd, containerd communication

------------------------------------------------------------------------

## Project Structure

    server.go   # Go server that listens on /tmp/ipc.sock
    client.js   # Node.js client that connects and sends a message

------------------------------------------------------------------------

## How It Works

### 1. Go Server

-   Creates `/tmp/ipc.sock`
-   Accepts connections
-   Reads incoming messages
-   Replies with `"Go received: <message>"`

### 2. Node.js Client

-   Connects to the Unix socket
-   Sends `"Hello from Node.js"`
-   Reads the reply
-   Closes the connection

This demonstrates full-duplex IPC.

------------------------------------------------------------------------

## Running the Example

### Start the Go server

``` bash
go run server.go
```

### Run the Node client

``` bash
node client.mjs
```

Expected output:

**Node.js**

    Connected to Go server
    Response from Go: Go received: Hello from Node.js

**Go**

    Received from Node: Hello from Node.js
