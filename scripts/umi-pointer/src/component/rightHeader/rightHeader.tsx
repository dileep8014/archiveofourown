import React, { useRef, useState } from 'react';
import { useModel, Link } from 'umi';
import { AutoComplete, Avatar, Button, Dropdown, Input, Menu, Modal } from 'antd';
import styles from './rightHeader.less';
import { SelectProps } from 'antd/es/select';
import SignForm from '@/component/sign/sign';

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


  if (user) {
    return (
      <div className={styles.right}>
        {search}
        <Dropdown overlay={
          <Menu onClick={signOut}>
            <Menu.Item>登出</Menu.Item>
          </Menu>}>
          <Link to='/userCenter' className={styles.userAvatar}>
            <Avatar src={user?.avatar}/><span>{user?.name}</span>
          </Link>
        </Dropdown>
        <Button className={styles.rightBtn}>创作中心</Button>
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
        <SignForm tab={tab}/>
      </Modal>}
    </div>
  );

};


export default RightHeader;
