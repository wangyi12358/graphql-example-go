# Go gin

## 快速开始
安装依赖
```bash
go download

```
启动服务
```bash
make start-service
```
打包服务
```bash
make build-servicez
```

## 生成数据库 Model
首先修改 `gen.yaml` 中的数据库配置，然后执行以下命令
```bash
$ make gen-db-model
```

## Postgres 自增ID
```sql
CREATE SEQUENCE lov_table_id_seq START 1;

ALTER TABLE lov
ALTER COLUMN id SET DEFAULT nextval('lov_table_id_seq');
```

## 支持的功能
- [x] 生成数据库 Model
- [x] Restful API
- [x] GraphQL API
- [x] validator 参数校验
- [ ] 全局 panic 捕获
- [ ] 定时任务