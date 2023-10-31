

import { id as idRegexp } from '@ttn-lw/lib/regexp'

export const encodeAttributes = formValue =>
  (Array.isArray(formValue) &&
    formValue.reduce(
      (result, { key, value }) => ({
        ...result,
        [key]: value,
      }),
      {},
    )) ||
  undefined

export const decodeAttributes = attributesType =>
  (attributesType &&
    Object.keys(attributesType).reduce(
      (result, key) =>
        result.concat({
          key,
          value: attributesType[key],
        }),
      [],
    )) ||
  []

export const attributesCountCheck = object =>
  object === undefined ||
  object === null ||
  (object instanceof Object && Object.keys(object).length <= 10)
export const attributeValidCheck = object =>
  object === undefined ||
  object === null ||
  (object instanceof Object && Object.values(object).every(attribute => Boolean(attribute)))

export const attributeTooShortCheck = object =>
  object === undefined ||
  object === null ||
  (object instanceof Object && Object.keys(object).every(key => RegExp(idRegexp).test(key)))

export const attributeKeyTooLongCheck = object =>
  object === undefined ||
  object === null ||
  (object instanceof Object && Object.keys(object).every(key => key.length <= 36))

export const attributeValueTooLongCheck = object =>
  object === undefined ||
  object === null ||
  (object instanceof Object &&
    Object.values(object).some(value => value !== undefined) &&
    Object.values(object).every(value => value.length <= 200))
