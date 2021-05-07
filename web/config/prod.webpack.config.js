const { CleanWebpackPlugin } = require('clean-webpack-plugin')
const getConfig = require('./webpack.config')

module.exports = (env, options) => {
  const config = getConfig(env, options)
  return {
    ...config,
    target: ['web', 'es5'],
    plugins: [
      ...config.plugins,
      new CleanWebpackPlugin(),
    ],
    optimization: {
      ...config.optimization,

    }
  }
}