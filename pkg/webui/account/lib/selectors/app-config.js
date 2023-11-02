
import { selectApplicationConfig } from '@ttn-lw/lib/selectors/env'

export const selectEnableUserRegistration = () => selectApplicationConfig().enable_user_registration

export const selectConsoleUrl = () => selectApplicationConfig().console_url
