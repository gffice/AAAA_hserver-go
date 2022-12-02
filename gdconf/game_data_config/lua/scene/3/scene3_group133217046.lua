-- 基础信息
local base_info = {
	group_id = 133217046
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
	{ config_id = 46001, gadget_id = 70900387, pos = { x = -4339.646, y = 200.908, z = -3794.887 }, rot = { x = 0.000, y = 244.503, z = 0.000 }, level = 30, isOneoff = true, persistent = true, area_id = 14 },
	{ config_id = 46002, gadget_id = 70900387, pos = { x = -4343.084, y = 200.858, z = -3791.616 }, rot = { x = 0.000, y = 214.176, z = 25.129 }, level = 30, isOneoff = true, persistent = true, area_id = 14 },
	{ config_id = 46003, gadget_id = 70900387, pos = { x = -4346.151, y = 201.567, z = -3792.688 }, rot = { x = 354.870, y = 29.296, z = 342.428 }, level = 30, isOneoff = true, persistent = true, area_id = 14 },
	{ config_id = 46004, gadget_id = 70900387, pos = { x = -4340.779, y = 199.939, z = -3796.834 }, rot = { x = 344.917, y = 32.756, z = 4.134 }, level = 30, isOneoff = true, persistent = true, area_id = 14 }
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
		monsters = { },
		gadgets = { 46001, 46002, 46003, 46004 },
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