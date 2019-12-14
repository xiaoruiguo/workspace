/*
 * app.js
 * Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
 *
 * Distributed under terms of the GPL license.
 */

async function asyncFunc1(){
  await console.log('asyncFunc content 1')
  await console.log('asyncFunc content 2')
  return 'asyncFunc1 return'
}

asyncFunc1().then(v => console.log(v))
