class HelloWorld {
  static hello(name?: string): string {
    if (name) {
      return `Hello, ${name}!`
    }

    return 'Hello, World!'
  }
}

export default HelloWorld