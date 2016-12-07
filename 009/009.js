function tripletProduct() {
  const s = 1000
  const r = Math.floor(Math.sqrt(s))
  for (let m = r; m > 1; m--)
    for (let n = m; n > 2; n--) {
      const a = m * m - n * n
      const b = 2 * (m * n)
      const c = m * m + n * n
      if (a + b + c === s) {
        console.log(a * b * c)
        return
      }
    }
}

tripletProduct()
