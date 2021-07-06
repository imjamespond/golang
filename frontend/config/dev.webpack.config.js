const path = require('path')
const getConfig = require('./webpack.config')

module.exports = (env, options) => {
  const config = getConfig(env, options)
  return {
    ...config,
    devtool: 'source-map',
    devServer: {
      publicPath: config.output.publicPath.slice(0, -1),//prefix url 
      contentBase: path.join(__dirname, '../public'),
      contentBasePublicPath: config.output.publicPath.slice(0, -1),//add prefix to assets in public folder 
      port: 9000
    },
    optimization: {
      ...config.optimization,
      minimize: false
    }
  }
}