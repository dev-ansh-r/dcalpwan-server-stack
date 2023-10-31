

const selectIsStore = state => state.is

export const selectIsConfiguration = state => selectIsStore(state).configuration

export const selectUserRegistration = state => selectIsConfiguration(state).user_registration || {}
export const selectPasswordRequirements = state =>
  selectUserRegistration(state).password_requirements || {}
