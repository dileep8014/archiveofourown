import { UserSetting } from '@/pages/userCenter/userstyle';
import { UserModelState } from '@/models/user';
import request from './request';

export const userService = {
  UserCollege: function(params: { current: number, pageSize: number }) {
    return request('/api/v1/users/college', { data: params });
  },
  UserTopics: function(params: { current: number, pageSize: number | null }) {
    if (params.pageSize == null) {
      params.pageSize = 10;
    }
    return request('/api/v1/users/topics', { data: params });
  },
  CurrentUserSetting: function() {
    return request('/api/v1/currentUser/setting');
  },
  UpdateCurrentSetting: function(st: UserSetting) {
    console.log(st);
    return request('/api/v1/currentUser/setting', {
      method: 'post',
      data: st,
    });
  },
  Update: async function(user: { username?: string, avatar?: string, gender?: number, introduce?: string }) {
    return request('/api/v1/currentUser', {
      method: 'post',
      data: user,
    });
  },
  UpdatePassword: async function(params: { oldPassword: string, password: string }) {
    return request('/api/v1/currentUser/password', {
      method: 'post',
      data: params,
    });
  },
  UpdateEmail: async function(params: { email: string, password: string }) {
    return request('/api/v1/currentUser/email', {
      method: 'post',
      data: params,
    });
  },
  QueryCurrentUser: function() {
    return request('/api/v1/currentUser');
  },
  SignIn: function(params: { username: string, password: string, rememberMe: boolean }) {
    return request('/api/v1/login', {
      method: 'post',
      data: params,
    });
  },
  SignOut: function() {
    return request('/api/v1/logout', {
      method: 'post',
    });
  },
  SignUp: function(email: string) {
    return request('/api/v1/register', {
      method: 'post',
      data: { email: email },
    });
  },
  Identify: function(path: string) {
    return request(`/api/v1/register/identify`, { params: { path: path } });
  },
  Create: function(params: { email: string, password: string, username: string }) {
    return request('/api/v1/register/create', {
      method: 'post',
      data: params,
    });
  },
};
