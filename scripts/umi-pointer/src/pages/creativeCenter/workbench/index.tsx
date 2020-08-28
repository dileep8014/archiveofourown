import React from 'react';
import { Avatar, Button, Calendar, List, Skeleton, Space } from 'antd';
import { useModel, history } from 'umi';
import ProCard from '@ant-design/pro-card';
import Text from 'antd/es/typography/Text';
import { useRequest } from '@umijs/hooks';
import WorkItem from '@/pages/creativeCenter/workItem/workItem';
import moment from 'moment';
import { workService } from '@/service/work';


export default function WorkBench(props: { onCreate: () => void }) {

  const { user, userLoading } = useModel('user', model => ({ user: model.user, userLoading: model.loading }));

  const { data, loading } = useRequest(workService.MineWorks, { defaultParams: [{ current: 1, pageSize: 1 }] });
  const { data: calendarData, run } = useRequest(workService.Calendar, {
    defaultParams: [{
      year: moment(moment.now()).year(),
    }],
  });


  return (
    <div>
      <ProCard bordered title='个人信息' loading={userLoading} headerBordered>
        <ProCard colSpan={12}>
          <Space>
            <Avatar style={{ width: 100, height: 100 }} src={user?.avatar}/>
            <h1>{user?.name}</h1>
          </Space>
        </ProCard>
        <ProCard colSpan={12}>
          <List>
            <List.Item>
              <span>{user?.workDay}</span>
              <Text>创作天数</Text>
            </List.Item>
            <List.Item>
              <span>{user?.words}</span>
              <Text>累计字数</Text>
            </List.Item>
            <List.Item>
              <span>{user?.fans}</span>
              <Text>粉丝</Text>
            </List.Item>
          </List>
        </ProCard>
      </ProCard>
      <Skeleton loading={loading}>
        <ProCard title='最近作品' style={{ marginTop: 20 }} headerBordered bordered loading={loading}
                 extra={data && data.list?.length > 0 &&
                 <Button onClick={props.onCreate} type={'primary'}>新建作品</Button>}
        >
          {data && data.list.length > 0 && <WorkItem {...data.list?.[0]}/>}
          {data && data.list.length == 0 &&
          <Button onClick={props.onCreate} type={'primary'}>新建作品</Button>}
        </ProCard>
      </Skeleton>
      <ProCard title='创作日历' style={{ marginTop: 20 }} headerBordered bordered>

        <Calendar fullscreen={false}
                  dateFullCellRender={value => {
                    const able = calendarData?.list[value.month()].day[value.date() - 1];
                    return (
                      <div style={{ backgroundColor: (able && '#ffddda') || '#fff' }}>
                        <Text disabled={!able}>
                          {value.date()}
                        </Text>
                      </div>);
                  }}
                  monthFullCellRender={value => {
                    const able = calendarData?.list[value.month()].work;
                    return (
                      <div style={{ backgroundColor: (able && '#ffddda') || '#fff' }}>
                        <Text disabled={!able}>
                          {value.date()}
                        </Text>
                      </div>);
                  }}
                  onChange={value => {
                    run({ year: value.year() });
                  }}
        />
      </ProCard>
    </div>);
}
