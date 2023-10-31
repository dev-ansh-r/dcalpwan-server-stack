import React from 'react'
import classnames from 'classnames'

import Link from '@ttn-lw/components/link'

import PropTypes from '@ttn-lw/lib/prop-types'

import style from './logo.styl'

const LogoLink = ({ safe, to, ...rest }) => {
  if (safe) {
    return <a href={to} {...rest} />
  }

  return <Link to={to} {...rest} />
}

LogoLink.propTypes = {
  safe: PropTypes.bool,
  to: PropTypes.string.isRequired,
}

LogoLink.defaultProps = {
  safe: false,
}

const Logo = ({ className, logo, brandLogo, vertical, safe }) => {
  const classname = classnames(style.container, className, {
    [style.vertical]: vertical,
    [style.customBranding]: Boolean(brandLogo),
  })

  return (
    <div className={classname}>
      {Boolean(brandLogo) && (
        <div className={style.brandLogo}>
          <LogoLink safe={safe} to="/" id="brand-logo" className={style.brandLogoContainer}>
            {/* <img {...brandLogo} /> */}
          </LogoLink>
        </div>
      )}
      <div className={style.logo}>
        <LogoLink safe={safe} className={style.logoContainer} to="/">
          {/* <img {...logo} /> */}
        </LogoLink>
      </div>
    </div>
  )
}

const imgPropType = PropTypes.shape({
  src: PropTypes.string.isRequired,
  alt: PropTypes.string.isRequired,
})

Logo.propTypes = {
  brandLogo: imgPropType,
  className: PropTypes.string,
  logo: imgPropType.isRequired,
  safe: PropTypes.bool,
  vertical: PropTypes.bool,
}

Logo.defaultProps = {
  className: undefined,
  brandLogo: undefined,
  vertical: false,
  safe: false,
}

export default Logo
