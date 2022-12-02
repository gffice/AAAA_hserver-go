-- 基础信息
local base_info = {
	group_id = 133007016
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
	{ config_id = 32, monster_id = 20010901, pos = { x = 2866.177, y = 208.338, z = 138.884 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 30, drop_tag = "大史莱姆", disableWander = true, area_id = 4 },
	{ config_id = 33, monster_id = 20010801, pos = { x = 2875.191, y = 206.740, z = 134.159 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 30, drop_tag = "史莱姆", area_id = 4 },
	{ config_id = 34, monster_id = 20010801, pos = { x = 2863.074, y = 209.527, z = 134.689 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 30, drop_tag = "史莱姆", area_id = 4 }
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
}

-- 区域
regions = {
}

-- 触发器
triggers = {
}

-- 变量
variables = {
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
	rand_suite = true
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
		monsters = { 32, 33, 34 },
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