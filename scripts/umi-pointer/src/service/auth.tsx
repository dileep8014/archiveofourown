import request from '@/service/request';

export const authService={
  Auth: function() {
    return request('/auth');
  },
}
