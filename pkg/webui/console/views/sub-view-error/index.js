

import React from 'react'
import { Container, Row, Col } from 'react-grid-system'

import Icon from '@ttn-lw/components/icon'

import Message from '@ttn-lw/lib/components/message'
import ErrorMessage from '@ttn-lw/lib/components/error-message'

import PropTypes from '@ttn-lw/lib/prop-types'
import { isBackend, isNotFoundError, httpStatusCode } from '@ttn-lw/lib/errors/utils'
import errorMessages from '@ttn-lw/lib/errors/error-messages'
import statusCodeMessages from '@ttn-lw/lib/errors/status-code-messages'

import style from './sub-view.styl'

const SubViewErrorComponent = ({ error }) => {
  const isNotFound = isNotFoundError(error)
  const statusCode = httpStatusCode(error)
  let errorExplanation = errorMessages.subviewErrorExplanation
  let errorTitleMessage = errorMessages.subviewErrorTitle
  if (isNotFound) {
    errorExplanation = errorMessages.genericNotFound
  }
  if (statusCode) {
    errorTitleMessage = statusCodeMessages[statusCode]
  }

  return (
    <Container>
      <Row>
        <Col>
          <div className={style.title}>
            <Icon icon="error_outline" large />
            <Message component="h2" content={errorTitleMessage} />
          </div>
          <p>
            <Message component="span" content={errorExplanation} />
            <br />
            <Message component="span" content={errorMessages.contactAdministrator} />
          </p>
          {isBackend(error) && (
            <React.Fragment>
              <hr />
              <ErrorMessage content={error} />
            </React.Fragment>
          )}
        </Col>
      </Row>
    </Container>
  )
}

SubViewErrorComponent.propTypes = {
  error: PropTypes.error.isRequired,
}

const subViewErrorRender = error => <SubViewErrorComponent error={error} />

export default subViewErrorRender
