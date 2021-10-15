create table sys_role
(
    id               bigint auto_increment comment '编号'
        primary key,
    name             varchar(100)                        null comment '角色名称',
    value            varchar(100)                        null comment '角色值',
    sort             int                                 null comment '排序',
    remark           varchar(100)                        null comment '备注',
    create_by        varchar(50)                         null comment '创建人',
    create_time      timestamp default CURRENT_TIMESTAMP not null comment '创建时间',
    last_update_by   varchar(50)                         null comment '更新人',
    last_update_time datetime                            null on update CURRENT_TIMESTAMP comment '更新时间',
    del_flag         tinyint   default 0                 null comment '是否删除  -1：已删除  0：正常',
    status           int       default 1                 null comment '状态  1:启用,0:禁用'
)
    comment '角色管理';

INSERT INTO sys_role (id, name, value, sort, remark, create_by, create_time, last_update_by, last_update_time, del_flag, status) VALUES (1, '超级管理员', 'admin', 1, '超级管理员', 'admin', '2019-01-19 11:11:11', 'admin', '2021-10-14 17:46:50', 0, 1);
INSERT INTO sys_role (id, name, value, sort, remark, create_by, create_time, last_update_by, last_update_time, del_flag, status) VALUES (2, '项目经理', 'mng', 2, '项目经理', 'admin', '2019-01-19 11:11:11', 'admin', '2021-10-14 17:46:51', 0, 1);
INSERT INTO sys_role (id, name, value, sort, remark, create_by, create_time, last_update_by, last_update_time, del_flag, status) VALUES (3, '开发人员', 'dev', 3, '开发人员', 'admin', '2019-01-19 11:11:11', 'admin', '2021-10-14 17:46:51', 0, 1);
INSERT INTO sys_role (id, name, value, sort, remark, create_by, create_time, last_update_by, last_update_time, del_flag, status) VALUES (4, '测试人员1', 'test', 4, '测试人员1', 'admin', '2019-01-19 11:11:11', 'admin', '2021-10-14 17:46:52', 0, 1);