
import { defineMessages } from 'react-intl'

export default defineMessages({
  imageProvided: 'Image has been provided',
  profilePicture: 'Profile picture',
  successMessage: 'Profile updated',
  deleteAccount: 'Delete account',
  useGravatar: 'Use Gravatar',
  gravatarInfo:
    "If available, we're using the <link>Gravatar</link> image associated with your email address. You can upload a different profile picture by selecting the option above.",
  gravatarInfoGravatarOnly:
    "If available, we're using the <link>Gravatar</link> image associated with your email address. Please follow the instructions on the Gravatar website to change your profile picture.",
  primaryEmailAddressDescription: 'Primary email address associated with your account',
  deleteAccountError: 'There was an error and your account could not be deleted',
  imageRequired:
    'Please select a file to use as your profile picture or choose "Gravatar" as source',
  imageUpload: 'Image upload',
  chooseImage: 'Choose image…',
  changeImage: 'Change image…',
  gravatarImage: 'Gravatar image',
  profilePictureDisabled:
    'Setting a profile picture is currently disabled. Hence, only an administrator can change the profile picture.',
  deleteWarning:
    "This will <strong>PERMANENTLY DELETE YOUR ACCOUNT</strong> and <strong>LOCK THE USER ID AND EMAIL FOR RE-REGISTRATION</strong>. Associated entities (e.g. gateways, applications and end devices) owned by you that do not have any other collaborators will become <strong>UNACCESSIBLE</strong> and it will <strong>NOT BE POSSIBLE TO REGISTER ENTITIES WITH THE SAME ID OR EUI's AGAIN</strong>. Make sure you assign new collaborators to such entities if you plan to continue using them.",
  purgeWarning:
    "This will <strong>PERMANENTLY DELETE YOUR ACCOUNT</strong> and make associated entities (e.g. gateways, applications and end devices) owned by you that do not have other collaborators <strong>UNACCESSIBLE</strong> and it will also <strong>NOT BE POSSIBLE TO REGISTER ENTITIES WITH THE SAME ID OR EUI's</strong>. Make sure you assign new collaborators to such entities if you plan to continue using them.",
  deleteTitle: 'Are you sure you want to delete your account?',
  deleteConfirmMessage: 'Please enter your user ID to confirm the account deletion.',
})
