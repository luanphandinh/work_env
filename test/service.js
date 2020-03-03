const http = require("http");
const name = process.env.NAME || "World";
http
  .createServer(function(req, res) {
    res.write(`Hello ${name}`);
    res.end();
  })
  .listen(process.env.PORT || 8000);
