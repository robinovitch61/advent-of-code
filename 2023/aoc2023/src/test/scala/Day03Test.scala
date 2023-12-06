import org.scalatest.flatspec.AnyFlatSpec
import org.scalatest.matchers.should.Matchers

class Day03Test extends AnyFlatSpec with Matchers {
  "getNumber" should "work" in {
    Day03.getNumber("...123...", 3) shouldEqual Some(Num(123, 3, 6))
    Day03.getNumber("...123...", 4) shouldEqual Some(Num(123, 3, 6))
    Day03.getNumber("...123...", 5) shouldEqual Some(Num(123, 3, 6))
    Day03.getNumber("...123...", 6) shouldEqual None
  }
  it should "work at edges" in {
    Day03.getNumber("467..", 2) shouldEqual Some(Num(467, 0, 3))
    Day03.getNumber("..467", 2) shouldEqual Some(Num(467, 2, 5))
  }
  it should "handle negative nums" in {
    Day03.getNumber(".-467", 2) shouldEqual Some(Num(467, 2, 5))
    Day03.getNumber(".-...", 1) shouldEqual None
    Day03.getNumber("..46-", 2) shouldEqual Some(Num(46, 2, 4))
  }
  "Day03" should "pass example A" in {
    Day03.exampleA() shouldEqual 4361
  }
  it should "pass part A" in {
    Day03.a() shouldEqual 514969
  }
  it should "pass example B" in {
    Day03.exampleB() shouldEqual 467835
  }
  it should "pass part B" in {
    Day03.b() shouldEqual 78915902
  }
}
