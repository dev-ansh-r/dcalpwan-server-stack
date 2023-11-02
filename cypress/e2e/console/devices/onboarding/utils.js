

export const selectDevice = ({ brand_id, model_id, hw_version, fw_version, band_id }) => {
  cy.findByLabelText('End device brand').selectOption(brand_id)
  cy.findByLabelText('Model').selectOption(model_id)
  cy.findByLabelText('Hardware Ver.').selectOption(hw_version)
  cy.findByLabelText('Firmware Ver.').selectOption(fw_version)
  cy.findByLabelText('Profile (Region)').selectOption(band_id)
}

export const interceptDeviceRepo = appId => {
  cy.intercept(
    'GET',
    `/api/v3/dr/applications/${appId}/brands/test-brand-otaa/models/test-model4/1.0/EU_863_870/template`,
    { fixture: 'console/devices/repository/test-brand-otaa-model4.template.json' },
  )
  cy.intercept(
    'GET',
    `/api/v3/dr/applications/${appId}/brands/test-brand-otaa/models/test-model3/1.0.1/EU_863_870/template`,
    { fixture: 'console/devices/repository/test-brand-otaa-model3.template.json' },
  )
  cy.intercept(
    'GET',
    `/api/v3/dr/applications/${appId}/brands/test-brand-otaa/models/test-model2/1.0/EU_863_870/template`,
    { fixture: 'console/devices/repository/test-brand-otaa-model2.template.json' },
  )
  cy.intercept('GET', `/api/v3/dr/applications/${appId}/brands/test-brand-otaa/models*`, {
    fixture: 'console/devices/repository/test-brand-otaa.models.json',
  })
  cy.intercept('GET', `/api/v3/dr/applications/${appId}/brands*`, {
    fixture: 'console/devices/repository/brands.json',
  })
}

export const composeQRGeneratorParseResponse = ({ joinEui, devEui, cac, vendorId }) => ({
  format_id: 'tr005',
  end_device_template: {
    end_device: {
      ids: { dev_eui: devEui, join_eui: joinEui },
      claim_authentication_code: { value: cac },
      lora_alliance_profile_ids: {
        vendor_id: vendorId,
      },
    },
    field_mask: { paths: ['ids', 'claim_authentication_code'] },
  },
})

export const composeClaimResponse = ({ joinEui, devEui, id, appId }) => ({
  application_ids: { application_id: appId },
  device_id: id,
  dev_eui: devEui,
  join_eui: joinEui,
})

export const composeExpectedRequest = ({ joinEui, devEui, cac, id, appId }) => ({
  authenticated_identifiers: {
    join_eui: joinEui.toUpperCase(),
    dev_eui: devEui.toUpperCase(),
    authentication_code: cac,
  },
  target_device_id: id,
  target_application_ids: { application_id: appId },
})
