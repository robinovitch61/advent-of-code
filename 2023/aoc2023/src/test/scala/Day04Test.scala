import org.scalatest.flatspec.AnyFlatSpec
import org.scalatest.matchers.should.Matchers

class Day04Test extends AnyFlatSpec with Matchers {
  "Day04" should "pass example A" in {
    Day04.exampleA() shouldEqual 13
  }
  it should "pass part A" in {
    Day04.a() shouldEqual 28750
  }
  it should "pass example B" in {
    Day04.exampleB() shouldEqual 30
  }
  it should "pass part B" in {
    Day04.b() shouldEqual 10212704
  }
}
