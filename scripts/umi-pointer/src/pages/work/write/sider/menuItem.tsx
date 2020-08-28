import { Card, Typography } from 'antd';
import styles from '@/pages/work/write/index.less';
import moment from 'moment';
import React from 'react';

export interface ChapterItem {
  words: number;
  id: number;
  updatedAt: any;
  title: string;
  type: 'draft' | 'published' | 'recycle';
  subsection: number,
}

export interface SubsectionItem {
  id: number;
  seq: number,
  name: string,
  introduce: string,
  chapters: ChapterItem[],
}

const ChapterMenuItem = (item: { words: number; updatedAt: any; title: string, del?: boolean }) =>
  <Card size={'small'} bordered={false} style={{ backgroundColor: 'inherit' }}>
    <Card.Meta title={<Typography.Text style={{ fontWeight: 8, fontSize: 13 }}>{item.title}</Typography.Text>}
               description={<Typography.Text className={styles.smallText}>
                 {item.updatedAt && (moment(item.updatedAt).format('M-D H:m') + '  ')}
                 {item.words}字
               </Typography.Text>}
    />
  </Card>;

export default ChapterMenuItem;
