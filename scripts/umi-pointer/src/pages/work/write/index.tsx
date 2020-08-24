import React, { useState, useMemo, useEffect, useRef } from 'react';
import { Node, createEditor } from 'slate';
import { Slate, Editable, withReact } from 'slate-react';
import { withHistory } from 'slate-history';
import {
  Button,
  Card,
  Col,
  Divider,
  Dropdown, Form, Input,
  Layout,
  Menu,
  Modal,
  Row,
  Skeleton,
  Space,
  Tooltip,
  Typography,
} from 'antd';
import { history } from 'umi';
import ErrorPage from '@/component/errorpage/errorpage';
import { useRequest } from '@umijs/hooks';
import service from '@/component/service';
import styles from './index.less';
import {
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  CaretDownOutlined,
  QuestionOutlined,
  FileAddOutlined,
  SendOutlined,
  RestOutlined,
} from '@ant-design/icons';
import IconFont from '@/component/iconfont/iconfont';
import moment from 'moment';

const { Header, Content, Sider } = Layout;
const { Text } = Typography;

const initialValue = [
  {
    children: [
      { text: '' },
    ],
  },
];

interface ChapterItem {
  words: number;
  id: number;
  updatedAt: any;
  title: string;
}

interface SubItem {
  id: number;
  seq: number,
  name: string,
  introduce: string,
  chapters: ChapterItem[],
}

const ChapterMenuItem = (item: { words: number; updatedAt: any; title: string }) =>
  <Card size={'small'} bordered={false} style={{ backgroundColor: 'inherit' }}>
    <Card.Meta title={<Text style={{ fontWeight: 8, fontSize: 13 }}>{item.title}</Text>}
               description={<Text className={styles.smallText}>
                 {item.updatedAt && (moment(item.updatedAt).format('M-D H:M:S') + '  ')}
                 {item.words}字
               </Text>}
    />
  </Card>;

