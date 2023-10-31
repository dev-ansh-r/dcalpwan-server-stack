

import React from 'react'

import Breadcrumbs from './breadcrumbs'
import { BreadcrumbsConsumer } from './context'

export default props => (
  <BreadcrumbsConsumer>
    {({ breadcrumbs }) => <Breadcrumbs {...props} breadcrumbs={breadcrumbs} />}
  </BreadcrumbsConsumer>
)
