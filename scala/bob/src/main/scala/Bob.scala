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
      val isQuestioning = remark.statement.endsWith("?")
      if (isQuestioning) {
        Some("Sure.")
      } else {
        None
      }
    }
  }

  object ForcefulQuestion {
    def unapply(remark: Remark): Option[String] = {
      val isShouting = remark.statement.toUpperCase() == remark.statement
      val isQuestioning = remark.statement.endsWith("?")

      if (isQuestioning && isShouting) {
        Some("Calm down, I know what I'm doing!")
      } else {
        None
      }
    }
  }

  def response(statement: String): String = Remark(statement) match {
    case ForcefulQuestion(reply) => reply
    case Shout(reply)            => reply
    case Question(reply)         => reply
    case _                       => "Whatever."
  }
}
