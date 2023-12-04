import scala.io.Source

object Util {
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
