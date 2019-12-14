/*
 * app.js
 * Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
 *
 * Distributed under terms of the GPL license.
 */

{
  var a = 1;
  let b = 2;
}

console.log("a is defined: %d",a)
//console.log("b is not defined: %d",b)

for (let i = 0 ; i <  10; i++) {
    console.log(i);
}


//错误用法演示！
var array1 = [];
for (var i = 0; i < 10; i++) {
  console.log("var i value is: %d", i)
  array1[i] = function () {
    console.log(i);
  };
}
array1[6](); // 10
array1[7](); // 10

//正确用法演示！
var array2 = [];
for (let i = 0; i < 10; i++) {
  console.log("let i value is: %d", i)
  array2[i] = function () {
    console.log(i);
  };
}
array2[6](); // 6
array2[7](); // 7
