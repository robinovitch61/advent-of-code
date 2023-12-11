import org.scalatest.flatspec.AnyFlatSpec
import org.scalatest.matchers.should.Matchers

class Day05Test extends AnyFlatSpec with Matchers {
  "toDestInterval" should "work" in {
    // maps from 5 to 14 inclusive up 5 units
    val map = Day05.Map(5, 10, 10)
    // mapped, unMapped
    Day05.toDestInterval((10, 20), map) shouldEqual ((IndexedSeq((15, 19)), IndexedSeq((15, 20))))
    Day05.toDestInterval((0, 10), map) shouldEqual ((IndexedSeq((10, 15)), IndexedSeq((0, 4))))
    Day05.toDestInterval((5, 14), map) shouldEqual ((IndexedSeq((10, 19)), IndexedSeq()))
    Day05.toDestInterval((0, 4), map) shouldEqual ((IndexedSeq(), IndexedSeq((0, 4))))
    Day05.toDestInterval((15, 20), map) shouldEqual ((IndexedSeq(), IndexedSeq((15, 20))))
  }
  "mapInterval" should "work" in {
    // maps from 5 to 14 inclusive up 5 units
    val map1 = Day05.Map(5, 10, 10)
    // maps from 16 to 18 inclusive down 10 units
    val map2 = Day05.Map(16, 6, 3)
    Day05.mapInterval((0, 3), Array(map1, map2)) shouldEqual IndexedSeq((0, 3))
    Day05.mapInterval((19, 20), Array(map1, map2)) shouldEqual IndexedSeq((19, 20))
    Day05.mapInterval((5, 14), Array(map1, map2)) shouldEqual IndexedSeq((10, 19))
    Day05.mapInterval((16, 18), Array(map1, map2)) shouldEqual IndexedSeq((6, 8))
    Day05.mapInterval((4, 15), Array(map1, map2)).sortWith(_._1 < _._1) shouldEqual
      IndexedSeq((4, 4), (10, 19), (15, 15)).sortWith(_._1 < _._1)
    Day05.mapInterval((3, 20), Array(map1, map2)).sortWith(_._1 < _._1) shouldEqual
      IndexedSeq((3, 4), (10, 19), (15, 15), (6, 8), (19, 20)).sortWith(_._1 < _._1)
  }
  "Day05" should "pass example A" in {
    Day05.exampleA() shouldEqual 35
  }
  it should "pass part A" in {
    Day05.a() shouldEqual 278755257
  }
  it should "pass example B" in {
    Day05.exampleB() shouldEqual 46
  }
  it should "pass part B" in {
    Day05.b() shouldEqual 26829166
  }
}
