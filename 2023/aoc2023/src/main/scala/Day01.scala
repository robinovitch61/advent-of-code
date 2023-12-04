import Util.{CharTrie, FileIO}

object Day01 {
  val fileContent = FileIO.readFile("day01.txt")

  private def solveA(input: String): Int = {
    input
      .split("\n")
      .map(s => s.filter(c => c.isDigit))
      .map(s =>
        s(0).toString + s(s.length - 1).toString)
      .map(_.toInt)
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

  private def solveB(input: String): Int = {
    val letters = Set(
      "one",
      "two",
      "three",
      "four",
      "five",
      "six",
      "seven",
      "eight",
      "nine",
    ).union("123456789".map(_.toString).toSet)
    val revLetters = letters.map(s => s.reverse)
    val letterTrie = CharTrie.fromStrings(letters)
    val revLetterTrie = CharTrie.fromStrings(revLetters)
    input
      .split("\n")
      .map(s =>
        findDigit(s, letterTrie).toString + findDigit(s.reverse, revLetterTrie).toString)
      .map(_.toInt)
      .sum
  }

  private def findDigit(s: String, trie: CharTrie): Int = {
    0 until s.length foreach (idx => {
      trie.containsPrefix(s.slice(idx, s.length)) match {
        case Some(t) => {
          Array(t, t.reverse).foreach {
            case d if d.matches("\\d") => return d.toInt
            case "one" => return 1
            case "two" => return 2
            case "three" => return 3
            case "four" => return 4
            case "five" => return 5
            case "six" => return 6
            case "seven" => return 7
            case "eight" => return 8
            case "nine" => return 9
            case _ =>
          }
        }
        case _ =>
      }
    })
    throw new RuntimeException(s"no digit found for $s")
  }

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
