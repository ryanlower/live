$(function(){
  var ws = new WebSocket("ws://{{$}}/websocket");

  ws.onopen = function() {
    console.log("WebSocket connected...");
  }

  ws.onmessage = function(event) {
    data = JSON.parse(event.data);
    console.log("WebSocket message received...", data);
    $('ol.hits').append('<li>' + data.Code + '</li>');
  }

  ws.onclose = function() {
    console.log("WebSocket closed.");
  }
});
