
import { selectApplicationConfig } from '@ttn-lw/lib/selectors/env'

const selectAccountUrl = () => selectApplicationConfig().account_url

export default selectAccountUrl
