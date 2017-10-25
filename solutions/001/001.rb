def sum(n, k)
	return k * ((((n - 1) / k) * (((n - 1) / k) + 1)) / 2)
end

puts "#{sum(1000, 3) + sum(1000, 5) - sum(1000, 15)}"
