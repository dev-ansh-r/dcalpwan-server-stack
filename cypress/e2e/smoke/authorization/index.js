
import { defineSmokeTest } from '../utils'

const authorizeConsole = defineSmokeTest(
  'succeeds manually authorizing Account App clients',
  () => {
    const user = {
      ids: { user_id: 'authorize-test-user' },
      primary_email_address: 'create-app-test-user@example.com',
      password: 'ABCDefg123!',
      password_confirm: 'ABCDefg123!',
    }

    cy.createUser(user)

    // Set Console client's `skip_authorization` flag to false.
    cy.task('execSql', `UPDATE clients SET skip_authorization=false WHERE client_id='console';`)
    cy.visit(Cypress.config('consoleRootPath'))

    // Login.
    cy.findByLabelText('User ID').type(user.ids.user_id)
    cy.findByLabelText('Password').type(user.password)
    cy.findByRole('button', { name: 'Login' }).click()

    // Authorize.
    cy.location('pathname').should('contain', `${Cypress.config('accountAppRootPath')}/authorize`)
    cy.findByRole('button', { name: /Authorize Console/ }).click()

    // Check Console landing view.
    cy.location('pathname').should('eq', `${Cypress.config('consoleRootPath')}/`)
    cy.findByTestId('full-error-view').should('not.exist')
  },
)

const abortAuthorization = defineSmokeTest(
  'succeeds manually aborting Account App client authorization',
  () => {
    const user = {
      ids: { user_id: 'authorize-test-user' },
      primary_email_address: 'create-app-test-user@example.com',
      password: 'ABCDefg123!',
      password_confirm: 'ABCDefg123!',
    }

    cy.createUser(user)

    // Set Console client's `skip_authorization` flag to false.
    cy.task('execSql', `UPDATE clients SET skip_authorization=false WHERE client_id='console';`)
    cy.visit(Cypress.config('consoleRootPath'))

    // Login.
    cy.findByLabelText('User ID').type(user.ids.user_id)
    cy.findByLabelText('Password').type(user.password)
    cy.findByRole('button', { name: 'Login' }).click()

    // Deny authorization.
    cy.location('pathname').should('contain', `${Cypress.config('accountAppRootPath')}/authorize`)
    cy.findByRole('button', { name: /Cancel/ }).click()

    // Check Console error.
    cy.location('pathname').should('eq', `${Cypress.config('consoleRootPath')}/oauth/callback`)
    cy.findByTestId('full-error-view').should('exist')
    cy.findByText(/Login failed/).should('be.visible')
  },
)

export default [authorizeConsole, abortAuthorization]
