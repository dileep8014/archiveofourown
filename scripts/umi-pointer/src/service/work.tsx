import { ChapterItem, SubsectionItem } from '@/pages/work/write/sider/menuItem';
import request from '@/service/request';

export const workService = {
  PublishChapter: function(params: { id: number, type: 'draft' | 'published' | 'recycle', subsection: number }) {
    return request('/api/v1/works/chapter/publish', { method: 'post', data: params });
  },
  DeleteChapter: function(params: { id: number, type: 'draft' | 'published' | 'recycle', subsection: number }) {
    return request('/api/v1/works/chapter', { method: 'delete', data: params });
  },
  SaveChapter: function(params: {
    id: number;
    title: string;
    type: 'draft' | 'published' | 'recycle';
    subsection: number,
    content: string,
  }) {
    return request('/api/v1/works/chapter', { method: 'post', data: params });
  },
  Detail: function(params: { id: number, type: 'draft' | 'published' | 'recycle', subsection?: number }) {
    return request('/api/v1/works/detail', { params: params });
  },
  DeleteSubsection: function(params: { id: number }) {
    return request('/api/v1/works/subsection', {
      method: 'delete',
      data: params,
    });
  },
  UpdateSubsection: function(params: { id: number, name: string, introduce: string }) {
    return request('/api/v1/works/subsection', {
      method: 'post',
      data: params,
    });
  },
  WorkInfo: function(params: { id: number }) {
    return request('/api/v1/works/info', {
      params: {
        id: params.id,
      },
    });
  },
  NewWork: function(params: { name: string, category: number, topic: number, tags: string[] | undefined, introduce: string }) {
    console.log(params);
    return request('/api/v1/works/new', { method: 'post', data: params });
  },
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
};
