

import { defineSmokeTest } from '../utils'

const organizationSubpages = defineSmokeTest('check all organization sub-pages', () => {
  const user = {
    ids: { user_id: 'org-sub-pages-test-user' },
    primary_email_address: 'test-user@example.com',
    password: 'ABCDefg123!',
    password_confirm: 'ABCDefg123!',
    email: 'org-subpage-test-user@example.com',
  }
  const organization = {
    ids: { organization_id: 'org-subpages-test' },
  }
  cy.createUser(user)
  cy.createOrganization(organization, user.ids.user_id)
  cy.loginConsole({ user_id: user.ids.user_id, password: user.password })
  cy.visit(Cypress.config('consoleRootPath'))

  cy.get('header').within(() => {
    cy.findByRole('link', { name: /Organizations/ }).click()
  })
  cy.findByRole('cell', { name: organization.ids.organization_id }).click()

  cy.findByRole('link', { name: /Live data/ }).click()
  cy.findByText(/Waiting for events from/).should('be.visible')
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
  cy.findByLabelText('Organization ID').should('be.visible')
  cy.findByRole('button', { name: 'Save changes' }).should('be.visible')
  cy.findByTestId('error-notification').should('not.exist')
})

export default [organizationSubpages]
