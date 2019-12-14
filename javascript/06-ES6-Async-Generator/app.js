/*
 * app.js
 * Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
 *
 * Distributed under terms of the GPL license.
 */

function* Generator1() {
  yield 'generator content 1';
  yield 'generator content 2';
  return 'ending';
}

let iter1 = Generator1()
console.log(iter1.next())
console.log(iter1.next())
console.log(iter1.next())


let co = require('co');
function* Generator2() {
  console.log('generator content 1');
  console.log('generator content 2');
  return 'hello'
}
iter2 = Generator2()
co(iter2).then(function(v){
  console.log(v)
}).catch(function(err){
  console.log(err)
})
