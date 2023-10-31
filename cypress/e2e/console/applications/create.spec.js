

import { disableApplicationServer } from '../../../support/utils'

describe('Application create', () => {
  let user
  const applicationId = 'test-application'

  before(() => {
    cy.dropAndSeedDatabase()
  })

  beforeEach(() => {
    user = {
      ids: { user_id: 'create-app-test-user' },
      primary_email_address: 'create-app-test-user@example.com',
      password: 'ABCDefg123!',
      password_confirm: 'ABCDefg123!',
    }
  })

  it('displays UI elements in place', () => {
    cy.createUser(user)
    cy.loginConsole({ user_id: user.ids.user_id, password: user.password })
    cy.visit(`${Cypress.config('consoleRootPath')}/applications/add`)

    cy.findByText('Create application', { selector: 'h1' }).should('be.visible')
    cy.findByLabelText('Application ID')
      .should('be.visible')
      .and('have.attr', 'placeholder')
      .and('eq', 'my-new-application')
    cy.findByLabelText('Application name')
      .should('be.visible')
      .and('have.attr', 'placeholder')
      .and('eq', 'My new application')
    cy.findByLabelText('Description')
      .should('be.visible')
      .and('have.attr', 'placeholder')
      .and('eq', 'Description for my new application')
    cy.findDescriptionByLabelText('Description')
      .should(
        'contain',
        'Optional application description; can also be used to save notes about the application',
      )
      .and('be.visible')
    cy.findByRole('button', { name: 'Create application' }).should('be.visible')
  })

  it('validates before submitting an empty form', () => {
    cy.loginConsole({ user_id: user.ids.user_id, password: user.password })
    cy.visit(`${Cypress.config('consoleRootPath')}/applications/add`)

    cy.findByRole('button', { name: 'Create application' }).should('be.visible').click()

    cy.findErrorByLabelText('Application ID')
      .should('contain.text', 'Application ID is required')
      .and('be.visible')
    cy.location('pathname').should('eq', `${Cypress.config('consoleRootPath')}/applications/add`)
  })

  it('succeeds adding application', () => {
    cy.loginConsole({ user_id: user.ids.user_id, password: user.password })
    cy.visit(`${Cypress.config('consoleRootPath')}/applications/add`)
    cy.findByLabelText('Application ID').type(applicationId)

    cy.findByRole('button', { name: 'Create application' }).click()

    cy.findByTestId('error-notification').should('not.exist')
    cy.findByTestId('full-error-view').should('not.exist')
    cy.location('pathname').should(
      'eq',
      `${Cypress.config('consoleRootPath')}/applications/${applicationId}`,
    )
  })

  describe('when has no Application Server in the local cluster', () => {
    beforeEach(() => {
      user = {
        ids: { user_id: 'create-app-no-as-test-user' },
        primary_email_address: 'create-app-no-as-test-user@example.com',
        password: 'ABCDefg123!',
        password_confirm: 'ABCDefg123!',
      }
    })

    beforeEach(() => {
      cy.augmentStackConfig(disableApplicationServer)
    })

    it('displays UI elements in place', () => {
      cy.createUser(user)
      cy.loginConsole({ user_id: user.ids.user_id, password: user.password })
      cy.visit(`${Cypress.config('consoleRootPath')}/applications/add`)

      cy.findByText('Create application', { selector: 'h1' }).should('be.visible')
      cy.findByLabelText('Application ID')
        .should('be.visible')
        .and('have.attr', 'placeholder')
        .and('eq', 'my-new-application')
      cy.findByLabelText('Application name')
        .should('be.visible')
        .and('have.attr', 'placeholder')
        .and('eq', 'My new application')
      cy.findByLabelText('Description')
        .should('be.visible')
        .and('have.attr', 'placeholder')
        .and('eq', 'Description for my new application')
      cy.findDescriptionByLabelText('Description')
        .should(
          'contain',
          'Optional application description; can also be used to save notes about the application',
        )
        .and('be.visible')
      cy.findByRole('button', { name: 'Create application' }).should('be.visible')
    })
  })
})
