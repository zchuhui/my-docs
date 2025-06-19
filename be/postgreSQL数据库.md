# postgreSQL 数据库使用

## 安装（mac）

```cmd
# 安装 PostgreSQL
brew install postgresql
```

## 启动/停止服务

```cmd
# 启动
pg_ctl -D /usr/local/var/postgres -l /usr/local/var/postgres/server.log start

# 停止
pg_ctl -D /usr/local/var/postgres stop -s -m fast

# 或者使用 Homebrew 管理服务
brew services start postgresql
brew services stop postgresql
```

## 连接

```cmd
# 连接 PostgreSQL
psql -U postgres
```

