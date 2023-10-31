

export const STACK_COMPONENTS = ['as', 'is', 'ns', 'js', 'gs', 'edtc', 'qrg', 'gcs', 'dcs']

export const STACK_COMPONENTS_MAP = STACK_COMPONENTS.reduce((acc, curr) => {
  acc[curr] = curr
  return acc
}, {})

export const URI_PREFIX_STACK_COMPONENT_MAP = {
  as: 'as',
  ns: 'ns',
  js: 'js',
  gs: 'gs',
  edtc: 'edtc',
  qrg: 'qrg',
  gcs: 'gcs',
  edcs: 'dcs',
}

export const AUTHORIZATION_MODES = Object.freeze({
  KEY: 'key',
  SESSION: 'session',
})

export const RATE_LIMIT_RETRIES = 5
