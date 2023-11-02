
import { defineSmokeTest } from '../utils'

const validatePasswordLinkRegExp = `https?:\\/\\/[a-zA-Z0-9-_.:]+/[a-zA-Z0-9-_]+\\/validate\\?.+&token=[A-Z0-9]+`

const contactInfoValidation = defineSmokeTest('succeeds validating contact info', () => {
  const user1 = {
    ids: { user_id: 'test-user-id1' },
    primary_email_address: 'test-user1@example.com',
    password: 'ABCDefg123!',
    password_confirm: 'ABCDefg123!',
  }
  const user2 = {
    ids: { user_id: 'test-user-id2' },
    primary_email_address: 'test-user2@example.com',
    password: 'ABCDefg123!',
    password_confirm: 'ABCDefg123!',
  }

  cy.createUser(user1)

  cy.task('findInLatestEmail', validatePasswordLinkRegExp).then(validationUri => {
    cy.log(validationUri)
    cy.visit(validationUri)
    cy.findByTestId('notification').should('be.visible').should('contain', 'Validation successful')
  })

  cy.createUser(user2)

  cy.task('findInLatestEmail', validatePasswordLinkRegExp).then(validationUri => {
    cy.log(validationUri)
    cy.visit(validationUri)
    cy.findByTestId('notification').should('be.visible').should('contain', 'Validation successful')
  })

  cy.reload()
  cy.findByTestId('error-notification').should('be.visible').should('contain', 'token already used')
})

export default [contactInfoValidation]
