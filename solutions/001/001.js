function sum(n, k) {
  return Math.trunc(k * (Math.trunc((n - 1) / k) * (Math.trunc((n - 1) / k) + 1)) / 2)
}

const solution = sum(1000, 3) + sum(1000, 5) - sum(1000, 15)
console.log(solution)
