function isLeapYear(year: number): boolean {
  const divisibleBy = (base: number, divisor: number) => {
    return !(base % divisor)
  }

  return (divisibleBy(year, 4) && !divisibleBy(year, 100)) || divisibleBy(year, 400)
}

export default isLeapYear