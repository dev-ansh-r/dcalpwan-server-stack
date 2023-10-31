

import createRequestActions from '@ttn-lw/lib/store/actions/create-request-actions'

export const createGetRightsListActionType = name => `GET_${name}_RIGHTS_LIST`

export default name => createRequestActions(createGetRightsListActionType(name), id => ({ id }))
