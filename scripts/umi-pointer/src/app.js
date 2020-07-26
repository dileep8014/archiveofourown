import { createLogger } from 'redux-logger';
import { message } from 'antd';
import RightHeader from './component/rightHeader';
import React from 'react';

export const layout = {
  logout: () => {
  }, // do something
  rightRender: (initInfo) => {
    return <RightHeader/>;
  },
};


export const dva = {
  config: {
    onAction: createLogger(),
    onError(e) {
      message.error(e.message, 3);
    },
  },
};
