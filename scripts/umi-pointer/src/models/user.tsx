import { useCallback, useEffect, useState } from 'react';
import service from '@/component/service';
import { message } from 'antd';
import { useRequest } from '@umijs/hooks';

export type UserModelState = {
  id: number,
  name: string,
  email: string,
  avatar: string,
  gender: string,
  introduce: string,
  root: boolean,
} | null

export default function useCurrentUserModel(): {
  signin: (account: string, password: string) => void;
  signout: () => void;
  setUser: (value: UserModelState) => void;
  user: UserModelState,
  loading: boolean,
} {

  const { data, error, refresh, loading } = useRequest(service.QueryCurrentUser);
  const [user, update] = useState<UserModelState>(data);
  const { run: signIn } = useRequest(service.SignIn, {
    manual: true, onSuccess: res => {
      if (res == 'ok') {
        refresh().then(r => {
          update(r);
        });
      }
    },
  });
  const { run: signOut } = useRequest(service.SignOut, {
    manual: true, onSuccess: res => {
      if (res == 'ok') {
        update(null);
      }
    },
  });
  const { run: updateUserInfo } = useRequest(service.Update, {
    manual: true, onSuccess: (res, params) => {
      if (res == 'ok') {
        // @ts-ignore
        update(params[0]);
      }
    },
  });

  useEffect(() => {
    if (data) {
      update(data);
    }
    if (error) {
      message.error(error);
    }
  }, [data, error]);

  const setUser = useCallback((userInfo: UserModelState) => {
    updateUserInfo(userInfo);
  }, []);

  const signin = useCallback((account: string, password: string) => {
    // signin implementation
    signIn(account, password);
  }, []);

  const signout = useCallback(() => {
    signOut();
  }, []);

  return {
    user,
    setUser,
    signin,
    signout,
    loading,
  };
}
