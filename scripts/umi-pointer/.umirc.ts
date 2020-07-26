import { defineConfig } from 'umi';

export default defineConfig({

  layout: {
    'navTheme': 'light',
    'layout': 'top',
    'contentWidth': 'Fixed',
    'fixedHeader': true,
    'fixSiderbar': true,
    'menu': {
      'locale': true,
    },
    'logo': '',
    'title': 'Ant Design Pro',
    'pwa': false,
    'iconfontUrl': '',
    'splitMenus': false,
  },
  antd: {
    dark: false,
    compact: true, // 开启紧凑主题
  },
  dva: {
    immer: true,
    hmr: true,
  },
  nodeModulesTransform: {
    type: 'none',
  },
  theme: {
    'primary-color': '#FA541C',
    'link-color': '#FA541C',
  },
  routes: [
    { path: '/', component: '@/pages/index', menu: { name: '首页' } },
  ],
});
