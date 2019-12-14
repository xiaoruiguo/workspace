/*
 * app.js
 * Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
 *
 * Distributed under terms of the GPL license.
 */

// 对象的声明 1
let object1 = {
  x: 1,
  y: 2,
  sum: function(){
    return this.x + this.y
  }
}

console.log(object1.sum())


// 对象的声明 2
let [x, y]  = [11, 12]
let object2 = {
  x,
  y,
  sum(){
    return this.x + this.y
  }
}

console.log(object2.sum())

// 属性表达式
const keyA = {a: 1};
const keyB = {b: 2};

const object3 = {
  [keyA]: 'valueA',
  [keyB]: 'valueB'
};

console.log(object3[keyA])   // 输出 valueB
console.log(object3[keyB])   // 输出 valueB


// super
const proto = {
  foo: 'hello'
};

const obj = {
  foo: 'world',
  find() {
    return super.foo;
  }
};

Object.setPrototypeOf(obj, proto);
console.log(obj.find()) // "hello"


// ...
let z = { a:3, b:4 }
let n = { ...z };
console.log(n)  // { a: 3, b: 4 }
