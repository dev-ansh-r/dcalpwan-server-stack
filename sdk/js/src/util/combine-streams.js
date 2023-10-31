

/**
 * Combines multiple streams into a single subscription provider.
 *
 * @param {Array} streams - An array of (async) stream functions.
 * @returns {object} The stream subscription object with the `on` function for
 * attaching listeners and the `close` function to close the stream.
 */
const combinedStream = async streams => {
  if (!(streams instanceof Array) || streams.length === 0) {
    throw new Error('Cannot combine streams with invalid stream array.')
  } else if (streams.length === 1) {
    return streams[0]
  }

  const subscribers = await Promise.all(streams)

  return {
    open: () => {
      for (const subscriber of subscribers) {
        subscriber.open()
      }
    },
    on: (eventName, callback) => {
      for (const subscriber of subscribers) {
        subscriber.on(eventName, callback)
      }
    },
    close: () => {
      for (const subscriber of subscribers) {
        subscriber.close()
      }
    },
  }
}

export default combinedStream
