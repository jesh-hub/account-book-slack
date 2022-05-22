const path = require('path');

// react-scripts의 의존성에 포함된 webpack config를 config-overrides.js에서 덮어쓴다.
module.exports = {
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src/'),
    },
  }
}
