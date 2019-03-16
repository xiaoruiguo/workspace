#! /usr/bin/env lua
--
-- table_shrink.lua
-- Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
--
-- Distributed under terms of the GPL license.
--

a = {}
lim = 10000000

for i = 1, lim do a[i] = i end            --  create a huge table
print(collectgarbage("count"))            --> 196626

for i = 1, lim do a[i] = nil end          --  erase all its elements
print(collectgarbage("count"))            --> 196626，不会收缩

for i = lim + 1, 2*lim do a[i] = nil end  --  create many nil element
print(collectgarbage("count"))            --> 17，添加足够多nil之后才会触发rehash
