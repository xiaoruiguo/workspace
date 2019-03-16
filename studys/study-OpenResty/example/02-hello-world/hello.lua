#! /usr/bin/env lua
--
-- hello.lua
-- Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
--
-- Distributed under terms of the GPL license.
--
mysql=require "resty.mysql"
ngx.say("hello world")
