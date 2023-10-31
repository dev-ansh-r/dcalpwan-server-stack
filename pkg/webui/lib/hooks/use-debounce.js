

import { useEffect, useState } from 'react'

const useDebounce = (value, delay = 350, sideEffects) => {
  const [debouncedValue, setDebouncedValue] = useState(value)

  useEffect(() => {
    // Unless it's the first render, set the debounced value to the current value.
    if (debouncedValue === value) {
      return
    }
    const timer = setTimeout(() => {
      setDebouncedValue(value)
      if (sideEffects) {
        sideEffects()
      }
    }, delay)

    return () => {
      clearTimeout(timer)
    }
  }, [value, delay, sideEffects, debouncedValue])

  return debouncedValue
}

export default useDebounce
