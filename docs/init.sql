# 分类
create table `category`
(
    id         int(10)             not null auto_increment comment 'ID',
    name       varchar(100) unique not null comment '名称',
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

#用户
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
    value ('shyptr', 'shyptr14@gmial.com', '$2a$10$LNXhKGnLdjKZ2Kvd.l99KOaoI31dMchuoSbI7rFgYGCm.ofUwFc8C', true)

# 待认证用户表
create table `identify`
(
    id         int(10)      not null unique auto_increment comment 'ID',
    email      varchar(100) not null unique comment '邮箱',
    created_at datetime     not null default current_timestamp comment '创建时间',
    created_by varchar(100) not null default 'root' comment '创建人',
    primary key (id)
) engine = InnoDB comment ='认证表';

# 书签
drop table if exists `mark`;
create table `mark`
(
    id          int(10)      not null unique auto_increment comment 'ID',
    name        varchar(100) not null unique comment '名称',
    category_id int(10)      not null comment '分类ID',
    created_at  datetime     not null default current_timestamp comment '创建时间',
    created_by  varchar(100) not null default 'root' comment '创建人',
    updated_at  datetime     not null default current_timestamp comment '修改时间',
    updated_by  varchar(100) not null default 'root' comment '修改人',
    PRIMARY KEY (id)
) engine = InnoDB comment = '书签表';

# 标签
create table `tag`
(
    id         int(10)      not null unique auto_increment comment 'ID',
    name       varchar(100) not null unique comment '名称',
    created_at datetime     not null default current_timestamp comment '创建时间',
    created_by varchar(100) not null default 'root' comment '创建人',
    updated_at datetime     not null default current_timestamp comment '修改时间',
    updated_by varchar(100) not null default 'root' comment '修改人',
    PRIMARY KEY (id)
) engine = InnoDB comment = '标签表';

# 文章
create table `article`
(
    id                int(10)          not null unique auto_increment comment 'ID',
    title             varchar(100)     not null comment '标题',
    sub_title         varchar(200)     not null comment '简述',
    user_id           int(10)          not null comment '作者ID',
    mark_id           int(10) comment '书签ID',
    language          tinyint(1)       not null default 1 comment '1-中文，2-英语',
    words             int(10) unsigned not null default 0 comment '字数',
    view_nums         int(10) unsigned not null default 0 comment '点击数',
    talk_nums         int(10) unsigned not null default 0 comment '讨论数',
    share_nums        int(10) unsigned not null default 0 comment '分享数',
    download_nums     int(10) unsigned not null default 0 comment '下载数',
    chapter_nums      int(10) unsigned not null default 0 comment '计划总章节数',
    chapter_real_nums int(10) unsigned not null default 0 comment '当前总章节数',
    created_at        datetime         not null default current_timestamp comment '创建时间',
    created_by        varchar(100)     not null default 'root' comment '创建人',
    updated_at        datetime         not null default current_timestamp comment '修改时间',
    updated_by        varchar(100)     not null default 'root' comment '修改人',
    PRIMARY KEY (id),
    index find_title (title),
    index find_user (user_id),
    index find_mark (mark_id),
    fulltext key find_full_title (title, sub_title)
) engine = InnoDB comment = '文章表';

# 文章标签关系表
create table `article_tag`
(
    id         int(10)      not null unique auto_increment comment 'ID',
    article_id int(10)      not null comment '文章ID',
    tag_id     int(10)      not null comment '标签ID',
    created_at datetime     not null default current_timestamp comment '创建时间',
    created_by varchar(100) not null default 'root' comment '创建人',
    updated_at datetime     not null default current_timestamp comment '修改时间',
    updated_by varchar(100) not null default 'root' comment '修改人',
    PRIMARY KEY (id),
    index find_articleID (article_id),
    index find_tagID (tag_id)
) engine = InnoDB comment = '文章标签关系表';

# 文章章节表
create table `chapter`
(
    id         int(10)        not null unique auto_increment comment 'ID',
    article_id int(10)        not null comment '文章ID',
    title      varchar(100)   not null comment '标题',
    seq        int(10) unique not null default 1 comment '序号',
    content    text           not null comment '内容',
    created_at datetime       not null default current_timestamp comment '创建时间',
    created_by varchar(100)   not null default 'root' comment '创建人',
    updated_at datetime       not null default current_timestamp comment '修改时间',
    updated_by varchar(100)   not null default 'root' comment '修改人',
    PRIMARY KEY (id),
    index find_title (title),
    fulltext key find_full_title (title, content)
) engine = InnoDB comment = '章节表';