

const selectNsStore = state => state.ns

export const generateDefaultMacSettingsKey = (frequencyPlanId, lorawanPhyVersion) =>
  `${frequencyPlanId}/${lorawanPhyVersion}`

export const selectDefaultMacSettings = (state, frequencyPlanId, lorawanPhyVersion) => {
  if (!frequencyPlanId || !lorawanPhyVersion) {
    return undefined
  }
  const store = selectNsStore(state)
  const key = generateDefaultMacSettingsKey(frequencyPlanId, lorawanPhyVersion)

  return store.defaultMacSettings[key]
}
