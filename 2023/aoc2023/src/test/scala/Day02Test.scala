import org.scalatest.flatspec.AnyFlatSpec
import org.scalatest.matchers.should.Matchers

class Day02Test extends AnyFlatSpec with Matchers {
  "Day02" should "pass example A" in {
    Day02.exampleA() shouldEqual 8
  }
  //  it should "pass part A" in {
  //    Day02.a() shouldEqual 53334
  //  }
  //  it should "pass example B" in {
  //    Day02.exampleB() shouldEqual 281
  //  }
  //  it should "pass part B" in {
  //    Day02.b() shouldEqual 52834
  //  }
}
