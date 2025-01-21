#! /opt/homebrew/opt/lua/bin/lua
print("hello")
-- print(12)
print(b)
b = 10
print(b)

-- kv = {"a"=2,"s"=5,"l"=4}
-- for k, v in pairs(kv) do
--     print(k.."+"..v)
-- end

tab1 = { key1 = "val1", key2 = "val2", "val3" }
for k, v in pairs(tab1) do
    print(k .. " - " .. v)
end

tab1.key1 = nil
for k, v in pairs(tab1) do
    print(k .. " - " .. v)
end

print(a==nil)
print(a=="nil")

print(type(a)==nil)
print(type(a)=="nil")



function max(i, j)
    if i > j then
        print(i)
    else
        print(j)
    end
end

max(3,5)

s = "sad"
s = string.rep("asd",3)
print(s)

m = {}
print(m)
m[1] = "s"
m[2] = "a"
c = table.concat(m)
print(c, m)
print(m[1])
m = nil
print(m)

