INSERT INTO db_dhg_base.per_orgs 
    (`name`,`level`,`company_id`,`code`,`show_order`,`grade`,`enable_date`,`parent_origin_id`) 
VALUES 
  ('121212','223321122','IT流程中心',2,1, 'LDC0034', 1, 2, '2025-04-05 00:00:00', '');

-- company_id（集团，都为 1）
-- enable_date 生效日期，格式是 datetime，默认值是当前时间, 格式是 YYYY-MM-DD HH:MM:SS 
-- parent_origin_id 上级组织id（0）
-- name 组织名称
-- level 组织层级（1）
-- code 组织代码（LDC）
-- show_order 显示顺序（1）
-- grade 组织级别（1）

