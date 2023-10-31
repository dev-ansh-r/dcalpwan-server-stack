

/* eslint-disable no-console */
const originalConsoleError = console.error
console.error = (message, ...args) => {
  console.log(message)
  if (/(Invalid prop|Failed prop type|Failed context type)/gi.test(message)) {
    throw new Error(message)
  }

  originalConsoleError.apply(console, [message, ...args])
}
/* eslint-enable no-console */
