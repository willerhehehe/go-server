# README
go Demo Server

## 目录说明
1. cmd：编译入口
   1. server.main：http服务启动入口
2. internal：项目内部，不允许其他项目调用
   1. config: 项目内部配置
   2. domain: 业务模型相关（模型自身的业务逻辑放在这里，仅包含业务模型自身的能力）（可依赖于抽象，不依赖于具体的第三方工具）
   3. service: 业务服务相关（核心业务逻辑在这里，通常从repo中取出domain模型，执行业务操作后再将domain存回repo）
   4. inbound: 外部调用内部服务（提供给外部的接口），包块但不限于http，rpc，cli等
   5. outbound: 内部调用外部服务（第三方接口的客户端），包括数据库等
      1. repo：数据仓库（任何可以存储数据的地方），包括但不限于db，文件系统等。
3. pkg：公共模块，允许其他项目调用
4. test：测试

## 设计原则
1. 模块功能清晰明确，符合go标准
2. 依赖倒置：解耦核心模块domain及service对外部的依赖（将对原本外部具体功能的依赖，转为外部具体功能对domain及service内interface的依赖）
3. 依赖注入：解决2的实现问题，依赖通过inbound接口处组合注入

## 额外功能
1. http服务支持自动生成OpenApi标准文档(考虑直接使用go-restful)
2. http服务支持/docs及/redoc直接访问web文档
3. http服务Graceful stop
4. http服务PProf
5. http服务监控（待开发）
7. http服务基于uri的权限管理（待开发）

## 标准开发过程
1. 设计domain及service功能，外部依赖通过domain及service内interface定义
2. 写针对domain及service改功能的单元测试
3. 开发outbound，实现外部依赖的interface
4. 开发inbound，提供外部接口，及写文档注释
5. 测试inbound外部接口（最好具有单元测试）
