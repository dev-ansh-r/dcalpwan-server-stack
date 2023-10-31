

import { createLogic } from 'redux-logic'

import api from '@account/api'

import * as init from '@ttn-lw/lib/store/actions/init'
import { isPermissionDeniedError, isUnauthenticatedError } from '@ttn-lw/lib/errors/utils'
import { promisifyDispatch } from '@ttn-lw/lib/store/middleware/request-promise-middleware'
import attachPromise from '@ttn-lw/lib/store/actions/attach-promise'

import * as user from '@account/store/actions/user'

const accountAppInitLogic = createLogic({
  type: init.INITIALIZE,
  process: async (_, dispatch, done) => {
    let session_id
    try {
      const meResult = await api.account.me()
      // Using `store.dispatch` since redux logic's dispatch won't return
      // the (promisified) action result like regular dispatch does.
      await promisifyDispatch(dispatch)(
        attachPromise(
          user.getUser(meResult.data.user.ids.user_id, [
            'profile_picture',
            'name',
            'description',
            'primary_email_address',
            'admin',
          ]),
        ),
      )
      session_id = meResult.data.session_id
    } catch (error) {
      if (!isUnauthenticatedError(error) && !isPermissionDeniedError(error)) {
        const initError = error?.data || error
        dispatch(init.initializeFailure(initError))
        return done()
      }
      // Unauthenticated or forbidden errors mean that the user is logged out.
      // This is expected and should not make the initialization fail.
    }

    // eslint-disable-next-line no-console
    console.log('Account app initialization successful!')
    dispatch(init.initializeSuccess(session_id))
    done()
  },
})

export default accountAppInitLogic
