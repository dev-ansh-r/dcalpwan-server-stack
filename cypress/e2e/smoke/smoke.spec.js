

import tests from '.'

describe('Smoke tests', () => {
  beforeEach(() => {
    cy.dropAndSeedDatabase()
  })

  tests.forEach(({ description, options, test }) =>
    options ? it(description, options, test) : it(description, test),
  )
})
