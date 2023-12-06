import Util.FileIO

import scala.collection.mutable

case class Num(n: Int, start: Int, end: Int)

case class NumWithY(n: Int, start: Int, end: Int, y: Int)

object Day03 {
  private val fileContent = FileIO.readFile("day03.txt")
  private val example =
    """467..114..
      |...*......
      |..35..633.
      |......#...
      |617*......
      |.....+.58.
      |..592.....
      |......755.
      |...$.*....
      |.664.598..""".stripMargin

  private def solveA(input: String): Int = {
    case class Coord(x: Int, y: Int)
    val nums = mutable.Set[NumWithY]()
    val symbolLocations = mutable.Set[Coord]()
    val lines = input.split("\n")
    lines.zipWithIndex.foreach { case (line, y) =>
      (0 until line.length).foreach(x => {
        if (!line(x).isDigit && line(x) != '.') {
          symbolLocations += Coord(y, x)
        }
        getNumber(line, x) match {
          case Some(n) => nums += NumWithY(n.n, n.start, n.end, y)
          case None =>
        }
      })
    }

    nums.map(num => {
      var found = false
      (num.start until num.end).foreach(x => {
        Array(
          Array(-1, -1),
          Array(-1, 0),
          Array(-1, 1),
          Array(1, -1),
          Array(1, 0),
          Array(1, 1),
          Array(0, -1),
          Array(0, 1),
        ).foreach { case Array(dx, dy) =>
          if (symbolLocations.contains(Coord(num.y + dy, x + dx))) {
            found = true
          }
        }
      })
      if (found) num.n else 0
    }).sum
  }

  def getNumber(s: String, i: Int): Option[Num] = {
    if (i < 0 || i >= s.length || !s(i).isDigit) {
      return None
    }
    var start = i
    while (start >= 0 && s(start).isDigit) {
      start -= 1
    }
    var end = start + 1
    while (end < s.length && s(end).isDigit) {
      end += 1
    }
    Some(Num(s.slice(start + 1, end).toInt, start + 1, end))
  }


  def exampleA(): Int = {
    solveA(example)
  }

  def a(): Int = {
    solveA(fileContent)
  }

  private def solveB(input: String): Int = {
    -1
  }

  def exampleB(): Int = {
    solveB(example)
  }

  def b(): Int = {
    solveB(fileContent)
  }
}
