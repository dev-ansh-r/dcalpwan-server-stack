
export default (f, ms) => {
  let timer = null
  let cancelled = false
  return {
    debounced: (...args) => {
      const onComplete = () => {
        if (!cancelled) {
          f(...args)
        }
        timer = null
      }
      if (timer) {
        clearTimeout(timer)
      }
      timer = setTimeout(onComplete, ms)
    },
    cancel: () => {
      cancelled = true
    },
  }
}
