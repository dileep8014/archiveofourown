import request from '@/service/request';

export const newsService={
  News: async function(params: { current: number, pageSize: number }) {
    return request('/api/v1/news', {
      params: params,
    });
  },
}
