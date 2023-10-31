

import { defineSmokeTest } from '../utils'

const applicationCreate = defineSmokeTest('succeeds creating application', () => {
  const user = {
    ids: { user_id: 'app-create-test-user' },
    primary_email_address: 'test-user@example.com',
    password: 'ABCDefg123!',
    password_confirm: 'ABCDefg123!',
    email: 'app-create-test-user@example.com',
  }
  cy.createUser(user)
  cy.loginConsole({ user_id: user.ids.user_id, password: user.password })
  cy.visit(Cypress.config('consoleRootPath'))

  const application = {
    application_id: 'app-create-test-app',
    name: 'Application Create Test',
    description: 'Application used in smoke test to verify application creation',
  }
  cy.get('header').within(() => {
    cy.findByRole('link', { name: /Applications/ }).click()
  })
  cy.findByRole('link', { name: /Create application/ }).click()
  cy.findByLabelText('Application ID').type(application.application_id)
  cy.findByLabelText('Application name').type(application.name)
  cy.findByLabelText('Description').type(application.description)
  cy.findByRole('button', { name: 'Create application' }).click()

  cy.location('pathname').should(
    'eq',
    `${Cypress.config('consoleRootPath')}/applications/${application.application_id}`,
  )
  cy.findByTestId('full-error-view').should('not.exist')
})

export default [applicationCreate]
