const WebSocket = require('ws');

const wss = new WebSocket.Server({ port: 8000 });

wss.on('connection', function connection(ws) {
  console.log('A new client connected.');

  const stream = createStream();

  stream.on('data', function(chunk) {
    ws.send(chunk);
  });

  ws.on('close', function() {
    console.log('Client disconnected.');
    stream.destroy();
  });
});

function createStream() {
  const { spawn } = require('child_process');
  const args = ['-re', '-i', 'video.mp4', '-f', 'mpegts', '-codec:v', 'mpeg1video', '-s', '640x480', '-b:v', '800k', '-r', '30', '-'];

  const ffmpeg = spawn('ffmpeg', args, { stdio: ['ignore', 'pipe', 'ignore'] });

  return ffmpeg.stdout;
}
