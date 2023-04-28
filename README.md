## 基于goravel的高性能短连接跳转服务
相关组件   
golang >= 1.20   
redis   
mysql 8.0

### short id设置
如果你想产出的ID看起来更习惯一点，但是short id的长度相当于来说也更长   
`alter table jump_link AUTO_INCREMENT = 10000000;`

### sql文件
`sql\short-link-jump.sql`