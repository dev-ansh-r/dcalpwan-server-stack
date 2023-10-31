

import React from 'react'

import tts from '@console/api/tts'

export default Component => props => {
  const [state, setState] = React.useState({
    devAddr: '',
    loading: false,
    error: undefined,
  })
  const { devAddr, loading, error } = state

  const handleGenerateDevAddr = React.useCallback(async () => {
    setState(prev => ({ ...prev, loading: true }))

    try {
      const { dev_addr } = await tts.Ns.generateDevAddress()

      setState({ loading: false, error: undefined, devAddr: dev_addr })
    } catch (error) {
      setState(prev => ({ ...prev, loading: false, error }))
    }
  }, [])

  return (
    <Component
      {...props}
      onGenerate={handleGenerateDevAddr}
      generatedValue={devAddr}
      generatedError={Boolean(error)}
      generatedLoading={loading}
    />
  )
}
