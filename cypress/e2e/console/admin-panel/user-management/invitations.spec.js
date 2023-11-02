

describe('Send invite', () => {
  before(() => {
    cy.dropAndSeedDatabase()
  })

  beforeEach(() => {
    cy.loginConsole({ user_id: 'admin', password: 'admin' })
  })

  it('displays UI elements in place', () => {
    cy.visit(`${Cypress.config('consoleRootPath')}/admin-panel/user-management/invitations/add`)

    cy.findByText('Invite', { selector: 'h1' }).should('be.visible')
    cy.findByLabelText('Email address')
      .should('be.visible')
      .and('have.attr', 'placeholder')
      .and('eq', 'mail@example.com')
    cy.findByRole('button', { name: 'Invite' }).should('be.visible')
  })

  it('validates before submitting an empty form', () => {
    cy.visit(`${Cypress.config('consoleRootPath')}/admin-panel/user-management/invitations/add`)

    cy.findByRole('button', { name: 'Invite' }).should('be.visible').click()

    cy.findErrorByLabelText('Email address')
      .should('contain.text', 'Email address is required')
      .and('be.visible')
    cy.location('pathname').should(
      'eq',
      `${Cypress.config('consoleRootPath')}/admin-panel/user-management/invitations/add`,
    )
  })

  it('succeeds inviting a user', () => {
    cy.visit(`${Cypress.config('consoleRootPath')}/admin-panel/user-management/invitations/add`)
    cy.findByLabelText('Email address').type('mail@example.com')

    cy.findByRole('button', { name: 'Invite' }).click()

    cy.findByTestId('error-notification').should('not.exist')
    cy.findByTestId('full-error-view').should('not.exist')
    cy.location('pathname').should(
      'eq',
      `${Cypress.config('consoleRootPath')}/admin-panel/user-management`,
    )
    cy.findByText('User invitations').click()
    cy.findByRole('cell', { name: 'mail@example.com' }).should('be.visible')
  })
})
