create table k8s_service
(
    id               bigint auto_increment comment '编号'
        primary key,
    name             varchar(50)                         not null comment '服务名称',
    content          varchar(300)                        null comment '内容',
    create_time      timestamp default CURRENT_TIMESTAMP not null comment '创建时间',
    last_update_time datetime                            null on update CURRENT_TIMESTAMP comment '更新时间',
    constraint name
        unique (name)
)
    comment 'service服务';

