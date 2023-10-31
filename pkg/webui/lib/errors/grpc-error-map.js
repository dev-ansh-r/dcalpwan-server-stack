

/* eslint-disable quote-props */

// Source: https://github.com/grpc/grpc/blob/master/doc/statuscodes.md

const errorMap = {
  0: '200',
  1: '499',
  2: '500',
  3: '400',
  4: '504',
  5: '404',
  6: '409',
  7: '403',
  8: '429',
  9: '400',
  10: '409',
  11: '400',
  12: '501',
  13: '500',
  14: '503',
  15: '500',
  16: '401',
}

export default rpcError => {
  if (typeof rpcError !== 'string' && typeof rpcError !== 'number') {
    return undefined
  }

  return errorMap[rpcError] || '520' // Fallback to 520 Unknown.
}
