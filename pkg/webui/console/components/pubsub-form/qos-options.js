

export const qosOptions = [
  { value: 'AT_MOST_ONCE', label: 'At most once (0)' },
  { value: 'AT_LEAST_ONCE', label: 'At least once (1)' },
  { value: 'EXACTLY_ONCE', label: 'Exactly once (2)' },
]

export const qosLevels = qosOptions.map(e => e.value)
