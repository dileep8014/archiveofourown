import styles from '@/pages/work/write/sider/index.less';
import ChapterMenuItem, { ChapterItem } from '@/pages/work/write/sider/menuItem';
import { Menu, Skeleton } from 'antd';
import React from 'react';
import { useSelector } from '@@/plugin-dva/exports';

export default function RecycleMenu() {

  const { data, loading } = useSelector((state: any) => {
    return { data: state.work.recycle, loading: state.loading.effects['work/query'] };
  });

  return (
    <Skeleton loading={loading}>
      <Menu>
        {data.map((item: ChapterItem) =>
          <Menu.Item className={styles.writeMenuItem} key={item.id} mode='vertical-left'>
            <ChapterMenuItem words={item.words} updatedAt={item.updatedAt} title={item.title} del={true}/>
          </Menu.Item>)
        }
      </Menu>
    </Skeleton>
  );
}
