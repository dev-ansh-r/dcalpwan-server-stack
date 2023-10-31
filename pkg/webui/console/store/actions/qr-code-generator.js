

import createRequestActions from '@ttn-lw/lib/store/actions/create-request-actions'

const PARSE_QR_CODE_BASE = 'PARSE_QR_CODE'
export const [
  { request: PARSE_QR_CODE, success: PARSE_QR_CODE_SUCCESS, failure: PARSE_QR_CODE_FAILURE },
  { request: parseQRCode, success: parseQRCodeSuccess, failure: parseQRCodeFailure },
] = createRequestActions(PARSE_QR_CODE_BASE, qrCode => ({ qrCode }))
