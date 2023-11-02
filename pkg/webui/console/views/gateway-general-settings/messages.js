
import { defineMessages } from 'react-intl'

const messages = defineMessages({
  basicTitle: 'Basic settings',
  basicDescription: 'General settings, gateway updates and metadata',
  lorawanDescription: 'LoRaWAN network-layer settings',
  updateSuccess: 'Gateway updated',
  deleteSuccess: 'Gateway deleted',
  deleteFailure: 'Gateway delete error',
  deleteGateway: 'Delete gateway',
  modalWarning:
    'Are you sure you want to delete "{gtwName}"? This action cannot be undone and it will not be possible to reuse the gateway ID.',
  disablePacketBrokerForwarding:
    'Disable forwarding uplink messages received from this gateway to the Packet Broker',
  adminContactDescription:
    'Administrative contact information for this gateway. Typically used to indicate who to contact with administrative questions about the gateway.',
  techContactDescription:
    'Technical contact information for this gateway. Typically used to indicate who to contact with technical/security questions about the gateway.',
})

export default messages
