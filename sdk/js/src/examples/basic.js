

/* eslint-disable no-console */

import TTS from '..'

const token = 'access-token-or-api-key'
const tts = new TTS({
  authorization: {
    mode: 'key',
    key: token,
  },
  connectionType: 'http',
  baseURL: 'http://localhost:1885/api/v3',
  defaultUserId: 'testuser',
})

const hash = new Date().valueOf()
const appName = `first-app-${hash}`
const appData = {
  ids: {
    application_id: appName,
  },
  name: 'Test App',
  description: `Some description ${hash}`,
}

const createApplication = async () => {
  const firstApp = await tts.Applications.create('testuser', appData)
  console.log(firstApp)
}

const getApplication = async () => {
  const firstApp = await tts.Applications.getById(appName)
  console.log(firstApp)
}

const listApplications = async () => {
  const apps = await tts.Applications.getAll()
  console.log(apps)
}

const updateApplication = async () => {
  const patch = { description: 'New description' }
  const res = await tts.Applications.updateById(appName, patch)
  console.log(res)
}

const deleteApplication = async () => {
  await tts.Applications.deleteById(appName)
  await tts.Applications.deleteById(`second-app-${hash}`)
}

const main = async () => {
  try {
    await createApplication()
    await getApplication()
    await listApplications()
    await updateApplication()
    await deleteApplication()
  } catch (err) {
    console.log(err)
  }
}

main()
