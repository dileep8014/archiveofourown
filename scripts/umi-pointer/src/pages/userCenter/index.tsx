import React, { useEffect, useState } from 'react';
import { Avatar, message, Space, Spin, Tabs, Upload, Radio, Descriptions, Button } from 'antd';
import ProCard from '@ant-design/pro-card';
import { useModel } from '@@/plugin-model/useModel';
import styles from './index.less';
import { LoadingOutlined, UploadOutlined } from '@ant-design/icons';
import ImgCrop from 'antd-img-crop';
import Field, { ProFieldFCMode } from '@ant-design/pro-field';
import { UserModelState } from '@/models/user';
import ButtonGroup from 'antd/es/button/button-group';
import moment from 'moment';

function getBase64(img: Blob, callback: { (imageUrl: any): void; (arg0: string | ArrayBuffer | null): any; }) {
  const reader = new FileReader();
  reader.addEventListener('load', () => callback(reader.result));
  reader.readAsDataURL(img);
}

export default () => {

  const { user, setUser } = useModel('user', model => ({
    user: model.user,
    setUser: model.setUser,
  }));

  const [userInfo, setUserInfo] = useState<UserModelState>(null);
  const [imgBtn, setImgBtn] = useState(false);
  const [imgLoading, setImgLoading] = useState(false);
  const [state, setState] = useState<ProFieldFCMode>('read');

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

  const handleChange = (info: any) => {
    if (info.file.status === 'uploading') {
      setImgLoading(true);
      return;
    }
    if (info.file.status === 'done') {
      getBase64(info.file.originFileObj, (imageUrl: any) => {
          setImgLoading(false);
          setUser(user && { ...user, avatar: imageUrl });
        },
      );
    }
  };

  const antIcon = <LoadingOutlined style={{ fontSize: 24 }} spin/>;
  const upIcon = <UploadOutlined style={{ fontSize: 24 }}/>;


  return (
    <Tabs tabPosition='left'>
      <Tabs.TabPane tab='个人资料' key='1'>
        <ProCard title={
          <div onClick={() => setImgBtn(false)}
               onMouseEnter={() => setImgBtn(!imgLoading)}
               onMouseLeave={() => setImgBtn(false)}>
            <ImgCrop rotate modalOk='确定' modalCancel='取消' modalTitle='头像裁剪'>
              <Upload name='头像'
                      accept='image/*'
                      showUploadList={false}
                      action="https://www.mocky.io/v2/5cc8019d300000980a055e76"
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
                 extra={<h1 className={styles.userName}>{user?.name}</h1>} headerBordered>
          <Radio.Group onChange={(e) => setState(e.target.value as ProFieldFCMode)} value={state}>
            <Radio value="read">只读</Radio>
            <Radio value="edit">编辑</Radio>
          </Radio.Group>
          <ButtonGroup>
            <Button disabled={state == 'read'} onClick={() => setUserInfo(user)}>取消</Button>
            <Button disabled={state == 'read'} type={'primary'} onClick={() => setUser(userInfo)}>修改</Button>
          </ButtonGroup>
          <Space align={'end'} style={{ float: 'right' }}>
            <Button style={{ marginLeft: 10 }} type={'primary'}>修改密码</Button>
            <Button style={{ marginLeft: 10 }} type={'primary'}>修改邮箱账户</Button>
          </Space>
          <br/>
          <br/>
          <Descriptions size='middle' column={2}>
            <Descriptions.Item label="笔名">
              <Field text={user?.name} value={userInfo?.name} mode={state}
                     formItemProps={{ maxlength: 10 }}
                     onChange={(e) => setUserInfo(userInfo && { ...userInfo, name: e.target.value })}/>
            </Descriptions.Item>
            <Descriptions.Item label="邮箱">
              <Field text={user?.email} mode={'read'}/>
            </Descriptions.Item>
            <Descriptions.Item label="性别">
              <Field text={user?.gender} value={userInfo?.gender} mode={state}
                     valueEnum={{ secret: { text: '保密' }, man: { text: '男' }, woman: { text: '女' } }}
                     onChange={(e) => {
                       if (!e) {
                         setUserInfo(userInfo && { ...userInfo, gender: 'secret' });
                       } else {
                         setUserInfo(userInfo && { ...userInfo, gender: e.target.value });
                       }
                     }}
              />
            </Descriptions.Item>
            <Descriptions.Item label="个人简介">
              <Field text={user?.introduce} value={userInfo?.introduce} mode={state} valueType={'textarea'}
                     formItemProps={{
                       style: { width: 400 },
                       autoSize: { minRows: 4, maxRows: 4 },
                       maxlength: 200,
                     }}
                     onChange={(e) => setUserInfo(userInfo && { ...userInfo, introduce: e.target.value })}
              />
            </Descriptions.Item>
            <Descriptions.Item label="注册时间">
              <Field text={moment().format('YYYY年 MM月 DD日')} mode={'read'}/>
            </Descriptions.Item>
          </Descriptions>
        </ProCard>
      </Tabs.TabPane>
      <Tabs.TabPane tab='偏好设置' key='2'>

      </Tabs.TabPane>
    </Tabs>
  );
}
