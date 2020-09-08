import { useCallback, useEffect, useState } from 'react';
import { useRequest } from '@umijs/hooks';
import moment from 'moment';
import { userService } from '@/service/user';

export type UserModelState = {
  id: number,
  username: string,
  email: string,
  avatar: string,
  gender: number,
  introduce: string,
  worksNums: number,
  root: boolean,
  workDay: number,
  words: number,
  fansNums: number,
  createdAt: moment.Moment,
} | null

export default function useCurrentUserModel(): {
  signin: (params: { username: string, password: string, rememberMe: boolean }) => void;
  signout: () => void;
  setUser: (value: { username?: string, avatar?: string, gender?: number, introduce?: string }) => void;
  user: UserModelState,
  loading: boolean,
} {

  const { data, refresh, loading } = useRequest(userService.QueryCurrentUser);
  const [user, update] = useState<UserModelState>(null);
  const { run: signIn } = useRequest(userService.SignIn, {
    manual: true, onSuccess: (res, params) => {
      if (res.code == 0) {
        refresh();
        localStorage.setItem('currentUser', params[0].username);
        localStorage.setItem('currentPass', params[0].password);
      }
    },
  });
  const { run: signOut } = useRequest(userService.SignOut, {
    manual: true, onSuccess: res => {
      update(null);
      localStorage.removeItem('Authorization');
    },
  });
  const { run: updateUserInfo } = useRequest(userService.Update, {
    manual: true, onSuccess: (res, params) => {
      if (res.code == 0) {
        // @ts-ignore
        update({ ...user, ...params[0] });
      }
    },
  });

  useEffect(() => {
    if (data && data.code == 0) {
      update(data.data);
    }
  }, [data]);

  const setUser = useCallback((userInfo: { username?: string, avatar?: string, gender?: number, introduce?: string }) => {
    updateUserInfo(userInfo);
  }, []);

  const signin = useCallback((params: { username: string, password: string, rememberMe: boolean }) => {
    // signin implementation
    signIn(params);
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
