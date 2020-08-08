import React, { ReactText, useEffect, useState } from 'react';
import './index.less';
import { useRequest } from '@umijs/hooks';
import service from '@/component/service';
import ErrorPage from '@/component/errorpage/errorpage';
import { Col, Row, Skeleton, Space, Tag } from 'antd';
import ProList from '@ant-design/pro-list';
import { useModel } from '@@/plugin-model/useModel';
import { Random } from 'mockjs';
import IconText from '@/component/iconfont/icontext';


export default () => {
  const { user } = useModel('user', model => ({ user: model.user }));
  const { run: fetchTags, data: tags } = useRequest(service.HotsTopic, { manual: true });
  const { run: fetchWorks, data, error, loading, pagination } = useRequest(
    ({ current, pageSize }) => service.SubWorks({ current, pageSize }), {
      paginated: true,
      manual: true,
    });

  const [expandedRowKeys, setExpandedRowKeys] = useState<ReactText[]>([]);

  useEffect(() => {
    if (user) {
      fetchWorks({ current: 1, pageSize: 10 });
    } else {
      fetchTags();
    }
  }, [user]);


  if (error) {
    return <ErrorPage title='首页错误' subTitle={error.message}/>;
  }

  return (
    <Row className='indexPage'>
      <Col xxl={16} xl={16} lg={18} md={20} sm={24} xs={24}>
        {user && <Skeleton loading={loading}>
          <ProList
            title="我的订阅"
            rowKey="subTitle"
            expandable={{ expandedRowKeys, onExpandedRowsChange: setExpandedRowKeys }}
            split={false}
            itemLayout='vertical'
            dataSource={data?.list || []}
            pagination={{
              ...pagination,
              onChange: (page, pageSize) => {
                return pagination.onChange(page, pageSize || 10);
              },
            }}
            renderItem={(item) => ({
              title: (<span>
                {item.title}
                <Space size={5}>
                  <Tag color={Random.color()} style={{ marginLeft: 8 }}>{item.topic}</Tag>
                </Space>
              </span>),
              children: (<span>{item.subTitle}</span>),
              avatar: item?.user?.avatar,
              actions: [
                <IconText icon={'字数'} text={item.words} key="words"/>,
                <IconText icon={'章节'} text={item.chapters} key="chapters"/>,
                <IconText icon={'评论'} text={item.comments} key="comments"/>,
                <IconText icon={'收藏'} text={item.collection} key="collection"/>,
                <IconText icon={'订阅'} text={item.subscribe} key="subscribe"/>,
                <IconText icon={'点击'} text={item.hits} key="hits"/>,
              ],
            })}
          />
        </Skeleton>}
        {!user &&
        <div>
          <h4>热门专题</h4>
          {tags && tags.list?.map((item: string) =>
            <Tag
                 color={Random.color()}
                 style={{ fontSize: item.length + 5, margin: item.length }}>
              {item}
            </Tag>,
          )}
        </div>}
      </Col>
      <Col xxl={8} xl={8} lg={6} md={4} sm={0} xs={0}>
        <Skeleton loading={false}>
          <ProList
            title="公告"
            locale={{ emptyText: '暂无公告' }}
            split={false}
            itemLayout='vertical'
            dataSource={[]}
            renderItem={(item) => ({})}
          />
        </Skeleton>
      </Col>
    </Row>
  );
}