export default function() {
  document.title = '写作';

  const pathname = history.location.pathname.split('/');
  const workId = pathname.length == 4 ? pathname[3] : undefined;

  const { data, loading, error, run, refresh } = useRequest(service.WorkInfo, { manual: true });

  const [collapsed, setCollapsed] = useState<boolean>(false);
  const [subManage, setSubManage] = useState<boolean>(false);
  const [subSeq, setSubSeq] = useState<SubItem | null>(null);
  const [currentSub, setCurrentSub] = useState<SubItem | null>(null);
  const [draftChapter, setDraftChapter] = useState<{ hide: boolean, words: number }>({ hide: true, words: 0 });
  const [value, setValue] = useState<Node[]>(initialValue);
  const [title, setTitle] = useState<Node[]>(initialValue);
  const editor = useMemo(() => withHistory(withReact(createEditor())), []);
  const titleEditor = useMemo(() => withHistory(withReact(createEditor())), []);
  const [form] = Form.useForm();

  useEffect(() => {
    if (workId) {
      run({ id: workId });
    }
  }, [workId]);

  useEffect(() => {
    if (data) {
      if ((!data.chapters.draft || data.chapters.draft.length == 0) && draftChapter.hide) {
        setDraftChapter({ hide: false, words: 0 });
      }
      setSubSeq(data.chapters.published.subsection[data.chapters.published.subsection.length - 1]);
      setCurrentSub(data.chapters.published.subsection[data.chapters.published.subsection.length - 1]);
    }
  }, [data]);

  if (workId == undefined) {
    return <ErrorPage title='错误的作品' subTitle='您要编辑的作品不存在'/>;
  }

  if (error) {
    return <ErrorPage title={error.name} subTitle={error.message}/>;
  }

  return (
    <Skeleton loading={loading}>
      {data &&
      <Layout className={styles.writeLayout}>
        <Sider theme={'light'}
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
            <Modal visible={subManage}
                   title='分卷管理'
                   footer={null}
                   centered
                   bodyStyle={{ height: 400 }}
                   onCancel={() => setSubManage(false)}>
              <Space style={{ width: '100%', height: '100%' }}>
                <div style={{ display: 'grid' }}>
                  <Menu className={styles.writeModalMenu}
                        selectedKeys={[subSeq?.id + '']}
                        onSelect={({ key }) => {
                          data.chapters.published.subsection.forEach((item: SubItem) => {
                            if (item.id == parseInt(key.toString())) {
                              setSubSeq(item);
                            }
                          });
                        }}
                        mode='inline'>
                    {data.chapters.published.subsection.map((item: SubItem) =>
                      <Menu.Item className={styles.writeModalMenuItem} key={item.id} mode='vertical-left'>
                        <Card size={'small'} bordered={false} style={{ backgroundColor: 'inherit' }}>
                          <Card.Meta title={<Text style={{ fontWeight: 8, fontSize: 13 }}>第{item.seq}卷</Text>}
                                     description={<Text className={styles.smallText}>本卷共{item.chapters.length}章</Text>}
                          />
                        </Card>
                      </Menu.Item>)
                    }
                  </Menu>
                  <Button style={{ border: 0, boxShadow: 'none', color: '#0067E6', marginTop: 20 }}
                          icon={<FileAddOutlined/>}>
                    新建分卷
                  </Button>
                </div>
                <Card bordered={false} style={{ height: 300, width: 350 }}>
                  <Card.Meta title={'第' + subSeq?.seq + '卷'}
                             description={<Form form={form}>
                               <Form.Item label='分卷名称'>
                                 <Input value={subSeq?.name} placeholder='非必填' onChangeCapture={event => {
                                   setSubSeq(subSeq && { ...subSeq, name: event.currentTarget.value });
                                 }}/>
                               </Form.Item>
                               <Form.Item label='分卷简介'>
                                 <Input.TextArea value={subSeq?.introduce}
                                                 autoSize={{ maxRows: 6, minRows: 6 }}
                                                 placeholder='非必填'
                                                 onChangeCapture={event => {
                                                   setSubSeq(subSeq && {
                                                     ...subSeq,
                                                     introduce: event.currentTarget.value,
                                                   });
                                                 }}
                                 />
                               </Form.Item>
                               <Form.Item style={{ float: 'right' }}>
                                 <Button style={{ marginRight: 10 }}>删除分卷</Button>
                                 <Button type={'primary'}>保存</Button>
                               </Form.Item>
                             </Form>}
                  />
                </Card>
              </Space>
            </Modal>
          </Space>
          <Menu className={styles.writeMenu} mode='inline'>
            <Menu.Item className={styles.writeMenuItem} key='newChapter'>
              <FileAddOutlined className={styles.writeIcon}/>新建章节
            </Menu.Item>
            <Menu.SubMenu className={styles.writeMenuItem}
                          key='draft'
                          icon={<IconFont type='icon-caogaoxiang' className={styles.writeIcon}/>}
                          title={<>草稿箱
                            <Text className={styles.smallText}>
                              (共{(data.chapters.draft && data.chapters.draft.length) || 0}章)
                            </Text></>}
            >
              {!draftChapter.hide &&
              <Menu.Item className={styles.writeMenuItem} key='draftChapter' mode='vertical-left'>
                <ChapterMenuItem words={draftChapter.words} updatedAt={null} title={'未命名'}/>
              </Menu.Item>}
              {data.chapters.draft && data.chapters.draft.length > 0 &&
              data.chapters.draft.map((item: ChapterItem) =>
                <Menu.Item className={styles.writeMenuItem} key={item.id} mode='vertical-left'>
                  <ChapterMenuItem words={item.words} updatedAt={item.updatedAt} title={item.title}/>
                </Menu.Item>)
              }
            </Menu.SubMenu>
            <Menu.SubMenu className={styles.writeMenuItem}
                          key='publish'
                          icon={<SendOutlined className={styles.writeIcon}/>}
                          title={<>已发布
                            <Text className={styles.smallText}>
                              共{data.chapters.published.subsectionNum}卷 {data.chapters.published.chapterNum}章
                            </Text></>}
            >
              {data.chapters.published.subsection.map((item: SubItem) =>
                <Menu.SubMenu key={item.id} title={<Text>第{item.seq}卷 {item.name}</Text>}>
                  {item.chapters && item.chapters.map((subItem: ChapterItem) =>
                    <Menu.Item className={styles.writeMenuItem} key={subItem.id} mode='vertical-left'>
                      <ChapterMenuItem words={subItem.words} updatedAt={subItem.updatedAt} title={subItem.title}/>
                    </Menu.Item>,
                  )}
                </Menu.SubMenu>)
              }
            </Menu.SubMenu>
            <Menu.SubMenu className={styles.writeMenuItem}
                          key='recycle'
                          icon={<RestOutlined className={styles.writeIcon}/>}
                          title={<>回收站
                            <Text className={styles.smallText}>
                              (共{(data.chapters.recycle && data.chapters.recycle.length) || 0}章)
                            </Text></>}
            >
              {data.chapters.recycle && data.chapters.recycle.map((item: ChapterItem) =>
                <Menu.Item className={styles.writeMenuItem} key={item.id} mode='vertical-left'>
                  <ChapterMenuItem words={item.words} updatedAt={item.updatedAt} title={item.title}/>
                </Menu.Item>)
              }
            </Menu.SubMenu>
          </Menu>
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
                  <Dropdown trigger={['click']}
                            overlay={<Menu>
                              {data.chapters.published.subsection.map((item: SubItem) =>
                                <Menu.Item key={item.id} onClick={() => setCurrentSub(item)}>第{item.seq}卷</Menu.Item>)}
                              <Menu.Item key='new'>新建分卷</Menu.Item>
                            </Menu>}>
                    <Button style={{ border: 0, boxShadow: 'none' }}>
                      第{currentSub?.seq}卷
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
                  <Button shape={'round'} type={'primary'} size={'small'}>删除</Button>
                  <Button shape={'round'} size={'small'}>保存</Button>
                  <Button shape={'round'} size={'small'}
                          style={{ backgroundColor: '#0067E6', color: 'white' }}>
                    发布
                  </Button>
                </Space>
              </Col>
            </Row>
          </Header>
          <Content className={styles.writeContent}>
            <div className={styles.chapterBody}>
              <div className={styles.scrollDiv}>
                <Slate editor={titleEditor} value={title} onChange={value => setTitle(value)}>
                  <Editable style={{ fontSize: 20, backgroundColor: 'inherit', borderBottom: '1px solid #e0e0e0' }}
                            placeholder='请输入章节号和章节名。示例："第一章 起始"'/>
                </Slate>
                <Slate editor={editor} value={value} onChange={value => setValue(value)}>
                  <Editable onKeyDown={event => {
                    if (event.keyCode == 9) {
                      editor.insertText('\t');
                      event.preventDefault();
                    }
                    if (event.keyCode == 13) {
                      editor.insertText('\n\t');
                      event.preventDefault();
                    }
                  }}
                            className={styles.writeArea}
                            placeholder="输入正文..."
                  />
                </Slate>
                <Divider type={'horizontal'}/>
                <div style={{ width: '100%', textAlign: 'center' }}>
                  <a style={{ fontSize: 13, color: '#bfbfbf', alignSelf: 'center' }}>+ 作者的话</a>
                </div>
              </div>
            </div>
          </Content>
        </Layout>
      </Layout>
      }
    </Skeleton>
  );
}
