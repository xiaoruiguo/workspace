#! /usr/bin/env lua
--
-- table_size_predefined.lua
-- Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
--
-- Distributed under terms of the GPL license.
--

for i = 1, 10000000 do
  local a = {true, true, true}   -- table变量a的size在创建是确定
  a[1] = 1; a[2] = 2; a[3] = 3   -- 不会触发容量重设
end
