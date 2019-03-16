#! /usr/bin/env lua
--
-- table_size_predefined.lua
-- Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
--
-- Distributed under terms of the GPL license.
--

for i = 1, 10000000 do
  local a = {}                   -- table变量a的size为0
  a[1] = 1; a[2] = 2; a[3] = 3   -- 触发3次容量重设
end
