import styles from './index.less';
import { Tabs } from 'antd';
import React from 'react';
import MyWork from '@/pages/creativeCenter/work';
import WorkBench from '@/pages/creativeCenter/workbench';
import TopicList from '@/pages/creativeCenter/topic';

export default () => {
  return (
    <Tabs tabPosition='left' size='large' style={{ paddingLeft: 100, paddingRight: 100 }}>
      <Tabs.TabPane tab='工作台' key='all'>
        <WorkBench/>
      </Tabs.TabPane>
      <Tabs.TabPane tab='作品' key='work'>
        <MyWork/>
      </Tabs.TabPane>
      <Tabs.TabPane tab='专题' key='topic'>
        <TopicList/>
      </Tabs.TabPane>
    </Tabs>);
}
