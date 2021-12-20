# create database wkg;

# drop database wkg;
use wkg;


# delete from websites where id=7;
#
# INSERT INTO `websites` (`cid`,`favicon`,`website`,`title`,`headers`,`finger`,`screenshot`,`updateTime`) VALUES (5,'1558343002','http://mock.coding.bjf.yun.unionpay.com:80','CODING - 一站式�','','','','2021-11-18 10:50:21');



# ALTER DATABASE wkg CHARACTER SET gbk COLLATE gbk;

drop table users;
create table users
(
    id       int primary key not null auto_increment,
    username varchar(20),
    password varchar(32)
) default charset = gbk  ;

insert into users
values (1, 'gelen', 'gelen');
insert into users
values (2, 'admin', '123456');
insert into users
values (3, 'test', '123456');

drop table company;
create table company
(
    id             int primary key not null auto_increment,
    projectType    varchar(10)     not null default '-', #项目类型，第三方，公益，CNVD，SRC
    companyName    text     ,
    domain         text    , #待搜集的域名信息
    keyWord        varchar(500)    not null default '-',
    srcUrl         varchar(50)     not null default '-', #SRC网址
    monitorStatus  bool            not null default true,
    monitorRate    int             not null default 24,
    vulnScanStatus bool            not null default false,
    vulnScanRate   int             not null default 24,  #以小时为单位
    lastUpdateTime varchar(20)     not null default '-'
) default charset = gbk ;

insert into company
values (2, 'SRC', '陌陌', 'momo.com', '-', 'https://sec.momo.com/#/', true, 24, false, 24, '-');
insert into company
values (3, 'SRC', '顺丰', 'sf-express.com', '-', 'https://sec.sf.com/#/', true, 24, false, 24, '-');

insert into company
values ( 10,'SRC', '陌陌', 'momo.com', '-', 'https://sec.momo.com/#/', true, 24, false, 24, '-');
insert into company
values ( 11,'SRC', '陌陌', 'momo.com', '-', 'https://sec.momo.com/#/', true, 24, false, 24, '-');
insert into company
values (12,'SRC', '陌陌', 'momo.com', '-', 'https://sec.momo.com/#/', true, 24, false, 24, '-');
insert into company
values ( 13,'SRC', '陌陌', 'momo.com', '-', 'https://sec.momo.com/#/', true, 24, false, 24, '-');
insert into company
values ( 14,'SRC', '陌陌', 'momo.com', '-', 'https://sec.momo.com/#/', true, 24, false, 24, '-');
insert into company
values ( 15,'SRC', '陌陌', 'momo.com', '-', 'https://sec.momo.com/#/', true, 24, false, 24, '-');
insert into company
values ( 16,'SRC', '陌陌', 'momo.com', '-', 'https://sec.momo.com/#/', true, 24, false, 24, '-');
insert into company
values (17,'SRC', '陌陌', 'momo.com', '-', 'https://sec.momo.com/#/', true, 24, false, 24, '-');
insert into company
values (18, 'SRC', '陌陌', 'momo.com', '-', 'https://sec.momo.com/#/', true, 24, false, 24, '-');
# insert into company
# values (4, 'SRC', '平安', 'pingan.com', '-', 'https://sec.pingan.com/#/', true, 24, false, 24, '-');
# insert into company
# values (5, 'SRC', '银联', 'unionpay.com', '-', 'https://sec.jd.com/#/', true, 24, false, 24, '-');
# insert into company values (6,'SRC','网易','sf.jd.com','-','https://sec.jd.com/#/',true,24,false,24,'-');
# insert into company values (7,'SRC','金山','sf.jd.com','-','https://sec.jd.com/#/',true,24,false,24,'-');
# insert into company values (8,'SRC','阿里巴巴','sf.jd.com','-','https://sec.jd.com/#/',true,24,false,24,'-');


drop table domains;
create table domains
(
    id         int primary key auto_increment,
    cid        int          not null,
    domain     text ,
    type       varchar(10)  not null default '-',
    source     text ,
    updateTime varchar(20)  not null default '-',
    isNew      bool
) default charset = gbk ;




drop table websites;
create table websites
(
    id         int primary key auto_increment,
    cid        int           not null,
    domain      varchar(150) ,
    ips         varchar(300),
    website    varchar(150)  not null default '-',
    favicon    varchar(32)            default '-', #favicon.ico
    faviconUrl    text , #favicon.ico
    title      text ,
    CDN         bool,
    headers    text ,
    cert    text ,
    finger     varchar(500)   not null default '-',
    screenshot  longblob   ,
    updateTime varchar(20)   not null default '-',
    isNew      bool
) default charset = gbk ;




drop table ips;
create table ips
(
    id         int primary key auto_increment,
    cid        int           not null,
    ip         varchar(46)   not null default '-', #ipv6最长46位
    os         varchar(50)   not null default '-',
    indomains  varchar(3000) not null default '-',
    geo        varchar(30)   not null default '-',
    updateTime varchar(20)   not null default '-',
    isNew      bool
) default charset = gbk ;



drop table services;
create table services
(
    id         int primary key auto_increment,
    cid        int           not null,
    service    varchar(100)   not null default '-',
    ipport     varchar(500) not null default '-',
    product    varchar(500)   not null default '-',
    updateTime varchar(20)   not null default '-',
    isNew      bool
) default charset = gbk ;


drop table apps;
create table apps
(
    id         int primary key auto_increment,
    cid        int         not null,
    appname    text not null ,
    notice     text ,
    updateTime varchar(20) not null default '-',
    isNew      bool
) default charset = gbk ;



drop table webchatOfficeAccount;
create table webchatOfficeAccount
(
    id         int primary key auto_increment,
    cid        int         not null,
    name       text not null ,
    notice     text not null ,
    updateTime varchar(20) not null default '-',
    isNew      bool
) default charset = gbk ;


drop table miniProgram;
create table miniProgram
(
    id         int primary key auto_increment,
    cid        int         not null,
    name       text,
    notice     text  ,
    updateTime varchar(20) not null default '-',
    isNew      bool
) default charset = gbk ;


drop table systemConfig;
create table systemConfig
(
    id                  int primary key,
    emailNotifyEnable   bool                 default false,
    emailServerAddr     varchar(30),
    emailUserName       varchar(30),
    emailPassword       varchar(30),
    weChatNotify_Enable bool                 default false,
    weChatKey           varchar(45),
    dingtalkNotfyEnable bool                 default false,
    updateTime          varchar(20) not null default '-',
    dingtalkAccessToken varchar(80)
) default charset = gbk ;


drop table knowledgeCategories;
create table knowledgeCategories (
    id          int primary key auto_increment,
    parentId    int,
    title   text,
    isLeaf      bool,
    `key`        bool
)default charset = gbk ;


insert into knowledgeCategories(id,parentId,title,`key`)
values (1,0,'思科',1);
insert into knowledgeCategories(id,parentId,title,`key`)
values (2,0,'ORACLE',2);
insert into knowledgeCategories(id,parentId,title,`key`)
values (3,0,'内网',3);

insert into knowledgeCategories(id,parentId,title,`key`)
values (4,1,'test',4);
insert into knowledgeCategories(id,parentId,title,`key`)
values (5,1,'xxx',5);

insert into knowledgeCategories(id,parentId,title,`key`)
values (6,3,'334',6);
insert into knowledgeCategories(id,parentId,title,`key`)
values (7,3,'xxx',7);