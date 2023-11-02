
describe('Skip payload crypto', () => {
  const applicationId = 'spc-test-application'
  const application = { ids: { application_id: applicationId } }
  const userId = 'edit-app-spc-test-user'
  const user = {
    ids: { user_id: userId },
    primary_email_address: 'edit-spc-test-user@example.com',
    password: 'ABCDefg123!',
    password_confirm: 'ABCDefg123!',
  }

  let endDeviceId
  before(() => {
    cy.dropAndSeedDatabase()
    cy.createUser(user)
    cy.createApplication(application, userId)
    cy.createMockDeviceAllComponents(applicationId).then(body => {
      endDeviceId = body.end_device.ids.device_id
    })
  })

  describe('Uplink', () => {
    describe('application link skips payload crypto', () => {
      beforeEach(() => {
        cy.loginConsole({ user_id: userId, password: user.password })
      })

      it('disables messaging when not using a SPC overwrite', () => {
        const adminApiKey = Cypress.config('adminApiKey')
        const linkRequestBody = {
          field_mask: { paths: ['skip_payload_crypto'] },
          link: {
            default_formatters: {},
            skip_payload_crypto: true,
          },
        }
        cy.request({
          method: 'PUT',
          url: `api/v3/as/applications/${applicationId}/link`,
          body: linkRequestBody,
          headers: {
            Authorization: `Bearer ${adminApiKey}`,
          },
        })

        const endDeviceRequestBody = {
          field_mask: { paths: ['skip_payload_crypto_override'] },
          end_device: { skip_payload_crypto_override: null },
        }
        cy.request({
          method: 'PUT',
          url: `api/v3/as/applications/${applicationId}/devices/${endDeviceId}`,
          body: endDeviceRequestBody,
          headers: {
            Authorization: `Bearer ${adminApiKey}`,
          },
        })

        cy.visit(
          `${Cypress.config(
            'consoleRootPath',
          )}/applications/${applicationId}/devices/${endDeviceId}/messaging/uplink`,
        )

        cy.findByTestId('notification')
          .should('be.visible')
          .findByText(`Simulation is disabled for devices that skip payload crypto`)
          .should('be.visible')

        cy.findByRole('button', { name: 'Simulate uplink' }).should('be.disabled')
      })

      it('allows messaging when device disabled SPC via overwrite', () => {
        const adminApiKey = Cypress.config('adminApiKey')
        const linkRequestBody = {
          field_mask: { paths: ['skip_payload_crypto'] },
          link: {
            default_formatters: {},
            skip_payload_crypto: true,
          },
        }
        cy.request({
          method: 'PUT',
          url: `api/v3/as/applications/${applicationId}/link`,
          body: linkRequestBody,
          headers: {
            Authorization: `Bearer ${adminApiKey}`,
          },
        })

        const endDeviceRequestBody = {
          field_mask: { paths: ['skip_payload_crypto_override'] },
          end_device: { skip_payload_crypto_override: false },
        }
        cy.request({
          method: 'PUT',
          url: `api/v3/as/applications/${applicationId}/devices/${endDeviceId}`,
          body: endDeviceRequestBody,
          headers: {
            Authorization: `Bearer ${adminApiKey}`,
          },
        })

        cy.visit(
          `${Cypress.config(
            'consoleRootPath',
          )}/applications/${applicationId}/devices/${endDeviceId}/messaging/uplink`,
        )

        cy.findByRole('button', { name: 'Simulate uplink' }).should('be.enabled')
      })
    })

    describe('application link does not skip payload crypto', () => {
      beforeEach(() => {
        cy.loginConsole({ user_id: userId, password: user.password })
      })

      it('allows messaging when not using a SPC overwrite', () => {
        const adminApiKey = Cypress.config('adminApiKey')
        const linkRequestBody = {
          field_mask: { paths: ['skip_payload_crypto'] },
          link: {
            default_formatters: {},
            skip_payload_crypto: false,
          },
        }
        cy.request({
          method: 'PUT',
          url: `api/v3/as/applications/${applicationId}/link`,
          body: linkRequestBody,
          headers: {
            Authorization: `Bearer ${adminApiKey}`,
          },
        })

        const endDeviceRequestBody = {
          field_mask: { paths: ['skip_payload_crypto_override'] },
          end_device: { skip_payload_crypto_override: null },
        }
        cy.request({
          method: 'PUT',
          url: `api/v3/as/applications/${applicationId}/devices/${endDeviceId}`,
          body: endDeviceRequestBody,
          headers: {
            Authorization: `Bearer ${adminApiKey}`,
          },
        })

        cy.visit(
          `${Cypress.config(
            'consoleRootPath',
          )}/applications/${applicationId}/devices/${endDeviceId}/messaging/uplink`,
        )

        cy.findByRole('button', { name: 'Simulate uplink' }).should('be.enabled')
      })

      it('disables messaging when device disabled SPC via overwrite', () => {
        const adminApiKey = Cypress.config('adminApiKey')
        const linkRequestBody = {
          field_mask: { paths: ['skip_payload_crypto'] },
          link: {
            default_formatters: {},
            skip_payload_crypto: false,
          },
        }
        cy.request({
          method: 'PUT',
          url: `api/v3/as/applications/${applicationId}/link`,
          body: linkRequestBody,
          headers: {
            Authorization: `Bearer ${adminApiKey}`,
          },
        })

        const endDeviceRequestBody = {
          field_mask: { paths: ['skip_payload_crypto_override'] },
          end_device: { skip_payload_crypto_override: true },
        }
        cy.request({
          method: 'PUT',
          url: `api/v3/as/applications/${applicationId}/devices/${endDeviceId}`,
          body: endDeviceRequestBody,
          headers: {
            Authorization: `Bearer ${adminApiKey}`,
          },
        })

        cy.visit(
          `${Cypress.config(
            'consoleRootPath',
          )}/applications/${applicationId}/devices/${endDeviceId}/messaging/uplink`,
        )

        cy.findByTestId('notification')
          .should('be.visible')
          .findByText(`Simulation is disabled for devices that skip payload crypto`)
          .should('be.visible')

        cy.findByRole('button', { name: 'Simulate uplink' }).should('be.disabled')
      })
    })
  })
})
