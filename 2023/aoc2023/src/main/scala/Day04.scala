import Util.FileIO

import scala.collection.mutable

object Day04 {
  private val fileContent = FileIO.readFile("day04.txt")
  private val example =
    """Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11""".stripMargin

  private case class Card(n: Int, winning: Set[Int], has: Set[Int]) {
    def matching(): Int = {
      winning.intersect(has).size
    }

    def score(): Int = {
      var score = 0
      (0 until matching()).foreach(_ => {
        if (score == 0) score = 1 else score *= 2
      })
      score
    }
  }

  private object Card {
    def fromLine(line: String): Card = {
      line.split(": ") match {
        case Array(id: String, nums: String) =>
          nums.split(" \\| ") match {
            case Array(win: String, has: String) => Card(
              n = id.replace("Card ", "").replace(" ", "").toInt,
              winning = win.split("\\s+").filter(_.nonEmpty).map(_.toInt).toSet,
              has = has.split(" ").filter(_.nonEmpty).map(_.toInt).toSet,
            )
            case _ => throw new RuntimeException(s"invalid nums $nums")
          }
        case _ => throw new RuntimeException(s"invalid line $line")
      }
    }
  }

  private def solveA(input: String): Int = {
    input.split("\n").map(Card.fromLine(_).score()).sum
  }

  def exampleA(): Int = {
    solveA(example)
  }

  def a(): Int = {
    solveA(fileContent)
  }

  private def solveB(input: String): Int = {
    val cards = input.split("\n").map(Card.fromLine)
    val cardIdToCount = mutable.Map[Int, Int]().withDefaultValue(1)
    cards.foreach(card => {
      (0 until cardIdToCount(card.n)).foreach(_ => {
        (card.n + 1 until card.n + card.matching() + 1).foreach(n => {
          cardIdToCount.update(n, cardIdToCount(n) + 1)
        })
      })
    })
    cards.map(c => cardIdToCount(c.n)).sum
  }

  def exampleB(): Int = {
    solveB(example)
  }

  def b(): Int = {
    solveB(fileContent)
  }
}
