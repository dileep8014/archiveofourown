import React, { useEffect, useMemo, useState } from 'react';
import {
  Button,
  Col, Collapse,
  Divider,
  Dropdown,
  Layout,
  Menu, message,
  Row,
  Space, Spin,
  Tooltip,
  Typography,
} from 'antd';
import styles from './index.less';
import {
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  CaretDownOutlined,
  QuestionOutlined,
  FileAddOutlined,
  SendOutlined,
  RestOutlined,
  InboxOutlined,
  RollbackOutlined,
  ArrowRightOutlined,
  HistoryOutlined,
} from '@ant-design/icons';
import IconFont from '@/component/iconfont/iconfont';
import numberParseChina from '@/component/numconvert/numconvert';
import SubsectionManager from '@/pages/work/write/subsectionManager';
import Editor from '@/pages/work/write/editor';
import { ChapterItem, SubsectionItem } from '@/pages/work/write/sider/menuItem';
import DraftMenu from '@/pages/work/write/sider/draft';
import { useDispatch, useSelector } from '@@/plugin-dva/exports';
import SubsectionMenu from '@/pages/work/write/sider/subsection';
import RecycleMenu from '@/pages/work/write/sider/recycle';
import ErrorPage from '@/component/errorpage/errorpage';
import { createEditor, Node } from 'slate';
import { useInterval, useMount, useUnmount } from 'ahooks';
import { Prompt } from 'react-router-dom';
import { withHistory } from 'slate-history';
import { withReact } from 'slate-react';

const { Header, Content, Sider } = Layout;
const { Text } = Typography;

export interface WriteState {
  title: string,
  value: Node[]
}

export const initValue = [{ children: [{ text: '输入正文...' }] }];

