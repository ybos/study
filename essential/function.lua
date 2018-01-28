function avg(...) 
	local args = {...}
	
	local total = 0
	
	for k, v in pairs(args) do 
		total = total + v
		print(k, v)
	end
	
	print("----------------------------------")
	
	return total, #args, (total / #args)
end

print(avg(1, 2, 32, 4, 3245, 32, 54, 325, 432, 43, 214, 321))

