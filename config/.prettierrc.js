

const yaml = require('js-yaml')
const fs = require('fs')
const path = require('path')

module.exports = yaml.load(fs.readFileSync(path.join(__dirname, 'prettierrc.yaml'), 'utf8'))
