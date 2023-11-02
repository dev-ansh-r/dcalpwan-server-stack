
export default function getByPath(ob, path) {
  const paths = path.split('.')
  if (paths.length === 1) {
    return ob[path]
  }

  let current = ob

  for (const level of paths) {
    const isObject = typeof current[level] === 'object' && !(current[level] instanceof Array)
    if (!isObject) {
      return current[level]
    }
    current = current[level]
  }

  return current
}
