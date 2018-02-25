object SumOfMultiples {
  def sum(factors: Set[Int], limit: Int): Int = {
    val multiples = for  (factor     <- factors;
                          multiplier <- 1 to limit;
                          multiple = multiplier * factor;
                    if    multiple < limit)
                    yield multiple

    multiples.sum
  }
}
