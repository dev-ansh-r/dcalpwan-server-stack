

const composeDataUri = (data, mimeType = 'data:application/json;charset=utf-8') =>
  `${mimeType},${encodeURIComponent(data)}`

const downloadDataUriAsFile = (dataUri, filename) => {
  const node = document.createElement('a')
  node.setAttribute('href', dataUri)
  node.setAttribute('download', filename)
  document.body.appendChild(node) // Required for Firefox.
  node.click()
  node.remove()
}

export { composeDataUri, downloadDataUriAsFile }
