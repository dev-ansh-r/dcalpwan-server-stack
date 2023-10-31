

import axios from 'axios'

import api from '@console/api'

import * as accessToken from '@ttn-lw/lib/access-token'
import createRequestLogic from '@ttn-lw/lib/store/logics/create-request-logic'
import { isUnauthenticatedError } from '@ttn-lw/lib/errors/utils'
import { selectApplicationRootPath } from '@ttn-lw/lib/selectors/env'

import * as user from '@console/store/actions/logout'

const logoutSequence = async () => {
  const response = await api.console.logout()
  accessToken.clear()
  window.location = response.data.op_logout_uri
}

export default [
  createRequestLogic({
    type: user.LOGOUT,
    process: async () => {
      try {
        await logoutSequence()
      } catch (err) {
        if (isUnauthenticatedError(err)) {
          // If there was an Unauthenticated Error, it either means that the
          // console client or the OAuth app session is no longer valid.
          // In this situation, it's best to try initializing the OAuth
          // roundtrip again. This might provide a new console session cookie
          // with which the propagated logout can be retried. If not, it can
          // be assumed that both console and OAuth app sessions are already
          // terminated, equalling a logged out state. In that case the request
          // logic will perform a page refresh which will initialize the auth
          // flow again.
          await axios.get(
            `${selectApplicationRootPath()}/login/ttn-stack?next=${window.location.pathname}`,
          )
          await logoutSequence()
        } else {
          throw err
        }
      }
    },
  }),
]
