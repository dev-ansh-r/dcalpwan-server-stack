

export default string => {
  let hash = 0
  let chr = 0
  if (string.length === 0) return hash
  for (let i = 0; i < string.length; i++) {
    chr = string.charCodeAt(i)
    hash = (hash << 5) - hash + chr
    hash |= 0 // Convert to 32bit integer.
  }
  return hash
}
