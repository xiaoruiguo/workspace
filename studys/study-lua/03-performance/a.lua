#! /usr/bin/env lua
--
-- a.lua
-- Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
--
-- Distributed under terms of the GPL license.
--

function memoize (f)
	local mem = {}                       -- memoizing table
	setmetatable(mem, {__mode = "kv"})   -- make it weak
	return function (x)       -- new version of ’f’, with memoizing
		local r = mem[x]
		if r == nil then    -- no previous result?
			r = f(x)          -- calls original function
			mem[x] = r        -- store result for reuse
		end
		return r
	end
end
