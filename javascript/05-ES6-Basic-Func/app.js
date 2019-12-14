/*
 * app.js
 * Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
 *
 * Distributed under terms of the GPL license.
 */

function f() { console.log('I am outside!'); }

(function () {
  if (false) {
    // 重复声明一次函数f
    function f1() { console.log('I am inside!'); }
  }

  console.log("typeof f is %s", typeof f)
  f();
}());


// 参数默认值
function log(x, y = 'World') {
  console.log(x, y);
}
log('hello')

//解构赋值
function foo({x, y = 5}) {
  console.log(x, y);
}
foo({x: 1, y: 2}) 
console.log(foo.length)

//参数作用域
var x = 1;
function f(x, y = x) {
  console.log(y);
}
f(2)

//不可缺省的参数
function throwIfMissing() {
  throw new Error('Missing parameter');
}
function mustFunc(mustBeProvided = throwIfMissing()) {
  return mustBeProvided;
}

//mustFunc()   // 没有传入不可缺省的参数，报错

//不定长参数
function restFunc(...values) {
  let sum = 0;

  for (var val of values) {
    sum += val;
  }

  return sum;
}

console.log(restFunc(2, 5, 3)) // 10

function cannotStrictFunc(a, b = a) {
//  'use strict';     // 报错
  // code
}

//箭头函数
var arrowFunc1 = v => v;
var arrowFunc2 = () => 5;
var arrowFunc3 = (num1, num2) => num1 + num2;
var arrowFunc4 = (num1, num2) => { let i = 10; return i + num1 + num2};

console.log(arrowFunc1("Func1"))
console.log(arrowFunc2())
console.log(arrowFunc3(1, 2))
console.log(arrowFunc4(1, 2))

