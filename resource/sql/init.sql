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
    comment '用户信息表' auto_increment = 8;

create index user_status_index
    on user (status);

