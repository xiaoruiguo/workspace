/*
 * start.js
 * Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
 *
 * Distributed under terms of the GPL license.
 */

require('babel-register') ({
    presets: [ 'env' ]
})

module.exports = require('./app.js')
