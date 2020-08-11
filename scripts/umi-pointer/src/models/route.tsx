import { Subscription } from 'umi';
import ProgressOpt from '@/component/progress/progress';

export interface SubRouteModelType {
  namespace: 'subRoute';
  subscriptions: { setup: Subscription };
}


const SubRouteModel: SubRouteModelType = {
  namespace: 'subRoute',
  subscriptions: {
    setup({ dispatch, history }) {
      return history.listen((location, action) => {
        ProgressOpt(() => {});
      });
    },
  },
};

export default SubRouteModel;
