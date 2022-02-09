package define

const ONE_MIN_SECOND = 60
const ONE_HOUR_MIN = 60
const ONE_HOUR_SECOND = ONE_MIN_SECOND * ONE_HOUR_MIN
const ONE_DAY_SECOND = ONE_HOUR_SECOND * 24
const ONE_WEEK_SECOND = ONE_DAY_SECOND * 7
const ONE_MONTH_SECOND = ONE_DAY_SECOND * 30

const MAX_TABLE_PLAYER_NUM = 2

//万分比是比例的默认基数
const DEFAULT_RATIO_BASE = 10000

const (
	ARENA_JUNIOR = iota + 1
	ARENA_MIDDLE
	ARENA_HIGH
	ARENA_PRO
	ARENA_SPECIAL
)

type ItemType int32

//物品类型
const (
	ITEM_TYPE_BUFF ItemType = iota + 1
	ITEM_TYPE_BLACK_HOLE
	ITEM_TYPE_POKER
	//转盘
	ITEM_TYPE_ROULETTE
	//宝箱(卡牌)
	ITEM_TYPE_TRESURE
)

//炮台形态
const (
	FORM_TYPE_NORMAL = iota
	FORM_TYPE_FLY
)

//退出类型
type LogOutType int32

const (
	LOG_OUT_TYPE_CLIENT        LogOutType = 0
	LOG_OUT_TYPE_BEKICK        LogOutType = 1
	LOG_OUT_TYPE_SOCKET_CLOSED LogOutType = 2
	LOG_OUT_TYPE_REPEAT_LOGIN  LogOutType = 3
	LOG_OUT_TYPE_LOGIN_FAIL    LogOutType = 4
)

//子弹类型
const (
	//普通
	BULLET_TYPE_NORMAL = iota
	//免费
	BULLET_TYPE_FREE
	//黑洞
	BULLET_TYPE_BLACKHOLE
	//爆炸
	BULLET_TYPE_EXPLODE
	//连锁
	BULLET_TYPE_CHAIN
	//武器
	BULLET_TYPE_WEAPON
	//机关枪
	BULLET_TYPE_GUN
)

//buff效果类型
const (
	//散射
	EFFECT_SPLINTER = iota + 1
	//闪电
	EFFECT_LIGHTNING
	//暴击
	EFFECT_CRIT
	//加速
	EFFECT_SPEEDUP
)

//触发子弹类型
const (
	TRIGGER_BULLET_TYPE_SPLINTER = iota + 1
	TRIGGER_BULLET_TYPE_LIGHTNING
)

//小数 用于比较浮点型数
const SMALL_NUM = 0.00001

//怪物移动检测间隔
const MONSTER_BASE_SPEED = 0.1

//定时器类型
type TableTimerType int32

const (
	TIMER_TYPE_CREATE_MONSTER TableTimerType = iota + 1
	TIMER_TYPE_MONSTER_MOVE
	TIMER_TYPE_SAVE_PLAYER_INFO
	TIMER_TYPE_KICK_TIMEOUT
	TIMER_TYPE_PLAYER_NOTIFY
	TIMER_TYPE_SYNC_FREEZE_BULLET
)

//怪物刷新间隔毫秒
const MONSTER_CREATE_DURATION_MS = 1000

//怪物移动间隔毫秒
const MONSTER_MOVE_DURATION_MS = 100

const (
	//未命中子弹返还
	GAIN_TYPE_MISS_BULLET = iota + 1
	//退出时剩余子弹、buff等结算
	GAIN_TYPE_SETTLE_ON_EXIT
	//切换后台结算
	GAIN_TYPE_SETTLE_ON_BACKGROUND
	//玩家充值
	GAIN_TYPE_RECHARGE
	//购买比赛加成消耗
	GAIN_TYPE_BUY_MATCH_GIFT
	//比赛奖励金币
	GAIN_TYPE_MATCH_REWARD_MONEY
)

type StateType int32

//状态类型
const (
	//buff
	STATE_TYPE_BUFF StateType = iota + 1
	//黑洞
	STATE_TYPE_BLACK_HOLE
	//武器
	STATE_TYPE_WEAPON
	//特殊武器
	STATE_TYPE_SPEC_WEAPON
	//机关枪
	STATE_TYPE_GUN
)

const ROOM_CENTER_HEART_BEAT_MS = 5 * 1000
const ROOM_CENTER_REPORT_MS = 60 * 1000
const RELOAD_GIVEN_CONFIG_MS = 60 * 1000
const RELOAD_PERSONAL_POOL_CONFIG_MS = 60 * 1000
const CLEAR_LOGIN_FAIL_SESSION_MS = 5 * 1000

//切后台时间长了把身上buff清掉
const CLEAR_BUFF_TIME_MS = 30 * 1000

//切后台30s清除特殊武器
const CLEAR_SPEC_WEAPON_MS = 30 * 1000

const CLEAR_GUN_WEAPON_MS = 30 * 1000

//切后台30s清除该玩家所有动画
const CLEAR_ANIMATION_MS =  30 * 1000

