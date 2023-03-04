package constant

const (
	QUEST_STATE_NONE       = 0
	QUEST_STATE_UNSTARTED  = 1
	QUEST_STATE_UNFINISHED = 2
	QUEST_STATE_FINISHED   = 3
	QUEST_STATE_FAILED     = 4
)

const (
	QUEST_LOGIC_TYPE_NONE              = 0
	QUEST_LOGIC_TYPE_AND               = 1
	QUEST_LOGIC_TYPE_OR                = 2
	QUEST_LOGIC_TYPE_NOT               = 3
	QUEST_LOGIC_TYPE_A_AND_ETCOR       = 4
	QUEST_LOGIC_TYPE_A_AND_B_AND_ETCOR = 5
	QUEST_LOGIC_TYPE_A_OR_ETCAND       = 6
	QUEST_LOGIC_TYPE_A_OR_B_OR_ETCAND  = 7
	QUEST_LOGIC_TYPE_A_AND_B_OR_ETCAND = 8
)

const (
	QUEST_ACCEPT_COND_TYPE_NONE                           = 0
	QUEST_ACCEPT_COND_TYPE_STATE_EQUAL                    = 1
	QUEST_ACCEPT_COND_TYPE_STATE_NOT_EQUAL                = 2
	QUEST_ACCEPT_COND_TYPE_PACK_HAVE_ITEM                 = 3
	QUEST_ACCEPT_COND_TYPE_AVATAR_ELEMENT_EQUAL           = 4
	QUEST_ACCEPT_COND_TYPE_AVATAR_ELEMENT_NOT_EQUAL       = 5
	QUEST_ACCEPT_COND_TYPE_AVATAR_CAN_CHANGE_ELEMENT      = 6
	QUEST_ACCEPT_COND_TYPE_CITY_LEVEL_EQUAL_GREATER       = 7
	QUEST_ACCEPT_COND_TYPE_ITEM_NUM_LESS_THAN             = 8
	QUEST_ACCEPT_COND_TYPE_DAILY_TASK_START               = 9
	QUEST_ACCEPT_COND_TYPE_OPEN_STATE_EQUAL               = 10
	QUEST_ACCEPT_COND_TYPE_DAILY_TASK_OPEN                = 11
	QUEST_ACCEPT_COND_TYPE_DAILY_TASK_REWARD_CAN_GET      = 12
	QUEST_ACCEPT_COND_TYPE_DAILY_TASK_REWARD_RECEIVED     = 13
	QUEST_ACCEPT_COND_TYPE_PLAYER_LEVEL_REWARD_CAN_GET    = 14
	QUEST_ACCEPT_COND_TYPE_EXPLORATION_REWARD_CAN_GET     = 15
	QUEST_ACCEPT_COND_TYPE_IS_WORLD_OWNER                 = 16
	QUEST_ACCEPT_COND_TYPE_PLAYER_LEVEL_EQUAL_GREATER     = 17
	QUEST_ACCEPT_COND_TYPE_SCENE_AREA_UNLOCKED            = 18
	QUEST_ACCEPT_COND_TYPE_ITEM_GIVING_ACTIVED            = 19
	QUEST_ACCEPT_COND_TYPE_ITEM_GIVING_FINISHED           = 20
	QUEST_ACCEPT_COND_TYPE_IS_DAYTIME                     = 21
	QUEST_ACCEPT_COND_TYPE_CURRENT_AVATAR                 = 22
	QUEST_ACCEPT_COND_TYPE_CURRENT_AREA                   = 23
	QUEST_ACCEPT_COND_TYPE_QUEST_VAR_EQUAL                = 24
	QUEST_ACCEPT_COND_TYPE_QUEST_VAR_GREATER              = 25
	QUEST_ACCEPT_COND_TYPE_QUEST_VAR_LESS                 = 26
	QUEST_ACCEPT_COND_TYPE_FORGE_HAVE_FINISH              = 27
	QUEST_ACCEPT_COND_TYPE_DAILY_TASK_IN_PROGRESS         = 28
	QUEST_ACCEPT_COND_TYPE_DAILY_TASK_FINISHED            = 29
	QUEST_ACCEPT_COND_TYPE_ACTIVITY_COND                  = 30
	QUEST_ACCEPT_COND_TYPE_ACTIVITY_OPEN                  = 31
	QUEST_ACCEPT_COND_TYPE_DAILY_TASK_VAR_GT              = 32
	QUEST_ACCEPT_COND_TYPE_DAILY_TASK_VAR_EQ              = 33
	QUEST_ACCEPT_COND_TYPE_DAILY_TASK_VAR_LT              = 34
	QUEST_ACCEPT_COND_TYPE_BARGAIN_ITEM_GT                = 35
	QUEST_ACCEPT_COND_TYPE_BARGAIN_ITEM_EQ                = 36
	QUEST_ACCEPT_COND_TYPE_BARGAIN_ITEM_LT                = 37
	QUEST_ACCEPT_COND_TYPE_COMPLETE_TALK                  = 38
	QUEST_ACCEPT_COND_TYPE_NOT_HAVE_BLOSSOM_TALK          = 39
	QUEST_ACCEPT_COND_TYPE_IS_CUR_BLOSSOM_TALK            = 40
	QUEST_ACCEPT_COND_TYPE_QUEST_NOT_RECEIVE              = 41
	QUEST_ACCEPT_COND_TYPE_QUEST_SERVER_COND_VALID        = 42
	QUEST_ACCEPT_COND_TYPE_ACTIVITY_CLIENT_COND           = 43
	QUEST_ACCEPT_COND_TYPE_QUEST_GLOBAL_VAR_EQUAL         = 44
	QUEST_ACCEPT_COND_TYPE_QUEST_GLOBAL_VAR_GREATER       = 45
	QUEST_ACCEPT_COND_TYPE_QUEST_GLOBAL_VAR_LESS          = 46
	QUEST_ACCEPT_COND_TYPE_PERSONAL_LINE_UNLOCK           = 47
	QUEST_ACCEPT_COND_TYPE_CITY_REPUTATION_REQUEST        = 48
	QUEST_ACCEPT_COND_TYPE_MAIN_COOP_START                = 49
	QUEST_ACCEPT_COND_TYPE_MAIN_COOP_ENTER_SAVE_POINT     = 50
	QUEST_ACCEPT_COND_TYPE_CITY_REPUTATION_LEVEL          = 51
	QUEST_ACCEPT_COND_TYPE_CITY_REPUTATION_UNLOCK         = 52
	QUEST_ACCEPT_COND_TYPE_LUA_NOTIFY                     = 53
	QUEST_ACCEPT_COND_TYPE_CUR_CLIMATE                    = 54
	QUEST_ACCEPT_COND_TYPE_ACTIVITY_END                   = 55
	QUEST_ACCEPT_COND_TYPE_COOP_POINT_RUNNING             = 56
	QUEST_ACCEPT_COND_TYPE_GADGET_TALK_STATE_EQUAL        = 57
	QUEST_ACCEPT_COND_TYPE_AVATAR_FETTER_GT               = 58
	QUEST_ACCEPT_COND_TYPE_AVATAR_FETTER_EQ               = 59
	QUEST_ACCEPT_COND_TYPE_AVATAR_FETTER_LT               = 60
	QUEST_ACCEPT_COND_TYPE_NEW_HOMEWORLD_MOUDLE_UNLOCK    = 61
	QUEST_ACCEPT_COND_TYPE_NEW_HOMEWORLD_LEVEL_REWARD     = 62
	QUEST_ACCEPT_COND_TYPE_NEW_HOMEWORLD_MAKE_FINISH      = 63
	QUEST_ACCEPT_COND_TYPE_HOMEWORLD_NPC_EVENT            = 64
	QUEST_ACCEPT_COND_TYPE_TIME_VAR_GT_EQ                 = 65
	QUEST_ACCEPT_COND_TYPE_TIME_VAR_PASS_DAY              = 66
	QUEST_ACCEPT_COND_TYPE_HOMEWORLD_NPC_NEW_TALK         = 67
	QUEST_ACCEPT_COND_TYPE_PLAYER_CHOOSE_MALE             = 68
	QUEST_ACCEPT_COND_TYPE_HISTORY_GOT_ANY_ITEM           = 69
	QUEST_ACCEPT_COND_TYPE_LEARNED_RECIPE                 = 70
	QUEST_ACCEPT_COND_TYPE_LUNARITE_REGION_UNLOCKED       = 71
	QUEST_ACCEPT_COND_TYPE_LUNARITE_HAS_REGION_HINT_COUNT = 72
	QUEST_ACCEPT_COND_TYPE_LUNARITE_COLLECT_FINISH        = 73
	QUEST_ACCEPT_COND_TYPE_LUNARITE_MARK_ALL_FINISH       = 74
	QUEST_ACCEPT_COND_TYPE_NEW_HOMEWORLD_SHOP_ITEM        = 75
	QUEST_ACCEPT_COND_TYPE_SCENE_POINT_UNLOCK             = 76
	QUEST_ACCEPT_COND_TYPE_SCENE_LEVEL_TAG_EQ             = 77
)

