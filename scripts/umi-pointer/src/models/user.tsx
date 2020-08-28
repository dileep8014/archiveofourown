import { useCallback, useEffect, useState } from 'react';
import { message } from 'antd';
import { useRequest } from '@umijs/hooks';
import moment from 'moment';
import { userService } from '@/service/user';

export type UserModelState = {
  id: number,
  name: string,
  email: string,
  avatar: string,
  gender: string,
  birthday: moment.Moment | null
  phone: number | null,
  introduce: string,
  root: boolean,
  workDay: number,
  words: number,
  fans: number,
} | null

export default function useCurrentUserModel(): {
  signin: (account: string, password: string) => void;
  signout: () => void;
  setUser: (value: UserModelState) => void;
  user: UserModelState,
  loading: boolean,
} {

  const { data, error, refresh, loading } = useRequest(userService.QueryCurrentUser);
  const [user, update] = useState<UserModelState>(data);
  const { run: signIn } = useRequest(userService.SignIn, {
    manual: true, onSuccess: res => {
      if (res == 'ok') {
        refresh().then(r => {
          update(r);
        });
      }
    },
  });
  const { run: signOut } = useRequest(userService.SignOut, {
    manual: true, onSuccess: res => {
      if (res == 'ok') {
        update(null);
      }
    },
  });
  const { run: updateUserInfo } = useRequest(userService.Update, {
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
