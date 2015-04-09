var conn = new WebSocket("ws://localhost:8080/ws/session?lastMod=13d370f412804962");
conn.onclose = function(evt) {
  data.textContent = 'Connection closed';
}
conn.onmessage = function(evt) {
  console.log('message received');
  console.log(evt.data);
}
