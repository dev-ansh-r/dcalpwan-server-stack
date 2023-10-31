

import { useEffect } from 'react'

export default (className, id = 'app') => {
  useEffect(() => {
    const containerClasses = document.getElementById(id).classList
    const classNamesList = className.split(' ')
    for (const cls of classNamesList) {
      containerClasses.add(cls)
    }
    return () => {
      for (const cls of classNamesList) {
        containerClasses.remove(cls)
      }
    }
    // Disabling deps check because in this case we want the effect hook to
    // run only once on component mount. See also:
    // https://github.com/facebook/create-react-app/issues/6880
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])
}
