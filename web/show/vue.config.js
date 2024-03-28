const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  lintOnSave:false,
  assetsDir:'static',
  productionSourceMap:false,
  devServer: {
    port:8033,
    proxy: {
      '/api': {
        target: 'http://8.218.92.69:8080/api',
        // ws: true,
        changeOrigin: true, //是否跨域
        pathRewrite: {
          '^/api': ''
        }
      },
      '/tool':{
        target: 'http://8.218.92.69:8099/tool',
        // ws: true,
        changeOrigin: true, //是否跨域
        pathRewrite: {
          '^/tool': ''
        }
      }
    },
    
    historyApiFallback: true //增加这个选项
  },
})
