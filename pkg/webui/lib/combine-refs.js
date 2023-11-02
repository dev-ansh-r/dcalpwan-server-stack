
import { isPlainObject } from 'lodash'

/**
 * Combines multiple refs.
 *
 * @param {Array<object>} refs - A list of refs to be merged.
 * @returns {Function} - The ref callback with the DOM element that is assigned to every ref in `refs`.
 */
const combineRefs = refs => val => {
  refs.forEach(ref => {
    if (isPlainObject(ref)) {
      ref.current = val
    } else if (typeof ref === 'function') {
      ref(val)
    }
  })
}

export default combineRefs
