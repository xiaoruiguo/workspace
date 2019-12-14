/*
 * app.js
 * Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
 *
 * Distributed under terms of the GPL license.
 */

var gvar1 = "is global.gvar1";

var gFunc = () =>{
  console.log("is global.gFunc");
}

let gvar2 = "isn't global.gvar2"

const gvar3 = "isn't global.gvar3"

console.log(this.gvar2)
console.log(global)
