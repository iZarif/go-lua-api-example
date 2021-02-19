print(my.sin(5))

local items = my.sqlSflow()

for i, v in ipairs(items) do
      print(i, v)
end

print(my.isItem(items[1]))
print(my.isItem(1))

print(my.getItemAddress(items[3]))