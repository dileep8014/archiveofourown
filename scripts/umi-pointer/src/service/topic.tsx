import request from '@/service/request';

export const topicService={
  ListTopicByCategory: function(params: { categoryId: number }) {
    return request('/api/v1/topics', { params: params });
  },
}
