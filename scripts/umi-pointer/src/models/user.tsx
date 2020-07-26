import { Effect, ImmerReducer, Reducer, Subscription } from 'umi';
import * as apis from '@/component/api.tsx';

export interface UserModelState {
  id: number,
  name: string,
  email: string,
}

export interface UserModelType {
  state: UserModelState;
  effects: {
    query: Effect;
  };
  reducers: {
    save: ImmerReducer<UserModelState>;
  };
  subscriptions: { setup: Subscription };
}

const UserModel: UserModelType = {

  state: {
    id: 0,
    name: '',
    email: '',
  },

  effects: {
    * query({ payload }, { call, put }) {
      payload = yield call(apis.QueryCurrentUser);
      yield put({ type: 'save', payload });
    },
  },
  reducers: {
    save(state, action) {
      state.id = action.payload.id;
      state.name = action.payload.name;
      state.email = action.payload.email;
    },
  },
  subscriptions: {
    setup({ dispatch, history }) {
      return history.listen(({ pathname }) => {
        if (pathname === '/') {
          dispatch({
            type: 'query',
          });
        }
      });
    },
  },
};

export default UserModel;
