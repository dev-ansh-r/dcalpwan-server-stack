

import tts from '@console/api/tts'

import createRequestLogic from '@ttn-lw/lib/store/logics/create-request-logic'

import * as QRCodeGenerator from '@console/store/actions/qr-code-generator'

const parseQRCodeLogic = createRequestLogic({
  type: QRCodeGenerator.PARSE_QR_CODE,
  process: async ({ action }) => {
    const { qrCode } = action.payload
    return await tts.QRCodeGenerator.parse(qrCode)
  },
})

export default [parseQRCodeLogic]
