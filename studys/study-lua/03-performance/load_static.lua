#! /usr/bin/env lua
--
-- load.lua
-- Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
--
-- Distributed under terms of the GPL license.
--

function fk (k)
	return function () return k end
end

local lim = 10000000
local a = {}
for i = 1, lim do
	a[i] = fk(i)
end

print(a[10]())  --> 10
