

export const notify = (listener, ...args) => {
  if (typeof listener === 'function') {
    listener(...args)
  }
}

export const EVENTS = Object.freeze({
  START: 'start',
  CHUNK: 'chunk',
  ERROR: 'error',
  CLOSE: 'close',
})
