import { Button, Checkbox, Form, Input, Tabs } from 'antd';
import React from 'react';
import { useModel } from '@@/plugin-model/useModel';
import './sign.less';
import { useRequest } from '@umijs/hooks';
import service from '@/component/service';
import Text from 'antd/es/typography/Text';
import { Link } from 'umi';
import ProgressOpt from '@/component/progress/progress';

const LoginForm = (props: { finish: () => void }) => {
  const [form] = Form.useForm();
  const { signIn } = useModel('user', model => ({
    signIn: model.signin,
  }));

  const onFinish = (values: any) => {
    if (values.rememberme === undefined) {
      values.rememberme = false;
    }
    ProgressOpt(() => signIn(values.username, values.password));
    props.finish();
  };

  return (
    <Form form={form}
          style={{ padding: 20, textAlign: 'unset' }}
          name="login"
          initialValues={{ remember: true }}
          scrollToFirstError
          onFinish={onFinish}>
      <Form.Item name="username" rules={
        [{ required: true, message: '用户名不能为空' }]}>
        <Input size={'large'} className="input-prepend restyle" placeholder="用户名 / 邮箱"/>
      </Form.Item>
      <Form.Item name="password" rules={
        [{ required: true, message: '密码不能为空' }]
      }>
        <Input type={'password'} className="input-prepend" size={'large'} placeholder="密码"/>
      </Form.Item>
      <Form.Item>
        <Form.Item name="rememberme" valuePropName="checked" className="remember-btn">
          <Checkbox>记住我</Checkbox>
        </Form.Item>
        <div className="forgot-btn">
          <a href="/?">
            忘记密码
          </a>
        </div>
      </Form.Item>
      <Form.Item>
        <Button htmlType="submit" className="login-btn">登录</Button>
      </Form.Item>
    </Form>
  );
};

const RegisterForm = (props: { finish: () => void }) => {
  const [form] = Form.useForm();
  const { run } = useRequest(
    ({ account, email, password }) => service.SignUp(account, email, password), {
      manual: true,
    });

  const onFinish = (values: any) => {
    run(values);
    props.finish();
  };


  return (
    <div style={{ textAlign: 'center' }}>
      <Form form={form}
            style={{ padding: 20, textAlign: 'unset' }}
            name="signUp"
            initialValues={{ remember: true }}
            scrollToFirstError
            onFinish={onFinish}>
        <Form.Item name="email" rules={
          [{ required: true, message: '邮箱不能为空' }]}>
          <Input size={'large'} className="input-prepend restyle" type={'email'} placeholder="邮箱"/>
        </Form.Item>
        <Form.Item>
          <Button htmlType="submit" className="register-btn">注册</Button>
        </Form.Item>
      </Form>
      <div>
        <Text>点击注册后Pointer会向你发送注册邮件，请根据邮件指示注册账户</Text>
      </div>
      <div>
        <Text>注册即代表你同意<Link to='/'>《用户协议》</Link>和<Link to='/'>《隐私政策》</Link></Text>
      </div>
    </div>
  );
};

interface SignProps {
  tab: string,
  finish: () => void
}

const SignForm: React.FC<SignProps> = ({ tab, finish }) => {

  return (
    <div className='sign'>
      <h2 style={{ textAlign: 'center' }}>Pointer</h2>
      <Tabs centered defaultActiveKey={tab}>
        <Tabs.TabPane tab='登录' key='signIn'>
          <LoginForm finish={finish}/>
        </Tabs.TabPane>
        <Tabs.TabPane tab='注册' key='signUp'>
          <RegisterForm finish={finish}/>
        </Tabs.TabPane>
      </Tabs>

    </div>
  );
};

export default SignForm;
