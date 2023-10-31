

import * as Yup from 'yup'

const StringSchema = Yup.string

/**
 * `NullableStringSchemaType` is an extension for the default `yup.string`
 * schema type. It transforms the value to `null` if it is empty and skips
 * validation.
 */
class NullableStringSchemaType extends StringSchema {
  constructor() {
    super()

    const self = this

    /* eslint-disable prefer-arrow/prefer-arrow-functions */
    self.withMutation(() => {
      self
        .transform(value => {
          if (self.isType(value) && Boolean(value)) {
            return value
          }

          return null
        })
        .nullable(true)
    })
    /* eslint-enable prefer-arrow/prefer-arrow-functions */
  }
}

const nullableString = () => new NullableStringSchemaType()

const passValues = message => values => ({
  message,
  values,
})

Yup.addMethod(Yup.string, 'emptyOrLength', function (length, message) {
  let m = message
  if (typeof message === 'function') {
    m = message({ length })
  }
  // eslint-disable-next-line no-invalid-this
  return this.test('empty-or-length', m, value => !Boolean(value) || value.length === length)
})

export default {
  ...Yup,
  nullableString,
  passValues,
}
