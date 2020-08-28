import { Menu, Skeleton } from 'antd';
import styles from '@/pages/work/write/sider/index.less';
import React from 'react';
import ChapterMenuItem, { ChapterItem } from '@/pages/work/write/sider/menuItem';
import { useSelector } from '@@/plugin-dva/exports';

export default function DraftMenu(props: {
  newDraft: boolean,
  draftChapter: ChapterItem,
  currentItem: ChapterItem | null,
  setCurrentItem: (item: ChapterItem) => void
}) {

  const { data, loading } = useSelector((state: any) => {
    return {
      data: state.work.draft,
      loading: state.loading.effects['work/query'],
    };
  });

  const { newDraft, currentItem, setCurrentItem, draftChapter } = props;


  return (
    <Skeleton loading={loading}>
      <Menu className={styles.writeMenu} selectedKeys={[currentItem?.id + '']}>
        {newDraft &&
        <Menu.Item onClick={() => setCurrentItem(draftChapter)}
                   className={styles.writeMenuItem} key={draftChapter.id} mode='vertical-left'>
          <ChapterMenuItem words={0} updatedAt={null} title={'未命名'}/>
        </Menu.Item>}
        {data.map((item: ChapterItem) =>
          <Menu.Item onClick={() => setCurrentItem(item)}
                     className={styles.writeMenuItem} key={item.id} mode='vertical-left'>
            <ChapterMenuItem words={item.words} updatedAt={item.updatedAt} title={item.title}/>
          </Menu.Item>)
        }
      </Menu>
    </Skeleton>
  );

}
