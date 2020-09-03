# 分类
drop table if exists `category`;
create table `category`
(
    id         int(10)             not null auto_increment comment 'ID',
    name       varchar(100) unique not null comment '名称',
    work_num   int(10)             not null default 0 comment '类别内作品数',
    created_at datetime            not null default current_timestamp comment '创建时间',
    created_by varchar(100)        not null default 'root' comment '创建人',
    updated_at datetime            not null default current_timestamp comment '修改时间',
    updated_by varchar(100)        not null default 'root' comment '修改人',
    PRIMARY KEY (id)
) engine = InnoDB comment ='类别表';

# 初始化种类数据
insert into `category`(name)
values ('动漫'),
       ('书籍'),
       ('漫画'),
       ('人物'),
       ('电影'),
       ('音乐'),
       ('其他媒体'),
       ('戏剧'),
       ('电视'),
       ('游戏'),
       ('未分类');

# 专题
drop table if exists `topic`;
create table `topic`
(
    id             int(10)      not null unique auto_increment comment 'ID',
    title          varchar(100) not null unique comment '专题名',
    category_id    int(10)      not null comment '所属分类ID',
    user_id        int(10)      not null comment '创建人ID',
    work_num       int(10)      not null default 0 comment '专题内作品数',
    subscribe_nums int(10)      not null default 0 comment '订阅数',
    original       varchar(100) not null default '' comment '专题原作',
    original_url   varchar(100) not null default '' comment '原作链接',
    description    varchar(500) not null comment '原作描述',
    created_at     datetime     not null default current_timestamp comment '创建时间',
    created_by     varchar(100) not null default 'root' comment '创建人',
    updated_at     datetime     not null default current_timestamp comment '修改时间',
    updated_by     varchar(100) not null default 'root' comment '修改人',
    PRIMARY KEY (id)
) engine = InnoDB comment = '专题表';

# 标签
drop table if exists `tag`;
create table `tag`
(
    id         int(10)      not null unique auto_increment comment 'ID',
    name       varchar(100) not null unique comment '标签名称',
    created_at datetime     not null default current_timestamp comment '创建时间',
    created_by varchar(100) not null default 'root' comment '创建人',
    updated_at datetime     not null default current_timestamp comment '修改时间',
    updated_by varchar(100) not null default 'root' comment '修改人',
    PRIMARY KEY (id)
) engine = InnoDB comment = '标签表';

# 作品
drop table if exists `work`;
create table `work`
(
    id          int(10)      not null unique auto_increment comment 'ID',
    type        tinyint(1)   not null default 0 comment '类型：0-同人，1-原创',
    title       varchar(100) not null comment '标题',
    introduce   varchar(200) not null comment '简述',
    user_id     int(10)      not null comment '作者ID',
    category_id int(10)      not null default 0 comment '分类ID(同人)',
    topic_id    int(10)      not null default 0 comment '专题ID(同人)',
    `lock`      boolean      not null default false comment '锁住不允观看',
    created_at  datetime     not null default current_timestamp comment '创建时间',
    created_by  varchar(100) not null default 'root' comment '创建人',
    updated_at  datetime     not null default current_timestamp comment '修改时间',
    updated_by  varchar(100) not null default 'root' comment '修改人',
    PRIMARY KEY (id)
) engine = InnoDB comment = '作品表';

# 作品扩展表
drop table if exists `work_ex`;
create table `work_ex`
(
    work_id         int(10)      not null unique comment '作品ID',
    words           int(10)      not null default 0 comment '字数',
    view_nums       int(10)      not null default 0 comment '点击数',
    talk_nums       int(10)      not null default 0 comment '讨论数',
    college_nums    int(10)      not null default 0 comment '收藏数',
    subscribe_nums  int(10)      not null default 0 comment '订阅数',
    chapter_nums    int(10)      not null default 0 comment '已发布章节数',
    subsection_nums int(10)      not null default 1 comment '卷数,默认有1卷',
    draft_nums      int(10)      not null default 0 comment '草稿数',
    recycle_nums    int(10)      not null default 0 comment '回收章节数',
    created_at      datetime     not null default current_timestamp comment '创建时间',
    created_by      varchar(100) not null default 'root' comment '创建人',
    updated_at      datetime     not null default current_timestamp comment '修改时间',
    updated_by      varchar(100) not null default 'root' comment '修改人',
    primary key (work_id)
) engine = InnoDB comment = '作品扩展表';

# 作品标签关系表
drop table if exists `work_tag`;
create table `work_tag`
(
    id         int(10)      not null unique auto_increment comment 'ID',
    work_id    int(10)      not null comment '作品ID',
    tag_id     int(10)      not null comment '标签ID',
    created_at datetime     not null default current_timestamp comment '创建时间',
    created_by varchar(100) not null default 'root' comment '创建人',
    updated_at datetime     not null default current_timestamp comment '修改时间',
    updated_by varchar(100) not null default 'root' comment '修改人',
    PRIMARY KEY (id),
    unique key work_tag_unique (work_id, tag_id)
) engine = InnoDB comment = '作品标签关系表';

