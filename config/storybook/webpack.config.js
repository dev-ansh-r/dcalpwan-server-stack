

/* eslint-env node */
/* eslint-disable import/no-commonjs */
const MiniCssExtractPlugin = require('mini-css-extract-plugin')
require('@babel/register')

const { default: bundleConfig, styleConfig } = require('../webpack.config.babel')

// List of allowed plugins.
const allow = [MiniCssExtractPlugin]

module.exports = async ({ config, mode }) => {
  if (mode === 'PRODUCTION') {
    const webpack = require('webpack')
    allow.push(webpack.DllReferencePlugin)
  }

  // Filter plugins on allowed type.
  const filteredPlugins = bundleConfig.plugins.filter(plugin =>
    allow.reduce((ok, klass) => ok || plugin instanceof klass, false),
  )

  // Compose storybook config, making use of stack webpack config.
  const cfg = {
    ...config,
    resolve: {
      alias: {
        ...config.resolve.alias,
        ...bundleConfig.resolve.alias,
      },
    },
    output: {
      ...config.output,
      publicPath: '',
    },
    module: {
      rules: [...config.module.rules, styleConfig],
    },
    plugins: [...config.plugins, ...filteredPlugins],
  }

  return cfg
}
