import React, { useState } from 'react';
import { useModel } from 'umi';
import { Avatar, Button, Checkbox, Dropdown, Form, Input, Menu, Modal } from 'antd';
import NProgress from 'nprogress';
import 'nprogress/nprogress.css';
import { UserOutlined, LockOutlined } from '@ant-design/icons';
import styles from './rightHeader.less';

const RightHeader: React.FC = () => {

  const { user, signIn, signOut } = useModel('user', model => ({
    user: model.user,
    signIn: model.signin,
    signOut: model.signout,
  }));

  const [loginVisible, setLoginVisible] = useState(false);
  const [registerVisible, setRegisterVisible] = useState(false);
  const [top, setTop] = useState(0);

  const onFinish = (values: any) => {
    if (values.rememberme === undefined) {
      values.rememberme = false;
    }
    NProgress.start();
    signIn(values.username, values.password);
    NProgress.done();
  };

  const visit = () => {
    if (!loginVisible) {
      setTop(100);
      setLoginVisible(true);
    } else {
      setTop(0);
      setLoginVisible(false);
    }
  };

  if (user) {
    return (
      <div className={styles.right}>
        <Dropdown overlay={
          <Menu onClick={signOut}>
            <Menu.Item>登出</Menu.Item>
          </Menu>}>
          <div className={styles.userAvatar}>
            <Avatar src={user?.avatar}/><span>{user?.name}</span>
          </div>
        </Dropdown>
        <Button className={styles.createBtn}>发布</Button>
      </div>
    );
  }

  return (
    <div>
      <Button onClick={visit}>登录</Button>
      <Modal
        centered={true}
        width={600}
        footer={null}
        onCancel={visit}
        visible={loginVisible}
      >
        <LoginForm onFinish={onFinish}/>
      </Modal>
      <Button>注册</Button>
    </div>
  );

};

const LoginForm = (props: {
  onFinish: (value: any) => void,
}) => {
  const [form] = Form.useForm();

  return (
    <Form form={form} style={{ padding: 20}} name="login" initialValues={{ remember: true }} scrollToFirstError
          onFinish={props.onFinish}>
      <Form.Item name="username" className="input-prepend restyle" rules={
        [{ required: true, message: '用户名不能为空' }]}>
        <Input size={'large'} prefix={<UserOutlined className="icon"/>}
               placeholder="用户名 / 邮箱"/>
      </Form.Item>
      <Form.Item name="password" className="input-prepend" rules={
        [{ required: true, message: '密码不能为空' }]
      }>
        <Input.Password prefix={<LockOutlined/>} size={'large'} placeholder="Password"/>
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

export default RightHeader;
