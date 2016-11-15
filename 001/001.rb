s = 0
for i in 0..999
    if i % 3 == 0 || i % 5 == 0
        s += i
    end
end
puts "#{s}"
