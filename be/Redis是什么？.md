# Redis 是什么？
Redis 就像是一个超快的「临时笔记本」，专门帮程序记住一些需要快速存取的数据。

### 1. 内存优先 = 闪电速度
- 传统数据库（如MySQL）：像从硬盘文件柜里翻资料，虽然能存大量数据，但速度慢。
- Redis：直接放在内存里，读写速度是微秒级，适合高频访问的场景（比如缓存用户会话）。
- 内存：像一个超级小的电脑，速度超快，但容量有限。
- Redis：像一个超级大的电脑，容量无限，但速度慢。

### 2. 适用 vs 不适用
试用 Redis：
- 高频读写的热点数据（如社交App的点赞数）
- 需要复杂数据结构的场景（如排行榜、限流器）

不适用 Redis：
- 大数据量持久化存储（用MySQL或MongoDB）
- 复杂关联查询（需要关系型数据库）

> Redis 和数据库（如 MySQL、PostgreSQL）的交互通常是为了互补短板——Redis 负责高频读写的热点数据，数据库负责持久化存储和复杂查询。

### 3. 读取流程

#### 3.1 读取流程（缓存优先）
```plaintext
用户请求 → 读 Redis → 是否存在？ → 是：直接返回数据
                          ↓ 否
                          读数据库 → 数据存入 Redis（下次命中缓存） → 返回数据
```

场景举例：用户查询商品详情页，优先从 Redis 缓存读取商品信息，若不存在则查数据库并回填缓存。
代码逻辑：
```python
def get_product(product_id):
    # 1. 先查 Redis
    data = redis.get(f"product:{product_id}")
    if data:
        return data
    # 2. 查数据库
    data = db.query("SELECT * FROM products WHERE id = ?", product_id)
    # 3. 回填 Redis（设置过期时间，避免冷数据堆积）
    redis.setex(f"product:{product_id}", 3600, data)
    return data

```

#### 3.2 写入数据（保持一致性）
```plaintext
用户写请求 → 更新数据库 → 删除/更新 Redis 对应缓存 → 返回成功
```
关键点：先更新数据库，再操作缓存（避免脏数据）。

场景举例：用户修改个人昵称，先更新数据库，再让 Redis 缓存失效（下次读取时自动回填新数据）。

代码逻辑：
```python
def update_nickname(user_id, new_nickname):
    # 1. 更新数据库
    db.execute("UPDATE users SET name = ? WHERE id = ?", new_name, user_id)
    # 2. 删除 Redis 缓存（下次读取时会自动回填）
    redis.delete(f"user:{user_id}")
```



#### 总结
- Redis 是缓存层：负责高频读写，减轻数据库压力。

- 数据库是持久层：负责最终数据存储和复杂查询。

- 交互核心：通过缓存回填、失效策略、异步同步等手段平衡性能与一致性。