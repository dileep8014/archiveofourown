import mockjs from 'mockjs';

export default {
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
}
