-- 基础信息
local base_info = {
	group_id = 133307459
}

-- Trigger变量
local defs = {
	gadget_VP = 459001,
	pointarray_id = 330700019,
	minPoint = 1,
	maxPoint = 3
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
	{ config_id = 459001, gadget_id = 70330200, pos = { x = -126.338, y = 56.668, z = 4909.434 }, rot = { x = 0.000, y = 172.034, z = 0.000 }, level = 1, is_use_point_array = true, area_id = 32 }
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
		gadgets = { 459001 },
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

require "V3_0/VisualizationPlantOne"