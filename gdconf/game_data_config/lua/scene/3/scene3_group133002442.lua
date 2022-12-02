-- 基础信息
local base_info = {
	group_id = 133002442
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
	{ config_id = 442001, gadget_id = 70220042, pos = { x = 1728.927, y = 244.937, z = -708.259 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 2, area_id = 3 }
}

-- 区域
regions = {
}

-- 触发器
triggers = {
}

-- 点位
points = {
	{ config_id = 442004, pos = { x = 1731.637, y = 246.243, z = -712.788 }, rot = { x = 0.000, y = 315.673, z = 0.000 }, area_id = 3, tag = 2 },
	{ config_id = 442005, pos = { x = 1725.434, y = 246.472, z = -711.265 }, rot = { x = 0.000, y = 59.277, z = 0.000 }, area_id = 3, tag = 2 },
	{ config_id = 442006, pos = { x = 1733.152, y = 244.750, z = -704.834 }, rot = { x = 0.000, y = 232.097, z = 0.000 }, area_id = 3, tag = 2 },
	{ config_id = 442007, pos = { x = 1728.240, y = 244.925, z = -701.740 }, rot = { x = 0.000, y = 186.155, z = 0.000 }, area_id = 3, tag = 2 },
	{ config_id = 442008, pos = { x = 1723.770, y = 244.923, z = -703.787 }, rot = { x = 0.000, y = 151.129, z = 0.000 }, area_id = 3, tag = 2 }
}

-- 变量
variables = {
}

-- 怪物随机池
monster_pools = {
	{ pool_id = 1004, rand_weight = 100 },
	{ pool_id = 1005, rand_weight = 100 },
	{ pool_id = 1006, rand_weight = 100 },
	{ pool_id = 1007, rand_weight = 100 }
}

--================================================================
-- 
-- 初始化配置
-- 
--================================================================

-- 初始化时创建
init_config = {
	suite = 1,
	end_suite = 0,
	rand_suite = false
}

--================================================================
-- 
-- 小组配置
-- 
--================================================================

suites = {
	{
		-- suite_id = 1,
		-- description = ,
		monsters = { },
		gadgets = { },
		regions = { },
		triggers = { },
		rand_weight = 100
	}
}

--================================================================
-- 
-- 触发器
-- 
--================================================================

require "TreasureMapEvent"