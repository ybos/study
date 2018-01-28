i = 1
repeat
	print("first: " .. i)
	i = i + 2
	print("second: " .. i)
until(i > 11)

print("now i: " .. i)

code = [[

i = 1
repeat
	print("first: " .. i)
	i = i + 2
	print("second: " .. i)
until(i > 11)
]]

print(code)