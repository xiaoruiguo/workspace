#! /usr/bin/env lua
--
-- print.lua
-- Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
--
-- Distributed under terms of the GPL license.
--


local number=0.2
local string="this is a string"
local table = {first="value1",second="value2"}
local array = {"a","b","c"}

print("parameter is nil(", type(nil),")",nil)
print("parameter is boolen(", type(true),")",true)
print("parameter is number(", type(number),")", number)
print("parameter is string(", type(string),")", string)
print("parameter is function(", type(print),")",print)
print("parameter is userdata(", type(io.stdin),")", io.stdin)
print("parameter is table(", type(table),")", table)
print("parameter is table(", type(table),")", tostring(table))
print("parameter is array(", type(array),")", array)

print(table[1], table.first)
print("parameter is table(", type(table),")", table)
print("parameter is array(", type(array),")", array)
print(array[1],array.a)

print("--- iterate table by pairs  ")
for i, v in pairs(table) do
    print(i ,v)
end
print("--- iterate table by ipairs  ")
for i, v in ipairs(table) do
    print(i ,v)
end

print("--- iterate array by pairs ")
for i, v in pairs(array) do
    print(i ,v)
end
print("--- iterate array by ipairs ")
for i, v in ipairs(array) do
    print(i ,v)
end

print("-- format output")
v = string.format('%q', 'a string with "quotes" and \n new line')
print(v)


