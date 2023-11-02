
import { defineSmokeTest } from '../utils'

const organizationCreate = defineSmokeTest('succeeds creating organization', () => {
  const user = {
    ids: { user_id: 'org-create-test-user' },
    primary_email_address: 'test-user@example.com',
    password: 'ABCDefg123!',
    password_confirm: 'ABCDefg123!',
    email: 'org-create-test-user@example.com',
  }
  cy.createUser(user)
  cy.loginConsole({ user_id: user.ids.user_id, password: user.password })
  cy.visit(Cypress.config('consoleRootPath'))

  const organization = {
    organization_id: 'org-create-test',
    name: 'Organization Create Test',
    description: 'Organization used in smoke test to verify organization creation',
  }

  cy.get('header').within(() => {
    cy.findByRole('link', { name: /Organizations/ }).click()
  })
  cy.findByRole('link', { name: /Create organization/ }).click()
  cy.findByLabelText('Organization ID').type(organization.organization_id)
  cy.findByLabelText('Name').type(organization.name)
  cy.findByLabelText('Description').type(organization.description)
  cy.findByRole('button', { name: 'Create organization' }).click()

  cy.location('pathname').should(
    'eq',
    `${Cypress.config('consoleRootPath')}/organizations/${organization.organization_id}`,
  )

  cy.findByTestId('error-notification').should('not.exist')
  cy.findByTestId('full-error-view').should('not.exist')
})

export default [organizationCreate]
