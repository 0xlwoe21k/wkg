
#create database wkg;
use wkg;
drop table users;

create table users (
    id int primary key not null auto_increment,
    username varchar(20),
    password varchar(32)
);

insert into users values (1,'gelen','gelen');


drop table company;
create table company (
    id int  primary key not null auto_increment,
    projectType varchar(10) not null default '-', #项目类型，第三方，公益，CNVD，SRC
    companyName varchar(30)  not null default '-',
    domain varchar(300) not null default  '-',            #待搜集的域名信息
    keyWord varchar(500) not null default '-',
    srcUrl varchar(30) not null default '-',             #SRC网址
    monitorStatus  bool not null default true,
    monitorRate int not null default 24,
    vulnScanStatus  bool not null default false,
    vulnScanRate   int not null default 24,   #以小时为单位
    lastUpdateTime varchar(20) not null default '-'
);

# insert into company values (1,'SRC','小米公司','xiaomi.com','-','https://sec.xiaomi.com/#/',true,24,false,24,'-');
insert into company values (2,'SRC','陌陌','momo.com','-','https://sec.momo.com/#/',true,24,false,24,'-');
insert into company values (3,'SRC','顺丰','sf-express.com','-','https://sec.sf.com/#/',true,24,false,24,'-');
insert into company values (4,'SRC','平安','pingan.com','-','https://sec.pingan.com/#/',true,24,false,24,'-');
insert into company values (5,'SRC','银联','unionpay.com','-','https://sec.jd.com/#/',true,24,false,24,'-');
# insert into company values (6,'SRC','网易','sf.jd.com','-','https://sec.jd.com/#/',true,24,false,24,'-');
# insert into company values (7,'SRC','金山','sf.jd.com','-','https://sec.jd.com/#/',true,24,false,24,'-');
# insert into company values (8,'SRC','阿里巴巴','sf.jd.com','-','https://sec.jd.com/#/',true,24,false,24,'-');
# insert into company values (9,'SRC','360SRC','sf.jd.com','-','https://sec.jd.com/#/',true,24,false,24,'-');
# insert into company values (10,'SRC','萤石','sf.jd.com','-','https://sec.jd.com/#/',true,24,false,24,'-');
# insert into company values (11,'SRC','携程','sf.jd.com','-','https://sec.jd.com/#/',true,24,false,24,'-');
# insert into company values (12,'SRC','欢聚时代','sf.jd.com','-','https://sec.jd.com/#/',true,24,false,24,'-');
# insert into company values (13,'SRC','新浪','sf.jd.com','-','https://sec.jd.com/#/',true,24,false,24,'-');
# insert into company values (14,'SRC','去哪儿','sf.jd.com','-','https://sec.jd.com/#/',true,24,false,24,'-');
# insert into company values (15,'SRC','唯品会','sf.jd.com','-','https://sec.jd.com/#/',true,24,false,24,'-');
# insert into company values (16,'SRC','苏宁','sf.jd.com','-','https://sec.jd.com/#/',true,24,false,24,'-');
# insert into company values (17,'SRC','滴滴出行','sf.jd.com','-','https://sec.jd.com/#/',true,24,false,24,'-');
# insert into company values (18,'SRC','饿了么','sf.jd.com','-','https://sec.jd.com/#/',true,24,false,24,'-');


drop table domains;
create table domains(
                        id int primary key auto_increment,
                        cid int not null ,
                        domain  varchar(150) not null default '-',
                        title   varchar(100) default '',
                        type    varchar(10) not null default '-',
                        ip      varchar(46) not null default '-',
                        source    varchar(20) not null default '-',
                        updateTime varchar(20) not null default '-',
                        isNew   bool ,
                        foreign key(cid) references company(id)
);
insert into domains values (1,1,'test.xiaomi.com','adfs','A','23.33.56.11','altdns','-',true);
insert into domains values (2,2,'test.momo.com','ssssss','A','23.33.56.11','altdns','-',false);
insert into domains values (3,2,'xcva.momo.com','234sddf','A','23.33.56.11','altdns','-',true);
insert into domains values (12,2,'sfsrc.sf-express.com','','','','','-',true);
insert into domains values (5,3,'kn.sf-tech.com','sdfsdfsdfd','A','23.33.56.11','altdns','-',true);
insert into domains values (6,3,'ielf.sf-express.com','','A','23.33.56.11','altdns','-',false);
insert into domains values (7,3,'abc.sf-express.com','','A','23.33.56.11','altdns','-',false);
insert into domains values (8,4,'abc.pingan.com','','A','23.33.56.11','altdns','-',false);
insert into domains values (9,4,'glaxy.pingan.com','','A','23.33.56.11','altdns','-',false);
insert into domains values (10,4,'test.pingan.com','','A','23.33.56.11','altdns','-',false);
insert into domains values (11,5,'test.jd.com','','A','23.33.56.11','altdns','-',false);

