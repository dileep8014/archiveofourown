import mockjs from 'mockjs';

export default {
  'GET /api/v1/topics': (req: any, res: any) => {
    setTimeout(() => res.send(mockjs.mock({ 'list|10-20': [{ 'id|+1': 1, name: '@ctitle' }] }).list), 1000);
  },
}
