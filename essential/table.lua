local tab1 = {}
tab1["name"] = "Neil"
tab1["age"]  = 28
tab1["sex"]  = "male"

for k, v in pairs(tab1) do
	print(k .. ":" .. v)
end

print("-----------------")

local tab2 = {"apple", "grape", "orange"}
-- tab2[] = "banana"

for k, v in pairs(tab2) do 
	print(k .. ":" .. v)
end

print("-----------------")

print("table tab2 has " .. #tab2 .. " items, the last one is: " .. tab2[#tab2])

print("-----------------")

print(tab2[22])
print(tab3)
