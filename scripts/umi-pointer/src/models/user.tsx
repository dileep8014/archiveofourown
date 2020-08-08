import { useCallback, useEffect, useState } from 'react';
import service from '@/component/service';
import { message } from 'antd';
import { useRequest } from '@umijs/hooks';

export type UserModelState = {
  id: number,
  name: string,
  email: string,
  avatar: string,
  root: boolean,
} | null

export default function useCurrentUserModel(): { user: UserModelState, signin: (account: string, password: string) => void, signout: () => void } {
  const [user, setUser] = useState(null);

  const { data, error, refresh } = useRequest(service.QueryCurrentUser);
  const { run: signIn } = useRequest(service.SignIn, {
    manual: true, onSuccess: res => {
      if (res == 'ok') {
        refresh().then(r => {
          setUser(r);
        });
      }
    },
  });
  const { run: signOut } = useRequest(service.SignOut, {
    manual: true, onSuccess: res => {
      if (res == 'ok') {
        setUser(null);
      }
    },
  });

  useEffect(() => {
    if (data) {
      setUser(data);
    }
    if (error) {
      message.error(error);
    }
  }, [data, error]);

  const signin = useCallback((account: string, password: string) => {
    // signin implementation
    signIn(account, password);
  }, []);

  const signout = useCallback(() => {
    signOut();
  }, []);

  return {
    user,
    signin,
    signout,
  };
}
