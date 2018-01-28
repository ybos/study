str1 = "string to upper"
print(str1)
str1 = string.upper(str1)
print(str1)

print("---------------------------")

str2 = "String to Lower"
print(str2)
str2 = string.lower(str2)
print(str2)

print("---------------------------")

str3 = "find Neil in this string, Neil is a name"
print(str3)
str3_start, str3_end = string.find(str3, "Neil")
print([[str3_start, str3_end = string.find(str3, "Neil")]])
print(str3_start, str3_end)

str3_start, str3_end = string.find(str3, "Neil", 20)
print([[str3_start, str3_end = string.find(str3, "Neil", 20)]])
print(str3_start, str3_end)
print("third param use for start index")

print("---------------------------")

str4 = "please replace this string"
print(str4)
str4 = string.gsub(str4, "ea", "b")
print([[str4 = string.gsub(str4, "ea", "b")]])
print(str4);

print("---------------------------")

str5 = "Can you reverse this string?"
print(str5)
str5 = string.reverse(str5)
print([[str5 = string.reverse(str5)]])
print(str5)

print("---------------------------")

str6 = "Format a string"
print([[string.format("the value is:%d",4)]])
print("it likes printf")

print("---------------------------")

str7 = "Calculate this string's length"
print(str7)
print([[string.len(str7)]])
print(string.len(str7))

print("---------------------------")

str8 = "Repeat this string?"
print(str8)
print([[string.rep(str8, 2)]])
print(string.rep(str8, 2))

print("---------------------------")

str9 = "Could you tell me how many e in this string?"
print(str9)
for character in string.gmatch(str9, "e") do
	print(character)
end
print([[for character in string.gmatch(str9, "e") do
	print(character)
end]]

print("---------------------------")

str10 = "match as like as gmatch, but match is only once."
print(str10)
)