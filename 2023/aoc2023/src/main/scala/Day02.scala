import Util.FileIO

import scala.collection.mutable

object Day02 {
  private val fileContent = FileIO.readFile("day02.txt")
  private val example =
    """Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
      |Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
      |Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
      |Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
      |Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green""".stripMargin


  private def solveA(input: String): Int = {
    input.split("\n").map(line => {
      line.split(": ") match {
        case Array(gameString, roundString) =>
          val gameID = gameString.replace("Game ", "").toInt
          val maxTileCounts = getMaxTileCounts(roundString)
          if (maxTileCounts.getOrElse("red", 0) <= 12 && maxTileCounts.getOrElse("green", 0) <= 13 && maxTileCounts.getOrElse("blue", 0) <= 14) gameID else 0
        case _ => throw new RuntimeException(s"invalid line $line")
      }
    }).sum
  }

  private def getMaxTileCounts(roundString: String): mutable.Map[String, Int] = {
    val rounds: Array[String] = roundString.split("; ")
    val maxTileCounts = mutable.Map[String, Int]()
    rounds.foreach(round => {
      val colors = round.split(", ")
      colors.foreach(color => {
        color.split(" ") match {
          case Array(count, c) => maxTileCounts.update(c, Math.max(count.toInt, maxTileCounts.getOrElse(c, 0)))
          case _ => throw new RuntimeException(s"invalid color $color")
        }
      })
    })
    maxTileCounts
  }

  def exampleA(): Int = {
    solveA(example)
  }

  def a(): Int = {
    solveA(fileContent)
  }

  private def solveB(input: String): Int = {
    input.split("\n").map(line => {
      line.split(": ") match {
        case Array(_, roundString) =>
          val maxTileCounts = getMaxTileCounts(roundString)
          maxTileCounts("red") * maxTileCounts("blue") * maxTileCounts("green")
        case _ => throw new RuntimeException(s"invalid line $line")
      }
    }).sum
  }

  def exampleB(): Int = {
    solveB(example)
  }

  def b(): Int = {
    solveB(fileContent)
  }
}
