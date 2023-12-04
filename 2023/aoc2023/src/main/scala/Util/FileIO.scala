package Util

import scala.io.Source

object FileIO {
  def readFile(fileName: String): String = {
    val filePath = getClass.getClassLoader.getResource(fileName).getPath
    val source = Source.fromFile(filePath)
    try {
      source.mkString
    } finally {
      source.close()
    }
  }
}
