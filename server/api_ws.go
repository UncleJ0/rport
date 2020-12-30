package chserver

import (
	"html/template"
	"net/http"
)

func (al *APIListener) home(w http.ResponseWriter, r *http.Request) {
	p := "ws://"
	if al.config.API.CertFile != "" && al.config.API.KeyFile != "" {
		p = "wss://"
	}
	_ = homeTemplate.Execute(w, p+r.Host+"/api/v1/test/commands/ws")
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>
window.addEventListener("load", function(evt) {
   var output = document.getElementById("output");
   var input = document.getElementById("input");
   var ws;
   var print = function(message) {
       var d = document.createElement("div");
       d.textContent = message;
       output.appendChild(d);
   };
   document.getElementById("open").onclick = function(evt) {
       if (ws) {
           return false;
       }
       ws = new WebSocket("{{.}}");
       ws.onopen = function(evt) {
           print("OPEN");
       }
       ws.onclose = function(evt) {
           print("CLOSE");
           ws = null;
       }
       ws.onmessage = function(evt) {
           print("RESPONSE: " + evt.data);
       }
       ws.onerror = function(evt) {
           print("ERROR: " + evt.data);
       }
       return false;
   };
   document.getElementById("send").onclick = function(evt) {
       if (!ws) {
           return false;
       }
       print("SEND: " + input.value);
       ws.send(input.value);
       return false;
   };
   document.getElementById("close").onclick = function(evt) {
       if (!ws) {
           return false;
       }
       ws.close();
       return false;
   };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server,
"Send" to send a message to the server and "Close" to close the connection.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p>
<textarea id="input" rows="5" cols="60" placeholder="Enter JSON request here...">
</textarea>
<p>
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<pre id="output"></pre>
</td></tr></table>
</body>
</html>
`))