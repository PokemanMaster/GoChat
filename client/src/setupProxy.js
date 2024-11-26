const {createProxyMiddleware} = require('http-proxy-middleware');

module.exports = function (app) {
    // https://www.lvyouwang.xyz/
    app.use(createProxyMiddleware('/api/v1', {
        target: 'http://localhost:9000', // 确保使用 HTTPS 协议
        changeOrigin: true,
        ws: true,
        pathRewrite: {
            '^/api/v1': ''
        },
    }));
};
