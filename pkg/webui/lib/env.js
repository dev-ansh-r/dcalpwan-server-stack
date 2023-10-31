

import * as envSelector from './selectors/env'

const env = {
  appRoot: envSelector.selectApplicationRootPath(),
  assetsRoot: envSelector.selectAssetsRootPath(),
  config: {
    stack: {
      as: envSelector.selectAsConfig(),
      gs: envSelector.selectGsConfig(),
      is: envSelector.selectIsConfig(),
      js: envSelector.selectJsConfig(),
      ns: envSelector.selectNsConfig(),
    },
    language: envSelector.selectLanguageConfig(),
    supportLink: envSelector.selectSupportLinkConfig(),
    documentationBaseUrl: envSelector.selectDocumentationUrlConfig(),
    pageStatusBaseUrl: envSelector.selectPageStatusBaseUrlConfig(),
  },
  pageData: envSelector.selectPageData(),
  sentryDsn: envSelector.selectSentryDsnConfig(),
  csrfToken: envSelector.selectCSRFToken(),
  siteName: envSelector.selectApplicationSiteName(),
  siteTitle: envSelector.selectApplicationSiteTitle(),
  siteSubTitle: envSelector.selectApplicationSiteSubTitle(),
  devEUIConfig: envSelector.selectDevEUIConfig(),
}

export default env
