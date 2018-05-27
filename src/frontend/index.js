console.log("Hello World")
var ws = new WebSocket("ws://" + document.location.host + "/v1/ws")

ws.addEventListener("message", function(e) {
    console.log(e)
});

setTimeout(function() {
    ws.send("foo")
}, 2000)
