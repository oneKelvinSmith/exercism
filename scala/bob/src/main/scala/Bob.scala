object Bob {
  case class Remark(statement: String) {
    def isQuestioning: Boolean = statement.endsWith("?")
    def isShouting: Boolean = {
      statement.indexWhere(char => char.isLetter) >= 0 &&
        statement.toUpperCase == statement
    }
    def isSilence: Boolean = statement.isEmpty
  }

  object ForcefulQuestion {
    def unapply(remark: Remark): Boolean = remark.isQuestioning && remark.isShouting
  }

  object Question {
    def unapply(remark: Remark): Boolean = remark.isQuestioning
  }

  object Shout {
    def unapply(remark: Remark): Boolean = remark.isShouting
  }

  object Silence {
    def unapply(remark: Remark): Boolean = remark.isSilence
  }

  def response(statement: String): String = Remark(statement.trim) match {
    case ForcefulQuestion() => "Calm down, I know what I'm doing!"
    case Question()         => "Sure."
    case Shout()            => "Whoa, chill out!"
    case Silence()          => "Fine. Be that way!"
    case _                  => "Whatever."
  }
}
