
/* eslint-disable import/no-commonjs */

module.exports = {
  stories: ['../../pkg/webui/**/story.js', '../../pkg/webui/**/*.stories.js'],
  addons: ['@storybook/addon-actions/register', '@storybook/addon-info'],
  staticDirs: ['../../public'],
}
