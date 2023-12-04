package Util

import scala.collection.mutable

case class CharTrie(children: mutable.Map[Char, CharTrie]) {
  /**
   * containsPrefix takes a long string and returns an optional word in the trie
   * if the prefix of the word matches it.
   *
   * Example:
   *
   * @ CharTrie.fromStrings(Set("one", "two")).containsPrefix("onetwo")
   * "one"
   *
   * @param s The string
   * @return optional matching prefix
   */
  def containsPrefix(s: String): Option[String] = {
    var curr = children
    var word = ""
    s.foreach(c => {
      if (!curr.contains(c)) {
        return None
      }
      curr = curr(c).children
      word += c
      if (curr.isEmpty) {
        return Some(word)
      }
    })
    if (curr.isEmpty) Some(word) else None
  }
}

object CharTrie {
  def fromStrings(strings: Set[String]): CharTrie = {
    val res = CharTrie(children = mutable.Map())
    strings.foreach(s => {
      var curr = res
      s.foreach(c => {
        curr = curr.children.getOrElseUpdate(c, CharTrie(children = mutable.Map()))
      })
    })
    res
  }
}

