import { defineConfig } from 'umi';

export default defineConfig({
  locale: {
    default: 'zh-CN',
    useLocalStorage: false,
    antd: true,
    title: true,
    baseNavigator: true,
    baseSeparator: '-',
  },
  mock: false,
  styles: ['./src/app.css'],
  layout: {
    'navTheme': 'light',
    'layout': 'top',
    'fixedHeader': false,
    'fixSiderbar': false,
    'menu': {
      'locale': true,
    },
    'logo': 'https://img.alicdn.com/tfs/TB1zomHwxv1gK0jSZFFXXb0sXXa-200-200.png',
    'title': 'Pointer',
    'pwa': false,
    'iconfontUrl': '',
    'splitMenus': false,
    'contentStyle': { backgroundColor: 'white', padding: 30 },
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
    'primary-color': '#d02525',
    'link-color': '#534545',
  },
  routes: [
    { exact: true, path: '/', component: '@/pages/index' },
    {
      exact: true,
      path: '/userCenter',
      title: '个人中心',
      component: '@/pages/userCenter/index',
      wrappers: ['@/component/wrapper/auth'],
    },
    {
      exact: true,
      path: '/creativeCenter',
      title: '创作中心',
      component: '@/pages/creativeCenter/index',
      wrappers: ['@/component/wrapper/auth'],
    },
    {
      exact: false,
      path: '/works/write',
      title: '写作',
      component: '@/pages/work/write/index',
      layout: { hideNav: true, hideMenu: true },
      wrappers: ['@/component/wrapper/auth'],
    },
    {
      exact: true,
      path: '/register/:path',
      title: '注册信息',
      component: '@/pages/register/index',
      wrappers: ['@/component/wrapper/notAuth'],
    },
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
    {
      menu: { name: '浏览' }, routes: [
        { path: '/works', menu: { name: '作品' } },
        { path: '/topics', menu: { name: '专题' } },
        { path: '/tags', menu: { name: '标签' } },
        { path: '/college', menu: { name: '书单' } },
      ],
    },
    {
      menu: { name: '搜索' }, routes: [
        { path: '/search/works', menu: { name: '作品' } },
        { path: '/search/topics', menu: { name: '专题' } },
        { path: '/search/tags', menu: { name: '标签' } },
        { path: '/search/users', menu: { name: '创作者' } },
      ],
    },
  ],
});
