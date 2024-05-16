const express = require('express');
const os = require('os');

const app = express();
const port = 3000;

app.use(express.static('public'));

app.get('/api/info', (req, res) => {
  const hostname = os.hostname();
  const ip = getIPAddress();

  res.json({ hostname, ip });
});

app.listen(port, () => {
  console.log(`Server is running on http://localhost:${port}`);
});

function getIPAddress() {
  const interfaces = os.networkInterfaces();
  for (const key in interfaces) {
    const iface = interfaces[key];
    for (let i = 0; i < iface.length; i++) {
      const { family, address, internal } = iface[i];
      if (family === 'IPv4' && !internal) {
        return address;
      }
    }
  }
  return 'N/A';
}