# 分卷表
drop table if exists `subsection`;
create table `subsection`
(
    id         int(10)      not null unique auto_increment comment 'ID',
    work_id    int(10)      not null unique comment '作品ID',
    name       varchar(100) comment '卷名',
    introduce  varchar(200) comment '分卷介绍',
    seq        int(10)      not null comment '序号',
    work_num   int(10)      not null default 0 comment '卷内作品数',
    created_at datetime     not null default current_timestamp comment '创建时间',
    created_by varchar(100) not null default 'root' comment '创建人',
    updated_at datetime     not null default current_timestamp comment '修改时间',
    updated_by varchar(100) not null default 'root' comment '修改人'
) engine = InnoDB comment ='分卷表';

# 章节表
drop table if exists `chapter`;
create table `chapter`
(
    id            int(10)      not null unique auto_increment comment 'ID',
    work_id       int(10)      not null comment '作品ID',
    title         varchar(100) not null comment '标题',
    content       text         not null comment '内容',
    status        tinyint(1)   not null default 0 comment '状态:(0-草稿,1-发布,2-回收)',
    seq           int(10)      not null default 1 comment '章节序号',
    version       int(10)      not null default 1 comment '发布版本',
    subsection_id int(10)      not null default 1 comment '分卷ID',
    `lock`        boolean      not null default false comment '锁住不允观看',
    created_at    datetime     not null default current_timestamp comment '创建时间',
    created_by    varchar(100) not null default 'root' comment '创建人',
    updated_at    datetime     not null default current_timestamp comment '修改时间',
    updated_by    varchar(100) not null default 'root' comment '修改人',
    PRIMARY KEY (id),
    unique unique_seq (work_id, seq, version)
) engine = InnoDB comment = '章节表';

#用户
drop table if exists `user`;
create table `user`
(
    id         int(10)      not null unique auto_increment comment 'ID',
    username   varchar(100) not null unique comment '用户名',
    email      varchar(100) not null unique comment '邮箱',
    password   varchar(200) not null comment '密码',
    root       bool         not null default false comment '是否管理员',
    created_at datetime     not null default current_timestamp comment '创建时间',
    created_by varchar(100) not null default 'root' comment '创建人',
    updated_at datetime     not null default current_timestamp comment '修改时间',
    updated_by varchar(100) not null default 'root' comment '修改人',
    PRIMARY KEY (id)
) engine = InnoDB comment = '用户表';

insert into `user`(username, email, password, root)
    value ('shyptr', 'shyptr14@gmial.com', '$2a$10$LNXhKGnLdjKZ2Kvd.l99KOaoI31dMchuoSbI7rFgYGCm.ofUwFc8C', true);

# 用户扩展表
drop table if exists `user_ex`;
create table `user_ex`
(
    user_id    int(10)      not null unique comment '用户ID',
    works_nums int(10)      not null default 0 comment '作品数',
    work_day   int(10)      not null default 0 comment '创作天数',
    words      int(10)      not null default 0 comment '累计字数',
    fans_nums  int(10)      not null default 0 comment '粉丝数',
    created_at datetime     not null default current_timestamp comment '创建时间',
    created_by varchar(100) not null default 'root' comment '创建人',
    updated_at datetime     not null default current_timestamp comment '修改时间',
    updated_by varchar(100) not null default 'root' comment '修改人',
    PRIMARY KEY (user_id)
) engine = InnoDB comment = '用户表';

# 待认证用户表
drop table if exists `identify`;
create table `identify`
(
    id         int(10)      not null unique auto_increment comment 'ID',
    email      varchar(100) not null unique comment '邮箱',
    created_at datetime     not null default current_timestamp comment '创建时间',
    created_by varchar(100) not null default 'root' comment '创建人',
    primary key (id)
) engine = InnoDB comment ='认证表';

# 创作日历表
drop table if exists `calendar`;
create table `calendar`
(
    id         int(10)      not null unique auto_increment comment 'ID',
    user_id    int(10)      not null unique comment '用户ID',
    year       int(10)      not null comment '创作年份',
    month      int(10)      not null comment '创作月份',
    day        int(10)      not null comment '创作日',
    words      int(10)      not null comment '创作字数',
    created_at datetime     not null default current_timestamp comment '创建时间',
    created_by varchar(100) not null default 'root' comment '创建人',
    updated_at datetime     not null default current_timestamp comment '修改时间',
    updated_by varchar(100) not null default 'root' comment '修改人',
    PRIMARY KEY (id),
    unique key unique_user_calendar (user_id, year, month, day),
    index find_calendar (user_id, year, month, day, words)
) engine = InnoDB comment ='创作日历表';

