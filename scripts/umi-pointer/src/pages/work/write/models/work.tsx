import { Effect, ImmerReducer, Subscription } from 'umi';
import { ChapterItem, SubsectionItem } from '@/pages/work/write/sider/menuItem';
import { workService } from '@/service/work';
import { useRequest } from '@umijs/hooks';

export interface WorkModelState {
  data: {
    id: number,
    title: string,
    cover: string,
    chapters: number,
    drafts: number,
    recycle: number,
    subsection: number,
    lastChapter: ChapterItem | null,
  } | null;
  subsection: SubsectionItem[]
  draft: ChapterItem[]
  recycle: ChapterItem[]
  error: Error | null
}

export interface WorkModelType {
  namespace: 'work';
  state: WorkModelState;
  effects: {
    query: Effect;
    save: Effect;
    delete: Effect;
    publish: Effect;
  };
  reducers: {
    setState: ImmerReducer<WorkModelState>;
  };
  subscriptions: { setup: Subscription };
}

const WorkModel: WorkModelType = {
  namespace: 'work',

  state: { data: null, subsection: [], draft: [], recycle: [], error: null },

  effects: {
    * query({ payload }, { call, put }) {
      const data = yield call(workService.WorkInfo, payload);
      yield put({ type: 'setState', payload: data });
    },
    * save({ payload }, { call, select, put }) {
      const data = yield call(workService.SaveChapter, payload);
      if (data == 'ok') {
        const id = select((state: any) => state.work.data.id);
        yield put({ type: 'query', payload: { id: id } });
      }
    },
    * delete({ payload }, { call, select, put }) {
      const data = yield call(workService.DeleteChapter, payload);
      if (data == 'ok') {
        const id = select((state: any) => state.work.data.id);
        yield put({ type: 'query', payload: { id: id } });
      }
    },
    * publish({ payload }, { call, select, put }) {
      const data = yield call(workService.PublishChapter, payload);
      if (data == 'ok') {
        const id = select((state: any) => state.work.data.id);
        yield put({ type: 'query', payload: { id: id } });
      }
    },
  },
  reducers: {
    setState(state, action) {
      if (action.payload.constructor == Response) {
        state.error = new Error(action.payload.statusText);
      } else {
        state.data = action.payload.data;
        state.draft = action.payload.draft;
        state.recycle = action.payload.recycle;
        state.subsection = action.payload.subsection;
      }
    },
    // 启用 immer 之后
    // save(state, action) {
    //   state.name = action.payload;
    // },
  },
  subscriptions: {
    setup({ dispatch, history }) {
      return history.listen(({ pathname }) => {
        if (pathname.startsWith('/works/write')) {
          let path = pathname.split('/');
          dispatch({
            type: 'query',
            payload: {
              id: path.length == 4 ? parseInt(path[3]) : -1,
            },
          });
        }
      });
    },
  },
};

export default WorkModel;
