
describe('Account App forgot password', () => {
  before(() => {
    cy.dropAndSeedDatabase()
  })

  it('displays UI elements in place', () => {
    cy.visit(`${Cypress.config('accountAppRootPath')}/forgot-password`)
    cy.findByText('Reset password').should('be.visible')
    cy.findByLabelText('User ID').should('be.visible')
    cy.findByRole('button', { name: 'Send' }).should('be.visible')
    cy.findByRole('link', { name: 'Cancel' }).should('be.visible')
  })

  it('validates before submitting the form', () => {
    cy.visit(`${Cypress.config('accountAppRootPath')}/forgot-password`)
    cy.findByRole('button', { name: 'Send' }).click()
    cy.findErrorByLabelText('User ID')
      .should('contain.text', 'User ID is required')
      .and('be.visible')
  })

  it('succeeds when submitting the form', () => {
    const user = {
      ids: { user_id: 'test-user-id1' },
      primary_email_address: 'test-user1@example.com',
      password: 'ABCDefg123!',
      password_confirm: 'ABCDefg123!',
    }
    cy.createUser(user)
    cy.visit(`${Cypress.config('accountAppRootPath')}/forgot-password`)
    cy.findByLabelText('User ID').type(user.ids.user_id)
    cy.findByRole('button', { name: 'Send' }).click()
    cy.findByTestId('notification').should('be.visible').should('contain', 'reset instruction')
    cy.location('pathname').should('include', `${Cypress.config('accountAppRootPath')}/login`)
  })
})
