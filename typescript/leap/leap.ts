function isLeapYear(year: number): boolean {
  const divisibleBy = (base: number, divisor: number) => {
    return base % divisor === 0
  }

  return divisibleBy(year, 400) || !divisibleBy(year, 100) && divisibleBy(year, 4)
}

export default isLeapYear