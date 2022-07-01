create table comment
(
    id         bigint unsigned auto_increment comment '编号'
        primary key,
    pid        bigint unsigned  not null comment '父级ID，为0表示顶级',
    post_id    bigint unsigned  not null comment '帖子ID',
    user_id    bigint unsigned  not null comment '评论人ID',
    status     tinyint unsigned not null comment '评论状态 0表示屏蔽 1表示正常',
    content    text             null comment '评论内容',
    created_at timestamp        not null comment '创建时间',
    updated_at timestamp        null comment '更新时间'
)
    comment '评论表';

create index comment_pid_index
    on comment (pid);

create index comment_post_id_index
    on comment (post_id);

create index comment_user_id_index
    on comment (user_id);

create table post
(
    id          bigint unsigned auto_increment comment '主键编号'
        primary key,
    title       varchar(80)                  not null comment '标题',
    description varchar(300)                 null comment '文章描述',
    user_id     bigint unsigned              not null comment '作者编号',
    category_id bigint unsigned              not null comment '分类编号',
    content     longtext                     not null comment '文章内容',
    status      tinyint unsigned default '1' not null comment '文章状态:0禁用1正常',
    created_at  timestamp                    not null comment '创建时间',
    updated_at  timestamp                    null comment '更新时间'
)
    comment '文章帖子';

create index post_category_id_index
    on post (category_id);

create index post_status_index
    on post (status);

create index post_user_id_index
    on post (user_id);

create table role
(
    id             bigint unsigned auto_increment comment '主键编号'
        primary key,
    avatar         varchar(300)                 null comment '头像',
    name           varchar(80)                  not null comment '名称',
    nickname       varchar(80)                  null comment '别名',
    survival_stage varchar(80)                  null comment '生存阶段',
    nationality    varchar(100)                 null comment '国籍',
    achievement    varchar(300)                 null comment '主要成就',
    gender         tinyint unsigned             not null comment '是否展示:0女1男',
    content        longtext                     not null comment '介绍',
    if_show        tinyint unsigned default '1' not null comment '是否展示:0否1是',
    created_at     timestamp                    not null comment '创建时间',
    updated_at     timestamp                    null comment '更新时间',
    constraint role_name_uindex
        unique (name)
)
    comment '故事角色' auto_increment = 15;

create index role_if_show_index
    on role (if_show);

create index role_name_index
    on role (name);

create table user
(
    id         bigint auto_increment comment '用户编号'
        primary key,
    avatar     varchar(300)                 null comment '头像',
    nickname   varchar(80)                  not null comment '用户昵称',
    phone      varchar(11)                  null comment '手机号',
    password   varchar(200)                 not null comment '密码',
    status     tinyint unsigned default '1' not null comment '状态:0禁用1正常',
    created_at timestamp                    not null comment '创建时间',
    constraint user_phone_uindex
        unique (phone)
)
    comment '用户信息表' auto_increment = 24;

create index user_status_index
    on user (status);

