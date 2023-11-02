

/* eslint-disable import/prefer-default-export */

export class TokenError extends Error {
  constructor(message, cause) {
    super(message)
    this.cause = cause
    this.name = 'TokenError'
  }
}
