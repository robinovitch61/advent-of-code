import Util.FileIO

import scala.collection.mutable

object Day05 {
  private val fileContent = FileIO.readFile("day05.txt")
  private val example =
    """seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37""".stripMargin

  private case class Maps(maps: Array[Array[Map]]) {
    def seedToLocation(s: Long): Long = {
      var res = s
      maps.foreach(stage => {
        // assumption that only one map in a given stage can map a given source point
        // if none can, keep that source point
        res = stage.filter(m => m.canMap(res)) match {
          case Array(m) => m.toDest(res)
          case Array() => res
          case _ => throw new RuntimeException(s"multiple maps in stage can map $res")
        }
      })
      res
    }
  }

  case class Map(src: Long, dest: Long, range: Long) {
    def canMap(s: Long): Boolean = src <= s && s < (src + range)

    def toDest(s: Long): Long = {
      if (src <= s && s < (src + range)) dest + s - src else throw new RuntimeException(s"$s cannot be converted with src $src, dest $dest, range $range")
    }
  }

  private def solveA(input: String): Long = {
    val (seeds, allMaps) = parseInput(input)
    seeds.map(s => allMaps.seedToLocation(s)).min
  }

  def exampleA(): Long = {
    solveA(example)
  }

  def a(): Long = {
    solveA(fileContent)
  }

  private def solveB(input: String): Long = {
    val (seeds, allMaps) = parseInput(input)
    var intervals = (seeds.indices by 2).map(idx => {
      (seeds(idx), seeds(idx) + seeds(idx + 1) - 1)
    })
    allMaps.maps.foreach(mapSet => {
      var mappedIntervalsThisStage = IndexedSeq[(Long, Long)]()
      var unMappedIntervalsThisStage = IndexedSeq[(Long, Long)]()
      intervals.foreach(interval => {
        var anyMapped = false
        mapSet.foreach(map => {
          val (mapped, unMapped) = toDestIntervals(interval, map)
          if (mapped.nonEmpty) {
            anyMapped = true
          }
          //          println(s"mapped $mapped")
          //          println(s"unMapped $unMapped")
          mappedIntervalsThisStage = mappedIntervalsThisStage ++ mapped
          if (mapped.nonEmpty && unMapped.nonEmpty) {
            unMappedIntervalsThisStage = unMappedIntervalsThisStage ++ unMapped
          }
        })
        if (!anyMapped) {
          unMappedIntervalsThisStage = unMappedIntervalsThisStage ++ IndexedSeq(interval)
        }
      })
      intervals = (mappedIntervalsThisStage ++ unMappedIntervalsThisStage).distinct
      //      println(s"\n$intervals\n")
    })
    intervals.map(int => if (int._1 < int._2) int._1 else int._2).min
  }

  def exampleB(): Long = {
    solveB(example)
  }

  def b(): Long = {
    solveB(fileContent)
  }

  def toDestIntervals(interval: (Long, Long), map: Map): (IndexedSeq[(Long, Long)], IndexedSeq[(Long, Long)]) = {
    val mapped = mutable.ArrayBuffer[(Long, Long)]()
    val unMapped = mutable.ArrayBuffer[(Long, Long)]()
    val (iStart, iEnd) = (interval._1, interval._2)
    val (mStart, mEnd) = (map.src, map.src + map.range - 1)
    println(s"mapping ($iStart, $iEnd) in ($mStart, $mEnd), range ${map.range}")
    val start = math.max(iStart, mStart)
    val end = math.min(iEnd, mEnd)
    println(s"$start to $end")
    if (start > end) {
      // no overlap - all unmapped
      return (IndexedSeq(), IndexedSeq(interval))
    } else {
      // overlap - calculate mapped
      mapped += Tuple2(map.toDest(start), map.toDest(end))
      // get unmapped if overflow on either end
      if (iStart < mStart) {
        unMapped += Tuple2(iStart, mStart - 1)
      }
      if (iEnd > mEnd) {
        unMapped += Tuple2(mEnd + 1, iEnd)
      }
    }
    (mapped.toIndexedSeq, unMapped.toIndexedSeq)
  }

  private def parseInput(input: String): (Array[Long], Maps) = {
    val entries = input.split("\n\n")
    val seeds = entries(0).split("seeds: ") match {
      case Array(_, s) => s.split("\\s+").map(_.toLong)
      case _ => throw new RuntimeException(s"seeds error ${entries(0)}")
    }
    val maps: Array[Array[Map]] = entries.slice(1, entries.size).map(entry => {
      val lines = entry.split("\n")
      lines.slice(1, lines.size).map(line => line.split("\\s+") match {
        case Array(dest, src, range) => Map(src.toLong, dest.toLong, range.toLong)
        case _ => throw new RuntimeException(s"invalid line $line")
      })
    })
    (seeds, Maps(maps = maps))
  }

}
