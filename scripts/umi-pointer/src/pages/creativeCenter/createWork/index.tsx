import {
  Button,
  Card,
  Col,
  Divider,
  Form,
  Input,
  PageHeader,
  Radio, Result,
  Row,
  Select,
  Skeleton, Spin,
  Steps,
} from 'antd';
import React, { useState } from 'react';
import { Link, history } from 'umi';
import { useRequest } from '@umijs/hooks';
import { Option } from 'antd/es/mentions';
import Text from 'antd/es/typography/Text';
import ProgressOpt from '@/component/progress/progress';
import { categoryService } from '@/service/category';
import { topicService } from '@/service/topic';
import { tagService } from '@/service/tag';
import { workService } from '@/service/work';

const { Step } = Steps;

export default function NewWork(props: { cancel: () => void }) {

  const { data: category, run: fetchCategory, loading: categoryLoad } = useRequest(categoryService.Category, { manual: true });
  const { data: topics, run: fetchTopics, loading: topicsLoad } = useRequest(topicService.ListTopicByCategory, { manual: true });
  const { data: tags, run: fetchTags, loading: tagsLoad } = useRequest(tagService.SimilarTags, { manual: true });
  const { run } = useRequest(workService.NewWork, { manual: true, onSuccess: () => next() });


  const [current, setCurrent] = useState(0);
  const [type, setType] = useState<number>(0);

  const [form] = Form.useForm();

  const next = () => setCurrent(current + 1);
  const previous = () => setCurrent(current - 1);

  const onFinish = (values: any) => {
    run(values);
  };


  const radioStyle = {
    display: 'flex',
    marginTop: 5,
  };

  return (
    <PageHeader title='创建作品'
                breadcrumb={{
                  routes: [{ path: 'dashboard', breadcrumbName: '工作台' },
                    { path: 'newWork', breadcrumbName: '新的作品' }],
                  itemRender: (route, params, routes, paths) => {
                    if (route.breadcrumbName == '新的作品') {
                      return <Link to={'/creativeCenter'}>{route.breadcrumbName}</Link>;
                    }
                    return <Link onClick={props.cancel} to='/creativeCenter'>
                      {route.breadcrumbName}
                    </Link>;
                  },
                }}
    >
      <Steps current={current}>
        <Step key='1' title='选择类型'/>
        <Step key='2' title='完善作品信息'/>
        <Step key='3' title='创建成功'/>
      </Steps>
      <div style={{ paddingTop: 20, textAlign: 'center' }}>
        {current == 0 && <Radio.Group onChange={e => setType(e.target.value)} value={type}>
          <Card bordered={false} style={{ textAlign: 'left' }}>
            <Card.Meta title={<Radio style={radioStyle} value={1}>同人</Radio>}
                       description='利用原有的漫画、动画、小说、影视作品中的人物角色、故事情节或背景设定等元素进行的二次创作小说。'/>
          </Card>
          <Card bordered={false} style={{ textAlign: 'left' }}>
            <Card.Meta title={<Radio style={radioStyle} value={2}>原创</Radio>}
                       description='以玄幻/历史/都市/灵异/仙侠/游戏/二次元/武侠/军事/言情/青春/悬疑/科幻等题材创作的类型小说。'/>
          </Card>
          <Button type="primary" onClick={() => {
            ProgressOpt(() => fetchCategory({ type: type }));
            next();
          }} disabled={type == 0}>
            下一步
          </Button>
        </Radio.Group>}
        {current == 1 && <Skeleton loading={categoryLoad}>
          <Form form={form} onFinish={onFinish}>
            <Form.Item label='作品名称' name='name' required rules={[{ required: true }, {
              pattern: new RegExp('^[a-zA-Z0-9\u4e00-\u9fa5]+$'),
              message: '不允许输入特殊符号',
            }]}>
              <Input maxLength={15} placeholder='请填写'/>
            </Form.Item>
            <Text disabled>15字以内，允许输入中/英文和数字，请勿添加书名号等特殊符号</Text>
            <br/>
            <br/>
            <Form.Item label='作品类型' required>
              <Row gutter={20}>
                <Col span={12}>
                  <Form.Item name='category' rules={[{ required: true, message: '请选择分类' }]}>
                    <Select showSearch placeholder="请选择"
                            onSelect={(v: string) => ProgressOpt(() => fetchTopics({ categoryId: parseInt(v) }))}>
                      {category && category.map((item: { id: number; name: string; }) =>
                        <Option key={item.id} value={item.id.toString()}>{item.name}</Option>)}
                    </Select>
                  </Form.Item>
                </Col>
                <Col span={12}>
                  <Form.Item name='topic' rules={[{ required: true, message: '请选择类别' }]}>
                    <Select showSearch placeholder="请选择" loading={topicsLoad}>
                      {topics && topics.map((item: { id: number; name: string; }) =>
                        <Option key={item.id} value={item.id.toString()}>{item.name}</Option>)}
                    </Select>
                  </Form.Item>
                </Col>
              </Row>
            </Form.Item>
            <Form.Item label='作品标签(多)' name='tags'>
              <Select mode={'tags'} placeholder='请填写'
                      onSearch={(v: string) => ProgressOpt(() => fetchTags({ similar: v }))}
                      notFoundContent={tagsLoad ? <Spin size="small"/> : null}
                      filterOption={false}
              >
                {tags && tags.map((item: { name: string; }) =>
                  <Option key={item.name} value={item.name}>{item.name}</Option>)}
              </Select>
            </Form.Item>
            <Divider type={'horizontal'}/>
            <Form.Item label='作品介绍' name='introduce' required rules={[{ required: true }]}>
              <Input.TextArea maxLength={200} rows={4} autoSize={{ maxRows: 4, minRows: 4 }} placeholder='请填写,200字以内'/>
            </Form.Item>
            <Form.Item>
              <Button onClick={previous}>
                选择类型
              </Button>
              <Button style={{ marginLeft: 20 }} type={'primary'} htmlType='submit'>
                创建作品
              </Button>
            </Form.Item>
          </Form></Skeleton>}
        {current == 2 && <Result title='创建作品成功' status={'success'}>
          <Button onClick={() => history.push('/works/new')}>
            开始写作
          </Button>
        </Result>}
      </div>
    </PageHeader>
  );
}
