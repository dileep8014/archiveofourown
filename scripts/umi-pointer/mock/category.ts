export default {
  'GET /api/v1/category': (req: any, res: any) => {
    if (req.query.type == 1) {
      res.send([{ id: 1, name: '动漫' }, { id: 2, name: '文学' }, { id: 3, name: '影视' }, { id: 4, name: '戏剧' },
        { id: 5, name: '音乐' }, { id: 6, name: '游戏' }, { id: 7, name: '其他' }]);
      return;
    }
    res.send([{ id: 1, name: '玄幻' }, { id: 2, name: '仙侠' }, { id: 3, name: '言情' },
      { id: 4, name: '武侠' }, { id: 5, name: '都市' },
      { id: 6, name: '军事' }, { id: 7, name: '悬疑' }, { id: 8, name: '文学' }, { id: 9, name: '灵异' }]);
  },
}
