import { Button, Card, Form, Input, Menu, message, Modal, Space } from 'antd';
import styles from '@/pages/work/write/subsectionManager/index.less';
import React, { useState } from 'react';
import numberParseChina from '@/component/numconvert/numconvert';
import Text from 'antd/es/typography/Text';
import { FileAddOutlined } from '@ant-design/icons';
import { useRequest } from '@umijs/hooks';
import { SubsectionItem } from '@/pages/work/write/sider/menuItem';
import { workService } from '@/service/work';
import { useDispatch, useSelector } from '@@/plugin-dva/exports';
import { history } from 'umi';

export default function SubsectionManager(props: {
  subManage: boolean, setSubManage: (arg: boolean) => void
}) {

  const { subManage, setSubManage } = props;
  const dispatch = useDispatch();
  const { subsection } = useSelector((state: any) => {
    return { subsection: state.work.subsection };
  });

  const { run: updateSubsection } = useRequest(workService.UpdateSubsection, {
    manual: true, onSuccess: () => {
      if (currentSubsection?.id == -1) {
        setNewSub(null);
      }
      dispatch({ type: 'work/query', payload: { id: parseInt(history.location.pathname.split('/')[4]) } });
    },
  });

  const { run: deleteSubsection } = useRequest(workService.DeleteSubsection, {
    manual: true, onSuccess: () => {
      dispatch({ type: 'work/query', payload: { id: parseInt(history.location.pathname.split('/')[4]) } });
    },
  });

  const [currentSubsection, setCurrentSubsection] = useState<SubsectionItem | null>(subsection[subsection.length - 1]);
  const [newSub, setNewSub] = useState<SubsectionItem | null>(null);
  const [form] = Form.useForm();

  const onNewSub = () => {
    let last = subsection[subsection.length - 1];
    if (last.chapters == 0) {
      message.warning('分卷' + numberParseChina(last.seq) + '尚无章节，无法新建分卷');
      return;
    }
    let next = { id: -1, seq: last.seq + 1, name: '', introduce: '', chapters: [] };
    setNewSub(next);
    setCurrentSubsection(next);
  };

  const onSubCancel = () => {
    setSubManage(false);
    if (newSub) {
      setNewSub(null);
      setCurrentSubsection(subsection[subsection.length - 1]);
    }
  };

  const onDelete = () => {
    if (currentSubsection) {
      if (currentSubsection?.id == -1) {
        setNewSub(null);
        setCurrentSubsection(subsection[0]);
        return;
      }
      if (currentSubsection.seq == 1) {
        message.error('第一卷无法删除');
        return;
      }
      if (currentSubsection.chapters.length != 0) {
        message.error('请先删除卷内章节');
        return;
      }
      deleteSubsection({ id: currentSubsection?.id });
    }
  };

  return (<Modal visible={subManage}
                 title='分卷管理'
                 footer={null}
                 centered
                 bodyStyle={{ height: 400 }}
                 onCancel={onSubCancel}>
    <Space style={{ width: '100%', height: '100%' }}>
      <div style={{ display: 'grid' }}>
        <Menu className={styles.writeModalMenu}
              selectedKeys={[currentSubsection?.id + '']}
              onSelect={({ key }) => {
                if (key == '-1') {
                  setCurrentSubsection(newSub);
                  return;
                }
                subsection.forEach((item: SubsectionItem) => {
                  if (item.id == parseInt(key.toString())) {
                    setCurrentSubsection(item);
                  }
                });
              }}
              mode='inline'>
          {subsection.map((item: SubsectionItem) =>
            <Menu.Item className={styles.writeModalMenuItem} key={item.id} mode='vertical-left'>
              <Card size={'small'} bordered={false} style={{ backgroundColor: 'inherit' }}>
                <Card.Meta
                  title={<Text style={{ fontWeight: 8, fontSize: 13 }}>第{numberParseChina(item.seq)}卷</Text>}
                  description={<Text className={styles.smallText}>本卷共{item.chapters.length}章</Text>}
                />
              </Card>
            </Menu.Item>)
          }
          <Menu.Item hidden={newSub == null} className={styles.writeModalMenuItem}
                     key={newSub?.id + ''}
                     mode='vertical-left'>
            <Card size={'small'} bordered={false} style={{ backgroundColor: 'inherit' }}>
              <Card.Meta
                title={
                  <Text style={{ fontWeight: 8, fontSize: 13 }}>第{numberParseChina(newSub?.seq || 0)}卷</Text>}
                description={<Text className={styles.smallText}>新分卷</Text>}
              />
            </Card>
          </Menu.Item>
        </Menu>
        <Button style={{ border: 0, boxShadow: 'none', color: '#0067E6', marginTop: 20 }}
                icon={<FileAddOutlined/>}
                onClick={onNewSub}
        >
          新建分卷
        </Button>
      </div>
      <Card bordered={false} style={{ height: 300, width: 350 }}>
        <Card.Meta title={'第' + numberParseChina(currentSubsection?.seq || 0) + '卷'}
                   description={<Form form={form}>
                     <Form.Item label='分卷名称'>
                       <Input value={currentSubsection?.name} placeholder='非必填' onChangeCapture={event => {
                         setCurrentSubsection(currentSubsection && {
                           ...currentSubsection,
                           name: event.currentTarget.value,
                         });
                       }}/>
                     </Form.Item>
                     <Form.Item label='分卷简介'>
                       <Input.TextArea value={currentSubsection?.introduce}
                                       autoSize={{ maxRows: 6, minRows: 6 }}
                                       placeholder='非必填'
                                       onChangeCapture={event => {
                                         setCurrentSubsection(currentSubsection && {
                                           ...currentSubsection,
                                           introduce: event.currentTarget.value,
                                         });
                                       }}
                       />
                     </Form.Item>
                     <Form.Item style={{ float: 'right' }}>
                       {currentSubsection?.id != -1 &&
                       <Button style={{ marginRight: 10 }} onClick={onDelete}>删除分卷</Button>}
                       {currentSubsection?.id == -1 &&
                       <Button onClick={onSubCancel} style={{ marginRight: 10 }}>取消</Button>}
                       <Button type={'primary'} onClick={() => {
                         if (currentSubsection) {
                           updateSubsection(currentSubsection);
                         }
                       }}>保存</Button>
                     </Form.Item>
                   </Form>}
        />
      </Card>
    </Space>
  </Modal>);
}
