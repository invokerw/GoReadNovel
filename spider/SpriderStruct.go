package spider

type Novel struct {
	Index         int    `json:"index"`     //索引
	NovelName     string `json:"novelname"` //章节名称
	NovelUrl      string `json:"url"`       //地址
	NovelType     string `json:"noveltype"` //小说类型  这里这个没用用到过
	LatestChpName string `json:"newchp"`    //最新章节名字
	LatestChpUrl  string `json:"newurl"`    //最新章节地址
	Author        string `json:"author"`    //作者
	Status        string `json:"status"`    //状态连载还是完结
	Desc          string `json:"desc"`      //描述
}

const (
	URL = "http://www.huanyue123.com"
)
