

import React from 'react'

/**
 * React hook that allows subscribing to `window.onbeforeunload` and control it inside
 * a react functional component.
 *
 * @param {Function} handler - The event handler function.
 * @example
 *
 * // Just prompt the user if page unload is indeed needed.
 * useBeforeUnload(event => {
 *  if (shouldBlock) {
 *    event.preventDefault()
 *  }
 * })
 *
 * // Use custom message when prompting the user. Note: this is not support by most browsers.
 * useBeforeUnload(event => {
 *  if (shouldBlock) {
 *    return "Some text"
 *  }
 * })
 */
const useBeforeUnload = (handler = () => {}) => {
  const handlerRef = React.useRef(handler)

  React.useEffect(() => {
    handlerRef.current = handler
  }, [handler])

  React.useEffect(() => {
    const handleBeforeunload = event => {
      let returnValue

      if (typeof handlerRef.current === 'function') {
        returnValue = handlerRef.current(event)
      }

      if (event.defaultPrevented) {
        event.returnValue = ''
      }

      if (typeof returnValue === 'string') {
        event.returnValue = returnValue
        return returnValue
      }
    }

    window.addEventListener('beforeunload', handleBeforeunload)

    return () => {
      window.removeEventListener('beforeunload', handleBeforeunload)
    }
  }, [])
}

export default useBeforeUnload
