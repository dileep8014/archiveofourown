import React, { ReactNode, useEffect, useState } from 'react';
import { Modal, Skeleton } from 'antd';
import SignForm from '@/component/sign/sign';
import { useModel, history } from 'umi';

export default (props: { children: ReactNode }) => {

  const [visit, setVisit] = useState(false);

  const { user, loading } = useModel('user', model => ({
    user: model.user,
    loading: model.loading,
  }));

  useEffect(() => {
    if (user) {
      setVisit(false);
    } else {
      setVisit(true);
    }
  }, [user]);

  const cancel = () => {
    setVisit(false);
    history.goBack();
  };


  return <Skeleton loading={loading}>
    {user && !visit && props.children}
    {!user &&
    <Modal
      centered={true}
      width={400}
      footer={null}
      onCancel={cancel}
      visible={visit}
    >
      <SignForm tab={'signIn'} finish={() => {
        setVisit(false);
        history.refresh();
      }}/>
    </Modal>}
  </Skeleton>;
}
