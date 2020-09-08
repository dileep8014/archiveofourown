import React from 'react';
import { Button, Checkbox, Form, Input, notification, Select, Skeleton, Tooltip } from 'antd';
import { QuestionCircleOutlined } from '@ant-design/icons';
import { history, useRouteMatch } from 'umi';
import { useRequest } from 'ahooks';
import { userService } from '@/service/user';
import ErrorPage from '@/component/errorpage/errorpage';

const { Option } = Select;

const formItemLayout = {
  labelCol: {
    xs: { span: 8 },
    sm: { span: 8 },
  },
  wrapperCol: {
    xs: { span: 8 },
    sm: { span: 8 },
  },
};

const tailFormItemLayout = {
  wrapperCol: {
    xs: {
      span: 24,
      offset: 0,
    },
    sm: {
      span: 16,
      offset: 8,
    },
  },
};

export default function() {

  // @ts-ignore
  const path = useRouteMatch().params.path;
  const { data, loading, error } = useRequest(userService.Identify, { defaultParams: [path] });
  const { run } = useRequest(userService.Create, {
    manual: true, onSuccess: (res) => {
      if (res.code == 0) {
        history.push('/');
      }
    },
  });

  const [form] = Form.useForm();

  const onFinish = (values: any) => {
    run({ email: data.data, username: values.username, password: values.password });
  };

  if (error) {
    return <ErrorPage title={error.name} subTitle={error.message}/>;
  }

  return (<Skeleton loading={loading}>{data &&
  <Form
    {...formItemLayout}
    style={{ textAlign: 'left' }}
    form={form}
    name="register"
    onFinish={onFinish}
    scrollToFirstError
  >
    <Form.Item
      name="email"
      label="邮箱"
    >
      <Input defaultValue={data.data} disabled/>
    </Form.Item>

    <Form.Item
      name="password"
      label="密码"
      rules={[
        {
          required: true,
        },
        {
          min: 8,
          max: 20,
        },
      ]}
      hasFeedback
    >
      <Input.Password/>
    </Form.Item>

    <Form.Item
      name="confirm"
      label="确认密码"
      dependencies={['password']}
      hasFeedback
      rules={[
        {
          required: true,
          message: '请确认你的密码',
        },
        {
          min: 8,
          max: 20,
        },
        ({ getFieldValue }) => ({
          validator(rule, value) {
            if (!value || getFieldValue('password') === value) {
              return Promise.resolve();
            }
            return Promise.reject('两次密码不一致');
          },
        }),
      ]}
    >
      <Input.Password/>
    </Form.Item>

    <Form.Item
      name="username"
      label={
        <span>
            用户名&nbsp;
          <Tooltip title="这将会是您在Pointer中的昵称以及笔名">
              <QuestionCircleOutlined/>
            </Tooltip>
          </span>
      }
      rules={[
        { required: true, whitespace: true, message: '请输入用户名' },
        { max: 20 },
      ]}
    >
      <Input/>
    </Form.Item>

    <Form.Item
      name="agreement"
      valuePropName="checked"
      rules={[
        { validator: (_, value) => value ? Promise.resolve() : Promise.reject('您必须接受协议后，才能正式注册Pointer！') },
      ]}
      {...tailFormItemLayout}
    >
      <Checkbox>
        我已经阅读了 <a href="">《用户协议》</a>和<a href="">《隐私政策》</a>
      </Checkbox>
    </Form.Item>
    <Form.Item {...tailFormItemLayout}>
      <Button type="primary" htmlType="submit">
        注册
      </Button>
    </Form.Item>
  </Form>}
  </Skeleton>);
}
