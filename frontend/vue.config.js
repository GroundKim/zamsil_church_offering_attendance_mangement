module.exports = {
  transpileDependencies: [
    'vuetify'
  ],
  chainWebpack: config => {
    config
      .plugin('html')
      .tap(args => {
          args[0].title = "잠실교회 소년부";
          return args;
      })
    }

}
