

/* eslint-disable react/prop-types */

import React from 'react'
import bind from 'autobind-decorator'

import CodeEditor from '.'

const containerStyles = { height: '500px' }

@bind
class Example extends React.Component {
  constructor(props) {
    super(props)

    this.state = {
      value: props.placeholder,
    }
  }

  onChange(value) {
    this.setState({ value })
  }

  render() {
    const { value } = this.state
    return (
      <div style={containerStyles}>
        <CodeEditor {...this.props} onChange={this.onChange} value={value} />
      </div>
    )
  }
}

const code = `
// Decode raw data.
function Decoder(bytes, port) {
  var colors = ["red", "green", "blue", "yellow", "cyan", "magenta", "white", "black"];
  var decoded = {
    light: (bytes[0] << 8) | bytes[1],
    temperature: ((bytes[2] << 8) | bytes[3]) / 100,
    state: {
      color: colors[bytes[4]]
    }
  };

  return decoded;
}
`

export default {
  title: 'CodeEditor',
}

export const Default = () => (
  <Example language="javascript" name="storybook-code-editor" placeholder={code} />
)

export const Readonly = () => (
  <Example language="javascript" name="storybook-code-editor" placeholder={code} readOnly />
)

export const WithWarning = () => (
  <Example language="javascript" name="storybook-code-editor" placeholder={`${code}}`} />
)

WithWarning.story = {
  name: 'With warning',
}
