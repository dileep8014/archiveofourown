import React, { ReactNode } from 'react';
import { Modal, notification, Skeleton } from 'antd';
import { useModel, history } from 'umi';

export default (props: { children: ReactNode }) => {

  const { user, loading } = useModel('user', model => ({
    user: model.user,
    loading: model.loading,
  }));

  if (user) {
    notification['warning']({ message: '请先退出当前登录账户' });
    history.goBack();
  }

  return <Skeleton loading={loading}>
    {props.children}
  </Skeleton>;
}
