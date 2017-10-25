fibo = [0, 1]
for i in 2..34
    fibo << fibo[i - 2] + fibo[i - 1]
end
fibo = fibo.select { |x| x % 2 == 0 }
s = 0
fibo.each do |x|
    s += x
end
puts "#{s}"
