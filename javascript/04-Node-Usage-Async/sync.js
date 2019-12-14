/*
 * sync.js
 * Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
 *
 * Distributed under terms of the GPL license.
 */

const fs = require('fs');
const data = fs.readFileSync('./file.md'); // blocks here until file is read
console.log(data);
console.log("finish"); // will run after console.log
