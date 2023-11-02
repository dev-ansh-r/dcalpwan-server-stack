

import { defineSmokeTest } from '../utils'

const gatewaySubpages = defineSmokeTest('check all gateway sub-pages', () => {
  const user = {
    ids: { user_id: 'gtw-sub-pages-test-user' },
    primary_email_address: 'test-user@example.com',
    password: 'ABCDefg123!',
    password_confirm: 'ABCDefg123!',
    email: 'gtw-subpage-test-user@example.com',
  }
  const gateway = {
    ids: { gateway_id: 'gtw-subpages-test' },
  }
  cy.createUser(user)
  cy.createGateway(gateway, user.ids.user_id)
  cy.loginConsole({ user_id: user.ids.user_id, password: user.password })
  cy.visit(Cypress.config('consoleRootPath'))

  cy.get('header').within(() => {
    cy.findByRole('link', { name: /Gateways/ }).click()
  })
  cy.findByRole('cell', { name: gateway.ids.gateway_id }).click()

  cy.findByRole('link', { name: /Live data/ }).click()
  cy.findByText(/Waiting for events from/).should('be.visible')
  cy.findByTestId('error-notification').should('not.exist')

  cy.findByRole('link', { name: /Location/ }).click()
  cy.findByLabelText('Location privacy').should('exist')
  cy.findByRole('button', { name: 'Save changes' }).should('be.visible')
  cy.findByTestId('error-notification').should('not.exist')

  cy.findByRole('link', { name: /Collaborators/ }).click()
  cy.findByText('Collaborators (1)').should('be.visible')
  cy.findByRole('link', { name: /Add collaborator/ }).should('be.visible')
  cy.findByTestId('error-notification').should('not.exist')

  cy.findByRole('link', { name: /API keys/ }).click()
  cy.findByText('API keys (0)').should('be.visible')
  cy.findByRole('link', { name: /Add API key/ }).should('be.visible')
  cy.findByTestId('error-notification').should('not.exist')

  cy.findByRole('link', { name: /General settings/ }).click()
  cy.findByText('Basic settings')
    .should('be.visible')
    .closest('[data-test-id="collapsible-section"]')
    .within(() => {
      cy.findByLabelText('Gateway ID').should('be.visible')
      cy.findByRole('button', { name: /Save changes/ }).should('be.visible')
      cy.findByTestId('error-notification').should('not.exist')
      cy.findByRole('button', { name: 'Collapse' }).click()
    })
  cy.findByText('LoRaWAN options')
    .should('be.visible')
    .closest('[data-test-id="collapsible-section"]')
    .within(() => {
      cy.findByRole('button', { name: 'Expand' }).click()
      cy.findByText('Frequency plan').should('be.visible')
      cy.findByRole('button', { name: /Save changes/ }).should('be.visible')
      cy.findByTestId('error-notification').should('not.exist')
      cy.findByRole('button', { name: 'Collapse' }).click()
    })

  cy.findByTestId('error-notification').should('not.exist')
})
export default [gatewaySubpages]
