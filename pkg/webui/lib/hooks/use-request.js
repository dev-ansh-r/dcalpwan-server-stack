

import { useState, useEffect } from 'react'
import { useDispatch } from 'react-redux'

import attachPromise from '@ttn-lw/lib/store/actions/attach-promise'

const useRequest = requestAction => {
  const dispatch = useDispatch()
  const [fetching, setFetching] = useState(true)
  const [error, setError] = useState('')
  const [result, setResult] = useState()

  useEffect(() => {
    if (requestAction) {
      const promise =
        requestAction instanceof Array
          ? Promise.all(requestAction.map(req => dispatch(attachPromise(req))))
          : typeof requestAction === 'function'
          ? requestAction(dispatch)
          : dispatch(attachPromise(requestAction))

      promise
        .then(() => {
          setResult(result)
          setFetching(false)
        })
        .catch(error => {
          setError(error)
          setFetching(false)
        })
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  return [fetching, error, result]
}

export default useRequest