const (
	QUEST_FINISH_COND_TYPE_NONE                              = 0
	QUEST_FINISH_COND_TYPE_KILL_MONSTER                      = 1
	QUEST_FINISH_COND_TYPE_COMPLETE_TALK                     = 2
	QUEST_FINISH_COND_TYPE_MONSTER_DIE                       = 3
	QUEST_FINISH_COND_TYPE_FINISH_PLOT                       = 4
	QUEST_FINISH_COND_TYPE_OBTAIN_ITEM                       = 5
	QUEST_FINISH_COND_TYPE_TRIGGER_FIRE                      = 6
	QUEST_FINISH_COND_TYPE_CLEAR_GROUP_MONSTER               = 7
	QUEST_FINISH_COND_TYPE_NOT_FINISH_PLOT                   = 8
	QUEST_FINISH_COND_TYPE_ENTER_DUNGEON                     = 9
	QUEST_FINISH_COND_TYPE_ENTER_MY_WORLD                    = 10
	QUEST_FINISH_COND_TYPE_FINISH_DUNGEON                    = 11
	QUEST_FINISH_COND_TYPE_DESTROY_GADGET                    = 12
	QUEST_FINISH_COND_TYPE_OBTAIN_MATERIAL_WITH_SUBTYPE      = 13
	QUEST_FINISH_COND_TYPE_NICK_NAME                         = 14
	QUEST_FINISH_COND_TYPE_WORKTOP_SELECT                    = 15
	QUEST_FINISH_COND_TYPE_SEAL_BATTLE_RESULT                = 16
	QUEST_FINISH_COND_TYPE_ENTER_ROOM                        = 17
	QUEST_FINISH_COND_TYPE_GAME_TIME_TICK                    = 18
	QUEST_FINISH_COND_TYPE_FAIL_DUNGEON                      = 19
	QUEST_FINISH_COND_TYPE_LUA_NOTIFY                        = 20
	QUEST_FINISH_COND_TYPE_TEAM_DEAD                         = 21
	QUEST_FINISH_COND_TYPE_COMPLETE_ANY_TALK                 = 22
	QUEST_FINISH_COND_TYPE_UNLOCK_TRANS_POINT                = 23
	QUEST_FINISH_COND_TYPE_ADD_QUEST_PROGRESS                = 24
	QUEST_FINISH_COND_TYPE_INTERACT_GADGET                   = 25
	QUEST_FINISH_COND_TYPE_DAILY_TASK_COMP_FINISH            = 26
	QUEST_FINISH_COND_TYPE_FINISH_ITEM_GIVING                = 27
	QUEST_FINISH_COND_TYPE_SKILL                             = 107
	QUEST_FINISH_COND_TYPE_CITY_LEVEL_UP                     = 109
	QUEST_FINISH_COND_TYPE_PATTERN_GROUP_CLEAR_MONSTER       = 110
	QUEST_FINISH_COND_TYPE_ITEM_LESS_THAN                    = 111
	QUEST_FINISH_COND_TYPE_PLAYER_LEVEL_UP                   = 112
	QUEST_FINISH_COND_TYPE_DUNGEON_OPEN_STATUE               = 113
	QUEST_FINISH_COND_TYPE_UNLOCK_AREA                       = 114
	QUEST_FINISH_COND_TYPE_OPEN_CHEST_WITH_GADGET_ID         = 115
	QUEST_FINISH_COND_TYPE_UNLOCK_TRANS_POINT_WITH_TYPE      = 116
	QUEST_FINISH_COND_TYPE_FINISH_DAILY_DUNGEON              = 117
	QUEST_FINISH_COND_TYPE_FINISH_WEEKLY_DUNGEON             = 118
	QUEST_FINISH_COND_TYPE_QUEST_VAR_EQUAL                   = 119
	QUEST_FINISH_COND_TYPE_QUEST_VAR_GREATER                 = 120
	QUEST_FINISH_COND_TYPE_QUEST_VAR_LESS                    = 121
	QUEST_FINISH_COND_TYPE_OBTAIN_VARIOUS_ITEM               = 122
	QUEST_FINISH_COND_TYPE_FINISH_TOWER_LEVEL                = 123
	QUEST_FINISH_COND_TYPE_BARGAIN_SUCC                      = 124
	QUEST_FINISH_COND_TYPE_BARGAIN_FAIL                      = 125
	QUEST_FINISH_COND_TYPE_ITEM_LESS_THAN_BARGAIN            = 126
	QUEST_FINISH_COND_TYPE_ACTIVITY_TRIGGER_FAILED           = 127
	QUEST_FINISH_COND_TYPE_MAIN_COOP_ENTER_SAVE_POINT        = 128
	QUEST_FINISH_COND_TYPE_ANY_MANUAL_TRANSPORT              = 129
	QUEST_FINISH_COND_TYPE_USE_ITEM                          = 130
	QUEST_FINISH_COND_TYPE_MAIN_COOP_ENTER_ANY_SAVE_POINT    = 131
	QUEST_FINISH_COND_TYPE_ENTER_MY_HOME_WORLD               = 132
	QUEST_FINISH_COND_TYPE_ENTER_MY_WORLD_SCENE              = 133
	QUEST_FINISH_COND_TYPE_TIME_VAR_GT_EQ                    = 134
	QUEST_FINISH_COND_TYPE_TIME_VAR_PASS_DAY                 = 135
	QUEST_FINISH_COND_TYPE_QUEST_STATE_EQUAL                 = 136
	QUEST_FINISH_COND_TYPE_QUEST_STATE_NOT_EQUAL             = 137
	QUEST_FINISH_COND_TYPE_UNLOCKED_RECIPE                   = 138
	QUEST_FINISH_COND_TYPE_NOT_UNLOCKED_RECIPE               = 139
	QUEST_FINISH_COND_TYPE_FISHING_SUCC                      = 140
	QUEST_FINISH_COND_TYPE_ENTER_ROGUE_DUNGEON               = 141
	QUEST_FINISH_COND_TYPE_USE_WIDGET                        = 142
	QUEST_FINISH_COND_TYPE_CAPTURE_SUCC                      = 143
	QUEST_FINISH_COND_TYPE_CAPTURE_USE_CAPTURETAG_LIST       = 144
	QUEST_FINISH_COND_TYPE_CAPTURE_USE_MATERIAL_LIST         = 145
	QUEST_FINISH_COND_TYPE_ENTER_VEHICLE                     = 147
	QUEST_FINISH_COND_TYPE_SCENE_LEVEL_TAG_EQ                = 148
	QUEST_FINISH_COND_TYPE_LEAVE_SCENE                       = 149
	QUEST_FINISH_COND_TYPE_LEAVE_SCENE_RANGE                 = 150
	QUEST_FINISH_COND_TYPE_IRODORI_FINISH_FLOWER_COMBINATION = 151
	QUEST_FINISH_COND_TYPE_IRODORI_POETRY_REACH_MIN_PROGRESS = 152
	QUEST_FINISH_COND_TYPE_IRODORI_POETRY_FINISH_FILL_POETRY = 153
)
