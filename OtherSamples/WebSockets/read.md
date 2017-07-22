# WebSockets

This folder is code as the tutorial:


## At browser:

For V1, create a socket with the browser's console:
``` javascript
var websocket = new WebSocket("ws://localhost:8070/v1/ws")

// Event listener
websocket.addEventListener("message", function(e){
  console.log(e)
})

// Message to send
websocket.send("Write a message to send")
```
