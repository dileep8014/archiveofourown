import mockjs from 'mockjs';

export default {
  'GET /api/v1/tags': (req: any, res: any) => {
    res.send(mockjs.mock({ 'list|0-10': [{ name: '@ctitle' }] }).list);
  },
}
