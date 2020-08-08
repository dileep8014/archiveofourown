import { UserModelState } from '@/models/user';
import mockjs, { Random } from 'mockjs';

let currentUser: UserModelState = null;

let users = [{
  id: 1,
  name: 'shyptr',
  email: 'xxx@qq.com',
  avatar: 'https://assets.leetcode-cn.com/aliyun-lc-upload/users/yun-yan-7/avatar_1575262045.png',
  root: true,
}];

export default {
  // works
  'GET /api/v1/works/subscribe': (req: any, res: any) => {
    let data = {
      'id|+1': 1,
      title: '@ctitle',
      subTitle: '@csentence',
      topic: '@ctitle',
      'tags|3-12': '@last',
      words: '@natural(100,1000000)',
      chapters: '@natural(0,1000)',
      comments: '@natural(0,100000)',
      collection: '@natural(0,10000)',
      subscribe: '@natural(1,1000)',
      hits: '@natural(10,1000000)',
      user: users[Random.natural(1, users.length) - 1],
    };
    switch (req.query.pageSize) {
      case '10':
        res.send(mockjs.mock({
          'list|10': [data],
          total: 100,
        }));
        break;
      case '20':
        res.send(mockjs.mock({
          'list|20': [data],
          total: 100,
        }));
        break;
      case '50':
        res.send(mockjs.mock({
          'list|50': [data],
          total: 100,
        }));
        break;
      case '100':
        res.send(mockjs.mock({
          'list|100': [data],
          total: 100,
        }));
        break;
    }
  },

  // tag
  'GET /api/v1/topic/hots': mockjs.mock({ 'list|50-100': ['@last'] }),

  //Auth
  'GET /auth': (req: any, res: any) => {
    res.setHeader('x-auth-token', 'token-token-token');
    res.send('');
  },

  // 支持值为 Object 和 Array
  'GET /api/v1/users/currentUser': (req: any, res: any) => {
    res.setHeader('x-auth-token', 'token-token-token');
    res.send(currentUser);
  },

  // GET 可忽略
  '/api/v1/users/:id': (req: any, res: any) => {
    users.forEach(item => {
      if (item.id == req.match.id) {
        res.send(item);
      }
    });
  },

  'POST /api/v1/users/login': (req: any, res: any) => {
    users.forEach(item => {
      if (item.name == req.body.account || item.email == req.body.account) {
        currentUser = item;
      }
    });
    res.end('ok');
  },

  'POST /api/v1/users/logout': (req: any, res: any) => {
    currentUser = null;
    res.end('ok');
  },

  'POST /api/v1/users/create': (req: any, res: any) => {
    users.concat({
      id: users.length,
      name: req.account,
      email: req.email,
      avatar: 'https://assets.leetcode-cn.com/aliyun-lc-upload/users/yun-yan-7/avatar_1575262045.png',
      root: false,
    });
    res.end('ok');
  },
};
