const WebSocket = require("ws");

const wss = new WebSocket.Server({ port : 8000 });

wss.on("connection", ws => {
    console.log("New Student connected");
    ws.on("message", message => {
        try {
            const date = JSON.parse(message);

            console.log(data.x, data.y);
        } catch(e) {
            console.log('Something went wrong: ${e.message} ');
        }
//        ws.send(data.toUppercase());
    });

    ws.on("close", () => {
        console.log("Student has disconnected")
    })
});
