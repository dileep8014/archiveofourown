import React, { useEffect, useState } from 'react';
import './index.less';
import { useRequest } from '@umijs/hooks';
import service from '@/component/service';
import ErrorPage from '@/component/errorpage/errorpage';
import { Button, Card, Col, Divider, Modal, Row, Skeleton, Space, Tag } from 'antd';
import ProList from '@ant-design/pro-list';
import { useModel } from '@@/plugin-model/useModel';
import { Random } from 'mockjs';
import { Link } from 'umi';
import moment from 'moment';
import Text from 'antd/es/typography/Text';
import SignForm from '@/component/sign/sign';
import Paragraph from 'antd/es/typography/Paragraph';


export default () => {
  const { user } = useModel('user', model => ({ user: model.user }));
  const { run: fetchWorks, data, error, loading } = useRequest(
    ({ current, pageSize }) => service.SubWorks({ current, pageSize }), {
      manual: true,
    });
  const { data: news } = useRequest(
    ({ current, pageSize }) => service.News({ current, pageSize }), {
      defaultParams: [{ current: 1, pageSize: 3 }],
    });
  const [canVisit, setCanVisit] = useState(false);

  const visit = (cancel: boolean) => {
    if (cancel) {
      setCanVisit(false);
    } else {
      setCanVisit(true);
    }
  };

  useEffect(() => {
    if (user) {
      fetchWorks({ current: 1, pageSize: 10 });
    }
  }, [user]);


  if (error) {
    return <ErrorPage title='首页错误' subTitle={error.message}/>;
  }

  return (
    <Row className='indexPage'>
      {canVisit && <Modal
        centered={true}
        width={400}
        footer={null}
        onCancel={() => visit(true)}
        visible={true}
      >
        <SignForm tab={'signUp'} finish={() => visit(true)}/>
      </Modal>}
      <Col xxl={16} xl={16} lg={18} md={20} sm={24} xs={24}>
        <Skeleton loading={loading}>
          {user && <ProList
            size='large'
            title={<div><span>最新订阅</span> <Divider type={'horizontal'}/></div>}
            split={false}
            itemLayout='vertical'
            dataSource={data?.list || []}
            footer={<Link style={{ float: 'right', fontSize: 16 }} to='/subscribe'><Text
              underline>>>更多订阅</Text></Link>}
            renderItem={(item) => ({
              title: (<span>
                <Link to={'/works/' + item.id}>{item.title}</Link>
                <Space size={5}>
                  <Tag color={Random.color()} style={{ marginLeft: 8 }}>{item.topic}</Tag>
                </Space>
              </span>),
              description: (item.tags && <Paragraph ellipsis={{ rows: 1, expandable: true }}>
                {item.tags.map((item: React.ReactNode) => (<Text style={{marginLeft:10}} underline>{item}</Text>))}
              </Paragraph>),
              children: (
                <div>
                  <Paragraph ellipsis={{ rows: 2, expandable: true }}>{item.newChapterDesc}</Paragraph>
                  <span style={{ float: 'right' }}>最新章节：{item.newChapter}</span>
                </div>),
              avatar: { src: item?.user?.avatar, size: 'large' },
            })}
          />}
          {!user && <div style={{ fontSize: 16 }}>
            <span>Pointer旨在创建一个非营业的开放性的自由创作平台，目前Pointer还处于初步创建时期，我们更多的内容会聚焦在同人作品上。</span>
            <Card bordered style={{ marginTop: 20 }}
                  bodyStyle={{ fontSize: 14 }}
                  actions={[<Button type={'primary'} onClick={() => visit(false)}>立即注册</Button>]}
            >
              <Card.Meta
                description={<Text>
                  Pointer是非商业的一个项目。

                  使用Pointer，您可以：
                  <ul>
                    <li>分享自己的作品</li>
                    <li>在您最喜欢的作品，专题或用户更新时得到通知</li>
                    <li>参与线上创作活动</li>
                    <li>获得更专业更精确的作品查询筛选服务</li>
                  </ul>
                  您可以通过邮箱自由创建属于自己的Pointer账户。欢迎所有志同道合的同伴！
                </Text>}/>
            </Card>
          </div>}
        </Skeleton>
      </Col>
      <Col xxl={8} xl={8} lg={6} md={4} sm={0} xs={0}>
        <Skeleton loading={false}>
          <ProList
            size='large'
            title={<div>
              <span>通知</span>
              <Link style={{ float: 'right' }} to='/news'>所有通知</Link>
              <Divider type={'horizontal'}/>
            </div>}
            locale={{ emptyText: '暂无通知' }}
            split={false}
            itemLayout='vertical'
            dataSource={news?.list || []}
            renderItem={(item) => ({
              title: item.title,
              subTitle: (<span>发布时间: {moment().format('yyyy年M月D日')} 评论: {item.comments}</span>),
              children: (<div>
                <span>{item.description}</span>
                <br/>
                <Link style={{ float: 'right' }} to={'/news/' + item.id}><Text underline>阅读更多</Text></Link>
              </div>),
            })}
          />
        </Skeleton>
      </Col>
    </Row>
  );
}
