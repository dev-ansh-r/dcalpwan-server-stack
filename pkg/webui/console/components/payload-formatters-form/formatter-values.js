

const DEFAULT_UPLINK_JS_FORMATTER = `function decodeUplink(input) {
  return {
    data: {
      bytes: input.bytes
    },
    warnings: [],
    errors: []
  };
}`
const DEFAULT_DOWNLINK_JS_FORMATTER = `function encodeDownlink(input) {
  return {
    bytes: [],
    fPort: 1,
    warnings: [],
    errors: []
  };
}

function decodeDownlink(input) {
  return {
    data: {
      bytes: input.bytes
    },
    warnings: [],
    errors: []
  }
}`

export const getDefaultJavascriptFormatter = uplink =>
  uplink ? DEFAULT_UPLINK_JS_FORMATTER : DEFAULT_DOWNLINK_JS_FORMATTER

export const getDefaultGrpcServiceFormatter = () => undefined
