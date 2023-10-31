

import React from 'react'

import Button from '@ttn-lw/components/button'

import SubmitBar from '.'

export default {
  title: 'SubmitBar',
}

export const OnlySubmit = () => (
  <SubmitBar>
    <Button message="Save Changes" icon="done" />
  </SubmitBar>
)

export const SubmitAndReset = () => (
  <SubmitBar>
    <Button message="Save Changes" icon="done" />
    <Button message="Delete" icon="delete" naked danger />
  </SubmitBar>
)

SubmitAndReset.story = {
  name: 'Submit and Reset',
}

export const SubmitAndText = () => (
  <SubmitBar align="start">
    <Button message="Save Changes" icon="done" />
    <SubmitBar.Message content="Note: End device level message payload formats take precedence over application level message payload formats" />
  </SubmitBar>
)

SubmitAndText.story = {
  name: 'Submit and Text',
}
