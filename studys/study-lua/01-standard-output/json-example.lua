local json = require "json"

local table={first="value1", second="value2"}
v = json.encode(table)
print(v)

local array={"a","b","c",1,2,3}
v = json.encode(array)
print(v)

