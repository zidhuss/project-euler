function isPalindrome(toCheck) {
  return String(toCheck).split('').reverse().join('') === String(toCheck)
}

function largestPalindromeProduct(last, first) {
  let largest = 0
  for (let a = last; a >= first; a--)
    for (let b = last; b >= first; b--) {
      const calculated = a * b
      if (calculated > largest && isPalindrome(calculated))
        largest = calculated
    }
  return largest
}

console.log(largestPalindromeProduct(999, 100))
