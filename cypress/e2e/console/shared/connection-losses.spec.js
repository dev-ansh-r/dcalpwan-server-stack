

describe('Connection loss detection', () => {
  const userId = 'connection-loss-test-user'
  const user = {
    ids: { user_id: userId },
    primary_email_address: 'connection-loss-test-user@example.com',
    password: 'ABCDefg123!',
    password_confirm: 'ABCDefg123!',
  }

  before(() => {
    cy.dropAndSeedDatabase()
    cy.createUser(user)
  })

  it('detects connection losses and attempts reconnects', () => {
    cy.loginConsole({ user_id: user.ids.user_id, password: user.password })
    cy.visit(Cypress.config('consoleRootPath'))
    cy.findByText('Welcome to the Console!').should('be.visible')

    cy.intercept('/api/v3/application*', { forceNetworkError: true }).as('offlineIntercept')
    cy.intercept('/api/v3/auth_info', { times: 2 }, { forceNetworkError: true }).as(
      'reconnectionIntercept',
    )

    cy.get('header').within(() => {
      cy.findByRole('link', { name: /Applications/ }).click()
    })

    cy.get('footer').within(() => {
      cy.findByText(/Connection issues/).should('be.visible')
      cy.findByText(/Offline/).should('be.visible')
    })

    cy.findByTestId('toast-notification')
      .findByText(/offline/)
      .as('offlineToast')
      .should('be.visible')

    // Use an assertion to check that the 'offline' toast notification is no longer in the DOM.
    cy.get('@offlineToast').should('not.exist')

    // After the 'offline' toast has disappeared, wait for the reconnection intercept to resolve.
    cy.wait('@reconnectionIntercept')

    // Now the 'online' toast should appear.
    cy.findByTestId('toast-notification')
      .findByText(/online/)
      .should('be.visible')

    cy.get('footer').within(() => {
      cy.findByText(/Connection issues/).should('not.exist')
      cy.findByText(/Offline/).should('not.exist')
    })
  })
  it('does not see individual network errors as connection loss', () => {
    cy.loginConsole({ user_id: user.ids.user_id, password: user.password })
    cy.visit(Cypress.config('consoleRootPath'))
    cy.findByText('Welcome to the Console!').should('be.visible')
    cy.intercept('/api/v3/application*', { forceNetworkError: true })

    cy.get('header').within(() => {
      cy.findByRole('link', { name: /Applications/ }).click()
    })

    cy.get('footer').within(() => {
      // Connection issue note will appear in the footer and
      // dissappear shortly thereafter.
      cy.findByText(/Connection issues/).should('be.visible')
      cy.findByText(/Connection issues/).should('not.exist')

      cy.findByText(/Offline/).should('not.exist')
    })

    cy.findByTestId('toast-notification').should('not.exist')

    // The error will be displayed by the consuming view.
    cy.findByTestId('error-notification').should('be.visible')
  })
})
