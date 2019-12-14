/*
 * async.js
 * Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
 *
 * Distributed under terms of the GPL license.
 */
const fs = require('fs');
fs.readFile('./file.md', (err, data) => {
  if (err) throw err;
  console.log(data);
});
console.log("finish") // will run before console.log
