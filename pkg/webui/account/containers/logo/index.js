

import React from 'react'

import LogoComponent from '@ttn-lw/components/logo'

import {
  selectAssetsRootPath,
  selectBrandingRootPath,
  selectApplicationSiteName,
} from '@ttn-lw/lib/selectors/env'

const logo = {
  src: `${selectAssetsRootPath()}/account.svg`,
  alt: `${selectApplicationSiteName()} Logo`,
}
const hasCustomBranding = selectBrandingRootPath() !== selectAssetsRootPath()
const brandLogo = hasCustomBranding
  ? {
      src: `${selectBrandingRootPath()}/logo.svg`,
      alt: 'Logo',
    }
  : undefined

const Logo = props => <LogoComponent logo={logo} brandLogo={brandLogo} {...props} />

export default Logo
