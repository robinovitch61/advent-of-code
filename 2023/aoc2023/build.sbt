ThisBuild / scalaVersion     := "2.13.12"
ThisBuild / version          := "0.1.0-SNAPSHOT"
ThisBuild / organization     := "com.example"
ThisBuild / organizationName := "example"

lazy val root = (project in file("."))
  .settings(
    name := "aoc2023",
    libraryDependencies += "org.scalatest" %% "scalatest" % "3.2.15" % Test
  )
