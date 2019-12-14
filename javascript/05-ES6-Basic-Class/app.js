/*
 * app.js
 * Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
 *
 * Distributed under terms of the GPL license.
 */
//类的声明和实例的构造
class Point {
  constructor(x, y) {
    this.x = x;
    this.y = y;
  }

  toString() {
    return '(' + this.x + ', ' + this.y + ')';
  }
}

let ins1 = new Point(1,3)
console.log(ins1)
console.log(ins1.toString())

//类属性的存取器
class GetSet {
  constructor(){}
  get A(){
    console.log('get for A is called');
    return 'This is A';
  }
  set A(value){
    console.log('set for A is called, value is %s', value);
  }
}

let ins2 = new GetSet();
ins2.A               // get A() 被调用
ins2.A = "abc"       // set A() 被调用

//类的迭代器方法
class Iter {
  constructor(num){
    this.num = num;
  }
  * next() {
    for (let i = 0; i < this.num ; i++) {
      yield i;
    }
  }
}

let ins3 = new Iter(4)
console.log(ins3.next())
for (let x of ins3.next()) {
  console.log(x)
}

//类的静态方法
class Static {
  constructor (){}
  static staticFunc(){
    return 'Static';
  }
}

let ins4 = new Static()
console.log(Static.staticFunc())
//console.log(ins4.staticFunc())   // 报错：TypeError: ins4.staticFunc is not a function

// 类的继承
class Father {
  constructor (){}
  static classMethod1(){
    return 'Father class method 1';
  }
  static classMethod2(){
    return 'Father class method 2';
  }
}

class Child extends Father {
  constructor(){
    super()        // 子类的构造函数必须调用 super()
  }
  static classMethod2(){
    return super.classMethod2() + ', Child class method 2';
  }
}

console.log(Child.classMethod1())
console.log(Child.classMethod2())

// 类的属性
class Page{
  static classAttr = 'This is class Page'
  name = 'lijiaocn';
  constructor(title){
    this.title = title
  }
}

let ins5 = new Page('First Page')
console.log(ins5)

console.log(Page.classAttr)
