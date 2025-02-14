package tools

import "app-api/services/tools/text_generate"

// 任务处理器
var ToolHandlerList = map[string]ToolHandler{
	"tool:abc": text_generate.NewToolABC(),

	// sy-商业
	"tool:yingxiaocehua":      text_generate.NewToolYingXiaoCeHua(),      // 营销策划方案
	"tool:huodongcehua":       text_generate.NewToolHuoDongCeHua(),       // 活动策划方案
	"tool:anlifenxibaogao":    text_generate.NewToolAnLiFenXiBaoGao(),    // 案例分析报告
	"tool:chanpincuxiao":      text_generate.NewToolChanPinCuXiao(),      // 产品促销方案
	"tool:gongsijieshao":      text_generate.NewToolGongSiJieShao(),      // 公司简介
	"tool:chanpinjieshao":     text_generate.NewToolChanPinJieShao(),     // 产品介绍
	"tool:hetongxieyi":        text_generate.NewToolHeTongXieYi(),        // 合同协议
	"tool:diaochawenjuan":     text_generate.NewToolDiaoChaWenJuan(),     // 调查问卷
	"tool:toubiaoshu":         text_generate.NewToolTouBiaoShu(),         // 投标书
	"tool:hangyeyanjiubaogao": text_generate.NewToolHangYeYanJiuBaoGao(), // 行业研究报告
	"tool:chuangyejihuashu":   text_generate.NewToolChuangYeJiHuaShu(),   // 创业计划书
	"tool:zhiyeguihua":        text_generate.NewToolZhiYeGuiHua(),        // 职业规划书
	"tool:yaoqinghan":         text_generate.NewToolYaoQingHan(),         // 邀请函

	// jy-教育
	"tool:jiaoxuesheji":              text_generate.NewToolJiaoXueSheJi(),              // 教学设计
	"tool:banhuizhuti":               text_generate.NewToolBanHuiZhuTi(),               // 班会主题
	"tool:shitishengcheng":           text_generate.NewToolShiTiShengCheng(),           // 试题生成
	"tool:monidabian":                text_generate.NewToolMoNiDaBian(),                // 模拟答辩
	"tool:zhichengpingxuan":          text_generate.NewToolZhiChengPingXuan(),          // 职称评选
	"tool:zhutiyanjiang":             text_generate.NewToolZhuTiYanJiang(),             // 主题演讲
	"tool:xuekejiaoan":               text_generate.NewToolXueKeJiaoAn(),               // 学科教案
	"tool:jiaoxuehuodong":            text_generate.NewToolJiaoXueHuoDong(),            // 教学活动
	"tool:jiaoshipeixunxinde":        text_generate.NewToolJiaoShiPeiXunXinDe(),        // 教师培训心得
	"tool:jiaoyangongzuozongjie":     text_generate.NewToolJiaoYanGongZuoZongJie(),     // 教研工作总结
	"tool:jiaoshishuzhibaogao":       text_generate.NewToolJiaoShiShuZhiBaoGao(),       // 教师述职报告
	"tool:jiaoxuemubiao":             text_generate.NewToolJiaoXueMuBiao(),             // 教学目标
	"tool:jiaoxuejianyi":             text_generate.NewToolJiaoXueJianYi(),             // 教学建议
	"tool:zhaoshengjihuawenan":       text_generate.NewToolZhaoShengJiHuaWenAn(),       // 招生计划文案
	"tool:youxiujiaoshidaibiaofayan": text_generate.NewToolYouXiuJiaoShiDaiBiaoFaYan(), // 优秀教师代表发言
	"tool:jiafangbaogao":             text_generate.NewToolJiaFangBaoGao(),             // 家访报告
	"tool:xinlishudaofangan":         text_generate.NewToolXinLiShuDaoFangAn(),         // 心理疏导方案
	"tool:banjiguanlizhidu":          text_generate.NewToolBanJiGuanLiZhiDu(),          // 班级管理制度

	// 媒体-w
	"tool:pengyouquan":  text_generate.NewToolPengYouQuan(),  // 朋友圈文案
	"tool:xiaohongshu":  text_generate.NewToolXiaoHongShu(),  // 小红书文案
	"tool:jinritoutiao": text_generate.NewToolJinRiTouTiao(), // 今日头条文章
	"tool:gongzhonghao": text_generate.NewToolGongZhongHao(), // 公众号文章
	"tool:zhihu":        text_generate.NewToolZhiHu(),        // 知乎问答
	"tool:xinwen":       text_generate.NewToolXinWen(),       // 新闻稿
	"tool:kepu":         text_generate.NewToolKePu(),         // 科普文案

	// 写作-x
	"tool:manfenzuowen":  text_generate.NewToolManFenZuoWen(),  // 满分作文
	"tool:duilian":       text_generate.NewToolDuiLian(),       // 对联
	"tool:yanjianggao":   text_generate.NewToolYanJiangGao(),   // 演讲稿
	"tool:bianlun":       text_generate.NewToolBianLun(),       // 辩论稿
	"tool:zhuchi":        text_generate.NewToolZhuChi(),        // 主持稿
	"tool:gushichi":      text_generate.NewToolGuShiCi(),       // 古诗词
	"tool:zhengwen":      text_generate.NewToolZhengWen(),      // 征文
	"tool:duhougan":      text_generate.NewToolDuHouGan(),      // 读后感
	"tool:guanhougan":    text_generate.NewToolGuanHouGan(),    // 观后感
	"tool:weiwenxin":     text_generate.NewToolWeiWenXin(),     // 慰问信
	"tool:banjiangci":    text_generate.NewToolBanJiangCi(),    // 颁奖词
	"tool:zhici":         text_generate.NewToolZhiCi(),         // 致辞
	"tool:xiandaishi":    text_generate.NewToolXianDaiShi(),    // 现代诗
	"tool:tuokouxiu":     text_generate.NewToolTuoKouXiu(),     // 脱口秀
	"tool:xiaopinjuben":  text_generate.NewToolXiaoPinJuBen(),  // 小品剧本
	"tool:wenyipinglun":  text_generate.NewToolWenYiPingLun(),  // 文艺评论
	"tool:weixiaoshuo":   text_generate.NewToolWeiXiaoShuo(),   // 微小说
	"tool:geci":          text_generate.NewToolGeCi(),          // 歌词
	"tool:sanwen":        text_generate.NewToolSanWen(),        // 散文
	"tool:minshiqisushu": text_generate.NewToolMinShiQiSuShu(), // 民事起诉书
	"tool:qingshugaobai": text_generate.NewToolQingShuGaoBai(), // 情书告白
	// 个人-y
	"tool:lvxingjihua":  text_generate.NewToolLvXingJiHua(),  // 旅行计划
	"tool:jiantaoshu":   text_generate.NewToolJianTaoShu(),   // 检讨书
	"tool:ziwojieshao":  text_generate.NewToolZiWoJieShao(),  // 自我介绍
	"tool:jianli":       text_generate.NewToolJianLi(),       // 简历
	"tool:shenqingshu":  text_generate.NewToolShenQingShu(),  // 申请书
	"tool:jianyishu":    text_generate.NewToolJianYiShu(),    // 建议书
	"tool:ziwopingjia":  text_generate.NewToolZiWoPingJia(),  // 自我评价
	"tool:zhuhexin":     text_generate.NewToolZhuHeXin(),     // 祝贺信
	"tool:gerenzizhuan": text_generate.NewToolGeRenZiZhuan(), // 个人自传
	// 工作类-z
	"tool:nianzhongzongjie": text_generate.NewToolNianZhongZongJie(), // 年终总结
	"tool:shuzhibaogao":     text_generate.NewToolShuZhiBaoGao(),     //  述职报告
	"tool:gongzuojihua":     text_generate.NewToolGongZuoJiHua(),     // 工作计划
	"tool:xindetihui":       text_generate.NewToolXinDeTiHui(),       // 心得体会
	"tool:zhuanyewenzhang":  text_generate.NewToolZhuanYeWenZhang(),  // 专业文章
	"tool:shiyanbaogao":     text_generate.NewToolShiYanBaoGao(),     // 实验报告
	"tool:shixunbaogao":     text_generate.NewToolShiXunBaoGao(),     // 实训报告
	"tool:shijianbaogao":    text_generate.NewToolShiJianBaoGao(),    // 实践报告
	"tool:shixibaogao":      text_generate.NewToolShiXiBaoGao(),      // 实习报告
	"tool:sixianghuibao":    text_generate.NewToolSiXiangHuiBao(),    // 思想汇报
	"tool:jingxuangao":      text_generate.NewToolJingXuanGao(),      // 竞选稿
	"tool:tuijianxin":       text_generate.NewToolTuiJianXin(),       // 推荐信
	"tool:gongzuozongjie":   text_generate.NewToolGongZuoZongJie(),   // 工作总结
	"tool:gongzuobaogao":    text_generate.NewToolGongZuoBaoGao(),    // 工作报告
	"tool:tongzhi":          text_generate.NewToolTongZhi(),          // 通知
	"tool:huiyijiyao":       text_generate.NewToolHuiYiJiYao(),       // 会议纪要
	"tool:diaoyanbaogao":    text_generate.NewToolDiaoYanBaoGao(),    // 调研报告
	"tool:jinshengzongjie":  text_generate.NewToolJinShengZongJie(),  // 晋升总结
	"tool:zhaopinxuqiu":     text_generate.NewToolZhaoPinXuQiu(),     // 招聘需求
	"tool:gongwen":          text_generate.NewToolGongWen(),          // 公文
	"tool:guizhangzhidu":    text_generate.NewToolGuiZhangZhiDu(),    // 规章制度
}
