

/* eslint-disable import/no-commonjs */

const { defineConfig } = require('cypress')
const { cypressBrowserPermissionsPlugin } = require('cypress-browser-permissions')
const cypressLogToOutput = require('cypress-log-to-output')

const tasks = require('../cypress/plugins/tasks')

module.exports = defineConfig({
  projectId: 'uqdraf',
  viewportHeight: 900,
  viewportWidth: 1440,
  defaultCommandTimeout: 10000,
  animationDistanceThreshold: 3,
  video: false,
  retries: {
    runMode: 1,
    openMode: 0,
  },
  env: {
    browserPermissions: {
      camera: 'allow',
    },
  },
  e2e: {
    setupNodeEvents: (on, config) => {
      const configWithPermissions = cypressBrowserPermissionsPlugin(on, config)

      tasks.stackConfigTask(on, configWithPermissions)
      tasks.sqlTask(on, configWithPermissions)
      tasks.fileExistsTask(on, configWithPermissions)
      tasks.emailTask(on, configWithPermissions)

      on('before:browser:launch', (browser = {}, launchOptions) => {
        // Log console log to output when debug mode is enabled.
        if (process.env.RUNNER_DEBUG) {
          launchOptions.args = cypressLogToOutput.browserLaunchHandler(browser, launchOptions.args)
        }
        if (browser.family === 'chromium' && browser.name !== 'electron') {
          launchOptions.args.push(
            '--use-file-for-fake-video-capture=cypress/fixtures/qr-code-mock-feed.y4m',
          )
        }

        if (browser.name === 'chrome' && browser.isHeadless) {
          launchOptions.args.push('--disable-gpu')
        }

        return launchOptions
      })

      return configWithPermissions
    },
    experimentalRunAllSpecs: true,
    excludeSpecPattern: '!*.spec.js',
    specPattern: 'cypress/e2e/**/*.{js,jsx,ts,tsx}',
  },
  asBaseUrl: 'http://localhost:1885/api/v3',
  asEnabled: true,
  nsBaseUrl: 'http://localhost:1885/api/v3',
  nsEnabled: true,
  jsBaseUrl: 'http://localhost:1885/api/v3',
  jsEnabled: true,
  isBaseUrl: 'http://localhost:1885/api/v3',
  isEnabled: true,
  gsBaseUrl: 'http://localhost:1885/api/v3',
  gsEnabled: true,
  edtcBaseUrl: 'http://localhost:1885/api/v3',
  edtcEnabled: true,
  qrgBaseUrl: 'http://localhost:1885/api/v3',
  qrgEnabled: true,
  consoleSiteName: 'The Things Stack for LoRaWAN',
  consoleSubTitle: 'Management platform for The Things Stack for LoRaWAN',
  consoleTitle: 'Console',
  consoleAssetsRootPath: '/assets',
  consoleRootPath: '/console',
  accountAppSiteName: 'The Things Stack for LoRaWAN',
  accountAppSubTitle: '',
  accountAppTitle: 'Account',
  accountAppRootPath: '/oauth',
  accountAppAssetsRootPath: '/assets',
})