//大概是1200多的长度 控制单包长度不超MTU
const MONSTER_PATH_SPLIT_LEN = 30

const (
	//捕获普通怪物
	MSG_ID_CAPTURE_NORMAL_MONSTER = 1
	//轮盘
	MSG_ID_CAPTURE_ROULETTE = 2
	//集装箱
	MSG_ID_CAPTURE_BUFF_BOX = 3
	//女王龙
	MSG_ID_CAPTURE_BUFF_DRAGON = 4
	//刮刮卡
	MSG_ID_USE_SCRATCH_CARD = 5
	//扑克
	MSG_ID_CAPTURE_POKER = 6
	//黑洞
	MSG_ID_CAPTURE_BLACK_HOLE = 7
	//机械鲨
	MSG_ID_CAPTURE_EXPLODE = 8
	//使用火箭炮
	MSG_ID_USE_MISSILE = 9
	//个人奖池
	MSG_ID_DRAW_PERSONAL_POOL = 10
	//武器收益
	MSG_ID_WEAPON_ERAN = 11
	//掉落卡牌 - 筹码
	MSG_ID_TREASURE_MONEY = 12
	//掉落卡牌 - 碎片
	MSG_ID_TREASURE_PIECE = 13
	//掉落卡牌 - 火箭炮
	MSG_ID_TREASURE_MISSILE = 14
	//蝎子固定奖励
	MSG_ID_CAPTURE_TREASURE = 15

	//发放红包
	MSG_ID_GRANT_RED_ENVELOPE = 16
	//开启公共红包
	MSG_ID_OPEN_COMMON_RED_ENVELOPE = 17

	//星际飞船特殊武器
	MSG_ID_SPEC_WEAPON_EARN = 18

	//破坏球
	MSG_ID_GUN = 19
	//幸运轮
	MSG_ID_FORTUNE = 20
)

type GiveType int32

const (
	//微信首次登录暗送
	GIVE_TYPE_WECHAT GiveType = 1
	//充值暗送
	GIVE_TYPE_RECHARGE GiveType = 2
	//损失价值暗送
	GIVE_TYPE_LOST    GiveType = 3
	GIVE_TYPE_INVALID GiveType = 4
)

type LoginType int32

const (
	LOGIN_TYPE_ACCOUNT LoginType = 0
	LOGIN_TYPE_WECHAT  LoginType = 1
	LOGIN_TYPE_PHONE   LoginType = 2
)

//检查玩家操作超时间隔
const CHECK_KICK_TABLE_PLAYER_SEC = 1
const PLAYER_NOTIFY_SEC = 3

type LoggerType string

const (
	LOGGER_TYPE_DEFAULT      LoggerType = "main"
	LOGGER_TYPE_PLAYER       LoggerType = "player"
	LOGGER_TYPE_GLOBAL       LoggerType = "global"
	LOGGER_TYPE_MONSTER      LoggerType = "monster"
	LOGGER_TYPE_TEST         LoggerType = "test"
	LOGGER_TYPE_PLAYER_EXTRA LoggerType = "player_extra"
	LOGGER_TYPE_OPERATION    LoggerType = "operation"
	LOGGER_TYPE_BIG_EARN     LoggerType = "big_earn"
	LOGGER_TYPE_CLIENT       LoggerType = "client_event"
)

//奖票价值系数
const TICKET_VALUE_FACTOR = 12

const (
	PIGGY_STATUS_NO_REWARD = iota
	PIGGY_STATUS_HAS_REWARD
	PIGGY_STATUS_FULL
)

//武器使用记录 redis key
const WEAPON_RECORD_LIST_KEY = "weapon_record"

const (
	TREASURE_SHOW_TYPE_NORMAL         = 1
	TREASURE_SHOW_TYPE_SPECIAL_EFFECT = 2
)

const (
	TREASURE_REWARD_TYPE_GOLD = iota + 1
	TREASURE_REWARD_TYPE_PIECE
	TREASURE_REWARD_TYPE_MISSILE
)

// 特殊模式类型
const (
	NORMAL_MODE = iota + 1
	ADVANCED_MODE
)

// 特殊武器
const (
	WEAPON_FIST = iota + 6
	WEAPON_STROM
	WEAPON_LASER
)

// 阻止生成的怪物
const (
	CRAFT_MONSTER = 4000
	GUN_MONSTER   = 5000
	FORTUNE_MONSTER = 6000
)

//红包活动任务
const (
	TaskStateFinish    = 1 //已完成可领取奖励
	TaskStateNotFinish = 2 //未完成
	TaskStateEnd       = 3 //奖励已领取，任务结束
	TaskNotBegin       = 4 //任务仍未开始
)

// 任务类型
const (
	TaskTypeMonster = iota + 1
	TaskTypeUseProp
	TaskTypeBuyProp
)

// 炮台副作用类型
const (
	BatteryCostMoney = iota + 1
)

// 炮台特殊效果
const (
	BatteryHitRate = iota + 1
	BatteryEarnMoney
)

// 动画id，要跟客户端约定
const (
	FortuneAnimation = 1
)