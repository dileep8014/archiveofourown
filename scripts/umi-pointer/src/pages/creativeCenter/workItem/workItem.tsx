import { Button, Card, Col, Row } from 'antd';
import React from 'react';

export interface WorkItemData {
  id: number
  title: string
  cover: string
  newChapter: string | null
  comments: number,
  subscribe: number,
  college: number,
  hits: number,
}

export default function WorkItem(props: WorkItemData) {
  return (
    <Card bordered={false} hoverable={false}>
      <Card.Meta title={props.title}
                 avatar={<img style={{ height: 120, width: 100 }} src={props.cover}/>}
                 description={<div>
                   <span>{'最新章节：' + props.newChapter || '无最新章节'}</span>
                   <br/>
                   <br/>
                   <br/>
                   <Row style={{ bottom: 0 }}>
                     <Col span={20}>
                       <span>收藏：{props.college} </span>
                       <span style={{ marginLeft: 10 }}>订阅：{props.subscribe}</span>
                       <span style={{ marginLeft: 10 }}>评论：{props.comments}</span>
                       <span style={{ marginLeft: 10 }}>点击：{props.hits}</span>
                     </Col>
                     <Col span={4}>
                       <Button>作品设置</Button>
                       <Button onClick={() => window.open('/works/write/'+props.id)} type={'primary'}
                               style={{ marginLeft: 10 }}>新章节</Button>
                     </Col>
                   </Row>
                 </div>}
      />
    </Card>
  );
}
