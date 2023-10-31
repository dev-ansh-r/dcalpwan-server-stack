

import { defineMessages } from 'react-intl'

const m = defineMessages({
  proceed: 'Proceed to end device list',
  retry: 'Retry from scratch',
  abort: 'Abort',
  converting: 'Converting templates…',
  creating: 'Creating end devices…',
  operationInProgress: 'Operation in progress',
  operationHalted: 'Operation halted',
  operationFinished: 'Operation finished',
  operationAborted: 'Operation aborted',
  errorTitle: 'There was an error and the operation could not be completed',
  conversionErrorTitle: 'Could not import devices',
  conversionErrorMessage:
    'An error occurred while processing the provided end device template. This could be due to invalid format, syntax or file encoding. Please check the provided template file and try again. See also our documentation on <DocLink>Importing End Devices</DocLink> for more information.',
  incompleteWarningTitle: 'Not all devices imported successfully',
  incompleteWarningMessage:
    '{count, number} {count, plural, one {end device} other {end devices}} could not be imported successfully, because {count, plural, one {its} other {their}} registration attempt resulted in an error',
  incompleteStatus:
    'The registration of the following {count, plural, one {end device} other {end devices}} failed:',
  noneWarningTitle: 'No end device was created',
  noneWarningMessage:
    'None of your specified end devices was imported, because each registration attempt resulted in an error',
  processLog: 'Process log',
  progress:
    'Successfully converted {errorCount, number} of {deviceCount, number} {deviceCount, plural, one {end device} other {end devices}}',
  successInfoTitle: 'All end devices imported successfully',
  successInfoMessage:
    'All of the specified end devices have been converted and imported successfully',
  documentationHint:
    'Please also see our documentation on <DocLink>Importing End Devices</DocLink> for more information and possible resolutions.',
  abortWarningTitle: 'Device import aborted',
  abortWarningMessage:
    'The end device import was aborted and the remaining {count, number} {count, plural, one {end device} other {end devices}} have not been imported',
  largeFileWarningMessage:
    'Providing files larger than {warningThreshold} can cause issues during the import process. We recommend you to split such files up into multiple smaller files and importing them one by one.',
})

export default m
