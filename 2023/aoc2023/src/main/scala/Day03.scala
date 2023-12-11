import Util.{Coord, FileIO, RangeNum, RangeNumWithY}

import scala.collection.mutable

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
    val nums = mutable.Set[RangeNumWithY]()
    val symbolLocations = mutable.Set[Coord]()
    val lines = input.split("\n")
    lines.zipWithIndex.foreach { case (line, y) =>
      (0 until line.length).foreach(x => {
        if (!line(x).isDigit && line(x) != '.') {
          symbolLocations += Coord(x, y)
        }
        getNumber(line, x) match {
          case Some(n) => nums += RangeNumWithY(n.n, n.start, n.end, y)
          case None =>
        }
      })
    }

    nums.toArray.map(num => {
      if ((num.start until num.end).exists(x => {
        Array(-1, 0, 1).exists(dx => {
          Array(-1, 0, 1).exists(dy =>
            symbolLocations.contains(Coord(x + dx, num.y + dy))
          )
        })
      })) num.n else 0
    }).sum
  }

  def exampleA(): Int = {
    solveA(example)
  }

  def a(): Int = {
    solveA(fileContent)
  }

  private def solveB(input: String): Int = {
    val starNums = mutable.ArrayBuffer[mutable.Set[RangeNum]]()
    val lines = input.split("\n")
    lines.zipWithIndex.foreach { case (line, y) =>
      (0 until line.length).foreach(x => {
        if (line(x) == '*') {
          val thisStarNums = mutable.Set[RangeNum]()
          if (y > 0) {
            Array(-1, 0, 1).foreach(dx => {
              getNumber(lines(y - 1), x + dx).exists(thisStarNums.add)
            })
          }
          if (y < lines.size - 1) {
            Array(-1, 0, 1).foreach(dx => {
              getNumber(lines(y + 1), x + dx).exists(thisStarNums.add)
            })
          }
          Array(-1, 0, 1).foreach(dx => {
            getNumber(lines(y), x + dx).exists(thisStarNums.add)
          })
          starNums += thisStarNums
        }
      })
    }
    starNums.filter(sn => sn.size == 2).map(_.toArray).map(sn => sn(0).n * sn(1).n).sum
  }

  def exampleB(): Int = {
    solveB(example)
  }

  def b(): Int = {
    solveB(fileContent)
  }

  def getNumber(s: String, i: Int): Option[RangeNum] = {
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
    Some(RangeNum(s.slice(start + 1, end).toInt, start + 1, end))
  }
}

