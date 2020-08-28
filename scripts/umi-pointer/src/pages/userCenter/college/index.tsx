import ProList from '@ant-design/pro-list';
import Field from '@ant-design/pro-field';
import React from 'react';
import { useModel } from 'umi';
import { useRequest } from '@umijs/hooks';
import { Button, Descriptions, Skeleton } from 'antd';
import DescriptionsItem from 'antd/es/descriptions/Item';
import Paragraph from 'antd/es/typography/Paragraph';
import { userService } from '@/service/user';

export default function CollegeList() {

  const { user, userLoading } = useModel('user', model => ({ user: model.user, userLoading: model.loading }));

  const { data, loading, pagination } = useRequest(userService.UserCollege, { paginated: true });

  return (
    <Skeleton loading={userLoading && loading}>
      <ProList title={user?.name + '的收藏夹'}
               itemLayout='vertical'
               split={false}
               dataSource={data?.list || []}
               renderItem={(item) => ({
                 title: item.title,
                 extra: (<Button type={'primary'}>修改</Button>),
                 children: (<Descriptions column={1}>
                   <DescriptionsItem label='作品数'>
                     <Field text={item.worksNum} mode={'read'}/>
                   </DescriptionsItem>
                   <DescriptionsItem label='描述'>
                     <Field text={item.description} mode={'read'}
                            render={(text) => <Paragraph ellipsis={{ rows: 2, expandable: true }}>{text}</Paragraph>}/>
                   </DescriptionsItem>
                 </Descriptions>),
               })}
               pagination={{
                 ...pagination,
                 onChange: (current, pageSize) => pagination.onChange(current, pageSize || 10),
               }}
      />
    </Skeleton>
  );
}
