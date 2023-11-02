

import { defineSmokeTest } from '../utils'

const organizationDelete = defineSmokeTest('succeeds deleting an organization', () => {
  const user = {
    ids: { user_id: 'org-delete-test-user' },
    primary_email_address: 'test-user@example.com',
    password: 'ABCDefg123!',
    password_confirm: 'ABCDefg123!',
    email: 'org-delete-test-user@example.com',
  }
  const organization = {
    ids: { organization_id: 'org-delete-test' },
  }
  cy.createUser(user)
  cy.createOrganization(organization, user.ids.user_id)
  cy.loginConsole({ user_id: user.ids.user_id, password: user.password })
  cy.visit(Cypress.config('consoleRootPath'))

  cy.get('header').within(() => {
    cy.findByRole('link', { name: /Organizations/ }).click()
  })
  cy.findByRole('rowgroup').within(() => {
    cy.findByRole('cell', { name: organization.ids.organization_id }).click()
  })
  cy.findByRole('link', { name: /General settings/ }).click()
  cy.findByRole('button', { name: /Delete organization/ }).click()
  cy.findByTestId('modal-window').within(() => {
    cy.findByRole('button', { name: /Delete organization/ }).click()
  })

  cy.findByTestId('full-error-view').should('not.exist')

  cy.location('pathname').should('eq', `${Cypress.config('consoleRootPath')}/organizations`)
  cy.findByRole('table').within(() => {
    cy.findByRole('cell', { name: organization.ids.organization_id }).should('not.exist')
  })
})

export default [organizationDelete]
