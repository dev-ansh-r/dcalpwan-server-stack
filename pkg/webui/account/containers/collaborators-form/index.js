

import React from 'react'

import tts from '@account/api/tts'

import CollaboratorForm from '@ttn-lw/containers/collaborator-form'

const AccountCollaboratorsForm = props => <CollaboratorForm {...props} tts={tts} />

export default AccountCollaboratorsForm
