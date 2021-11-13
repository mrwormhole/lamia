window.onload = reloadBrowser();

// This function is used for browser reloading.
function reloadBrowser() {
    let ws = new WebSocket("ws://localhost:8080/reload-browser");
    ws.addEventListener("open", function (event) {
        console.log("Connected to server!");
        ws.send("Frontend is refreshed!");
    });

    ws.addEventListener("close", function(evt) {
        console.log("Disconnected from the server!");
        let reloadTimes = 0;
        setInterval(function() {
            if (reloadTimes < 3) {
                reloadTimes++;
                location.reload();
            } else {
                console.log("Reload times is over 3! Use F5 button now!");
                clearInterval(this);
            }
            console.log("Reload times: " + reloadTimes);
        }, 2000);
    });
}