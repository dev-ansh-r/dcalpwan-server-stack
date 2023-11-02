

describe('Account App authorization management', () => {
  const user = {
    ids: { user_id: 'create-app-test-user' },
    primary_email_address: 'create-app-test-user@example.com',
    password: 'ABCDefg123!',
    password_confirm: 'ABCDefg123!',
  }

  beforeEach(() => {
    cy.dropAndSeedDatabase()
    cy.createUser(user)
    cy.visitConsoleAuthorizationScreen({ user_id: user.ids.user_id, password: user.password })
    cy.findByRole('button', { name: /Authorize Console/ }).click()
  })

  it('succeeds showing authorized clients', () => {
    cy.visit(`${Cypress.config('accountAppRootPath')}/client-authorizations`)

    cy.findByRole('rowgroup').within(() => {
      cy.findAllByRole('row').should('have.length', 1)
    })
    cy.findByRole('cell', { name: 'console' }).should('be.visible')
  })

  it('succeeds viewing an authorized client', () => {
    cy.loginAccountApp({ user_id: user.ids.user_id, password: user.password })
    cy.visit(`${Cypress.config('accountAppRootPath')}/client-authorizations/console`)
    cy.findByText('Authorization settings').should('be.visible')
    cy.findByText('Active access tokens').should('be.visible')
    cy.findByText('General information').should('be.visible')
    cy.findByText('Actions').should('be.visible')
  })

  it('succeeds showing active access tokens', () => {
    cy.loginAccountApp({ user_id: user.ids.user_id, password: user.password })
    cy.visit(`${Cypress.config('accountAppRootPath')}/client-authorizations/console/access-tokens`)
    cy.findByRole('rowgroup').within(() => {
      cy.findAllByRole('row').should('have.length', 1)
    })
  })

  it('succeeds revoking authorizations', () => {
    cy.loginAccountApp({ user_id: user.ids.user_id, password: user.password })
    cy.visit(`${Cypress.config('accountAppRootPath')}/client-authorizations/console`)
    cy.findByRole('button', { name: /Revoke authorization/ }).click()
    cy.findByTestId('modal-window')
      .should('be.visible')
      .within(() => {
        cy.findByText('Revoke authorization', { selector: 'h1' }).should('be.visible')
        cy.findByRole('button', { name: /Revoke authorization/ }).click()
      })
    cy.findByTestId('error-notification').should('not.exist')

    cy.location('pathname').should(
      'eq',
      `${Cypress.config('accountAppRootPath')}/client-authorizations`,
    )

    cy.findByRole('cell', { name: 'console' }).should('not.exist')
  })

  it('succeeds invalidating all tokens', () => {
    cy.loginAccountApp({ user_id: user.ids.user_id, password: user.password })
    cy.visit(`${Cypress.config('accountAppRootPath')}/client-authorizations/console/access-tokens`)
    cy.findByRole('button', { name: /Invalidate all access tokens/ }).click()
    cy.findByTestId('error-notification').should('not.exist')

    cy.visit(`${Cypress.config('accountAppRootPath')}/client-authorizations/console/access-tokens`)
    cy.findByRole('rowgroup').should('not.exist')
  })

  it('succeeds invalidating one token', () => {
    cy.loginAccountApp({ user_id: user.ids.user_id, password: user.password })
    cy.visit(`${Cypress.config('accountAppRootPath')}/client-authorizations/console/access-tokens`)
    cy.findByRole('button', { name: /Invalidate this access token/ }).click()
    cy.findByTestId('error-notification').should('not.exist')
    cy.findByRole('rowgroup').should('not.exist')
  })
})
