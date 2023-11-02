

import tts from '@account/api/tts'

import createRequestLogic from '@ttn-lw/lib/store/logics/create-request-logic'

import * as session from '@account/store/actions/sessions'

const getUserSessionsLogic = createRequestLogic({
  type: session.GET_USER_SESSIONS_LIST,
  process: async ({ action }) => {
    const { id, params } = action.payload

    const result = await tts.Sessions.getAllSessions(id, {
      page: params?.page,
      limit: params?.limit,
    })

    return { sessions: result.sessions, sessionsTotalCount: result.totalCount }
  },
})

const deleteUserSessionLogic = createRequestLogic({
  type: session.DELETE_USER_SESSION,
  process: async ({ action }) => {
    const { user, sessionId } = action.payload
    const result = await tts.Sessions.deleteSession(user, sessionId)

    return result
  },
})

export default [getUserSessionsLogic, deleteUserSessionLogic]
