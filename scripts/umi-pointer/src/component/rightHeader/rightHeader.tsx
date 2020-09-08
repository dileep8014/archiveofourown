import React, { ReactText, useState } from 'react';
import { useModel, history, Link } from 'umi';
import { AutoComplete, Avatar, Button, Dropdown, Input, Menu, Modal } from 'antd';
import styles from './rightHeader.less';
import { SelectProps } from 'antd/es/select';
import SignForm from '@/component/sign/sign';
import ProgressOpt from '@/component/progress/progress';

function getRandomInt(max: number, min: number = 0) {
  return Math.floor(Math.random() * (max - min + 1)) + min; // eslint-disable-line no-mixed-operators
}

const searchResult = (query: string) => {
  return new Array(getRandomInt(5))
    .join('.')
    .split('.')
    .map((item, idx) => {
      const category = `${query}${idx}`;
      return {
        value: category,
        label: (
          <div
            style={{
              display: 'flex',
              justifyContent: 'space-between',
            }}
          >
            <span>
              Found {query} on{' '}
              <a
                href={`https://s.taobao.com/search?q=${query}`}
                target="_blank"
                rel="noopener noreferrer"
              >
                {category}
              </a>
            </span>
            <span>{getRandomInt(200, 100)} results</span>
          </div>
        ),
      };
    });
};

const RightHeader: React.FC = () => {

  const { user, signOut } = useModel('user', model => ({
    user: model.user,
    signOut: model.signout,
  }));

  const [canVisit, setCanVisit] = useState(false);
  const [tab, setTab] = useState('signIn');

  const visit = (login: boolean | null, cancel?: boolean) => {
    if (cancel !== undefined && cancel) {
      setCanVisit(false);
    }
    if (login !== null) {
      if (login) {
        setTab('signIn');
      } else {
        setTab('signUp');
      }
      setCanVisit(true);
    }
  };

  const [options, setOptions] = useState<SelectProps<object>['options']>([]);

  const handleSearch = (value: string) => {
    setOptions(value ? searchResult(value) : []);
  };

  const onSelect = (value: string) => {
    console.log('onSelect', value);
  };

  const search = (
    <AutoComplete dropdownMatchSelectWidth={152}
                  options={options}
                  onSelect={onSelect}
                  onSearch={handleSearch}
                  className={styles.search}
    >
      <Input.Search style={{ border: 0 }} placeholder='在Pointer中搜索' size={'middle'}/>
    </AutoComplete>);

  const menuSelect = (key: ReactText) => {
    if (key === 'user') {
      history.push('/userCenter');
    }
    if (key === 'out') {
      ProgressOpt(signOut);
    }
  };


  if (user) {
    return (
      <div className={styles.right}>
        {search}
        <Dropdown overlay={
          <Menu onClick={({ key }) => menuSelect(key)}>
            <Menu.Item key='user'>个人中心</Menu.Item>
            <Menu.Item key='out'>登出</Menu.Item>
          </Menu>}>
          <span className={styles.userAvatar}>
            <Avatar src={user?.avatar}/><span> {user?.username}</span>
          </span>
        </Dropdown>
        <Link to='/msgCenter' className={styles.rightBtn}>消息</Link>
        <Link to='/creativeCenter' className={styles.rightBtn}>创作中心</Link>
      </div>
    );
  }

  return (
    <div className={styles.right}>
      {search}
      <Button onClick={() => visit(true)} className={styles.rightBtn}>登录</Button>
      <Button onClick={() => visit(false)} className={styles.rightBtn}>注册</Button>
      {canVisit && <Modal
        centered={true}
        width={400}
        footer={null}
        onCancel={() => visit(null, true)}
        visible={true}
      >
        <SignForm tab={tab} finish={() => visit(null, true)}/>
      </Modal>}
    </div>
  );

};


export default RightHeader;
