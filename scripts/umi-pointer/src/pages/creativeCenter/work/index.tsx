import React, { useState } from 'react';
import { Button, Card, Skeleton } from 'antd';
import { useModel, history } from 'umi';
import { useRequest } from '@umijs/hooks';
import ErrorPage from '@/component/errorpage/errorpage';
import ProList from '@ant-design/pro-list';
import ButtonGroup from 'antd/es/button/button-group';
import { UnorderedListOutlined, TableOutlined } from '@ant-design/icons';
import WorkItem from '@/pages/creativeCenter/workItem/workItem';
import { workService } from '@/service/work';

export default function MyWork() {

  const { user } = useModel('user', model => ({ user: model.user }));

  const { data, loading, error, pagination } = useRequest(workService.MineWorks, { paginated: true });

  const [list, unList] = useState(true);

  if (error) {
    return <ErrorPage title={error.name} subTitle={error.message}/>;
  }

  return (
    <Skeleton loading={loading}>
      <ProList
        title={
          <div>
            <h1>{user?.name}的作品</h1>
            <ButtonGroup style={{ float: 'right' }}>
              <Button disabled={list} onClick={() => unList(true)} icon={<UnorderedListOutlined/>}/>
              <Button disabled={!list} onClick={() => unList(false)} type={'primary'} icon={<TableOutlined/>}/>
            </ButtonGroup>
          </div>
        }
        itemLayout={(list && 'vertical') || 'horizontal'}
        grid={{ column: (list && 1) || 4 }}
        dataSource={data?.list || []}
        renderItem={(item) => ({
          children: (list &&
            <WorkItem {...item}/> ||
            <Card style={{ maxWidth: 200 }} onClick={() => history.push('/work/' + item.id)}
                  cover={<img src={item.cover} alt={item.title}/>}
                  actions={[<Button>作品设置</Button>, <Button type={'primary'}>新章节</Button>]}
            >
              <Card.Meta description={item.title}/>
            </Card>
          ),
        })}
        pagination={{
          ...pagination,
          onChange: (current, pageSize) => pagination.onChange(current, pageSize || 10),
        }}
      />
    </Skeleton>
  );
}
