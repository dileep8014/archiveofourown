import { defineConfig } from 'umi';

export default defineConfig({
  styles: ['./src/app.css'],
  layout: {
    'navTheme': 'light',
    'layout': 'top',
    'contentWidth': 'Fixed',
    'fixedHeader': true,
    'fixSiderbar': true,
    'menu': {
      'locale': true,
    },
    'logo': 'https://img.alicdn.com/tfs/TB1zomHwxv1gK0jSZFFXXb0sXXa-200-200.png',
    'title': 'Pointer',
    'pwa': false,
    'iconfontUrl': '',
    'splitMenus': false,
    'contentStyle': { backgroundColor: 'white',padding:30 },
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
    'primary-color': '#ffe072',
    'link-color': '#ffe072',
  },
  routes: [
    { exact: true, path: '/', component: '@/pages/index' },
    {
      menu: { name: '同人圈' }, routes: [
        { path: '/category', menu: { name: '所有同人' } },
        { path: '/category/Comic and Animation', menu: { name: '动漫' } },
        { path: '/category/Book and Literature', menu: { name: '文学' } },
        { path: '/category/Film and Television', menu: { name: '影视' } },
        { path: '/category/Drama', menu: { name: '戏剧' } },
        { path: '/category/Music', menu: { name: '音乐' } },
        { path: '/category/Game', menu: { name: '游戏' } },
        { path: '/category/Other', menu: { name: '其他' } },
        { path: '/category/Original', menu: { name: '原创' } },
      ],
    },
  ],
});