export default function() {
  document.title = '写作';

  const dispatch = useDispatch();
  const { data, subsection, error, saveLoading, draft } = useSelector((state: any) => {
    return {
      data: state.work.data,
      draft: state.work.draft,
      subsection: state.work.subsection,
      error: state.work.error,
      saveLoading: state.loading.effects['work/save'],
    };
  });

  const editor = useMemo(() => withHistory(withReact(createEditor())), []);

  const [draftChapter, setDraftChapter] = useState<ChapterItem>({
    words: 0,
    id: -1,
    updatedAt: null,
    title: '',
    type: 'draft',
    subsection: 1,
  });
  const [collapsed, setCollapsed] = useState<boolean>(false);
  const [subManage, setSubManage] = useState<boolean>(false);
  const [canUpdate, setCanUpdate] = useState<boolean>(false);
  const [newDraft, setNewDraft] = useState<boolean>(false);
  const [write, setWrite] = useState<WriteState>({ title: '', value: initValue });

  const [currentChapter, setCurrentChapter] = useState<ChapterItem | null>(null);

  const listener = (e: { preventDefault: () => void; returnValue: string; }) => {
    e.preventDefault();
    e.returnValue = '离开当前页后，所编辑的数据将不可恢复';
    // localStorage.removeItem('draft');
  };

  useEffect(() => {
    if (data) {
      if (data.draft == 0) {
        setNewDraft(true);
        if (!currentChapter) {
          setCurrentChapter(draftChapter);
        }
      } else if (!currentChapter) {
        setCurrentChapter(draft[draft.length - 1]);
      }
    }
  }, [data, draft]);

  useEffect(() => {
    if (subsection.length > 0) {
      setDraftChapter({ ...draftChapter, subsection: subsection[subsection.length - 1].seq });
    }
  }, [subsection]);

  useInterval(() => {
    if (write.title == '' || write.value == initValue || nodesConvert(write.value) == '') {
      return;
    }
    save();
    if (currentChapter?.id == -1) {
      setNewDraft(false);
      setCurrentChapter(null);
    }
  }, 30000);

  useMount(() => window.addEventListener('beforeunload', listener));

  useUnmount(() => {
    window.removeEventListener('beforeunload', listener);
    localStorage.removeItem('draft');
  });

  const save = () => {
    if (currentChapter?.type != 'draft' && !canUpdate) {
      return;
    }
    dispatch({
      type: 'work/save', payload: {
        id: currentChapter?.id,
        title: write.title,
        type: currentChapter?.type,
        subsection: currentChapter?.subsection,
        content: nodesConvert(write.value),
      },
    });
  };

  const onNewChapter = () => {
    if (newDraft) {
      if (data.draft > 0 && currentChapter?.id != -1) {
        setCurrentChapter(draftChapter);
        return;
      }
      if (write.title == '' || write.value == initValue || nodesConvert(write.value) == '') {
        message.error('请先完成当前草稿');
        return;
      }
      save();
      setWrite({ title: '', value: initValue });
    } else {
      setNewDraft(true);
      setCurrentChapter(draftChapter);
    }
  };

  const onChangeItem = (item: ChapterItem) => {
    if (currentChapter?.id != item.id) {
      setCurrentChapter(item);
      if (write.title == '' || write.value == initValue || nodesConvert(write.value) == '') {
        if (currentChapter?.id == -1) {
          localStorage.setItem('draft', JSON.stringify(write));
        }
        return;
      }
      save();
      if (currentChapter?.id == -1) {
        setNewDraft(false);
      }
    }
  };

  const onSave = () => {
    if (write.title == '' || write.value == initValue || nodesConvert(write.value) == '') {
      if (currentChapter?.id == -1) {
        message.error('请先完成当前草稿');
      } else {
        message.error('请先完成当前章节');
      }
      return;
    }
    save();
    if (currentChapter?.id == -1) {
      setNewDraft(false);
      setCurrentChapter(null);
    }
  };

  const onDelete = () => {
    if (currentChapter?.id == -1) {
      localStorage.removeItem('draft');
      if (data.draft == 0) {
        setCurrentChapter(draftChapter);
        return;
      }
      setNewDraft(false);
      setCurrentChapter(draft[draft.length - 1]);
      return;
    }
    dispatch({
      type: 'work/delete',
      payload: { id: currentChapter?.id, type: currentChapter?.type, subsection: currentChapter?.subsection },
    });
  };

  const onPublish = () => {
    if (currentChapter?.id == -1) {
      message.warning('请先保存当前草稿');
      return;
    }
    dispatch({
      type: 'work/publish',
      payload: { id: currentChapter?.id, type: currentChapter?.type, subsection: currentChapter?.subsection },
    });
    setCurrentChapter(null);
  };

  if (error) {
    return <ErrorPage title={error.name} subTitle={error.message}/>;
  }

  return (
    <div>
      <Prompt
        when={true}
        message={() => '离开当前页后，所编辑的数据将不可恢复'}
      />
      {data &&
      <Layout className={styles.writeLayout}>
        <Sider theme={'light'}
               width={264}
               className={styles.writeSider}
               collapsed={collapsed}
               collapsedWidth={0}
        >
          <Space className={styles.writeSpace}>
            <img alt='封面' className={styles.writeLogo} src={data.cover}/>
            <Divider type={'vertical'}/>
            <Text ellipsis className={styles.writeTitle}>{data.title}</Text>
            <Tooltip title='分卷管理'>
              <Button onClick={() => setSubManage(true)}
                      style={{ border: 0, boxShadow: 'none' }}
                      icon={<IconFont type='icon-anjuanguanli'/>}/>
            </Tooltip>
            <SubsectionManager subManage={subManage} setSubManage={setSubManage}/>
          </Space>
          <Button style={{ border: 0, boxShadow: 'none', width: '100%', height: 40, textAlign: 'left' }}
                  icon={<FileAddOutlined className={styles.writeIcon}/>}
                  onClick={onNewChapter}
          >
            新建章节
          </Button>
          <Collapse defaultActiveKey={[currentChapter?.type || 'draft']}
                    expandIconPosition={'right'} bordered={false} ghost>
            <Collapse.Panel key='draft'
                            header={<><InboxOutlined className={styles.writeIcon}/>草稿箱
                              <Text className={styles.smallText}>(共{data.draft}章)</Text></>}
            >
              <DraftMenu draftChapter={draftChapter} currentItem={currentChapter}
                         newDraft={newDraft} setCurrentItem={onChangeItem}/>
            </Collapse.Panel>
            <Collapse.Panel key='published'
                            header={<><SendOutlined className={styles.writeIcon}/>已发布
                              <Text className={styles.smallText}>(共{data.subsection}卷 {data.chapters}章)</Text></>}>
              <SubsectionMenu currentItem={currentChapter} setCurrentItem={onChangeItem}/>
            </Collapse.Panel>
            <Collapse.Panel key='recycle'
                            header={<><RestOutlined className={styles.writeIcon}/>回收站
                              <Text className={styles.smallText}>(共{data.recycle}章)</Text></>}>
              <RecycleMenu/>
            </Collapse.Panel>
          </Collapse>
        </Sider>
        <Layout>
          <Header className={styles.writeHeader}>
            <Row className={styles.writeRowCol}>
              <Col xxl={20} xl={20} lg={18} md={16} sm={12} xs={0} className={styles.writeRowCol}>
                <Space>
                  <Tooltip title={(collapsed && '退出沉浸模式') || '沉浸模式'}>
                    <Button style={{ border: 0, boxShadow: 'none' }}
                            icon={
                              (collapsed && <MenuUnfoldOutlined style={{ fontSize: 18 }}/>) ||
                              <MenuFoldOutlined style={{ fontSize: 18 }}/>}
                            onClick={() => setCollapsed(!collapsed)}
                    />
                  </Tooltip>
                  <Dropdown trigger={['click']} disabled={currentChapter?.type == 'published'}
                            overlay={<Menu>
                              {subsection.map((item: SubsectionItem) =>
                                <Menu.Item onClick={() => setCurrentChapter(currentChapter && {
                                  ...currentChapter,
                                  subsection: item.seq,
                                })}
                                           key={item.id}>第{numberParseChina(item.seq)}卷</Menu.Item>)}
                            </Menu>}>
                    <Button style={{ border: 0, boxShadow: 'none' }}>
                      第{numberParseChina(currentChapter?.subsection || 0)}卷
                      <CaretDownOutlined/>
                    </Button>
                  </Dropdown>
                </Space>
              </Col>
              <Col xxl={4} xl={4} lg={6} md={8} sm={12} xs={24} className={styles.writeRowCol}>
                <Space>
                  <Tooltip title='帮助'>
                    <Button shape={'circle'} icon={<QuestionOutlined/>} size={'small'}/>
                  </Tooltip>
                  {currentChapter?.type == 'draft' && <>
                    <Button shape={'round'} type={'primary'} size={'small'} onClick={onDelete}>删除</Button>
                    <Spin spinning={saveLoading == undefined ? false : saveLoading}>
                      <Button shape={'round'} size={'small'} onClick={onSave}>保存</Button>
                    </Spin>
                    <Spin spinning={saveLoading == undefined ? false : saveLoading}>
                      <Button shape={'round'} size={'small'} onClick={onPublish}
                              style={{ backgroundColor: '#0067E6', color: 'white' }}>
                        发布
                      </Button>
                    </Spin></>}
                  {currentChapter?.type != 'draft' && <>
                    <Button shape={'round'} type={'primary'} size={'small'} onClick={onDelete}>删除</Button>
                    <Spin spinning={saveLoading == undefined ? false : saveLoading}>
                      <Button shape={'round'} size={'small'} onClick={() => {
                        if (canUpdate) {
                          onSave();
                          setCanUpdate(false);
                        } else {
                          setCanUpdate(true);
                        }
                      }}>
                        {canUpdate && '保存' || '修改'}
                      </Button>
                    </Spin>
                    {currentChapter?.type == 'recycle' &&
                    <Spin spinning={saveLoading == undefined ? false : saveLoading}>
                      <Button shape={'round'} size={'small'} onClick={onPublish}
                              style={{ backgroundColor: '#0067E6', color: 'white' }}>
                        重新发布
                      </Button>
                    </Spin>}</>}
                </Space>
              </Col>
            </Row>
          </Header>
          <Content className={styles.writeContent}>
            {(currentChapter?.type == 'draft' || canUpdate) &&
            <div className={styles.writeOpGroup}>
              <Tooltip title='回退'>
                <Button icon={<RollbackOutlined/>} className={styles.opBtn} onClick={editor.undo}/>
              </Tooltip>
              <Tooltip title='前进'>
                <Button icon={<ArrowRightOutlined/>} className={styles.opBtn} onClick={editor.redo}/>
              </Tooltip>
              <Tooltip title='历史版本'>
                <Button icon={<HistoryOutlined/>} className={styles.opBtn}/>
              </Tooltip>
            </div>
            }
            <Editor editor={editor}
                    key={currentChapter?.id}
                    currentItem={currentChapter}
                    write={write}
                    disable={currentChapter?.type != 'draft' && !canUpdate}
                    setWrite={setWrite}/>
          </Content>
        </Layout>
      </Layout>
      }
    </div>
  );
}

export const nodesConvert = (nodes: Node[]) => {
  let content = '';
  nodes.forEach((item, index) => {
    // @ts-ignore
    content += item.children[0].text;
    if (index != nodes.length - 1) {
      content += '\n';
    }
  });
  return content;
};
