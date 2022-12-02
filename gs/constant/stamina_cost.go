package constant

var StaminaCostConst *StaminaCost

type StaminaCost struct {
	// 消耗耐力
	CLIMBING_BASE   int32 // 缓慢攀爬基数
	CLIMB_START     int32 // 攀爬开始
	CLIMB_JUMP      int32 // 攀爬跳跃
	DASH            int32 // 快速跑步
	FLY             int32 // 滑翔
	SKIFF_DASH      int32 // 浪船加速
	SPRINT          int32 // 冲刺
	SWIM_DASH_START int32 // 快速游泳开始
	SWIM_DASH       int32 // 快速游泳
	SWIMMING        int32 // 缓慢游泳
	// 恢复耐力
	POWERED_FLY   int32 // 滑翔加速(风圈等)
	POWERED_SKIFF int32 // 浪船加速(风圈等)
	RUN           int32 // 正常跑步
	SKIFF         int32 // 游艇行驶
	STANDBY       int32 // 站立
	WALK          int32 // 走路
}

func InitStaminaCostConst() {
	StaminaCostConst = new(StaminaCost)

	StaminaCostConst.CLIMBING_BASE = -100
	StaminaCostConst.CLIMB_START = -500
	StaminaCostConst.CLIMB_JUMP = -2500
	StaminaCostConst.DASH = -360
	StaminaCostConst.FLY = -60
	StaminaCostConst.SKIFF_DASH = -204
	StaminaCostConst.SPRINT = -1800
	StaminaCostConst.SWIM_DASH_START = -2000
	StaminaCostConst.SWIM_DASH = -204
	StaminaCostConst.SWIMMING = -400
	StaminaCostConst.POWERED_FLY = 500
	StaminaCostConst.POWERED_SKIFF = 500
	StaminaCostConst.RUN = 500
	StaminaCostConst.SKIFF = 500
	StaminaCostConst.STANDBY = 500
	StaminaCostConst.WALK = 500
}
