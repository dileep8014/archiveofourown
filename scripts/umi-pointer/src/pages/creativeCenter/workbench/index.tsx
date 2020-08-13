import React from 'react';
import { Avatar, Button, Calendar, Card, List, Skeleton, Space } from 'antd';
import { useModel } from '@@/plugin-model/useModel';
import ProCard from '@ant-design/pro-card';
import Text from 'antd/es/typography/Text';
import { useRequest } from '@umijs/hooks';
import service from '@/component/service';
import WorkItem from '@/pages/creativeCenter/workItem/workItem';
import locale from 'antd/es/date-picker/locale/zh_CN';
import moment from 'moment';


export default function WorkBench() {

  const { user, userLoading } = useModel('user', model => ({ user: model.user, userLoading: model.loading }));

  const { data, loading } = useRequest(service.MineWorks, { defaultParams: [{ current: 1, pageSize: 1 }] });
  const { data: calendarData, loading: calendarLoading, run } = useRequest(service.Calendar, {
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
                 extra={data && data.list?.length > 0 && <Button type={'primary'}>新建作品</Button>}
        >
          {data && data.list.length > 0 && <WorkItem {...data.list?.[0]}/>}
          {data && data.list.length == 0 && <Button type={'primary'}>新建作品</Button>}
        </ProCard>
      </Skeleton>
      <ProCard title='创作日历（红色为作品发布日）' style={{ marginTop: 20 }} headerBordered bordered>

        <Calendar fullscreen={false} locale={locale}
                  dateFullCellRender={value => {
                    return (
                      <Text type={calendarData?.list[value.month()].day[value.date() - 1] && 'danger'}>
                        {value.date()}
                      </Text>);
                  }}
                  monthFullCellRender={value => {
                    return (
                      <Text type={calendarData?.list[value.month()].work && 'danger'}>
                        {value.date()}
                      </Text>);
                  }}
                  onChange={value => {
                    run({ year: value.year() });
                  }}
        />
      </ProCard>
    </div>);
}
