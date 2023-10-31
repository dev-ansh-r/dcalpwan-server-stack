

describe('Network information section', () => {
  before(() => {
    cy.dropAndSeedDatabase()
  })

  beforeEach(() => {
    cy.loginConsole({ user_id: 'admin', password: 'admin' })
  })

  it('succeeds showing registry totals', () => {
    cy.visit(`${Cypress.config('consoleRootPath')}/admin-panel/network-information`)
    cy.findByText('Network information', { selector: 'h1' }).should('be.visible')
    // Shows registry totals.
    cy.findAllByText('Total applications').should('be.visible')
    cy.findAllByText('Total gateways').should('be.visible')
    cy.findAllByText('Registered users').should('be.visible')
    cy.findAllByText('Organizations').should('be.visible')
    cy.findByText('Deployment').should('be.visible')
    cy.findByText('Available components').should('be.visible')
  })
})
