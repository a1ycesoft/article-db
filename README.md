# 文章搜索库

## 需求描述

使用TRPC-GO，搭建一个小型文章库，可以通过query任意关键词获取对应文章列表。

## 实现要求

文章库内容可以由本地文件导入，也可以从网上论坛或新闻站点爬取，量级要求在5000+。关键词可以匹配标题和内容。

不需要页面，API方式实现最小demo即可

API耗时需要控制在150ms以下，可自行压测性能。

通过项目demo，掌握db、缓存等常用中间件

代码需要提交到[腾讯工蜂代码仓库](https://code.tencent.com/)

需要有可以访问的api接口用于体验。

## 参考方案

mysql存储文章库，需要考虑索引设计

redis实现查询缓存。需要考虑命中率和缓存与db一致性，设计好数据结构

数据量大、分词能力可以尝试使用ES

可以尝试使用消息队列kafka或者nsq等做导入任务

以上中间件，使用腾讯云的服务，产生的费用团队报销。

## 设计

### es索引

使用了ik分词器插件

~~~mapping
{
    "mappings":{
        "properties":{
            "title":{
                "type":"text",
                "analyzer": "ik_max_word",
                "copy_to": "all"
            },
            "content":{
                "type":"text",
                "analyzer": "ik_max_word",
                "copy_to": "all"
            },
            "all":{
                "type":"text",
                "analyzer": "ik_max_word"
            }
        }
    }
}
~~~

### mysql表结构

~~~
-- auto-generated definition
create table articles
(
    id         bigint auto_increment
        primary key,
    created_at datetime     null,
    updated_at datetime     null,
    deleted_at datetime     null,
    title      varchar(255) null,
    content    longtext     null
)
    charset = utf8mb4
    row_format = DYNAMIC;
~~~

## 启动

- 批量导入数据

  在client文件夹中有示例，需要改动部分代码

- conf中改配置

  start.bat启动服务器

## 测试

未命中缓存时：

![image-20240507233224593](README.assets\image-20240507233224593.png)

命中缓存时：

![image-20240507233236290](README.assets\image-20240507233236290.png)

redis：

![image-20240507233044368](README.assets\image-20240507233044368.png)
