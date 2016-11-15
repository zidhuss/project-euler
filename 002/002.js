const fibo = [0, 1]
for (let i = 2; i < 34; i++)
  fibo.push(fibo[i - 2] + fibo[i - 1])
console.log(fibo.filter(x => x % 2 === 0).reduce((a, b) => a + b, 0))
