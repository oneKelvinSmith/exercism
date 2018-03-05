object Bob {
  case class Remark(statement: String) {
    def isShouting = {
      statement.indexWhere(char => char.isLetter) >= 0 &&
        statement.toUpperCase == statement
    }
    def isQuestioning = statement.endsWith("?")
    def isSilence     = statement.isEmpty
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

  object Silence {
    def unapply(remark: Remark): Option[String] = {
      if (remark.isSilence) {
        Some("Fine. Be that way!")
      } else {
        None
      }
    }
  }

  def response(statement: String): String = Remark(statement.trim) match {
    case ForcefulQuestion(reply) => reply
    case Question(reply)         => reply
    case Shout(reply)            => reply
    case Silence(reply)          => reply
    case _                       => "Whatever."
  }
}
