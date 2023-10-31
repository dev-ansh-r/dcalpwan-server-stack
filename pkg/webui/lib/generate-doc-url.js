

import { selectApplicationConfig, selectDocumentationUrlConfig } from './selectors/env'

export default path => {
  if (selectApplicationConfig() && selectDocumentationUrlConfig()) {
    return selectDocumentationUrlConfig().concat(path)
  }
  return null
}
