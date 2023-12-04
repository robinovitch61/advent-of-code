package Util

import org.scalatest.flatspec.AnyFlatSpec
import org.scalatest.matchers.should.Matchers

import scala.collection.mutable

class CharTrieTest extends AnyFlatSpec with Matchers {
  "CharTrie" should "work with no strings" in {
    CharTrie.fromStrings(Set()) shouldEqual CharTrie(children = mutable.Map())
  }
  it should "work with one string" in {
    CharTrie.fromStrings(Set("t")) shouldEqual CharTrie(children = mutable.Map('t' -> CharTrie(children = mutable.Map())))
  }
  it should "contain one exact string" in {
    CharTrie.fromStrings(Set("testing")).containsPrefix("testing") shouldEqual Some("testing")
  }
  it should "contain one prefix string" in {
    CharTrie.fromStrings(Set("one", "two")).containsPrefix("onetwo") shouldEqual Some("one")
  }
  it should "not contain one prefix string" in {
    CharTrie.fromStrings(Set("one1", "two")).containsPrefix("onetwo") shouldEqual None
  }
  it should "not false match single letters" in {
    CharTrie.fromStrings(Set("one", "two")).containsPrefix("o") shouldEqual None
  }
  //  it should "not contain one string" in {
  //    CharTrie.fromStrings(Set("testing")).isPrefix("nope") shouldEqual false
  //  }
  //  it should "work with multiple strings" in {
  //    val ct = CharTrie.fromStrings(Set("testing", "testing2"))
  //    ct.isPrefix("testing") shouldEqual true
  //    ct.isPrefix("nope") shouldEqual false
  //  }
  //  it should "work with prefix of multiple strings" in {
  //    val ct = CharTrie.fromStrings(Set("testing", "testing2"))
  //    ct.isPrefix("testi") shouldEqual true
  //    ct.isPrefix("testing3") shouldEqual false
  //  }
}
