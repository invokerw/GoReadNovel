package spider

type Note struct {
	Index         int    `json:"index"`    //索引
	NoteName      string `json:"notename"` //章节名称
	NoteUrl       string `json:"url"`      //地址
	LatestChpName string `json:"newchp"`   //最新章节名字
	LatestChpUrl  string `json:"newurl"`   //最新章节地址
	Author        string `json:"author"`   //作者
	Status        string `json:"status"`   //状态连载还是完结
}

const (
	URL = "http://www.huanyue123.com"
)
