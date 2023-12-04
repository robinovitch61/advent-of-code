import scala.collection.mutable

object Day01 {
  val fileContent = Util.readFile("day01.txt")

  private def solveA(s: String): Int = {
    s
      .split("\n")
      .map(s => s.filter(c => c.isDigit))
      .map(s =>
        s(0).toString + s(s.length - 1).toString)
      .map(s => s.toInt)
      .sum
  }

  def exampleA(): Int = {
    val example =
      """1abc2
        |pqr3stu8vwx
        |a1b2c3d4e5f
        |treb7uchet""".stripMargin

    solveA(example)
  }

  def a(): Int = {
    solveA(fileContent)
  }

  private def solveB(s: String): Int = {
    val letters = Set(
      "one",
      "two",
      "three",
      "four",
      "five",
      "six",
      "seven",
      "eight",
      "nine"
    )
    val revLetters = letters.map(s => s.reverse)
    //    val letterTrie = CharTrie.fromSet(letters)
    //    val revLetterTrie = makeTrie(revLetters)
    -1
  }

  //  // TODO: just make map
  //  case class CharTrie(c: Option[Char], children: Map[Char, CharTrie])
  //
  //  private object CharTrie {
  //    def fromSet(strings Set[String]): CharTrie = {
  //      val res CharTrie(None, Map[Char, CharTrie]())
  ////      strings.foreach(s => s.foreach(c => ))
  //      res
  //    }
  //  }

  def exampleB(): Int = {
    val example =
      """two1nine
        |eightwothree
        |abcone2threexyz
        |xtwone3four
        |4nineeightseven2
        |zoneight234
        |7pqrstsixteen""".stripMargin

    solveB(example)
  }

  def b(): Int = {
    solveB(fileContent)
  }
}
