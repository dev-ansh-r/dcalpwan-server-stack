

const user = {
  ids: { user_id: 'test-user' },
  name: 'Test User',
  primary_email_address: 'test-user@example.com',
  password: 'ABCDefg123!',
  password_confirm: 'ABCDefg123!',
}

describe('Account App session management', () => {
  before(() => {
    cy.dropAndSeedDatabase()
    cy.createUser(user)
    cy.loginAccountApp({ user_id: user.ids.user_id, password: user.password })
    cy.visit(Cypress.config('accountAppRootPath'))
    cy.loginAccountApp({ user_id: user.ids.user_id, password: user.password })
    cy.visit(Cypress.config('accountAppRootPath'))
  })

  it('succeeds showing a list of sessions', () => {
    cy.visit(`${Cypress.config('accountAppRootPath')}/session-management`)

    cy.findByRole('rowgroup').within(() => {
      cy.findAllByRole('row').should('have.length', 2)
    })
  })

  it('succeeds deleting a session', () => {
    cy.loginAccountApp({ user_id: user.ids.user_id, password: user.password })
    cy.visit(`${Cypress.config('accountAppRootPath')}/session-management`)

    cy.findByRole('rowgroup').within(() => {
      cy.findAllByRole('row').should('have.length', 3)
    })

    cy.findByRole('rowgroup').within(() => {
      cy.get('button', { name: 'deleteRemove this session' }).first().click()
    })

    cy.findByTestId('error-notification').should('not.exist')
    cy.findByTestId('toast-notification')
      .findByText('Session removed successfully')
      .should('be.visible')

    cy.findByRole('rowgroup').within(() => {
      cy.findAllByRole('row').should('have.length', 2)
    })
  })
})
