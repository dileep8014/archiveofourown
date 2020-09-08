import RightHeader from './component/rightHeader/rightHeader';
import React from 'react';
import { BackTop } from 'antd';
import ErrorPage from '@/component/errorpage/errorpage';


export const layout = {
  logout: () => {
  }, // do something
  rightRender: (initInfo: any) => {
    return <RightHeader/>;
  },
  footerRender: () => {
    return <BackTop/>;
  },
  ErrorComponent: (error: Error) => <ErrorPage title={error.name} subTitle={error.message}/>,
};

// export const dva = {
//   config: {
//     onAction: createLogger(),
//     onError(e: Error) {
//       message.error(e.message, 3);
//     },
//   },
// };
