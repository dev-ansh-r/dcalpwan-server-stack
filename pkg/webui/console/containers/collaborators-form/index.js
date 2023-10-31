

import React from 'react'

import tts from '@console/api/tts'

import CollaboratorForm from '@ttn-lw/containers/collaborator-form'

const ConsoleCollaboratorsForm = props => <CollaboratorForm {...props} tts={tts} />

export default ConsoleCollaboratorsForm
