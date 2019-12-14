/*
 * app.js
 * Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
 *
 * Distributed under terms of the GPL license.
 */

const PI=3.1415;
//PI = 3;         // 修改常量值，报错
//const foo;      // 常量声明时没有初始化，报错

const array = [];

//array = [];       //更改常量的指向，报错
array[0] = 1;       //更改常量所指的对象的值，不报错

{
const array2 = Object.freeze([]);
//array2= 1;        //报错，对象被冻结不能修改
}

{
let author = {}

author.name = "lijiaocn"
author.age = 30

var constantize = (obj) => {
  Object.freeze(obj);
  Object.keys(obj).forEach( (key, i) => {
    if ( typeof obj[key] === 'object' ) {
      constantize( obj[key] );
    }
  });
};

constantize(author)
author.age = 32     // 修改不生效
author.grade= 3     // 修改不生效
console.log(author)
}
