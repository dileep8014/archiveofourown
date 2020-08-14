import { UserModelState } from '@/models/user';
import mockjs, { Random } from 'mockjs';

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

let users: UserModelState[] = [{
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
  // Topics
  // news
  'GET /api/v1/news': mockjs.mock({
    'list|4': [{
      'id|+1': 1,
      title: '@ctitle',
      createdAt: '@date',
      comments: '@natural(0,100)',
      description: '@cparagraph',
    }],
    'total': 5,
  }),
  // works
  'GET /api/v1/works/calendar': mockjs.mock({
    'list|12': [{
      work: '@boolean',
      'day|31': ['@boolean'],
    }],
  }),
  'GET /api/v1/works/mine': (req: any, res: any) => {
    res.send(mockjs.mock({
      'list|100': [{
        'id|+1': 1,
        title: '@ctitle',
        cover: '@image(200x240)',
        introduce: '@cparagraph',
        newChapter: '第@natural(1,1000)章 @ctitle',
        comments: '@natural(0,10000)',
        subscribe: '@natural(0,1000)',
        college: '@natural(0,1000)',
        hits: '@natural(0,100000)',
      }],
      total: 100,
    }));
  },
  'GET /api/v1/works/subscribe': (req: any, res: any) => {
    let data = {
      'id|+1': 1,
      title: '@ctitle',
      newChapterDesc: '@cparagraph',
      topic: '@ctitle',
      'tags|3-12': '@last',
      newChapter: '第@natural(1,1000)章 @ctitle',
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
  'GET /api/v1/topic/hots': mockjs.mock({ 'list|20-40': ['@ctitle'] }),

  //Auth
  'GET /auth': (req: any, res: any) => {
    res.setHeader('x-auth-token', 'token-token-token');
    res.send('');
  },
  // User
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
};
