i = 10
while (i >= 0) do
	print(i)
	
	i = i - 3
end

print("now i: " .. i)

code = [[

i = 10
while (i <= 0) do
	print(i)
	
	i = i - 3
end
]]

print(code)