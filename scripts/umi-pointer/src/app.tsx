import RightHeader from './component/rightHeader/rightHeader';
import service from './component/service';
import React from 'react';
import { useRequest } from '@umijs/hooks';
import { BackTop } from 'antd';
import ErrorPage from '@/component/errorpage/errorpage';


export async function getInitialState() {
  let token = localStorage.getItem('x-auth-token');
  if (!token) {
    useRequest(service.Auth);
  }
}


export const layout = {
  logout: () => {
  }, // do something
  rightRender: (initInfo: any) => {
    return <RightHeader/>;
  },
  footerRender: () => {
    return <BackTop/>;
  },
  onError: (error: Error, info: any) => <ErrorPage title={error.name} subTitle={error.message}/>,
};


// export const dva = {
//   config: {
//     onAction: createLogger(),
//     onError(e: Error) {
//       message.error(e.message, 3);
//     },
//   },
// };
