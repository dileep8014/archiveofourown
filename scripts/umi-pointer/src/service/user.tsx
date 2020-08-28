import { StyleState } from '@/pages/userCenter/userstyle';
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
  GetStyles: function() {
    return request('/api/v1/users/styles');
  },
  UpdateStyles: function(styles: StyleState) {
    return request('/api/v1/users/styles', {
      method: 'post',
      data: styles,
    });
  },
  Update: async function(user: UserModelState) {
    return request('/api/v1/users/update', {
      method: 'post',
      data: user,
    });
  },
  QueryCurrentUser: function() {
    return request('/api/v1/users/currentUser');
  },
  SignIn: function(account: string, password: string) {
    return request('/api/v1/users/login', {
      method: 'post',
      data: { account: account, password: password },
    });
  },
  SignOut: function() {
    return request('/api/v1/users/logOut', {
      method: 'post',
    });
  },
  SignUp: async function(account: string, email: string, password: string) {
    return request('/api/v1/users/create', {
      method: 'post',
      data: { account: account, email: email, password: password },
    });
  },
};
