import React from 'react';
import { connect, Loading, ConnectProps } from 'umi';
import { UserModelState } from '@/models/user';

interface UserProps extends ConnectProps {
  user: UserModelState;
  loading: boolean;
}

const RightHeader: React.FC<UserProps> = ({ user, loading }) => {
  const { name } = user;

  return (
    <span>
      {name}
    </span>
  );
};

export default connect(({ user, loading }: { user: UserModelState, loading: Loading }) => ({
  user,
  loading: loading.models.user,
}))(RightHeader);
