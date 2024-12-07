const { createProxyMiddleware } = require('http-proxy-middleware');

module.exports = function (app) {
    app.use(createProxyMiddleware('/api/v1', {
        // target: 'https://www.lvyouwang.xyz:9000', // 使用 HTTPS 协议并正确设置端口
        target: 'http://localhost:9000', // 使用 HTTPS 协议并正确设置端口
        changeOrigin: true,
        ws: true,
        pathRewrite: {
            '^/api/v1': ''  // 去掉 /api/v1 前缀
        },
    }));
};
