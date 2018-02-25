object SumOfMultiples {
  def sum(factors: Set[Int], limit: Int): Int = {
    factors.flatMap((factor) => { for (number <- 1 to limit; if number * factor < limit) yield number * factor })
           .sum
  }
}
