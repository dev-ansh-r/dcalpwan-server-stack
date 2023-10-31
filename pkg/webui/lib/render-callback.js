

/**
 * Allows react components to use the render prop design pattern without explicitly
 * specifying `render`, `children` and `children` as a function props.
 *
 * @param {object} props - The props of a component.
 * @param {Function} props.render - The render function.
 * @param {object} props.children - The component children.
 * @param {object} context - The context to be passed to children.
 *
 * @returns {*} - Processed children by the provided `render`, `children` function or just `children`
 * depending on the props.
 */
const renderCallback = ({ render, children }, context) => {
  if (render) {
    return render(context)
  } else if (typeof children === 'function') {
    return children(context)
  }

  return children
}

export default renderCallback
