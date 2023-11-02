
describe('OAuth Clients list', () => {
  const userId = 'list-clients-test-user'
  const user = {
    ids: { user_id: 'list-clients-test-user' },
    primary_email_address: 'list-clients-test-user@example.com',
    password: 'ABCDefg123!',
    password_confirm: 'ABCDefg123!',
  }

  const clients = [
    {
      ids: { client_id: 'xyz-test-client' },
      grants: ['GRANT_AUTHORIZATION_CODE'],
    },
    {
      ids: { client_id: 'other-test-client' },
      grants: ['GRANT_AUTHORIZATION_CODE'],
    },
    {
      ids: { client_id: 'nice-test-client' },
      grants: ['GRANT_AUTHORIZATION_CODE'],
    },
  ]

  before(() => {
    cy.dropAndSeedDatabase()
    cy.createUser(user)
    cy.createClient(clients[0], userId)
    cy.createClient(clients[1], userId)
    cy.createClient(clients[2], userId)
  })

  beforeEach(() => {
    cy.loginConsole({ user_id: user.ids.user_id, password: user.password })
    cy.visit(`${Cypress.config('accountAppRootPath')}/oauth-clients`)
  })

  it('succeeds searching by client id', () => {
    cy.findByRole('rowgroup').within(() => {
      cy.findAllByRole('row').should('have.length', 3)
    })
    cy.findByRole('cell', { name: clients[0].ids.client_id }).should('be.visible')
    cy.findByRole('cell', { name: clients[1].ids.client_id }).should('be.visible')
    cy.findByRole('cell', { name: clients[2].ids.client_id }).should('be.visible')

    cy.findByTestId('search-input').as('searchInput')
    cy.get('@searchInput').type('xyz')

    cy.findByRole('rowgroup').within(() => {
      cy.findAllByRole('row').should('have.length', 1)
    })
    cy.findByRole('cell', { name: clients[0].ids.client_id }).should('be.visible')
    cy.findByRole('cell', { name: clients[1].ids.client_id }).should('not.exist')
    cy.findByRole('cell', { name: clients[2].ids.client_id }).should('not.exist')

    cy.get('@searchInput').clear()
    cy.get('@searchInput').type('other')

    cy.findByRole('rowgroup').within(() => {
      cy.findByRole('row').should('have.length', 1)
    })
    cy.findByRole('cell', { name: clients[0].ids.client_id }).should('not.exist')
    cy.findByRole('cell', { name: clients[1].ids.client_id }).should('be.visible')
    cy.findByRole('cell', { name: clients[2].ids.client_id }).should('not.exist')

    cy.get('@searchInput').clear()
    cy.get('@searchInput').type('test-client')

    cy.findByRole('rowgroup').within(() => {
      cy.findAllByRole('row').should('have.length', 3)
    })
    cy.findByRole('cell', { name: clients[0].ids.client_id }).should('be.visible')
    cy.findByRole('cell', { name: clients[1].ids.client_id }).should('be.visible')
    cy.findByRole('cell', { name: clients[2].ids.client_id }).should('be.visible')
  })
})
