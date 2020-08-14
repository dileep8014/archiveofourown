import { extend } from 'umi-request';
import { notification } from 'antd';
import { UserModelState } from '@/models/user';
import { StyleState } from '@/pages/userCenter/userstyle';


const codeMessage = {
  200: '服务器成功返回请求的数据。',
  201: '新建或修改数据成功。',
  202: '一个请求已经进入后台排队（异步任务）。',
  204: '删除数据成功。',
  400: '发出的请求有错误，服务器没有进行新建或修改数据的操作。',
  401: '用户没有权限（令牌、用户名、密码错误）。',
  403: '用户得到授权，但是访问是被禁止的。',
  404: '发出的请求针对的是不存在的记录，服务器没有进行操作。',
  406: '请求的格式不可得。',
  410: '请求的资源被永久删除，且不会再得到的。',
  422: '当创建一个对象时，发生一个验证错误。',
  500: '服务器发生错误，请检查服务器。',
  502: '网关错误。',
  503: '服务不可用，服务器暂时过载或维护。',
  504: '网关超时。',
};

/**
 * 异常处理程序
 */
const errorHandler = (error: { response: { status: number, url: string, statusText: string } }) => {
  const { response } = error;

  if (response && response.status) {
    // @ts-ignore
    const errorText = codeMessage[response.status] || response.statusText;
    const { status, url } = response;
    notification.error({
      message: `请求错误 ${status}: ${url}`,
      description: errorText,
    });
  }

  return response;
};
const request = extend({
  errorHandler,
  // 默认错误处理
  credentials: 'include', // 默认请求是否带上cookie

});

// request拦截器, 改变url 或 options.
request.interceptors.request.use((url, options) => {

  let c_token = localStorage.getItem('x-auth-token');
  if (c_token) {
    const headers = {
      'Content-Type': 'application/json',
      'Accept': 'application/json',
      'x-auth-token': c_token,
    };
    return (
      {
        url: url,
        options: { ...options, headers: headers },
      }
    );
  } else {
    const headers = {
      'Content-Type': 'application/json',
      'Accept': 'application/json',
      'x-auth-token': c_token,
    };
    return (
      {
        url: url,
        options: { ...options },
      }
    );
  }

});

// response拦截器, 处理response
request.interceptors.response.use((response, options) => {
  let token = response.headers.get('x-auth-token');
  if (token) {
    localStorage.setItem('x-auth-token', token);
  }
  return response;
});

const service = {
  // Auth
  Auth: function() {
    return request('/auth');
  },
  // User
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
  // Topic
  HotsTopic: function() {
    return request('/api/v1/topic/hots');
  },
  // Work
  Calendar: function(params: { year: number }) {
    return request('/api/v1/works/calendar', { data: params });
  },
  MineWorks: function(params: { current: number, pageSize: number }) {
    return request('/api/v1/works/mine', {
      params: params,
    });
  },
  SubWorks: async function(params: { current: number, pageSize: number }) {
    return request('/api/v1/works/subscribe', {
      params: params,
    });
  },
  // News
  News: async function(params: { current: number, pageSize: number }) {
    return request('/api/v1/news', {
      params: params,
    });
  },
};

export default service;


