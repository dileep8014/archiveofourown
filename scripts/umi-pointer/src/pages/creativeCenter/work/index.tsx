import React, { useState } from 'react';
import { Anchor, Button, Card, Col, Row, Skeleton, Tooltip } from 'antd';
import { useModel, history } from 'umi';
import { useRequest } from '@umijs/hooks';
import service from '@/component/service';
import ErrorPage from '@/component/errorpage/errorpage';
import ProList from '@ant-design/pro-list';
import ButtonGroup from 'antd/es/button/button-group';
import { EditOutlined, UnorderedListOutlined, TableOutlined} from '@ant-design/icons';

export default function MyWork() {

  const { user } = useModel('user', model => ({ user: model.user }));

  const { data, loading, error, pagination } = useRequest(service.MineWorks, { paginated: true });

  const [list, unList] = useState(true);

  if (error) {
    return <ErrorPage title={error.name} subTitle={error.message}/>;
  }

  return (
    <Skeleton loading={loading}>
      <Row>
        <Col span={22}>
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
            split={false}
            renderItem={(item) => ({
              children: (list &&
                <Card hoverable onClick={() => history.push('/work/' + item.id)}>
                  <Card.Meta title={item.title}
                             avatar={<img style={{ height: 120, width: 100 }} src={item.cover} alt={item.title}/>}
                             description={item.introduce}
                  />
                </Card> ||
                <Card hoverable style={{ maxWidth: 200 }} onClick={() => history.push('/work/' + item.id)}
                      cover={<img src={item.cover} alt={item.title}/>}>
                  <Card.Meta title={item.title}/>
                </Card>
              ),
            })}
            pagination={{
              ...pagination,
              onChange: (current, pageSize) => pagination.onChange(current, pageSize || 10),
            }}
          />
        </Col>
        <Col span={2}>
          <Anchor offsetTop={window.innerHeight / 2}>
            <Tooltip placement={'left'} title='新的作品'>
              <Button shape={'circle'} icon={ <EditOutlined />}/>
            </Tooltip>
          </Anchor>
        </Col>
      </Row>
    </Skeleton>
  );
}
