object Leap {
  def leapYear(year: Int): Boolean = {
    val divisibleBy = (number: Int) => year % number == 0
    divisibleBy(4) && !divisibleBy(100) || divisibleBy(400)
  }
}
