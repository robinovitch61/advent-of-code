import org.scalatest.flatspec.AnyFlatSpec
import org.scalatest.matchers.should.Matchers

class Day05Test extends AnyFlatSpec with Matchers {
  //  "toDestIntervals" should "work" in {
  //    // maps from 5 to 14 inclusive up 5 units
  //    val map = Day05.Map(5, 10, 10)
  //    // mapped, unMapped
  //    Day05.toDestIntervals((10, 20), map) shouldEqual ((IndexedSeq((15, 19)), IndexedSeq((15, 20))))
  //    Day05.toDestIntervals((0, 10), map) shouldEqual ((IndexedSeq((10, 15)), IndexedSeq((0, 4))))
  //    Day05.toDestIntervals((5, 14), map) shouldEqual ((IndexedSeq((10, 19)), IndexedSeq()))
  //    Day05.toDestIntervals((0, 4), map) shouldEqual ((IndexedSeq(), IndexedSeq((0, 4))))
  //    Day05.toDestIntervals((15, 20), map) shouldEqual ((IndexedSeq(), IndexedSeq((15, 20))))
  //  }
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
    // 128414658 too high
    Day05.b() shouldEqual 10212705
  }
}
