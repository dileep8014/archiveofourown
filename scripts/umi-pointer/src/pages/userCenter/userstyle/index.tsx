import { Button, Card, Checkbox, List, message } from 'antd';
import React, { useEffect, useState } from 'react';
import { useRequest } from '@umijs/hooks';
import service from '@/component/service';
import { CheckboxChangeEvent } from 'antd/es/checkbox';

export type StyleState = {
  email: boolean,
  birthday: boolean,
  searchAll: boolean,
  shareWork: boolean,
  adult: boolean,
  grade: boolean,
  tag: boolean,
  subMsg: boolean,
  topicMsg: boolean,
  commentMsg: boolean,
  sysMsg: boolean,
} | null

export default function UserStyle() {

  const [styles, setStyles] = useState<StyleState>(null);
  const { data, loading, refresh } = useRequest(service.GetStyles);
  const { run } = useRequest(service.UpdateStyles, {
    manual: true, onSuccess: () => {
      refresh();
      message.success('修改成功');
    },
  });

  useEffect(() => {
    if (data) {
      setStyles(data);
    }
  }, [data]);

  const onChange = (e: CheckboxChangeEvent, field: string) => {
    if (styles) {
      let tmp = { ...styles };
      Object.defineProperty(tmp, field, {
        value: e.target.checked,    // default undefined
      });
      setStyles(tmp);
    }
  };

  return (
    <List itemLayout={'vertical'} split={false} loading={loading}
          footer={<Button onClick={() => run(styles)} style={{ float: 'right' }} type={'primary'}>更新偏好</Button>}>
      <List.Item>
        <Card title='隐私'>
          <Checkbox checked={styles?.email} onChange={e => onChange(e, 'email')}>向他人展示邮件</Checkbox>
          <br/>
          <Checkbox checked={styles?.birthday} onChange={e => onChange(e, 'birthday')}>向他人展示生日</Checkbox>
          <br/>
          <Checkbox checked={styles?.searchAll} onChange={e => onChange(e, 'searchAll')}>禁止搜索我的一切</Checkbox>
          <br/>
          <Checkbox checked={styles?.shareWork} onChange={e => onChange(e, 'shareWork')}>禁止分享我的作品</Checkbox>
        </Card>
      </List.Item>
      <List.Item>
        <Card title='显示'>
          <Checkbox checked={styles?.adult} onChange={e => onChange(e, 'adult')}>允许向我显示成年内容</Checkbox>
          <br/>
          <Checkbox checked={styles?.grade} onChange={e => onChange(e, 'grade')}>隐藏作品分级信息</Checkbox>
          <br/>
          <Checkbox checked={styles?.tag} onChange={e => onChange(e, 'tag')}>隐藏作品标签信息</Checkbox>
        </Card>
      </List.Item>
      <List.Item>
        <Card title='提醒'>
          <Checkbox checked={styles?.subMsg && styles.topicMsg && styles.sysMsg && styles.commentMsg}
                    onChange={e => setStyles(styles && {
                      ...styles,
                      topicMsg: e.target.checked,
                      sysMsg: e.target.checked,
                      commentMsg: e.target.checked,
                      subMsg: e.target.checked,
                    })}>
            打开所有邮件提醒
          </Checkbox>
          <br/>
          <Checkbox checked={styles?.subMsg} onChange={e => onChange(e, 'subMsg')}>打开订阅邮件提醒</Checkbox>
          <br/>
          <Checkbox checked={styles?.topicMsg} onChange={e => onChange(e, 'topicMsg')}>打开专题邮件提醒</Checkbox>
          <br/>
          <Checkbox checked={styles?.commentMsg} onChange={e => onChange(e, 'commentMsg')}>打开评论邮件提醒</Checkbox>
          <br/>
          <Checkbox checked={styles?.sysMsg} onChange={e => onChange(e, 'sysMsg')}>打开系统消息邮件提醒</Checkbox>
        </Card>
      </List.Item>
    </List>
  );
}
