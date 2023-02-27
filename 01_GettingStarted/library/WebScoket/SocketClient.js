const socket = new WebSocket('ws://localhost:8080');

socket.onopen = function(event) {
  console.log('WebSocket connection established');
};

socket.onerror = function(error) {
  console.error('WebSocket error: ' + error);
};

socket.onmessage = function(event) {
  console.log('Message received from server: ' + event.data);
};

socket.onclose = function(event) {
  console.log('WebSocket connection closed');
};
