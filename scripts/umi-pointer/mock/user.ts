import mockjs, { Random } from 'mockjs';
import { UserModelState } from '@/models/user';

let style = {
  email: true,
  birthday: false,
  searchAll: true,
  shareWork: true,
  adult: true,
  grade: false,
  tag: false,
  subMsg: true,
  topicMsg: false,
  commentMsg: false,
  sysMsg: true,
};

export let users: UserModelState[] = [{
  id: 1,
  name: 'shyptr',
  email: '3200828584@qq.com',
  avatar: 'https://assets.leetcode-cn.com/aliyun-lc-upload/users/yun-yan-7/avatar_1575262045.png',
  gender: 'man',
  birthday: null,
  phone: 13340210412,
  introduce: 'Pointer同人网创始人',
  root: true,
  workDay: 10,
  words: 10000,
  fans: 100,
}];

let currentUser: UserModelState = users[0];

export default {
  'GET /api/v1/users/college': mockjs.mock({
    'list|100': [{
      'id|+1': 1,
      title: '@ctitle',
      description: '@cparagraph',
      worksNum: '@natural(1,1000)',
    }],
    total: 100,
  }),
  'GET /api/v1/users/topics': mockjs.mock({
    'list|100': [{
      'id|+1': 1,
      title: '@ctitle',
      category: function() {
        let list = ['动漫', '文学', '影视', '戏剧', '音乐', '游戏', '其他'];
        return list[Random.integer(0, 6)];
      },
      description: '@cparagraph',
      original: '@ctitle',
      url: '@url(http)',
      worksNum: '@natural(1,1000)',
    }],
    total: 100,
  }),
  'GET /api/v1/users/styles': (req: any, res: any) => {
    res.send(style);
  },
  'POST /api/v1/users/styles': (req: any, res: any) => {
    style = req.body;
    res.send('ok');
  },
  'POST /api/v1/users/update': (req: any, res: any) => {
    if (currentUser) {
      currentUser = req.body;
    }
    res.end('ok');
  },
  'GET /api/v1/users/currentUser': (req: any, res: any) => {
    res.setHeader('x-auth-token', 'token-token-token');
    res.send(currentUser);
  },
  '/api/v1/users/:id': (req: any, res: any) => {
    users.forEach(item => {
      if (item?.id == req.match.id) {
        res.send(item);
      }
    });
  },
  'POST /api/v1/users/login': (req: any, res: any) => {
    users.forEach(item => {
      if (item?.name == req.body.account || item?.email == req.body.account) {
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
      gender: 'secret',
      birthday: null,
      introduce: '',
      phone: null,
      root: false,
      words: 0,
      workDay: 0,
      fans: 0,
    });
    res.end('ok');
  },
}
