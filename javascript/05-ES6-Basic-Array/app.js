/*
 * app.js
 * Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
 *
 * Distributed under terms of the GPL license.
 */

console.log([1,2,3])
console.log(...[1,2,3])


//数组复制
let a1 = [1, 2];
let a2 = a1          // 指向同一个数组
a2[1] = 22

let a3 = [...a1];    // 复制数组，写法1
a3[0] = 21

let [...a4] = a1;    // 复制数组，写法2
a4[0] = 31

console.log(a1)
console.log(a2)
console.log(a3)
console.log(a4)
