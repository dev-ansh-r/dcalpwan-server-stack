

describe('Collaborators', () => {
  const userId = 'main-collab-user'
  const user = {
    ids: { user_id: userId },
    primary_email_address: 'main-collab-user@example.com',
    password: 'ABCDefg123!',
    password_confirm: 'ABCDefg123!',
  }
  const collaboratorId = 'collab-test-user'
  const collaboratorUser = {
    ids: { user_id: collaboratorId },
    primary_email_address: 'collab-test-user@example.com',
    password: 'ABCDefg123!',
    password_confirm: 'ABCDefg123!',
  }
  const orgUserId = 'org-test-user'
  const orgUser = {
    ids: { user_id: orgUserId },
    primary_email_address: 'org-test-user@example.com',
    password: 'ABCDefg123!',
    password_confirm: 'ABCDefg123!',
  }
  const organizationId = 'test-collab-org'
  const organization = {
    ids: { organization_id: organizationId },
  }

  before(() => {
    cy.dropAndSeedDatabase()

    cy.createUser(user)
    cy.createUser(collaboratorUser)
    cy.createUser(orgUser)
    cy.createOrganization(organization, orgUserId)
  })

  describe('Application', () => {
    const applicationId = 'collab-test-app'
    const application = { ids: { application_id: applicationId } }

    before(() => {
      cy.createApplication(application, userId)
    })

    beforeEach(() => {
      cy.loginConsole({ user_id: userId, password: user.password })
      cy.visit(
        `${Cypress.config('consoleRootPath')}/applications/${applicationId}/collaborators/add`,
      )
    })

    it('succeeds adding user as a collaborator', () => {
      cy.findByLabelText('Collaborator').selectOption(collaboratorId)
      cy.findByLabelText('Grant all current and future rights').check()
      cy.findByRole('button', { name: 'Add collaborator' }).click()

      cy.findByTestId('error-notification').should('not.exist')
      cy.location('pathname').should(
        'eq',
        `${Cypress.config('consoleRootPath')}/applications/${applicationId}/collaborators`,
      )
      cy.findByText(collaboratorId).should('be.visible')
    })

    it('succeeds adding organization as a collaborator', () => {
      cy.findByLabelText('Collaborator').selectOption(organizationId)
      cy.findByLabelText('Grant all current and future rights').check()
      cy.findByRole('button', { name: 'Add collaborator' }).click()

      cy.findByTestId('error-notification').should('not.exist')
      cy.location('pathname').should(
        'eq',
        `${Cypress.config('consoleRootPath')}/applications/${applicationId}/collaborators`,
      )
      cy.findByText(organizationId).should('be.visible')
    })
  })

  describe('Gateway', () => {
    const gatewayId = 'collab-test-gtw'
    const gateway = { ids: { gateway_id: gatewayId } }

    before(() => {
      cy.createGateway(gateway, userId)
    })

    beforeEach(() => {
      cy.loginConsole({ user_id: userId, password: user.password })
      cy.visit(`${Cypress.config('consoleRootPath')}/gateways/${gatewayId}/collaborators/add`)
    })

    it('succeeds adding user as a collaborator', () => {
      cy.findByLabelText('Collaborator').selectOption(collaboratorId)
      cy.findByLabelText('Grant all current and future rights').check()
      cy.findByRole('button', { name: 'Add collaborator' }).click()

      cy.findByTestId('error-notification').should('not.exist')
      cy.location('pathname').should(
        'eq',
        `${Cypress.config('consoleRootPath')}/gateways/${gatewayId}/collaborators`,
      )
      cy.findByText(collaboratorId).should('be.visible')
    })

    it('succeeds adding organization as a collaborator', () => {
      cy.findByLabelText('Collaborator').selectOption(organizationId)

      cy.findByLabelText('Grant all current and future rights').check()
      cy.findByRole('button', { name: 'Add collaborator' }).click()

      cy.findByTestId('error-notification').should('not.exist')
      cy.location('pathname').should(
        'eq',
        `${Cypress.config('consoleRootPath')}/gateways/${gatewayId}/collaborators`,
      )
      cy.findByText(organizationId).should('be.visible')
    })
  })

  describe('Organization', () => {
    const testOrgId = 'collab-test-org-2'
    const testOrg = {
      ids: { organization_id: testOrgId },
    }

    before(() => {
      cy.createOrganization(testOrg, userId)
    })

    beforeEach(() => {
      cy.loginConsole({ user_id: userId, password: user.password })
      cy.visit(`${Cypress.config('consoleRootPath')}/organizations/${testOrgId}/collaborators/add`)
    })

    it('succeeds adding user as a collaborator', () => {
      cy.findByLabelText('Collaborator').selectOption(collaboratorId)
      cy.findByLabelText('Grant all current and future rights').check()
      cy.findByRole('button', { name: 'Add collaborator' }).click()

      cy.findByTestId('error-notification').should('not.exist')
      cy.location('pathname').should(
        'eq',
        `${Cypress.config('consoleRootPath')}/organizations/${testOrgId}/collaborators`,
      )
      cy.findByText(collaboratorId).should('be.visible')
    })

    it('fails adding organization as a collaborator', () => {
      cy.findByLabelText('Collaborator').type(organizationId)
      cy.findByText('No matching user or organization was found')
    })
  })
})