drop table websites;
create table websites(
                         id int primary key auto_increment,
                         cid int not null ,
                         website varchar(50) not null default '-',
                         title  varchar(50) not null default '-',
                         headers  varchar(3000) not null default '-',
                         finger    varchar(10) not null default '-',
                         screenshot      varchar(40) not null default '-',
                         updateTime varchar(20) not null default '-',
                         foreign key(cid) references company(id)
);
insert into websites values (1,1,'http://test.xiaomi.com','XX未知ssdf台','sdf','-','/static/xss.jpg','-');
insert into websites values (2,1,'http://test.adf.com','XXsdf联合sdf台','tesdfsdfst','-','/static/222x.jpg','-');
insert into websites values (3,3,'http://test.xxx.com','XX天帛sdf平台','tesdfsdst','-','/static/x333.jpg','-');
insert into websites values (4,5,'http://test.sererr.com','XX开sdf平台','sdfsdf','-','/static/x444.jpg','-');
insert into websites values (5,6,'http://test.bsfb.com','XX为是田sdfsdf台','sdfsf123123','-','/static/123213x.jpg','-');



drop table ips;
create table ips(
                    id int primary key auto_increment,
                    cid int not null ,
                    ip varchar(46) not null default '-',  #ipv6最长46位
                    os  varchar(30) not null default '-',
                    indomains  varchar(3000) not null default '-',
                    geo    varchar(10) not null default '-',
                    updateTime varchar(20) not null default '-',
                    foreign key(cid) references company(id)
);
insert into ips values (1,1,'1.23.66.2','windows','test.xiaomi.com,abc.xiaomi.com','-','-');


drop table services;
create table services(
                         id int primary key auto_increment,
                         cid int not null ,
                         service varchar(20) not null default '-',
                         ipport  varchar(1600) not null default '-',
                         product  varchar(20) not null default '-',
                         updateTime varchar(20) not null default '-',
                         foreign key(cid) references company(id)
);
insert into services values (1,1,'ssh','23.33.6.4:22','-','-');

drop table apps;
create table apps(
                     id int primary key auto_increment,
                     cid int not null ,
                     appname varchar(20) not null default '-',
                     notice  varchar(20) not null default '-',
                     updateTime varchar(20) not null default '-',
                     foreign key(cid) references company(id)
);
insert into apps values (1,1,'米家','-','-');


drop table webchatOfficeAccount;
create table webchatOfficeAccount(
                                     id int primary key auto_increment,
                                     cid int not null ,
                                     name varchar(20) not null default '-',
                                     notice  varchar(20) not null default '-',
                                     updateTime varchar(20) not null default '-',
                                     foreign key(cid) references company(id)
);
insert into webchatOfficeAccount values (1,1,'米家在线','-','-');


drop table miniProgram;
create table miniProgram(
                            id int primary key auto_increment,
                            cid int not null ,
                            name varchar(20) not null default '-',
                            notice  varchar(20) not null default '-',
                            updateTime varchar(20) not null default '-',
                            foreign key(cid) references company(id)
);
insert into miniProgram values (1,1,'米家小程序','-','-');


drop table systemConfig;
create table systemConfig(
                             id int primary key,
                             emailNotifyEnable bool default false,
                             emailServerAddr varchar(30),
                             emailUserName varchar(30),
                             emailPassword varchar(30),
                             weChatNotify_Enable bool default false,
                             weChatKey varchar(45),
                             dingtalkNotfyEnable bool default false,
                             updateTime varchar(20) not null default '-',
                             dingtalkAccessToken varchar(80)
);

insert into systemConfig values (1,true,'smtp.sina.com:25','like@mgtv.com','qwer1234',true,'asdf1-1123j1-123-123',true,'12312312312313213123','-')
