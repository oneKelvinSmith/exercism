object Bob {
  case class Remark(statement: String)

  object Shout {
    def unapply(remark: Remark): Option[String] = {
      val isShouting = remark.statement.toUpperCase() == remark.statement
      if (isShouting) {
        Some("Whoa, chill out!")
      } else {
        None
      }
    }
  }

  object Question {
    def unapply(remark: Remark): Option[String] = {
      if (remark.statement.endsWith("?")) {
        Some("Sure.")
      } else {
        None
      }
    }
  }

  def response(statement: String): String = Remark(statement) match {
    case Shout(reply)    => reply
    case Question(reply) => reply
    case _               => "Whatever."
  }
}
