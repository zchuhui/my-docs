# SQL

## 常用模板

### 1. 基础表结构
```sql
-- 创建表模板
CREATE TABLE `table_name` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `column1` VARCHAR(255) NOT NULL,
  `column2` DATETIME DEFAULT CURRENT_TIMESTAMP,
  `column3` DECIMAL(10,2),
  PRIMARY KEY (`id`),
  INDEX `index_name` (`column1` ASC)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

### 2. 数据操作
```sql
-- 插入数据（单条）
INSERT INTO `table_name` (`column1`, `column2`) 
VALUES ('value1', NOW());

-- 批量插入
INSERT INTO `table_name` (`column1`, `column2`)
VALUES 
  ('value1', '2024-01-01'),
  ('value2', '2024-01-02');

-- 更新数据
UPDATE `table_name`
SET `column1` = 'new_value', `column2` = NOW()
WHERE `id` = 1;

-- 删除数据
DELETE FROM `table_name`
WHERE `id` > 100
LIMIT 10;
```

### 3. 查询模板
```sql
-- 基础查询
SELECT 
  t1.`column1`,
  t2.`column2`,
  COUNT(*) AS total 
FROM `table1` t1
JOIN `table2` t2 ON t1.id = t2.table1_id
WHERE t1.`create_time` BETWEEN '2024-01-01' AND '2024-01-31'
GROUP BY t1.`column1`, t2.`column2`
HAVING total > 10
ORDER BY t1.`column1` DESC
LIMIT 20;
```

### 4. 索引管理
```sql
-- 创建索引
CREATE INDEX `idx_column` ON `table_name` (`column`);

-- 唯一索引
CREATE UNIQUE INDEX `uniq_idx` ON `table_name` (`column1`, `column2`);
```

### 5. 事务模板
```sql
START TRANSACTION;

UPDATE `accounts` 
SET balance = balance - 100 
WHERE user_id = 1;

UPDATE `accounts` 
SET balance = balance + 100 
WHERE user_id = 2;

COMMIT;

-- 异常回滚
ROLLBACK;
```

### 6. 存储过程
```sql
DELIMITER $$

CREATE PROCEDURE `get_user_orders`(IN userId INT)
BEGIN
  SELECT * 
  FROM `orders`
  WHERE `user_id` = userId
  ORDER BY `order_date` DESC;
END$$

DELIMITER ;
```