object Bob {
  case class Remark(statement: String) {
    def isShouting = {
      statement.indexWhere(char => char.isLetter) >= 0 &&
        statement.toUpperCase == statement
    }
    def isQuestioning = statement.endsWith("?")
  }

  object Shout {
    def unapply(remark: Remark): Option[String] = {
      if (remark.isShouting) {
        Some("Whoa, chill out!")
      } else {
        None
      }
    }
  }

  object Question {
    def unapply(remark: Remark): Option[String] = {
      if (remark.isQuestioning) {
        Some("Sure.")
      } else {
        None
      }
    }
  }

  object ForcefulQuestion {
    def unapply(remark: Remark): Option[String] = {
      if (remark.isQuestioning && remark.isShouting) {
        Some("Calm down, I know what I'm doing!")
      } else {
        None
      }
    }
  }

  def response(statement: String): String = Remark(statement) match {
    case ForcefulQuestion(reply) => reply
    case Question(reply)         => reply
    case Shout(reply)            => reply
    case _                       => "Whatever."
  }
}
