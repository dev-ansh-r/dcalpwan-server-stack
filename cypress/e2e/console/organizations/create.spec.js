
describe('Organization create', () => {
  const organizationId = 'test-organization'
  const userId = 'create-organization-test-user'
  const user = {
    ids: { user_id: userId },
    primary_email_address: 'edit-organization-test-user@example.com',
    password: 'ABCDefg123!',
    password_confirm: 'ABCDefg123!',
  }

  before(() => {
    cy.dropAndSeedDatabase()
    cy.createUser(user)
  })

  beforeEach(() => {
    cy.loginConsole({ user_id: user.ids.user_id, password: user.password })
    cy.visit(`${Cypress.config('consoleRootPath')}/organizations/add`)
  })

  it('displays UI elements in place', () => {
    cy.findByText('Create organization', { selector: 'h1' }).should('be.visible')
    cy.findByLabelText('Organization ID').should('be.visible')
    cy.findByLabelText('Name').should('be.visible')
    cy.findByLabelText('Description').should('be.visible')
    cy.findDescriptionByLabelText('Description')
      .should(
        'contain',
        'Optional organization description; can also be used to save notes about the organization',
      )
      .and('be.visible')
    cy.findByRole('button', { name: 'Create organization' }).should('be.visible')
  })

  it('validates before submitting an empty form', () => {
    cy.findByRole('button', { name: 'Create organization' }).should('be.visible').click()

    cy.findErrorByLabelText('Organization ID')
      .should('contain.text', 'Organization ID is required')
      .and('be.visible')
    cy.location('pathname').should('eq', `${Cypress.config('consoleRootPath')}/organizations/add`)
  })

  it('succeeds adding organization', () => {
    cy.findByLabelText('Organization ID').type(organizationId)

    cy.findByRole('button', { name: 'Create organization' }).should('be.visible').click()

    cy.findByTestId('error-notification').should('not.exist')
    cy.findByTestId('full-error-view').should('not.exist')
    cy.location('pathname').should(
      'eq',
      `${Cypress.config('consoleRootPath')}/organizations/${organizationId}`,
    )
  })
})
