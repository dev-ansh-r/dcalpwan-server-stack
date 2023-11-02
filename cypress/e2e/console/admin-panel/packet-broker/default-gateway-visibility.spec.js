

describe('Packet Broker routing policies', () => {
  before(() => {
    cy.dropAndSeedDatabase()
  })

  beforeEach(() => {
    cy.fixture('console/packet-broker/policies-home-network.json').as('homeNetworkPolicies')

    cy.intercept('/api/v3/pba/info', { fixture: 'console/packet-broker/info-registered.json' })

    cy.loginConsole({ user_id: 'admin', password: 'admin' })
  })

  it('succeeds setting default gateway visibility configuration', () => {
    cy.intercept('GET', '/api/v3/pba/home-networks/gateway-visibilities/default', {
      statusCode: 404,
    })
    cy.intercept('GET', '/api/v3/pba/home-networks/policies/default', { statusCode: 404 })
    cy.intercept('PUT', '/api/v3/pba/home-networks/gateway-visibilities/default', {}).as('putCall')
    cy.visit(
      `${Cypress.config('consoleRootPath')}/admin-panel/packet-broker/default-gateway-visibility`,
    )

    cy.findByLabelText('Location').check()
    cy.findByLabelText('Antenna placement').check()
    cy.findByLabelText('Antenna count').check()
    cy.findByLabelText('Fine timestamps').check()

    cy.findByRole('button', { name: 'Save default gateway visibility' }).click()

    cy.findByTestId('error-notification').should('not.exist')
    cy.findByTestId('toast-notification')
      .findByText('Default gateway visibility set')
      .should('be.visible')
    cy.get('@putCall').should('have.property', 'state', 'Complete')
  })

  it('succeeds unsetting default gateway visibility configuration', () => {
    cy.intercept('GET', '/api/v3/pba/home-networks/policies/default', { statusCode: 404 })
    cy.intercept('GET', '/api/v3/pba/home-networks/gateway-visibilities/default', {
      fixture: 'console/packet-broker/default-gateway-visibility.json',
    })
    cy.intercept('PUT', '/api/v3/pba/home-networks/gateway-visibilities/default', {}).as('putCall')
    cy.visit(
      `${Cypress.config('consoleRootPath')}/admin-panel/packet-broker/default-gateway-visibility`,
    )

    cy.findByLabelText('Location').uncheck()
    cy.findByLabelText('Antenna placement').uncheck()
    cy.findByLabelText('Antenna count').uncheck()
    cy.findByLabelText('Fine timestamps').uncheck()

    cy.findByRole('button', { name: 'Save default gateway visibility' }).click()

    cy.findByTestId('error-notification').should('not.exist')
    cy.findByTestId('toast-notification')
      .findByText('Default gateway visibility set')
      .should('be.visible')
    cy.get('@putCall').should('have.property', 'state', 'Complete')
  })
})
