const webpackConfig = require('./webpack.config.js');

module.exports = function override(config) {
  config.resolve.alias = webpackConfig.resolve.alias;
  return config;
}
