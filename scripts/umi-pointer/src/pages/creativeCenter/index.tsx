import styles from './index.less';
import { Tabs } from 'antd';
import React from 'react';
import MyWork from '@/pages/creativeCenter/work';

export default () => {
  return (
    <Tabs tabPosition='left' size='large'>
      <Tabs.TabPane tab='作品' key='work'>
        <MyWork/>
      </Tabs.TabPane>
      <Tabs.TabPane tab='草稿' key='draft'>

      </Tabs.TabPane>
      <Tabs.TabPane tab='专题' key='topic'>

      </Tabs.TabPane>
      <Tabs.TabPane tab='收藏夹' key='college'>

      </Tabs.TabPane>
    </Tabs>);
}
