local hello = require("hello")

print(hello.sin(5))

local items = hello.sqlSflow("statement")

for i, v in ipairs(items) do
      print(i, v)
end

print(hello.isItem(items[1]))
print(hello.isItem(1))

print(items[3].address)
