

import Overview from '@account/views/overview'
import ProfileSettings from '@account/views/profile-settings'
import Code from '@account/views/code'
import SessionManagement from '@account/views/session-management'
import { ValidateWithAuth } from '@account/views/validate'
import OAuthClients from '@account/views/oauth-clients'
import OAuthClientAuthorizations from '@account/views/oauth-client-authorizations'

export default [
  {
    path: '/',
    Component: Overview,
  },
  {
    path: '/profile-settings',
    Component: ProfileSettings,
  },
  {
    path: '/code',
    Component: Code,
  },
  {
    path: '/session-management',
    Component: SessionManagement,
  },
  {
    path: '/validate',
    Component: ValidateWithAuth,
  },
  {
    path: '/oauth-clients/*',
    Component: OAuthClients,
  },
  {
    path: '/client-authorizations/*',
    Component: OAuthClientAuthorizations,
  },
]
