import mockjs from 'mockjs';

export default {
  // 支持值为 Object 和 Array
  'GET /api/v1/users/currentUser': { id: 1, name: 'shyptr', email: 'xxx@qq.com' },

  // GET 可忽略
  '/api/v1/users/1': mockjs.mock({
    'list|1-100': [
      { 'id|+1': 2, 'name|1-10': 'string'},
    ],
  }),

  // 支持自定义函数，API 参考 express@4
  'POST /api/v1/users/create': (req: any, res: { setHeader: (arg0: string, arg1: string) => void; end: (arg0: string) => void; }) => {
    // 添加跨域请求头
    res.setHeader('Access-Control-Allow-Origin', '*');
    res.end('ok');
  },
};
