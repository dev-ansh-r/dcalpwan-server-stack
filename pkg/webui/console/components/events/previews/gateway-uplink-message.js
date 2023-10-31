

import React from 'react'

import Message from '@ttn-lw/lib/components/message'

import { getDataRate, getSignalInformation } from '@console/components/events/utils'

import PropTypes from '@ttn-lw/lib/prop-types'
import getByPath from '@ttn-lw/lib/get-by-path'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import messages from '../messages'

import DescriptionList from './shared/description-list'

const GatewayUplinkMessagePreview = React.memo(({ event }) => {
  const { data } = event
  let fPort, snr, devAddr, fCnt, joinEui, devEui, rssi, isConfirmed, dataRate

  if ('message' in data) {
    if ('payload' in data.message) {
      if ('mac_payload' in data.message.payload) {
        devAddr = getByPath(data, 'message.payload.mac_payload.f_hdr.dev_addr')
        fPort = getByPath(data, 'message.payload.mac_payload.f_port')
        fCnt = getByPath(data, 'message.payload.mac_payload.f_hdr.f_cnt')
      }

      if ('join_request_payload' in data.message.payload) {
        joinEui = getByPath(data, 'message.payload.join_request_payload.join_eui')
        devEui = getByPath(data, 'message.payload.join_request_payload.dev_eui')
      }

      isConfirmed = getByPath(data, 'message.payload.m_hdr.m_type') === 'CONFIRMED_UP'
    }

    const signalInfo = getSignalInformation(data.message)
    snr = signalInfo.snr
    rssi = signalInfo.rssi
    dataRate = getDataRate(data.message)
  }

  return (
    <DescriptionList>
      <DescriptionList.Byte title={messages.devAddr} data={devAddr} />
      <DescriptionList.Item title={messages.fCnt} data={fCnt} highlight />
      <DescriptionList.Item title={messages.fPort} data={fPort} />
      {isConfirmed && (
        <DescriptionList.Item>
          <Message content={messages.confirmedUplink} />
        </DescriptionList.Item>
      )}
      <DescriptionList.Byte title={sharedMessages.joinEUI} data={joinEui} />
      <DescriptionList.Byte title={sharedMessages.devEUI} data={devEui} />
      <DescriptionList.Item title={messages.dataRate} data={dataRate} />
      <DescriptionList.Item title={messages.snr} data={snr} />
      <DescriptionList.Item title={messages.rssi} data={rssi} />
    </DescriptionList>
  )
})

GatewayUplinkMessagePreview.propTypes = {
  event: PropTypes.event.isRequired,
}

export default GatewayUplinkMessagePreview
