-- 基础信息
local base_info = {
	group_id = 199001188
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
	{ config_id = 188001, monster_id = 23030101, pos = { x = 691.360, y = 189.493, z = 323.043 }, rot = { x = 0.000, y = 350.373, z = 0.000 }, level = 20, drop_tag = "召唤师", disableWander = true, area_id = 402 },
	{ config_id = 188002, monster_id = 23010601, pos = { x = 687.686, y = 189.566, z = 326.252 }, rot = { x = 0.000, y = 308.016, z = 0.000 }, level = 20, drop_tag = "先遣队", pose_id = 9001, area_id = 402 },
	{ config_id = 188003, monster_id = 21010401, pos = { x = 675.689, y = 173.483, z = 327.153 }, rot = { x = 0.000, y = 36.520, z = 0.000 }, level = 20, drop_tag = "远程丘丘人", disableWander = true, pose_id = 9003, area_id = 402 },
	{ config_id = 188004, monster_id = 23010601, pos = { x = 693.294, y = 180.543, z = 306.994 }, rot = { x = 0.000, y = 132.924, z = 0.000 }, level = 20, drop_tag = "先遣队", disableWander = true, area_id = 402 }
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
		monsters = { 188001, 188002, 188003, 188004 },
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