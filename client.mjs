import net from "net";
import fs from "fs";

const socketPath = "/tmp/ipc.sock";

// Ensure the socket exists
if (!fs.existsSync(socketPath)) {
  console.error("Socket not found, is the Go server running?");
  process.exit(1);
}

const client = net.createConnection(socketPath);

client.on("connect", () => {
  console.log("Connected to Go server");
  client.write("Hello from Node.js\n");
});

client.on("data", (data) => {
  console.log("Response from Go:", data.toString());
  client.end();
});

client.on("error", (err) => {
  console.error("Connection error:", err);
});
