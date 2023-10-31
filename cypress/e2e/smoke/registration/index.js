

import { defineSmokeTest } from '../utils'

const emailConfirmationRegex = `https?:\\/\\/[a-zA-Z0-9-_.:]+/[a-zA-Z0-9-_]+\\/validate\\?.+&token=[A-Z0-9]+`

const loginConsole = defineSmokeTest('succeeds registering and logging into the Console', () => {
  const user = {
    user_id: 'console-login-test-user',
    name: 'Console Login Test User',
    password: '123456QWERTY!',
    email: 'console-login-test-user@example.com',
  }
  cy.visit(Cypress.config('consoleRootPath'))

  cy.findByRole('link', { name: 'Create an account' }).click()
  cy.findByLabelText('User ID').type(user.user_id)
  cy.findByLabelText('Name').type(user.name)
  cy.findByLabelText('Email').type(user.email)
  cy.findByLabelText('Password').type(user.password)
  cy.findByLabelText('Confirm password').type(user.password)
  cy.findByRole('button', { name: 'Create account' }).click()

  cy.findByTestId('notification')
    .should('be.visible')
    .should('contain', 'You have successfully registered and can login now')

  // Check for the validation email (via the stack logs).
  cy.task('findInLatestEmail', emailConfirmationRegex).then(validationUri => {
    // eslint-disable-next-line jest/valid-expect, no-unused-expressions
    expect(validationUri).to.not.be.empty
  })

  // Login.
  // TODO: https://github.com/TheThingsNetwork/lorawan-stack/issues/2923
  cy.visit(Cypress.config('consoleRootPath'))
  cy.findByLabelText('User ID').type(user.user_id)
  cy.findByLabelText('Password').type(`${user.password}`)
  cy.findByRole('button', { name: 'Login' }).click()

  cy.location('pathname').should('eq', `${Cypress.config('consoleRootPath')}/`)
  cy.findByTestId('full-error-view').should('not.exist')
})

const loginAccountApp = defineSmokeTest(
  'succeeds registering and logging into the Account app',
  () => {
    const user = {
      user_id: 'account-app-login-test-user',
      name: 'Account App Login Test User',
      password: '123456QWERTY!',
      email: 'account-app-login-test-user@example.com',
    }
    cy.visit(Cypress.config('accountAppRootPath'))

    cy.findByRole('link', { name: 'Create an account' }).click()
    cy.findByLabelText('User ID').type(user.user_id)
    cy.findByLabelText('Name').type(user.name)
    cy.findByLabelText('Email').type(user.email)
    cy.findByLabelText('Password').type(user.password)
    cy.findByLabelText('Confirm password').type(user.password)
    cy.findByRole('button', { name: 'Create account' }).click()

    cy.findByTestId('notification')
      .should('be.visible')
      .should('contain', 'You have successfully registered and can login now')

    // Check for the validation email (via the stack logs).
    cy.task('findInLatestEmail', emailConfirmationRegex).then(validationUri => {
      // eslint-disable-next-line jest/valid-expect, no-unused-expressions
      expect(validationUri).to.not.be.empty
    })

    cy.findByLabelText('User ID').type(user.user_id)
    cy.findByLabelText('Password').type(`${user.password}`)
    cy.findByRole('button', { name: 'Login' }).click()

    cy.location('pathname').should('eq', `${Cypress.config('accountAppRootPath')}/`)
    cy.findByTestId('full-error-view').should('not.exist')
  },
)

export default [loginConsole, loginAccountApp]
