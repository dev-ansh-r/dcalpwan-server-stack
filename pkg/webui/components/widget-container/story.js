

import React from 'react'

import WidgetContainer from '.'

export default {
  title: 'WidgetContainer',
}

export const Default = () => (
  <div style={{ width: '500px' }}>
    <WidgetContainer title="Location" toAllUrl="#" linkMessage="Change location">
      <div
        style={{
          height: '300px',
          border: '1px solid gray',
          backgroundColor: '#eee',
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          color: 'gray',
        }}
      >
        Map placeholder as example
      </div>
    </WidgetContainer>
  </div>
)
