import Util.{CharTrie, FileIO}

object Day02 {
  private val fileContent = FileIO.readFile("day02.txt")

  private def solveA(input: String): Int = {
    input.split("\n").foreach(line => {
      line.split(": ") match {
        case Array(gameID, rounds) => {
          println(gameID)
          println(rounds)
        }
        case _ => throw new RuntimeException(s"invalid line $line")
      }
      return -1
    })
    -1
  }

  def exampleA(): Int = {
    val example =
      """Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
        |Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
        |Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
        |Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
        |Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green""".stripMargin

    solveA(example)
  }

  def a(): Int = {
    solveA(fileContent)
  }

  private def solveB(input: String): Int = {
    -1
  }

  def exampleB(): Int = {
    val example =
      """Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
        |Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
        |Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
        |Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
        |Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green""".stripMargin

    solveB(example)
  }

  def b(): Int = {
    solveB(fileContent)
  }
}
