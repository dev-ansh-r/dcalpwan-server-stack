

describe('Account App user registration', () => {
  beforeEach(() => {
    cy.dropAndSeedDatabase()
    cy.visit(`${Cypress.config('accountAppRootPath')}/register`)
  })

  it('displays UI elements in place', () => {
    cy.findByText(Cypress.config('accountAppSiteName'), {
      selector: 'h1',
    }).should('be.visible')
    cy.findByText('Create a new account').should('be.visible')
    cy.findByLabelText('User ID').should('be.visible')
    cy.findByLabelText('Name').should('be.visible')
    cy.findByLabelText('Email').should('be.visible')
    cy.findByLabelText('Password').should('be.visible')
    cy.findByLabelText('Confirm password').should('be.visible')
    cy.findByRole('button', { name: 'Create account' }).should('be.visible')
    cy.findByRole('link', { name: 'Login' }).should('be.visible')
  })

  it('validates before submitting an empty form', () => {
    cy.visit(`${Cypress.config('accountAppRootPath')}/register`)
    cy.findByRole('button', { name: 'Create account' }).click()

    cy.findErrorByLabelText('User ID')
      .should('contain.text', 'User ID is required')
      .and('be.visible')

    cy.findErrorByLabelText('Email').should('contain.text', 'Email is required').and('be.visible')

    cy.findErrorByLabelText('Password')
      .should('contain.text', 'Password is required')
      .and('be.visible')
    cy.findErrorByLabelText('Confirm password')
      .should('contain.text', 'Confirm password is required')
      .and('be.visible')

    cy.location('pathname').should('eq', `${Cypress.config('accountAppRootPath')}/register`)
  })

  it('succeeds when using valid user data for `approved` users', () => {
    cy.findByLabelText('User ID').type('test-user')
    cy.findByLabelText('Name').type('Test User')
    cy.findByLabelText('Email').type('mail@example.com')
    cy.findByLabelText('Password').type('ABCdefg123456!')
    cy.findByLabelText('Confirm password').type('ABCdefg123456!{enter}')

    cy.findByTestId('notification')
      .should('be.visible')
      .should('contain', 'You have successfully registered and can login now')
    cy.location('pathname').should('include', `${Cypress.config('accountAppRootPath')}/login`)
  })
})
