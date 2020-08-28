import styles from '@/pages/work/write/editor/index.less';
import { Input, Skeleton, Typography } from 'antd';
import { Editable, ReactEditor, Slate } from 'slate-react';
import React, { useEffect } from 'react';
import { ChapterItem } from '@/pages/work/write/sider/menuItem';
import { useRequest } from '@umijs/hooks';
import { workService } from '@/service/work';
import { initValue, nodesConvert, WriteState } from '@/pages/work/write';

const HDivider = () => <div style={{ backgroundColor: '#c4c4c4', height: 1, marginTop: 10 }}/>;


export default function Editor(props: {
  disable: boolean,
  currentItem: ChapterItem | null,
  write: WriteState,
  setWrite: (value: WriteState) => void,
  editor: ReactEditor
}) {

  const { currentItem, write, setWrite, editor, disable } = props;


  const { run: fetch, loading } = useRequest(workService.Detail, {
    manual: true, onSuccess: (res) => {
      setWrite({ title: res.title, value: convertNodes(res.content) });
    },
  });

  useEffect(() => {
    if (currentItem) {
      if (currentItem.id < 0) {
        let content = localStorage.getItem('draft');
        if (content) {
          setWrite(JSON.parse(content));
        } else {
          setWrite({ title: '', value: initValue });
        }
        return;
      }
      fetch({ id: currentItem.id, type: currentItem.type, subsection: currentItem.subsection });
    }
  }, [currentItem]);

  return (<Skeleton loading={loading}>
    <div className={styles.chapterBody}>
      <div className={styles.scrollDiv}>
        {!disable &&
        <>
          <Input value={write.title}
                 onChange={event => {
                   if (!disable) {
                     setWrite({ title: event.target.value, value: write.value });
                   }
                 }}
                 style={{ fontSize: 20, border: 0, boxShadow: 'none' }} placeholder='请输入章节号和章节名。示例："第一章 起始"'/>
          <HDivider/>
          <Slate editor={editor} value={write.value} onChange={v => setWrite({ title: write.title, value: v })}>
            <Editable className={styles.writeArea}
                      onKeyDown={event => {
                        if (event.keyCode == 9) {
                          editor.insertText('\t');
                          event.preventDefault();
                        }
                        if (event.keyCode == 13) {
                          editor.insertText('\n\t');
                          event.preventDefault();
                        }
                      }}
                      spellCheck
                      autoFocus/>
          </Slate>
        </>}
        {disable && <Typography style={{ paddingTop: 20 }}>
          <Typography.Title style={{ fontSize: 20, border: 0, boxShadow: 'none' }}>
            {write.title}
          </Typography.Title>
          <HDivider/>
          <div className={styles.writeArea}>
            {write.value.map((item) =>
              // @ts-ignore
              <Typography.Paragraph>{item.children[0].text}</Typography.Paragraph>,
            )}
          </div>
        </Typography>}
        {/*<HDivider/>*/}
        {/*<div style={{ width: '100%', textAlign: 'center' }}>*/}
        {/*  <a style={{ fontSize: 13, color: '#bfbfbf', alignSelf: 'center' }}>+ 作者的话</a>*/}
        {/*</div>*/}
      </div>
    </div>
  </Skeleton>);
}

const convertNodes = (content: string) => {
  let splits = content.split('\n');
  console.log(splits);
  let nodes: { children: { text: string; }[]; }[] = [];
  splits.forEach((item) => {
    nodes.push({ children: [{ text: item }] });
  });

  return nodes;
};

