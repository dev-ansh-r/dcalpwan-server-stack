

class Token {
  constructor(token) {
    /**
     * Make sure it is possible to instantiate an instance
     * of the class only once.
     */
    if (!!Token.instance) {
      return Token.instance
    }

    if (this.token) {
      return Token.instance
    }

    Token.instance = this
    this.token = token

    return this
  }

  get() {
    return this.token
  }
}

export default Token
