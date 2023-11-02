
import React from 'react'
import ReactDOM from 'react-dom'

import DeviceUplink from '../containers/device-uplink'

import 'normalize.css'
import '../styles/main.styl'

const { device } = window.config

const render = (id, component) => {
  ReactDOM.render(
    <React.StrictMode>
      {component}
    </React.StrictMode>,
    document.getElementById(id),
  )
}

render('device-events', <DeviceUplink model_id={device.modelid} brand_id={device.vendorid} device={device} />)
