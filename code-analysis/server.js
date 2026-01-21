const http = require("http");
const fs = require("fs");
const path = require("path");
const { execFile } = require("child_process");

http.createServer((req, res) => {

    if (req.url === "/") {
        fs.readFile(path.join(__dirname, "index.html"), (err, data) => {
            if (err) {
                res.writeHead(500, { "Content-Type": "text/plain" });
                res.end("Gagal membaca index.html\n" + err.toString());
                return;
            }
            res.writeHead(200, { "Content-Type": "text/html" });
            res.end(data);
        });
        return;
    }

    if (req.url.startsWith("/run")) {
        const url = new URL(req.url, "http://localhost");
        const n = url.searchParams.get("n");
        const scenario = url.searchParams.get("case");

        execFile(
            path.join(__dirname, "katalog.exe"),
            [n, scenario],
            (err, stdout, stderr) => {
                if (err) {
                    res.writeHead(500, { "Content-Type": "text/plain" });
                    res.end("Gagal menjalankan katalog.exe\n" + err.toString());
                    return;
                }
                res.writeHead(200, { "Content-Type": "application/json" });
                res.end(stdout);
            }
        );
        return;
    }

    res.writeHead(404, { "Content-Type": "text/plain" });
    res.end("404 Not Found");

}).listen(3000);

console.log("Server berjalan di http://localhost:3000");
