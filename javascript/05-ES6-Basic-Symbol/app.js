/*
 * app.js
 * Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
 *
 * Distributed under terms of the GPL license.
 */

let s = Symbol();
console.log(typeof s);

let s1 = Symbol();
let s2 = Symbol();
if (s1 === s2) {
  console.log("equal");
}else{
  console.log("not equal");  // 输出 not equal
}


let s3 = Symbol('this is symbol3')
console.log(s3.description)
console.log(typeof s3)
console.log(String(s3))