# 订阅表
drop table if exists `subscription`;
create table `subscription`
(
    id         int(10)      not null unique auto_increment comment 'ID',
    user_id    int(10)      not null unique comment '用户ID',
    obj_type   tinyint(1)   not null default 0 comment '订阅类型:(0-专题,1-作品,2-用户)',
    obj_id     int(10)      not null comment '订阅目标ID',
    created_at datetime     not null default current_timestamp comment '创建时间',
    created_by varchar(100) not null default 'root' comment '创建人',
    updated_at datetime     not null default current_timestamp comment '修改时间',
    updated_by varchar(100) not null default 'root' comment '修改人',
    PRIMARY KEY (id)
) engine = InnoDB comment ='订阅表';

# 公告表
drop table if exists `news`;
create table `news`
(
    id            int(10)      not null unique auto_increment comment 'ID',
    title         varchar(100) not null unique comment '标题',
    description   varchar(200) not null comment '简述',
    content       text         not null comment '内容',
    comments_nums int(10)      not null default 0 comment '评论数',
    created_at    datetime     not null default current_timestamp comment '创建时间',
    created_by    varchar(100) not null default 'root' comment '创建人',
    updated_at    datetime     not null default current_timestamp comment '修改时间',
    updated_by    varchar(100) not null default 'root' comment '修改人',
    PRIMARY KEY (id)
) engine = InnoDB comment ='公告表';


# 评论表
drop table if exists `comments`;
create table `comments`
(
    id         int(10)      not null unique auto_increment comment 'ID',
    user_id    int(10)      not null unique comment '用户ID',
    obj_type   tinyint(1)   not null default 0 comment '评论类型:(0-作品,1-章节,2-选段,3-公告)',
    obj_id     int(10)      not null comment '评论目标ID',
    content    varchar(300) not null comment '评论内容',
    created_at datetime     not null default current_timestamp comment '创建时间',
    created_by varchar(100) not null default 'root' comment '创建人',
    updated_at datetime     not null default current_timestamp comment '修改时间',
    updated_by varchar(100) not null default 'root' comment '修改人',
    PRIMARY KEY (id)
) engine = InnoDB comment ='评论表';

# 评论回复表
drop table if exists `reply`;
create table `reply`
(
    id         int(10)      not null unique auto_increment comment 'ID',
    user_id    int(10)      not null unique comment '用户ID',
    comment_id int(10)      not null comment '评论ID',
    content    varchar(300) not null comment '回复内容',
    created_at datetime     not null default current_timestamp comment '创建时间',
    created_by varchar(100) not null default 'root' comment '创建人',
    updated_at datetime     not null default current_timestamp comment '修改时间',
    updated_by varchar(100) not null default 'root' comment '修改人',
    PRIMARY KEY (id)
) engine = InnoDB comment ='评论回复表';

# 书单表
drop table if exists `college`;
create table `college`
(
    id         int(10)      not null unique auto_increment comment 'ID',
    user_id    int(10)      not null unique comment '用户ID',
    title      varchar(100) not null comment '书单名',
    introduce  varchar(300) not null comment '书单介绍',
    works_nums int(10)      not null default 0 comment '书单内作品数',
    created_at datetime     not null default current_timestamp comment '创建时间',
    created_by varchar(100) not null default 'root' comment '创建人',
    updated_at datetime     not null default current_timestamp comment '修改时间',
    updated_by varchar(100) not null default 'root' comment '修改人',
    PRIMARY KEY (id)
) engine = InnoDB comment ='书单表';

# 书单作品表
drop table if exists `college_work`;
create table `college_work`
(
    id         int(10)      not null unique auto_increment comment 'ID',
    college_id int(10)      not null comment '书单ID',
    work_id    int(10)      not null comment '作品ID',
    created_at datetime     not null default current_timestamp comment '创建时间',
    created_by varchar(100) not null default 'root' comment '创建人',
    updated_at datetime     not null default current_timestamp comment '修改时间',
    updated_by varchar(100) not null default 'root' comment '修改人',
    PRIMARY KEY (id)
) engine = InnoDB comment ='书单作品表';

# 私信表
drop table if exists `message`;
create table `message`
(
    id          int(10)      not null unique auto_increment comment 'ID',
    type        tinyint(1)   not null comment '私信类型:(0-系统,1-用户)',
    receiver_id int(10)      not null comment '接收用户ID',
    sender_id   int(10)      not null comment '发送用户ID',
    content     varchar(300) not null comment '私信内容',
    readed      boolean      not null default false comment '已读',
    created_at  datetime     not null default current_timestamp comment '创建时间',
    created_by  varchar(100) not null default 'root' comment '创建人',
    updated_at  datetime     not null default current_timestamp comment '修改时间',
    updated_by  varchar(100) not null default 'root' comment '修改人',
    PRIMARY KEY (id)
) engine = InnoDB comment ='私信表';