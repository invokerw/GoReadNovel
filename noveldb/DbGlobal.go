package noveldb

import (
	"GoReadNovel/logger"
	"database/sql"
)

var (
	dbhostsip  = "fsnsaber.cn:3306" //IP地址
	dbusername = "root"             //用户名
	dbpassword = "weifei"           //密码
	dbname     = "novel"            //数据库名
)

type Novel struct {
	ID            int    `json:"id"`         //ID
	NovelName     string `json:"novelname"`  //章节名称
	NovelUrl      string `json:"url"`        //地址
	NovelType     string `json:"noveltype"`  //小说类型
	LatestChpName string `json:"newchp"`     //最新章节名字
	LatestChpUrl  string `json:"newurl"`     //最新章节地址
	ImagesAddr    string `json:"imagesaddr"` //封面图片地址
	Author        string `json:"author"`     //作者
	Status        string `json:"status"`     //状态连载还是完结
	Desc          string `json:"desc"`       //描述
	UpdateTime    int64  `json:"uptime"`     //更新时间 时间戳
}

type User struct {
	UserID    string `json:"uid"`       //用户ID
	NikeName  string `json:"nickname"`  //昵称
	Gender    string `json:"gender"`    //性别
	City      string `json:"city"`      //城市
	Province  string `json:"province"`  //省份
	Country   string `json:"country"`   //国家
	AvatarUrl string `json:"avatarurl"` //头像地址
	JoinTime  int64  `json:"jointime"`  //进入本小程序的时间 时间戳
}

type BookShelf struct {
	ShelfID         int    `json:"bookshelfid"`     //书架ID
	UserID          string `json:"userid"`          //用户ID
	NovelID         int    `json:"novelid"`         //小说ID
	ReadChapterName string `json:"readchaptername"` //读到的章节名称
	ReadChapterUrl  string `json:"readchapterurl"`  //读到的章节地址
}

type AllVote struct {
	AllVoteID  int   `json:"allvoteid"` //总推荐ID
	NovelID    int   `json:"novelid"`   //小说ID
	UpdateTime int64 `json:"uptime"`    //更新时间  时间戳
}

type GoodNum struct {
	GoodNumID  int   `json:"goodnumid"` //总收藏ID
	NovelID    int   `json:"novelid"`   //小说ID
	UpdateTime int64 `json:"uptime"`    //更新时间  时间戳
}
type Comment struct {
	CommentID   int    `json:"commentid"`   //评论ID
	UserID      string `json:"userid"`      //用户ID
	NovelID     int    `json:"novelid"`     //小说ID
	Content     string `json:"content"`     //评论内容
	CommentTime int64  `json:"commenttime"` //评论时间
	Zan         int    `json:"zan"`         //赞同的数量
}
type Feedback struct {
	FeedbackID   int    `json:"feedbackid"`   //反馈ID
	UserID       string `json:"userid"`       //用户ID
	FeedbackType int    `json:"feedbacktype"` //反馈类型 0 书籍问题 1 操作问题 2 其他问题
	Content      string `json:"content"`      //反馈内容
	ContactInfo  string `json:"contactinfo"`  //联系邮箱
	Solve        int    `json:"solve"`        //是否解决 0未解决 1解决
	AddTime      int64  `json:"addtime"`      //添加时间  时间戳
}

const (
	DEFAULT_NOVEL_TYPE = "其他"
	DEFAULT_STATUS     = "连载中"
)

var db *sql.DB
var err error
var NovelTypeEtoC map[string]string //汉语拼音 转换为 数据库中的类型  xuanhuna -> 玄幻小说
//var NovelTypeCtoE map[string]string //数据库中的类型 转换为 汉语拼音  玄幻小说 -> xuanhuan  貌似用不上

func init() { //如果改成init()会自动运行
	//func mysqlInit() {
	logger.ALogger().Debug("Init Mysql DB Connect..")
	db, err = sql.Open("mysql", "root:weifei@tcp(ergou.vip:3306)/novel?charset=utf8")
	checkErr(err)
	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(500)
	initMap()
}
func initMap() {
	NovelTypeEtoC = make(map[string]string)

	NovelTypeEtoC["xuanhuan"] = "玄幻小说"
	NovelTypeEtoC["dushi"] = "都市小说"
	NovelTypeEtoC["xianxia"] = "仙侠小说"
	NovelTypeEtoC["yanqing"] = "言情小说"
	NovelTypeEtoC["wangyou"] = "网游小说"
	NovelTypeEtoC["kehuan"] = "科幻小说"
	NovelTypeEtoC["lishi"] = "历史小说"
	NovelTypeEtoC["lingyi"] = "灵异小说"
	NovelTypeEtoC["xuanhuan"] = "其他小说"

}
func checkErr(err error) bool {
	if err != nil {
		logger.ALogger().Error(err)
		return false
	}
	return true
}

func GetMysqlDB() *sql.DB {
	/*if db == nil {
		mysqlInit()
	}*/
	return db
}

func CloseMysqlDB() {
	if db != nil {
		db.Close()
	}
}
