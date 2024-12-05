local mysplit = function(inputstr, sep)
	if sep == nil then
		sep = "%s"
	end
	local t = {}
	for str in string.gmatch(inputstr, "([^" .. sep .. "]+)") do
		table.insert(t, str)
	end
	return t
end

local rawValues = io.open("../input.txt", "r")

if rawValues ~= nil then
	---@type string
	local input = rawValues:read("*a")
	local lines = mysplit(input, "\n")

	--- @type number[], number[]
	local lefts, rights = {}, {}

	for _, line in ipairs(lines) do
		local ids = mysplit(line, "   ")

		table.insert(lefts, tonumber(ids[1]))
		table.insert(rights, tonumber(ids[2]))
	end

	table.sort(lefts)
	table.sort(rights)

	local totalDistance = 0
	for i, l in ipairs(lefts) do
		local r = rights[i]
		totalDistance = totalDistance + math.abs(l - r)
	end

	print("Total distance:", totalDistance)
end
