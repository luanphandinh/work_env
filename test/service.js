const http = require("http");
const name = process.env.NAME || "World";
const say = process.env.SAY || "Hello";

http
  .createServer(function(req, res) {
    res.write(`${say} ${name}`);
    res.end();
  })
  .listen(process.env.PORT || 8000);
