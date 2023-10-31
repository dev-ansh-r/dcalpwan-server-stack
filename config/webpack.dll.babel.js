

/* eslint-env node */

import path from 'path'

import webpack from 'webpack'

const { CONTEXT = '.', CACHE_DIR = '.cache', PUBLIC_DIR = 'public' } = process.env
const mode = process.env.NODE_ENV === 'development' ? 'development' : 'production'
const WEBPACK_GENERATE_PRODUCTION_SOURCEMAPS =
  process.env.WEBPACK_GENERATE_PRODUCTION_SOURCEMAPS === 'true'

const context = path.resolve(CONTEXT)
const library = '[name]_[fullhash]'

const pkg = require(path.resolve(context, 'package.json'))
const excludeLibs = ['react-hot-loader', 'ttn-lw']
const libs = Object.keys(pkg.dependencies || {}).filter(lib => !excludeLibs.includes(lib))
const devtool =
  (mode === 'production' && WEBPACK_GENERATE_PRODUCTION_SOURCEMAPS) || mode === 'development'
    ? 'source-map'
    : false

export default {
  context,
  mode,
  target: 'web',
  stats: 'minimal',
  devtool,
  recordsPath: path.resolve(context, CACHE_DIR, '_libs_records'),
  entry: { libs },
  output: {
    filename: mode === 'production' ? '[name].[fullhash].bundle.js' : '[name].bundle.js',
    hashDigest: 'hex',
    hashDigestLength: 20,
    path: path.resolve(context, PUBLIC_DIR),
    library,
  },
  plugins: [
    new webpack.DllPlugin({
      name: library,
      path: path.resolve(context, CACHE_DIR, 'dll.json'),
    }),
  ],
  performance: {
    hints: false,
  },
  module: {
    rules: [
      {
        test: /\.(woff|woff2|ttf|eot|jpg|jpeg|png|svg)$/i,
        use: [
          {
            loader: 'file-loader',
            options: {
              name: '[name].[contenthash:20].[ext]',
            },
          },
        ],
      },
    ],
  },
}
