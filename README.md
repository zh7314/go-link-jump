## 基于goravel的高性能短连接跳转服务
相关组件   
golang >= 1.20   
redis   
mysql 8.0

### hash_key设置
可能会出现hash碰撞，如果系统请求量大，请更换其他算法


### 提供以下几个通用方法
- 请求统一返回
- 业务层panic业务处理
- token中间件

### sql文件
`doc\sql\short-link-jump.sql`