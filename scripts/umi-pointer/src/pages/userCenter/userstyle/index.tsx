import { Button, Card, Checkbox, List, message, Skeleton } from 'antd';
import React, { useEffect, useState } from 'react';
import { useRequest } from '@umijs/hooks';
import { CheckboxChangeEvent } from 'antd/es/checkbox';
import ProgressOpt from '@/component/progress/progress';
import { userService } from '@/service/user';

export type UserSetting = {
  showEmail: boolean
  disableSearch: boolean
  showAdult: boolean
  hiddenGrade: boolean
  hiddenTag: boolean
  subscriptionEmail: boolean
  topicEmail: boolean
  commentEmail: boolean
  systemEmail: boolean
} | null

export default function UserStyle() {

  const [setting, updateSetting] = useState<UserSetting>(null);
  const { data, loading, refresh } = useRequest(userService.CurrentUserSetting);
  const { run } = useRequest(userService.UpdateCurrentSetting, {
    manual: true, onSuccess: () => {
      ProgressOpt(refresh);
      message.success('修改成功');
    },
  });

  useEffect(() => {
    if (data && data.code == 0) {
      updateSetting(data.data);
    }
  }, [data]);

  const onChange = (e: CheckboxChangeEvent, field: string) => {
    if (setting) {
      let tmp = { ...setting };
      Object.defineProperty(tmp, field, {
        value: e.target.checked,    // default undefined
      });
      updateSetting(tmp);
    }
  };

  return (
    <Skeleton loading={loading}>
      <List itemLayout={'vertical'} split={false}
            footer={<Button onClick={() => run(setting)} style={{ float: 'right' }} type={'primary'}>更新偏好</Button>}>
        <List.Item>
          <Card title='隐私'>
            <Checkbox checked={setting?.showEmail} onChange={e => onChange(e, 'showEmail')}>向他人展示邮件</Checkbox>
            <br/>
            <Checkbox checked={setting?.disableSearch} onChange={e => onChange(e, 'disableSearch')}>禁止搜索我的一切</Checkbox>
          </Card>
        </List.Item>
        <List.Item>
          <Card title='显示'>
            <Checkbox checked={setting?.showAdult} onChange={e => onChange(e, 'showAdult')}>允许向我显示成年内容</Checkbox>
            <br/>
            <Checkbox checked={setting?.hiddenGrade} onChange={e => onChange(e, 'hiddenGrade')}>隐藏作品分级信息</Checkbox>
            <br/>
            <Checkbox checked={setting?.hiddenTag} onChange={e => onChange(e, 'hiddenTag')}>隐藏作品标签信息</Checkbox>
          </Card>
        </List.Item>
        <List.Item>
          <Card title='提醒'>
            <Checkbox
              checked={setting?.subscriptionEmail && setting.topicEmail && setting.commentEmail && setting.systemEmail}
              onChange={e => updateSetting(setting && {
                ...setting,
                subscriptionEmail: e.target.checked,
                topicEmail: e.target.checked,
                commentEmail: e.target.checked,
                systemEmail: e.target.checked,
              })}>
              打开所有邮件提醒
            </Checkbox>
            <br/>
            <Checkbox checked={setting?.subscriptionEmail}
                      onChange={e => onChange(e, 'subscriptionEmail')}>打开订阅邮件提醒</Checkbox>
            <br/>
            <Checkbox checked={setting?.topicEmail} onChange={e => onChange(e, 'topicEmail')}>打开专题邮件提醒</Checkbox>
            <br/>
            <Checkbox checked={setting?.commentEmail} onChange={e => onChange(e, 'commentEmail')}>打开评论邮件提醒</Checkbox>
            <br/>
            <Checkbox checked={setting?.systemEmail} onChange={e => onChange(e, 'systemEmail')}>打开系统消息邮件提醒</Checkbox>
          </Card>
        </List.Item>
      </List>
    </Skeleton>
  );
}
