

import { defineSmokeTest } from '../utils'

const applicationFeatureToggles = defineSmokeTest(
  'restricts access to restricted content correctly',
  () => {
    const user = {
      ids: { user_id: 'feature-toggle-test-user' },
      primary_email_address: 'test-user@example.com',
      password: 'ABCDefg123!',
      password_confirm: 'ABCDefg123!',
      email: 'feature-toggle-test-user@example.com',
    }
    const application = { ids: { application_id: 'feature-toggle-test-app' } }
    const rights = [
      'RIGHT_APPLICATION_DELETE',
      'RIGHT_APPLICATION_DEVICES_READ',
      'RIGHT_APPLICATION_DEVICES_READ_KEYS',
      'RIGHT_APPLICATION_DEVICES_WRITE',
      'RIGHT_APPLICATION_DEVICES_WRITE_KEYS',
      'RIGHT_APPLICATION_INFO',
      'RIGHT_APPLICATION_LINK',
      'RIGHT_APPLICATION_SETTINGS_BASIC',
      'RIGHT_APPLICATION_SETTINGS_COLLABORATORS',
      'RIGHT_APPLICATION_SETTINGS_PACKAGES',
      'RIGHT_APPLICATION_TRAFFIC_DOWN_WRITE',
      'RIGHT_APPLICATION_TRAFFIC_READ',
      'RIGHT_APPLICATION_TRAFFIC_UP_WRITE',
    ]
    const collaborator = {
      ids: { user_id: 'feature-toggle-test-collaborator' },
      primary_email_address: 'test-collaborator@example.com',
      password: 'ABCDefg123!',
      password_confirm: 'ABCDefg123!',
      email: 'feature-toggle-test-collaborator@example.com',
    }
    cy.createUser(user)
    cy.createUser(collaborator)
    cy.createApplication(application, user.ids.user_id)
    cy.setApplicationCollaborator(application.ids.application_id, collaborator.ids.user_id, rights)

    cy.loginConsole({ user_id: collaborator.ids.user_id, password: collaborator.password })
    cy.visit(`${Cypress.config('consoleRootPath')}/applications/${application.ids.application_id}`)

    cy.findByTestId('navigation-sidebar').within(() => {
      cy.findByText('API Keys').should('not.exist')
      cy.findByText('Collaborators', { selector: 'span' }).should('be.visible')
    })

    cy.visit(
      `${Cypress.config('consoleRootPath')}/applications/${
        application.ids.application_id
      }/api-keys`,
    )
    cy.location('pathname').should(
      'eq',
      `${Cypress.config('consoleRootPath')}/applications/${application.ids.application_id}`,
    )
  },
)

export default [applicationFeatureToggles]
