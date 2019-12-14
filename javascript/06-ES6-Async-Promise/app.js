/*
 * app.js
 * Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
 *
 * Distributed under terms of the GPL license.
 */

// async success
let promise1 = new Promise(function(resolve, reject){
  /* 处理代码 */
  let result = true
  let value = "promise1 success"
  let error = "promise1 fail"
  if(result){
    resolve(value);
  }else{
    reject(error);
  }
});

promise1.then(function(value){
  console.log(value)
},function(error){
  console.log(error)
});

// async fail
let promise2 = new Promise(function(resolve, reject){
  /* 处理代码 */
  let result = false
  let value = "promise2 success"
  let error = "promise2 fail"
  if(result){
    resolve(value);
  }else{
    reject(error);
  }
});

promise2.then(function(value){
  console.log(value)
},function(error){
  console.log(error)
});

// Promise.prototype.then
let promise3 = new Promise(function(resolve, reject){
  /* 处理代码 */
  let result = true
  let value = "promise1 success"
  let error = "promise1 fail"
  if(result){
    resolve(value);
  }else{
    reject(error);
  }
});

promise3.then(function(value){
  return "first then"
},function(error){
  console.log(error)
}).then(function(str){
  console.log("resolved: print first then return: %s",str)
});

