

export const hostname =
  /^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9-]*[A-Za-z0-9])$/
export const url = /^\b((http|https):\/\/?)[^\s()<>]+(?:\([\w\d]+\)|([^[:punct:]\s]|\/?))$/
export const id = /^[a-z0-9](?:[-]?[a-z0-9]){2,}$/
export const userId = /^[a-z0-9](?:[-]?[a-z0-9]){1,}$/
export const pathId = /([a-z0-9-]{3,})/
export const userPathId = /([a-z0-9-]{2,})/
