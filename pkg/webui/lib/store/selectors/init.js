

import { createErrorSelector } from '@ttn-lw/lib/store/selectors/error'
import { createFetchingSelector } from '@ttn-lw/lib/store/selectors/fetching'
import { INITIALIZE_BASE } from '@ttn-lw/lib/store/actions/init'

export const selectInitStore = state => state.init

export const selectIsInitialized = state => selectInitStore(state).initialized
export const selectInitFetching = createFetchingSelector(INITIALIZE_BASE)
export const selectInitError = createErrorSelector(INITIALIZE_BASE)
