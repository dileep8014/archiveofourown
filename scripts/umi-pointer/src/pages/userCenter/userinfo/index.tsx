import ImgCrop from 'antd-img-crop';
import {
  Avatar,
  Button,
  Descriptions,
  Form,
  Input,
  message,
  Modal,
  notification,
  Radio,
  Space,
  Spin,
  Upload,
} from 'antd';
import styles from './index.less';
import Field, { ProFieldFCMode } from '@ant-design/pro-field';
import ButtonGroup from 'antd/es/button/button-group';
import moment from 'moment';
import ProCard from '@ant-design/pro-card';
import React, { useEffect, useState } from 'react';
import { useModel } from '@@/plugin-model/useModel';
import { UserModelState } from '@/models/user';
import { LoadingOutlined, UploadOutlined } from '@ant-design/icons';
import { baseurl } from '@/service/request';
import { UploadChangeParam } from 'antd/lib/upload/interface';
import { useRequest } from 'ahooks';
import { userService } from '@/service/user';


const formItemLayout = {
  labelCol: {
    xs: { span: 6 },
    sm: { span: 6 },
  },
  wrapperCol: {
    xs: { span: 14 },
    sm: { span: 14 },
  },
};

export default function UserInfo() {

  const { user, setUser } = useModel('user', model => ({
    user: model.user,
    setUser: model.setUser,
  }));

  const { run: runPass } = useRequest(userService.UpdatePassword, {
    manual: true, onSuccess: res => {
      if (res.code == 0) {
        message.success("修改成功")
        setUpdatePass(false);
      }
    },
  });
  const { run: runEmail } = useRequest(userService.UpdateEmail, {
    manual: true, onSuccess: res => {
      if (res.code == 0) {
        message.success("修改成功")
        setUpdateEmail(false);
      }
    },
  });

  const [userInfo, setUserInfo] = useState<UserModelState>(null);
  const [imgBtn, setImgBtn] = useState(false);
  const [imgLoading, setImgLoading] = useState(false);
  const [updateEmail, setUpdateEmail] = useState(false);
  const [updatePass, setUpdatePass] = useState(false);
  const [state, setState] = useState<ProFieldFCMode>('read');

  const [form1] = Form.useForm();
  const [form2] = Form.useForm();


  useEffect(() => {
    if (user) {
      setUserInfo(user);
    }
  }, [user]);

  const beforeUpload = (file: { type: string; size: number; }) => {
    setImgBtn(false);
    const isLt2M = file.size / 1024 / 1024 < 2;
    if (!isLt2M) {
      message.error('图片大小必须小于 2MB!');
    }
    return isLt2M;
  };

  const handleChange = (info: UploadChangeParam) => {
    if (info.file.status === 'uploading') {
      setImgLoading(true);
      return;
    }
    if (info.file.status === 'done') {
      setImgLoading(false);
      setUser({ avatar: info.file.response?.data });
    }
    if (info.file.status == 'error') {
      setImgLoading(false);
      notification['error']({ message: info.file.response.msg, description: info.file.response.details });
    }
  };

  const antIcon = <LoadingOutlined style={{ fontSize: 24 }} spin/>;
  const upIcon = <UploadOutlined style={{ fontSize: 24 }}/>;


  return (
    <ProCard title={
      <div onClick={() => setImgBtn(false)}
           onMouseEnter={() => setImgBtn(!imgLoading)}
           onMouseLeave={() => setImgBtn(false)}>
        <ImgCrop rotate modalOk='确定' modalCancel='取消' modalTitle='头像裁剪'>
          <Upload name='file'
                  accept='image/*'
                  showUploadList={false}
                  action={`${baseurl}/api/v1/upload`}
                  beforeUpload={beforeUpload}
                  onChange={handleChange}
                  openFileDialogOnClick={imgBtn}
          >
            <Spin indicator={upIcon} spinning={imgBtn}>
              <Spin indicator={antIcon} spinning={imgLoading}>
                <Avatar src={user?.avatar} className={styles.userAvatar} shape={'circle'}/>
              </Spin>
            </Spin>
          </Upload>
        </ImgCrop>
      </div>}
             extra={<h1 className={styles.userName}>{user?.username}</h1>} headerBordered>
      <Radio.Group onChange={(e) => setState(e.target.value as ProFieldFCMode)} value={state}>
        <Radio value="read">只读</Radio>
        <Radio value="edit">编辑</Radio>
      </Radio.Group>
      <ButtonGroup>
        <Button disabled={state == 'read'} onClick={() => setUserInfo(user)}>取消</Button>
        <Button disabled={state == 'read'} type={'primary'} onClick={() => {
          setState('read');
          setUser({ username: userInfo?.username, introduce: userInfo?.introduce, gender: userInfo?.gender });
        }}>修改</Button>
      </ButtonGroup>
      <Space align={'end'} style={{ float: 'right' }}>
        <Button style={{ marginLeft: 10 }} type={'primary'} onClick={() => setUpdatePass(true)}>修改密码</Button>
        <Modal visible={updatePass} title='修改用户密码' onCancel={() => setUpdatePass(false)} footer={null}>
          <Form {...formItemLayout} form={form1}
                onFinish={values => runPass({ oldPassword: values.oldPassword, password: values.password })}>
            <Form.Item label='旧密码' name='oldPassword' rules={[{ required: true }]} hasFeedback>
              <Input.Password/>
            </Form.Item>
            <Form.Item label='新密码' name='password' rules={[{ required: true }, { min: 8, max: 20 }]} hasFeedback>
              <Input.Password/>
            </Form.Item>
            <Form.Item label='确认新密码' name='confirm' dependencies={['password']}
                       rules={[{ required: true }, { min: 8, max: 20 }, ({ getFieldValue }) => ({
                         validator(rule, value) {
                           if (!value || getFieldValue('password') === value) {
                             return Promise.resolve();
                           }
                           return Promise.reject('两次密码不一致');
                         },
                       })]} hasFeedback>
              <Input.Password/>
            </Form.Item>
            <Form.Item labelCol={{ span: 0 }} wrapperCol={{ span: 24 }} style={{ textAlign: 'center' }}>
              <Button type="primary" htmlType="submit" style={{ width: '40%' }}>确认修改</Button>
            </Form.Item>
          </Form>
        </Modal>
        <Button style={{ marginLeft: 10 }} type={'primary'} onClick={() => setUpdateEmail(true)}>修改邮箱账户</Button>
        <Modal visible={updateEmail} title='修改用户邮箱' onCancel={() => setUpdateEmail(false)} footer={null}>
          <Form {...formItemLayout} form={form2}
                onFinish={values => runEmail({ email: values.email, password: values.password })}>
            <Form.Item label='密码' name='password' rules={[{ required: true }]} hasFeedback>
              <Input.Password/>
            </Form.Item>
            <Form.Item label='新的邮箱' name='email' rules={[{ required: true }, { type: 'email' }]} hasFeedback>
              <Input type={'email'}/>
            </Form.Item>
            <Form.Item labelCol={{ span: 0 }} wrapperCol={{ span: 24 }} style={{ textAlign: 'center' }}>
              <Button type="primary" htmlType="submit" style={{ width: '40%' }}>确认修改</Button>
            </Form.Item>
          </Form>
        </Modal>
      </Space>
      <br/>
      <br/>
      <Descriptions size='middle' column={2}>
        <Descriptions.Item label="笔名">
          <Field text={user?.username} value={userInfo?.username} mode={state}
                 formItemProps={{ maxLength: 10 }}
                 onChange={(e) => setUserInfo(userInfo && { ...userInfo, username: e.target.value })}/>
        </Descriptions.Item>
        <Descriptions.Item label="邮箱">
          <Field text={user?.email} mode={'read'}/>
        </Descriptions.Item>
        <Descriptions.Item label="性别">
          <Field text={user?.gender.toString()} value={userInfo?.gender.toString()} mode={state}
                 valueEnum={{ '0': { text: '保密' }, '1': { text: '男' }, '2': { text: '女' } }}
                 onChange={(e) => {
                   console.log(e);
                   if (!e) {
                     setUserInfo(userInfo && { ...userInfo, gender: 0 });
                   } else {
                     setUserInfo(userInfo && { ...userInfo, gender: parseInt(e) });
                   }
                 }}
          />
        </Descriptions.Item>
        <Descriptions.Item label="个人简介">
          <Field text={user?.introduce} value={userInfo?.introduce} mode={state} valueType={'textarea'}
                 formItemProps={{
                   style: { width: 400 },
                   autoSize: { minRows: 4, maxRows: 4 },
                   maxLength: 200,
                 }}
                 onChange={(e) => setUserInfo(userInfo && { ...userInfo, introduce: e.target.value })}
          />
        </Descriptions.Item>
        <Descriptions.Item label="注册时间">
          <Field text={moment(user?.createdAt).format('YYYY年 MM月 DD日')} mode={'read'}/>
        </Descriptions.Item>
      </Descriptions>
    </ProCard>
  );
}
