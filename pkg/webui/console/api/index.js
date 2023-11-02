
import axios from 'axios'

import { selectApplicationRootPath, selectCSRFToken } from '@ttn-lw/lib/selectors/env'
import tokenCreator from '@ttn-lw/lib/access-token'

const appRoot = selectApplicationRootPath()

const csrf = selectCSRFToken()

const token = tokenCreator(() => axios.get(`${appRoot}/api/auth/token`))

export default {
  console: {
    token: () => axios.get(`${appRoot}/api/auth/token`),
    logout: async () => {
      const headers = token => ({
        headers: { 'X-CSRF-Token': token },
      })
      try {
        return await axios.post(`${appRoot}/api/auth/logout`, undefined, headers(csrf))
      } catch (error) {
        if (
          error.response &&
          error.response.status === 403 &&
          typeof error.response.data === 'string' &&
          error.response.data.includes('CSRF')
        ) {
          // If the CSRF token is invalid, it likely means that the CSRF cookie
          // has been deleted or became outdated. Making a new request to the
          // current path can then retrieve a fresh CSRF cookie, with which
          // the logout can be retried.
          const csrfResult = await axios.get(window.location)
          const freshCsrf = csrfResult.headers['x-csrf-token']
          if (freshCsrf) {
            return axios.post(`${appRoot}/api/auth/logout`, undefined, headers(freshCsrf))
          }
        }

        throw error
      }
    },
  },
}

export { token }
