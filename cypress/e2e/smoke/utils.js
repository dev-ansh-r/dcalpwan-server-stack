

/* eslint-disable import/prefer-default-export */

const defineSmokeTest = (description, ...rest) => {
  switch (rest.length) {
    case 1:
      return {
        description,
        test: rest[0],
      }
    case 2:
      return {
        description,
        options: rest[0],
        test: rest[1],
      }
  }
}

export { defineSmokeTest }
