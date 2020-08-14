import React from 'react';
import {  Tabs} from 'antd';
import UserInfo from '@/pages/userCenter/userinfo';
import UserStyle from '@/pages/userCenter/userstyle';


export default () => {


  return (
    <Tabs tabPosition='left' style={{ paddingLeft: 100, paddingRight: 100 }}>
      <Tabs.TabPane tab='个人资料' key='1'>
        <UserInfo/>
      </Tabs.TabPane>
      <Tabs.TabPane tab='偏好设置' key='2'>
        <UserStyle/>
      </Tabs.TabPane>
      <Tabs.TabPane tab='收藏夹' key='3'>

      </Tabs.TabPane>
    </Tabs>
  );
}
