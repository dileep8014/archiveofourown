import request from '@/service/request';

export const categoryService={
  Category: function(params: { type: number }) {
    return request('/api/v1/category', { params: params });
  },
}
