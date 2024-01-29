# Kratos-Blog 微服务博客系统

一个基于Golang的微服务博客系统

- 后端基于 [golang](https://go.dev/) + [go-kratos](https://go-kratos.dev/)
- 前端基于 [VUE3](https://vuejs.org/) 

## 项目结构

| 子项目名 | 项目路径                                                                                       |
|------|--------------------------------------------------------------------------------------------|
| 博文服务 | [/kratos-blog/app/blog](https://gitee.com/morning-ljn/kratos-blog/tree/dev/app/blog)       |
| 评论服务 | [/kratos-blog/app/comment](https://gitee.com/morning-ljn/kratos-blog/tree/dev/app/comment) |
| 用户服务 | [/kratos-blog/app/user](https://gitee.com/morning-ljn/kratos-blog/tree/dev/app/user)       |
| 管理前端 | [/kratos-blog/web/manager](https://gitee.com/morning-ljn/kratos-blog/web/manager)          |
| 展现前端 | [/kratos-blog/web/show](https://gitee.com/morning-ljn/kratos-blog/web/show)                |

## 功能特点

1. **文章管理**：用户可以创建、编辑和删除文章。支持 MarkDown 编辑器，可以插入图片、链接等多媒体元素。
2. **用户认证**：支持用户注册和登录功能，对用户的角色进行区分。
3. **评论功能**：用户可以在文章下发表评论，并查看其他用户的评论。
4. **分类和标签**：文章可以按照不同的分类和标签进行组织，方便用户查找和浏览相关的文章。
5. **搜索功能**：提供全文搜索功能，帮助用户快速找到所需的文章内容。
6. **响应式设计**：界面采用响应式设计，适应不同分辨率的设备，如手机、平板和桌面电脑。

## 技术架构

Kratos-blog 采用 go-kratos 微服务框架作为后端核心，实现了高并发、可扩展的服务端逻辑。具体来说，它包括以下几个部分：

1. **API 网关**：负责处理来自客户端的请求，将请求路由到相应的微服务上，并提供统一的入口地址。
2. **用户服务**：处理用户相关的操作，如注册、登录、用户信息管理等。
3. **文章服务**：负责文章的创建、编辑、删除等操作，以及文章的分类、标签管理。
4. **评论服务**：处理用户对文章的评论功能，包括添加评论、回复评论等。
5. **搜索服务**：提供全文搜索功能，帮助用户快速找到所需的文章内容。
6. **数据库**：使用高效的数据库存储用户数据、文章数据和评论数据，保证数据的持久化存储。
7. **缓存**：使用缓存技术提升系统的响应速度和性能，减少数据库的访问压力。
8. **日志和监控**：记录系统运行过程中的关键信息，并进行实时监控和告警。

## 部署方式

Kratos-blog 可以通过容器化的方式部署，支持 Docker 和 Kubernetes。具体的部署步骤如下：

1. 安装 Docker 或 Kubernetes。
2. 克隆 Kratos-blog 的源代码仓库。
3. 构建 Docker 镜像或 Kubernetes 部署文件。
4. 启动 Docker 容器或使用 Kubernetes 部署应用。
5. 配置数据库连接信息和必要的环境变量。
6. 访问 Kratos-blog 的入口地址即可开始使用。

## 总结

Kratos-blog 是一款基于 go-kratos 微服务框架的博客系统，具有丰富的功能和良好的扩展性。通过使用先进的技术和架构，它为用户提供了一个可靠、高效的博客平台，满足了不同用户的写作和分享需求。无论是个人博客还是团队协作，Kratos-blog 都是一个很好的选择。
