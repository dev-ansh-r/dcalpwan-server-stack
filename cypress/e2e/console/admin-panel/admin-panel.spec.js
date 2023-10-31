

describe('Admin Panel', () => {
  before(() => {
    cy.dropAndSeedDatabase()
    cy.intercept('/api/v3/pba/info', { fixture: 'console/packet-broker/info-registered.json' })
  })

  beforeEach(() => {
    cy.loginConsole({ user_id: 'admin', password: 'admin' })
  })

  it('succeeds displaying different views in the admin panel', () => {
    cy.visit(`${Cypress.config('consoleRootPath')}/admin-panel`)

    cy.findByText('User management').should('be.visible').click()
    cy.findByText('User management', { selector: 'h1' }).should('be.visible')
    cy.findByTestId('error-notification').should('not.exist')

    cy.findByText('Peering settings').should('be.visible').click()
    cy.findByText('Packet Broker', { selector: 'h1' }).should('be.visible')
    cy.findByTestId('error-notification').should('not.exist')
  })
})
