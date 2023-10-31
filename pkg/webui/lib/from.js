

/* eslint-disable jsdoc/require-description-complete-sentence */

/**
 * `from` returns an array that contains the values of the keys of a which have
 * a trueish value in b. If only is a passed, it is used as a whitelist for keys
 * that should be used.
 *
 * This can be used in conjunction with css modules and the classnames package
 * to effectively created dynamic classnames. Before:
 *
 * @example
 *     import classnames from "classnames"
 *     import styles from "./some/styles.css".
 *
 *     const { foo, bar } = props.
 *
 *     const classname = classnames(style.foo, {
 *       [style.bar]: bar,
 *       [style.baz]: baz,
 *     })
 * After:
 * @example
 *
 *     import classnames from "classnames"
 *     import styles from "./some/styles.css"
 *
 *     const { foo, bar } = props
 *
 *     const classname = classnames(style.foo, ...from(styles, { foo, bar }))
 *     // or
 *     const classname = classnames(style.foo, ...from(styles, props, [ "foo", "bar" ]))
 *
 *
 * @param {object} a - The object to take the values from.
 * @param {object} b - The object that controls which values will be taken.
 * @param {Array} only - Filter the keys by name.
 * @returns {Array} - An array of values from a for which the key in b had a
 * trueish value.
 */
export default (a = {}, b = {}, only) => {
  const res = []
  for (const key in b) {
    if (!b[key] || !(key in a) || (only && !only.includes(key))) {
      continue
    }

    res.push(a[key])
  }

  return res
}
