import org.scalatest.{Matchers, FunSuite}

class HelloWorldTest extends FunSuite with Matchers {
  test("Say Hi!") {
    HelloWorld.hello() should be ("Hello, World!")
  }
}
