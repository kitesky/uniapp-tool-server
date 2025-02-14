package consts

const (
	TaskRewardTypeScore   string = "score"   // 积分
	TaskRewardTypePoint   string = "point"   // 点数
	TaskRewardTypeBalance string = "balance" // 余额
	TaskRewardTypeReward  string = "reward"  // 奖励金
)

// 任务code
const (
	TaskDailySignIn      string = "check_in:daily"      // 每日签到
	TaskPostRelease      string = "post:release"        // 发布帖子
	TaskWechatSubscribe  string = "wechat:subscribe"    // 关注公众号
	TaskSettingEmail     string = "setting:email"       // 设置邮箱
	TaskUserRegister     string = "user:register"       // 注册
	TaskToolUsedTextOnce string = "tool:used-text:once" // 文本工具使用一次
	TaskUserInvite       string = "user:invite"         // 邀请用户
	TaskVIPGiftDaily     string = "vip:gift:daily"      // 会员每日福利
)
