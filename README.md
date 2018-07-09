## 项目目录

── goaction
    ├── crawler                 -- 单机并发爬虫
    │   ├── engine
    │   ├── fetcher
    │   ├── lagou
    │   │   └── parser
    │   ├── model
    │   ├── persist
    │   ├── redis
    │   ├── scheduler
    │   └── zhenai
    │       └── parser
    ├── crawler_distribute       -- 分布式并发爬虫
    │   ├── config
    │   ├── itemsaver
    │   ├── persist
    │   │   └── server
    │   ├── rpcsupport
    │   └── worker
    │       ├── client
    │       └── server
    └── frontend                 -- 前端
        ├── controller
        ├── model
        └── view
            ├── css
            └── js

