create table sys_dict
(
    id               bigint auto_increment comment '编号'
        primary key,
    value            varchar(100)                        not null comment '数据值',
    label            varchar(100)                        not null comment '标签名',
    type             varchar(100)                        not null comment '类型',
    description      varchar(100)                        not null comment '描述',
    sort             decimal                             not null comment '排序（升序）',
    create_by        varchar(50)                         null comment '创建人',
    create_time      timestamp default CURRENT_TIMESTAMP  comment '创建时间',
    last_update_by   varchar(50)                         null comment '更新人',
    last_update_time datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
    remarks          varchar(255)                        null comment '备注信息',
    del_flag         tinyint   default 0                 null comment '是否删除  -1：已删除  0：正常'
)
    comment '字典表';

INSERT INTO sys_dict (id, value, label, type, description, sort, create_by, create_time, last_update_by, last_update_time, remarks, del_flag) VALUES (1, 'male', '男', 'sex', '性别', 0, 'admin', '2018-09-23 19:52:54', 'admin', null, '性别', 0);
INSERT INTO sys_dict (id, value, label, type, description, sort, create_by, create_time, last_update_by, last_update_time, remarks, del_flag) VALUES (2, 'female', '女', 'sex', '性别', 1, 'admin', '2018-09-23 19:53:17', 'admin', null, '性别', 0);