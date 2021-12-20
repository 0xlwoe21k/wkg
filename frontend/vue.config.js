module.exports = {
  devServer: {
    proxy: {
      "/api": {
        target: "http://127.0.0.1:7788",
        changeOrigin: true,
        //   ws: true,
        // rewrite:(path) => path.replace(/^\/api/,'')
        pathRewrite: {
          "^/api": "",
        },
      },
    },
  },
};
