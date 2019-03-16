print("-- single capature")
s = "hello world from Lua"
for w in string.gmatch(s, "%a+") do
    print(w)
end

print("-- multi capature")
s = "from=world, to=Lua, abc=1"
for k, v in string.gmatch(s, "(%w+)=(%w+)") do
    print(k,v)
end

print("-- match at pos 5")
s= "123456789"
v=string.match(s,"%d*", 5)
print(v)

print("-- repeat each words")
x = string.gsub("hello world", "(%w+)", "%1 , %1 |")
print(x)

print("-- repeat first words")
x = string.gsub("hello world", "(%w+)", "%1 , %1 |",1 )
print(x)

print("-- repel is a function")
x = string.gsub("home = $HOME, user = $USER", "%$(%w+)", os.getenv)
print(x)

print("-- repel is table")
local t = {name="lua", version="5.3"}
x = string.gsub("$name-$version.tar.gz", "%$(%w+)", t)
print(x)

print("-- sub string 0-3, 1-3, 2-3")
s="123456789"
v=string.sub(s,0,3)
print(v)
v=string.sub(s,1,3)
print(v)
v=string.sub(s,2,3)
print(v)
v=string.sub(s,2,-1)
print(v)
v=string.sub(s,-2)
print(v)
