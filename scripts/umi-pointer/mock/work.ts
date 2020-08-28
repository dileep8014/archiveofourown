import mockjs, { Random } from 'mockjs';
import { ChapterItem, SubsectionItem } from '@/pages/work/write/sider/menuItem';
import { users } from './user';
import moment from 'moment';

let firstWork = {
  id: 1,
  title: '灵气是我',
  cover: Random.image('200x240'),
  chapters: 0,
  draft: 0,
  recycle: 0,
  subsection: 1,
};

let subsection: SubsectionItem[] = [{
  id: 1,
  seq: 1,
  name: '',
  introduce: '',
  chapters: [],
}];
let draft: ChapterItem[] = [];
let recycle: ChapterItem[] = [];
let idSeq = 1;

export default {
  'GET /api/v1/works/detail': (req: any, res: any) => {
    switch (req.query.type) {
      case 'draft': {
        draft.forEach((item) => {
          if (item.id == req.query.id) {
            res.send(item);
            return;
          }
        });
        break;
      }
      case 'published': {
        subsection.forEach((item) => {
          if (item.seq == req.query.subsection) {
            item.chapters.forEach((subItem) => {
              if (subItem.id == req.query.id) {
                res.send(subItem);
                return;
              }
            });
          }
        });
        break;
      }
      case 'recycle': {
        recycle.forEach((item) => {
          if (item.id == req.query.id) {
            res.send(item);
            return;
          }
        });
        break;
      }
    }
  },
  'POST /api/v1/works/chapter/publish': (req: any, res: any) => {
    let work = req.body;
    switch (work.type) {
      case 'draft': {
        draft.forEach((item, index) => {
          if (item.id == work.id) {
            draft = draft.slice(0, index).concat(draft.slice(index + 1));
            firstWork.draft--;
            firstWork.chapters++;
            work = item;
          }
        });
        break;
      }
      case 'recycle': {
        recycle.forEach((item, index) => {
          if (item.id == work.id) {
            recycle = recycle.slice(0, index).concat(recycle.slice(index + 1));
            work = item;
            firstWork.recycle--;
            firstWork.chapters++;
          }
        });
        break;
      }
    }
    work.type = 'published';
    subsection.forEach((item, i) => {
      if (item.seq == work.subsection) {
        subsection[i].chapters.push(work);
      }
    });
    res.end('ok');
  },
  'DELETE /api/v1/works/chapter': (req: any, res: any) => {
    let work = req.body;
    switch (work.type) {
      case 'draft': {
        draft.forEach((item, index) => {
          if (item.id == work.id) {
            draft = draft.slice(0, index).concat(draft.slice(index + 1));
            firstWork.draft--;
          }
        });
        break;
      }
      case 'published': {
        subsection.forEach((item, i) => {
          if (item.seq == work.subsection) {
            item.chapters.forEach((subItem, j) => {
              if (subItem.id == item.id) {
                subsection[i].chapters = subsection[i].chapters.slice(0, j).concat(subsection[i].chapters.slice(j + 1));
                firstWork.chapters--;
              }
            });
          }
        });
        break;
      }
      case 'recycle': {
        recycle.forEach((item, index) => {
          if (item.id == work.id) {
            recycle = recycle.slice(0, index).concat(recycle.slice(index + 1));
            firstWork.recycle--;
          }
        });
        break;
      }
    }
    res.end('ok');
  },
  'POST /api/v1/works/chapter': (req: any, res: any) => {
    let work = req.body;
    work.updatedAt = moment.now();
    work.words = work.content.length;
    if (work.id == -1) {
      work.id = idSeq;
      draft.push(work);
      firstWork.draft++;
      idSeq++;
    } else {
      switch (work.type) {
        case 'draft': {
          draft.forEach((item, index) => {
            if (item.id == work.id) {
              draft[index] = work;
            }
          });
          break;
        }
        case 'published': {
          subsection.forEach((item, i) => {
            if (item.seq == work.subsection) {
              item.chapters.forEach((subItem, j) => {
                if (subItem.id == item.id) {
                  subsection[i].chapters[j] = work;
                }
              });
            }
          });
          break;
        }
        case 'recycle': {
          recycle.forEach((item, index) => {
            if (item.id == work.id) {
              recycle[index] = work;
            }
          });
          break;
        }
      }
    }
    res.end('ok');
  },
  'DELETE /api/v1/works/subsection': (req: any, res: any) => {
    subsection.forEach((item, index) => {
      if (item.id == req.body.id) {
        subsection = subsection.slice(0, index).concat(subsection.slice(index + 1));
      }
    });
    res.end('ok');
  },
  'POST /api/v1/works/subsection': (req: any, res: any) => {
    let sub = req.body;
    if (sub.id == -1) {
      let index = subsection.length - 1;
      sub.id = subsection[index].id + 1;
      sub.seq = subsection[index].seq + 1;
      firstWork.subsection++;
      subsection.push(sub);
    } else {
      subsection.forEach((item, index) => {
        if (item.id == sub.id) {
          subsection[index].name = sub.name;
          subsection[index].introduce = sub.introduce;
        }
      });
    }
    res.end('ok');
  },
  'GET /api/v1/works/info': (req: any, res: any) => {
    if (req.query.id == -1) {
      res.status(500).end();
    } else {
      let data = { data: firstWork, draft: draft, subsection: subsection, recycle: recycle };
      res.send(data);
    }
  },
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
      'tags|10-100': ['@last'],
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
};
