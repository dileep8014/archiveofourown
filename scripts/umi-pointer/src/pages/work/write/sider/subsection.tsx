import styles from '@/pages/work/write/sider/index.less';
import ChapterMenuItem, { ChapterItem, SubsectionItem } from '@/pages/work/write/sider/menuItem';
import { Collapse, Menu, Skeleton, Typography } from 'antd';
import numberParseChina from '@/component/numconvert/numconvert';
import React from 'react';
import { useSelector } from '@@/plugin-dva/exports';
import { CaretRightOutlined } from '@ant-design/icons';

export default function SubsectionMenu(props: { currentItem: ChapterItem | null, setCurrentItem: (item: ChapterItem) => void }) {

  const { currentItem, setCurrentItem } = props;

  const { data, loading } = useSelector((state: any) => {
    return { data: state.work.subsection, loading: state.loading.effects['work/query'] };
  });

  return (
    <Skeleton loading={loading}>
      {currentItem &&
      <Collapse ghost bordered={false}
                defaultActiveKey={currentItem.type == 'published' && currentItem.subsection || -1}
                expandIcon={({ isActive }) => <CaretRightOutlined rotate={isActive ? 90 : 0}/>}>
        {data.map((item: SubsectionItem) =>
          <Collapse.Panel key={item.seq}
                          header={<Typography.Text>第{numberParseChina(item.seq)}卷 {item.name}</Typography.Text>}>
            <Menu className={styles.writeMenu} selectedKeys={[currentItem?.id + '']}>
              {item.chapters.map((subItem: ChapterItem) => (
                <Menu.Item onClick={() => setCurrentItem(subItem)}
                           className={styles.writeMenuItem} key={subItem.id} mode='vertical-left'>
                  <ChapterMenuItem words={subItem.words} updatedAt={subItem.updatedAt} title={subItem.title}/>
                </Menu.Item>),
              )}
            </Menu>
          </Collapse.Panel>)
        }
      </Collapse>
      }
    </Skeleton>
  );
}
