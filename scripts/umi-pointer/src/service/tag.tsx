import request from '@/service/request';

export const tagService={
  SimilarTags: function(params: { similar: string }) {
    return request('/api/v1/tags', { params: params });
  },
}
