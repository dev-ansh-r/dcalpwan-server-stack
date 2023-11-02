
import React from 'react'
import { useSelector } from 'react-redux'

import FooterComponent from '@ttn-lw/components/footer'

import { selectOnlineStatus } from '@ttn-lw/lib/store/selectors/status'
import {
  selectDocumentationUrlConfig,
  selectSupportLinkConfig,
  selectPageStatusBaseUrlConfig,
} from '@ttn-lw/lib/selectors/env'

const supportLink = selectSupportLinkConfig()
const documentationBaseUrl = selectDocumentationUrlConfig()
const statusPageBaseUrl = selectPageStatusBaseUrlConfig()

const Footer = props => {
  const onlineStatus = useSelector(selectOnlineStatus)

  return (
    <FooterComponent
      {...props}
      onlineStatus={onlineStatus}
      supportLink={supportLink}
      documentationLink={documentationBaseUrl}
      statusPageLink={statusPageBaseUrl}
    />
  )
}

const { onlineStatus: onlineStatusPropType, ...propTypes } = FooterComponent.propTypes

Footer.propTypes = propTypes

const { onlineStatus: onlineStatusDefaultProp, ...defaultProps } = FooterComponent.defaultProps

Footer.defaultProps = defaultProps

export default Footer
