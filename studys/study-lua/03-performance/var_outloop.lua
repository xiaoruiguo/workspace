#! /usr/bin/env lua
--
-- var_out_loop.lua
-- Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
--
-- Distributed under terms of the GPL license.
--

local t={0}
function foo ()
	for i = 1, 10000000 do
		t[1] = i
	end
end 

foo()
