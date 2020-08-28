import ProList from '@ant-design/pro-list';
import Field from '@ant-design/pro-field';
import React from 'react';
import { useModel } from 'umi';
import { useRequest } from '@umijs/hooks';
import { Button, Descriptions, List, Skeleton, Space } from 'antd';
import DescriptionsItem from 'antd/es/descriptions/Item';
import Paragraph from 'antd/es/typography/Paragraph';
import Text from 'antd/es/typography/Text';
import { userService } from '@/service/user';

export default function TopicList() {

  const { user, userLoading } = useModel('user', model => ({ user: model.user, userLoading: model.loading }));

  const { data, loading, pagination } = useRequest(userService.UserTopics, { paginated: true });

  return (
    <Skeleton loading={userLoading && loading}>
      <ProList title='专题'
               itemLayout='vertical'
               split={false}
               dataSource={data?.list || []}
               renderItem={(item) => ({
                 title: item.title,
                 extra: (<Button type={'primary'}>修改专题设置</Button>),
                 children: (<Descriptions column={1}>
                   <DescriptionsItem label='分类'>
                     <Field text={item.category} mode={'read'}/>
                   </DescriptionsItem>
                   <DescriptionsItem label='作品数'>
                     <Field text={item.worksNum} mode={'read'}/>
                   </DescriptionsItem>
                   <DescriptionsItem label='原作'>
                     <Field text={[item.original, item.url]} mode={'read'}
                            render={(text) => <Text underline><a target='_blank' href={text[1]}>{text[0]}</a></Text>}/>
                   </DescriptionsItem>
                   <DescriptionsItem label='原作描述'>
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
