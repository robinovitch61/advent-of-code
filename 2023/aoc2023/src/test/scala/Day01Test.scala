import org.scalatest.flatspec.AnyFlatSpec
import org.scalatest.matchers.should.Matchers

class Day01Test extends AnyFlatSpec with Matchers {
  "Day01" should "pass example A" in {
    Day01.exampleA() shouldEqual 142
  }
  it should "pass part A" in {
    Day01.a() shouldEqual 53334
  }
  it should "pass example B" in {
    Day01.exampleB() shouldEqual 281
  }
  it should "pass part B" in {
    Day01.b() shouldEqual 52834
  }
}
