package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestFloat(t *testing.T) {
	rs1 := float64(99)
	rs2 := float64(9)
	rs := rs1 / rs2 * 10
	t.Logf("%f", rs)
	t.Logf("%f", Bcdiv(rs1, rs2, 2))
	t.Log(Bcdiv(rs1*10, rs2, 2))
	t.Log(GetDiscount(rs1, rs2))
}

func GetDiscount(showPrice, basePrice float64) float64 {
	round := Bcdiv(showPrice*10, basePrice, 2)
	if round == 0 && showPrice > 0 && basePrice > 0 {
		return 0.1
	} else {
		return round
	}
}

// Bcdiv
// 2个任意精度的数字除法计算
func Bcdiv(x, y float64, scale int) float64 {
	if y == 0 {
		return 0
	}
	rs := x / y

	data := round(rs, scale)
	fmt.Println("data", data)
	inst, _ := strconv.ParseFloat(data, 64)
	return inst
}

// round
func round(f float64, m int) string {
	n := strconv.FormatFloat(f, 'f', -1, 64)
	if n == "" {
		return n
	}
	if m >= len(n) {
		return n
	}
	newn := strings.Split(n, ".")
	if len(newn) < 2 || m >= len(newn[1]) {
		return n
	}
	return newn[0] + "." + newn[1][:m]
}

// FormatFloat64
func FormatFloat64(data interface{}) (float64, error) {
	str, _ := FormatString(data)
	return strconv.ParseFloat(str, 64)
}

// FormatString
// 转换为字符串
func FormatString(data interface{}) (string, error) {
	switch data.(type) {
	case json.Number:
		i, err := data.(json.Number).Int64()
		if err != nil {
			return "", err
		}
		return strconv.FormatInt(i, 10), nil
	case float32, float64:
		fdata := reflect.ValueOf(data).Float()
		return strconv.FormatFloat(fdata, 'f', -1, 64), nil
	case int, int8, int16, int32, int64:
		idata := reflect.ValueOf(data).Int()
		return strconv.FormatInt(idata, 10), nil
	case uint, uint8, uint16, uint32, uint64:
		udata := reflect.ValueOf(data).Uint()
		return strconv.FormatUint(udata, 10), nil
	case string:
		return data.(string), nil
	}
	return "", errors.New("invalid value type")
}

func TestBool(t *testing.T) {
	str := `{
		"code": 0,
		"msg": "ok",
		"body": {
			"albumId": 15,
			"asset": true,
			"programAssetList": [
				{
					"pId": "50000045",
					"asset": false
				},
				{
					"pId": "50000046",
					"asset": false
				},
				{
					"pId": "50000047",
					"asset": false
				}
			]
		}
	}`

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		t.Error(err.Error())
	}
	var mapData map[string]interface{}
	var ok bool
	if mapData, ok = data["body"].(map[string]interface{}); !ok {
		t.Error("断言异常")
	}

	if asset, ok := mapData["asset"].(bool); !ok {
		t.Error("bool 断言异常")
	} else {
		t.Log(asset)
	}

}

func TestTime(t *testing.T) {
	dateTime := secToDuration(-1)
	t.Log(dateTime)
}

func secToDuration(second int64) string {
	result := `00'00"`
	if second > 0 {
		dateStr := strconv.FormatInt(second, 10) + "s"
		dateTime, err := time.ParseDuration(dateStr)
		if err != nil {
			return result
		}
		rs := dateTime.String()
		rs = strings.Replace(rs, "h", "'", 1)
		rs = strings.Replace(rs, "m", "'", 1)
		rs = strings.Replace(rs, "s", `"`, 1)
		return rs
	}
	return result
}

type AuthorInfoList struct {
	AuditStatus  int    `json:"auditStatus"`
	AuthorIcon   string `json:"authorIcon"`
	AuthorId     int    `json:"authorId"`
	AuthorName   string `json:"author_name"`
	CircleId     string `json:"circleId"`
	RelationName string `json:"relation_name"`
	RelationType int    `json:"relation_type"`
	Synopsis     string `json:"synopsis"`
	Usr          string `json:"usr"`
}

type ListenPriceInfo struct {
	Chapter_count    int     `json:"chapter_count"`
	Chapter_price    float64 `json:"chapter_price"`
	Discount         int     `json:"discount"`
	Fee_chapter_num  int     `json:"fee_chapter_num"`
	Fee_type         string  `json:"fee_type"`
	Limit_end        int64   `json:"limit_end"`
	Limit_price      float64 `json:"limit_price"`
	Limit_start      int64   `json:"limit_start"`
	Vip_sub_code     int     `json:"vip_sub_code"`
	Whole_book_price float64 `json:"whole_book_price"`
}

type ListenSectionBook struct {
	Author              string           `json:"author"`
	AuthorInfoList      []AuthorInfoList `json:"authorInfoList"`
	BookId              string           `json:"book_id"`
	BookLevel           string           `json:"book_level"`
	BookLevelName       string           `json:"book_level_name"`
	BookTypeId          string           `json:"book_type_id"`
	BookTypeName        string           `json:"book_type_name"`
	BookTypePicUrl      string           `json:"book_type_pic_url"`
	CTag                string           `json:"c_tag"`
	ChapterCount        string           `json:"chapter_count"`
	CommentCount        string           `json:"comment_count"`
	CommonChaptersUrl   string           `json:"common_chapters_url"`
	ConsumeCoinNum      string           `json:"consume_coin_num"`
	CopyrightBelong     string           `json:"copyright_belong"`
	CopyrightBelongName string           `json:"copyright_belong_name"`
	CopyrightType       string           `json:"copyright_type"`
	CopyrightTypeName   string           `json:"copyright_type_name"`
	CreateTime          string           `json:"create_time"`
	Description         string           `json:"description"`
	FeeDownloadNum      string           `json:"fee_download_num"`
	FeeUserNum          string           `json:"fee_user_num"`
	FileKey             string           `json:"file_key"`
	FinishState         string           `json:"finish_state"`
	FinishStateName     string           `json:"finish_state_name"`
	FreeDownloadNum     string           `json:"free_download_num"`
	FreeUserNum         string           `json:"free_user_num"`
	HasPic              string           `json:"has_pic"`
	IsFree              string           `json:"isFree"`
	IsVip               int              `json:"isVip"`
	IsInner             string           `json:"is_inner"`
	IsPreview           string           `json:"is_preview"`
	IsSingle            string           `json:"is_single"`
	Keywords            string           `json:"keywords"`
	LRecommend          string           `json:"l_recommend"`
	LastChapterTime     string           `json:"last_chapter_time"`
	Name                string           `json:"name"`
	OriPic              string           `json:"ori_pic"`
	PicUrl              string           `json:"pic_url"`
	Player              string           `json:"player"`
	PriceInfo           ListenPriceInfo  `json:"price_info"`
	Producer            string           `json:"producer"`
	PublishState        string           `json:"publish_state"`
	Pv                  string           `json:"pv"`
	Quality_128         string           `json:"quality_128"`
	Quality_32          string           `json:"quality_32"`
	ReadHardcoverBookId string           `json:"read_hardcover_book_id"`
	Read_ori_book_id    string           `json:"read_ori_book_id"`
	Relation_type       string           `json:"relation_type"`
	S_recommend         string           `json:"s_recommend"`
	Schemes             []int            `json:"schemes"`
	Show_name           string           `json:"show_name"`
	Status              string           `json:"status"`
	Subtitle            string           `json:"subtitle"`
	Tag_end             string           `json:"tag_end"`
	Tag_start           string           `json:"tag_start"`
	Tags                string           `json:"tags"`
	Tversion            string           `json:"tversion"`
	Update_time         string           `json:"update_time"`
	Uv                  int              `json:"uv"`
	Vip_end_time        int              `json:"vip_end_time"`
	Vip_start_time      int              `json:"vip_start_time"`
}

type PageInfo struct {
	TotalPage   int `json:"total_page"`
	TotalRecord int `json:"total_record"`
}

type ListenSection struct {
	BackupSection       int                 `json:"backupSection"`
	BookSource          string              `json:"bookSource"`
	Books               []ListenSectionBook `json:"books"`
	CampaignId          int                 `json:"campaignId"`
	CampaignType        int                 `json:"campaignType"`
	ClassifyBookNum     int                 `json:"classifyBookNum"`
	ClassifyTitle       string              `json:"classifyTitle"`
	CreateDate          string              `json:"createDate"`
	Customize           int                 `json:"customize"`
	Data_end_time       string              `json:"data_end_time"`
	Data_etime          string              `json:"data_etime"`
	Data_start_time     string              `json:"data_start_time"`
	Data_stime          string              `json:"data_stime"`
	Description         string              `json:"description"`
	Description2        string              `json:"description2"`
	DownloadChapterNum  int                 `json:"downloadChapterNum"`
	DownloadType        int                 `json:"downloadType"`
	EpubOnekeyDownload  int                 `json:"epubOnekeyDownload"`
	Ext                 string              `json:"ext"`
	ExtendAttr          string              `json:"extendAttr"`
	Fixed_bk_count      int                 `json:"fixed_bk_count"`
	Icon                string              `json:"icon"`
	Id                  int                 `json:"id"`
	IsDisplayTitle      int                 `json:"isDisplayTitle"`
	IsDrive             int                 `json:"isDrive"`
	IsMore              int                 `json:"isMore"`
	IsOnekeyDownload    int                 `json:"isOnekeyDownload"`
	IsRefreshChange     int                 `json:"isRefreshChange"`
	IsShare             string              `json:"isShare"`
	IsTop               int                 `json:"isTop"`
	Laisee              int                 `json:"laisee"`
	MoreTitle           string              `json:"moreTitle"`
	MoreUrl             string              `json:"moreUrl"`
	Name                string              `json:"name"`
	Number              int                 `json:"number"`
	ReadingNum          int                 `json:"readingNum"`
	RuleId              int                 `json:"ruleId"`
	SecColor            int                 `json:"secColor"`
	SectionUv           int                 `json:"sectionUv"`
	Section_update_time int64               `json:"section_update_time"`
	ShareContent        string              `json:"shareContent"`
	SharePic            string              `json:"sharePic"`
	ShareTitle          string              `json:"shareTitle"`
	SortId              int                 `json:"sortId"`
	Status              int                 `json:"status"`
	StyleId             int                 `json:"styleId"`
	StyleIdNew          int                 `json:"styleIdNew"`
	StyleName           string              `json:"styleName"`
	StyleSwitch         int                 `json:"styleSwitch"`
	TagList             []string            `json:"tagList"`
	Tags                string              `json:"tags"`
	ThemeImgUrl         string              `json:"themeImgUrl"`
	TopBookUv           int                 `json:"topBookUv"`
	TopStableNum        int                 `json:"topStableNum"`
	Sectiontype         string              `json:"type"`
	PageInfo            PageInfo            `json:"page_info"`
}

type result struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Body ListenSection `json:"body"`
}

func StructStr() (ListenSection, error) {
	str := `{
		"body": {
			"backupSection": 0,
			"bookSource": "",
			"books": [
				{
					"author": "我吃西红柿",
					"authorInfoList": [
						{
							"auditStatus": 0,
							"authorIcon": "",
							"authorId": 825741,
							"author_name": "我吃西红柿",
							"circleId": "",
							"relation_name": "作者",
							"relation_type": 10,
							"synopsis": "我吃西红柿",
							"usr": "i1728406806"
						},
						{
							"auditStatus": 1,
							"authorIcon": "",
							"authorId": 830933,
							"author_name": "流逝",
							"circleId": "",
							"relation_name": "",
							"relation_type": 140,
							"synopsis": "流逝",
							"usr": "i1919214638"
						}
					],
					"book_id": "30014524",
					"book_level": "0",
					"book_level_name": "",
					"book_type_id": "222",
					"book_type_name": "仙侠",
					"book_type_pic_url": "",
					"c_tag": "0",
					"chapter_count": "872",
					"comment_count": "0",
					"common_chapters_url": "group1/M00/5E/14/wKgKFFv76XmEHNZ8AAAAAGeVP1Y939857534.xml?v=NMX8pDDj&t=wKgKFFzqg-Q.",
					"consume_coin_num": "0",
					"copyright_belong": "2448",
					"copyright_belong_name": "华音文化",
					"copyright_type": "10",
					"copyright_type_name": "CP合作",
					"create_time": "2018-04-16 09:45:22",
					"description": "【编辑推荐】天天向上热荐作品 ，奇幻大神——我吃西红柿神作，唐家三少 、天蚕土豆、 骷髅精灵众大神联袂推荐！\r\n\r\n【内容简介】故事讲述的是男主角纪宁在地府因遇到六道轮回被袭击，未喝孟婆汤就投胎，而后所发生的一切。这里有为了生存，与天斗、与地斗、与妖斗的部落人们，有夸父逐日、后羿射金乌，还有为了逍遥长生，历三灾九劫、纵死无悔的修仙者……在他们的影响下，纪宁也成为了一名修仙者，开始了他传奇的修仙之路。",
					"fee_download_num": "0",
					"fee_user_num": "0",
					"file_key": "TWRer4h0+n+FPD8qWBT/f559/IFf+1J0MPuVmFeDjEwZmZ6e",
					"finish_state": "Y",
					"finish_state_name": "完结",
					"free_download_num": "0",
					"free_user_num": "0",
					"has_pic": "Y",
					"isFree": "0",
					"isVip": 0,
					"is_inner": "N",
					"is_preview": "N",
					"is_single": "N",
					"keywords": "莽荒纪,我吃西红柿",
					"l_recommend": "天天向上热荐作品 ，奇幻大神——我吃西红柿神作，唐家三少 、天蚕土豆、 骷髅精灵众大神联袂推荐！",
					"last_chapter_time": "2019-05-26 20:17:40",
					"name": "莽荒纪（华音流逝演播）",
					"ori_pic": "group1/M00/5A/AB/wKgKFFvyaIyEBXnsAAAAANRHEAA198412389.jpg?v=fKyBdFoH&t=wKgKFFxBT7M.",
					"pic_url": "group1/M00/35/1B/wKgKFVrT_-uEBGi7AAAAAEWFY7Y337885479.jpg?v=h3dbnZyf&t=wKgKFVxBT7M.",
					"player": "流逝",
					"price_info": {
						"chapter_count": 872,
						"chapter_price": 0.2,
						"discount": 1,
						"fee_chapter_num": 26,
						"fee_type": "1",
						"limit_end": 1585540800,
						"limit_price": 0,
						"limit_start": 1585058400,
						"vip_sub_code": 0,
						"whole_book_price": 0
					},
					"producer": "韩金珂",
					"publish_state": "N",
					"pv": "0",
					"quality_128": "Y",
					"quality_32": "Y",
					"read_hardcover_book_id": "0",
					"read_ori_book_id": "0",
					"relation_type": "0",
					"s_recommend": "",
					"schemes": [
						47,
						50,
						49,
						41,
						34,
						35,
						36,
						42,
						43,
						44,
						45,
						51,
						55,
						56,
						58,
						59,
						60,
						61,
						62,
						63,
						65,
						67,
						68,
						69,
						70,
						73,
						76,
						77,
						78,
						52
					],
					"show_name": "莽荒纪",
					"status": "Y",
					"subtitle": "",
					"tag_end": "1970-01-01 00:00:00",
					"tag_start": "1970-01-01 00:00:00",
					"tags": "",
					"tversion": "0",
					"update_time": "2020-03-02 03:03:31",
					"uv": 3329771,
					"vip_end_time": 0,
					"vip_start_time": 0
				},
				{
					"author": "素锦",
					"authorInfoList": [
						{
							"auditStatus": 1,
							"authorIcon": "",
							"authorId": 736525,
							"author_name": "素锦",
							"circleId": "author_736525",
							"relation_name": "作者",
							"relation_type": 10,
							"synopsis": "素锦，网络小说作者，著有《邪性老公枕边爱》《神医毒妃：嗜宠废材大小姐》。",
							"usr": "i843829171"
						},
						{
							"auditStatus": 0,
							"authorIcon": "",
							"authorId": 989718,
							"author_name": "盖拉",
							"circleId": "",
							"relation_name": "",
							"relation_type": 140,
							"synopsis": "",
							"usr": "i1947338150"
						},
						{
							"auditStatus": 1,
							"authorIcon": "",
							"authorId": 832712,
							"author_name": "郝赫奕",
							"circleId": "",
							"relation_name": "",
							"relation_type": 140,
							"synopsis": "",
							"usr": "i1919230890"
						}
					],
					"book_id": "30018828",
					"book_level": "0",
					"book_level_name": "",
					"book_type_id": "210",
					"book_type_name": "言情",
					"book_type_pic_url": "group1/M00/00/0E/wKgKFVPN9z2ELh-GAAAAANS2fao722368760.jpg",
					"c_tag": "0",
					"chapter_count": "425",
					"comment_count": "0",
					"common_chapters_url": "group2/M00/2A/33/wKgKNFy9JJGEdAzHAAAAAMPub8Q476968958.xml?v=7__uRJbT&t=wKgKNF2mmSk.",
					"consume_coin_num": "0",
					"copyright_belong": "2850",
					"copyright_belong_name": "红薯有声",
					"copyright_type": "10",
					"copyright_type_name": "CP合作",
					"create_time": "2019-04-22 10:18:58",
					"description": "一夜沉沦，苏蜜腹中多了一对龙凤胎。\r\n一夜强爱，千亿总裁多了夜夜梦见同一女人的病。\r\n五年后——\r\n他们再次重逢，命运的齿轮便开始运转，注定他们彼此纠缠，爱恨痴迷。\r\n第一面，拍卖会他拍下她，却道：“不是什么女人都能爬上我傅奕臣的床！”\r\n第二面，他压着她，“女人，你让我傅奕臣玩别的男人玩剩下的，好！有种！”\r\n第三面，他当众将她压在车盖上，“跟他离婚！以后你是我的女人！记住了！”\r\n……\r\n她怀着目的接近他，怕他惧他恨他，他占有她禁锢她纵容她，情感的纠葛，让冷清的她沉沦，令冷漠的他疯狂。\r\n却原来她竟是他多年前遗失的青梅竹马，缘分天定，相识相爱，注定他们将一生缠绵。",
					"fee_download_num": "0",
					"fee_user_num": "0",
					"file_key": "oM+r7odJi95E1tS/vwBbtQB55Fvbl4bIIAcSnWa9ajvulDmn",
					"finish_state": "Y",
					"finish_state_name": "完结",
					"free_download_num": "0",
					"free_user_num": "0",
					"has_pic": "Y",
					"isFree": "0",
					"isVip": 0,
					"is_inner": "N",
					"is_preview": "N",
					"is_single": "N",
					"keywords": "都市,言情,霸道总裁",
					"l_recommend": "",
					"last_chapter_time": "2019-10-16 12:14:33",
					"name": "一夜强宠，禁欲总裁强制爱",
					"ori_pic": "group2/M00/2A/48/wKgKNFy9VIyEB2GnAAAAAJOEcFo120816463.jpg?v=fUlpwwQy&t=wKgKNFy9VI4.",
					"pic_url": "group2/M00/2A/48/wKgKNVy9VI2ELpn8AAAAAGWN0Zk127429700.jpg?v=yO2fal1I&t=wKgKNVy9VI4.",
					"player": "盖拉，郝赫奕",
					"price_info": {
						"chapter_count": 425,
						"chapter_price": 0.2,
						"discount": 1,
						"fee_chapter_num": 36,
						"fee_type": "1",
						"limit_end": 1585540800,
						"limit_price": 0,
						"limit_start": 1585058400,
						"vip_sub_code": 0,
						"whole_book_price": 0
					},
					"producer": "杨少卿",
					"publish_state": "N",
					"pv": "0",
					"quality_128": "Y",
					"quality_32": "Y",
					"read_hardcover_book_id": "0",
					"read_ori_book_id": "0",
					"relation_type": "0",
					"s_recommend": "",
					"schemes": [
						47,
						50,
						49,
						41,
						34,
						35,
						36,
						42,
						43,
						44,
						45,
						51,
						55,
						56,
						58,
						59,
						60,
						61,
						62,
						63,
						65,
						67,
						68,
						69,
						70,
						72,
						73,
						76,
						77,
						78,
						52
					],
					"show_name": "邪性老公太霸道",
					"status": "Y",
					"subtitle": "",
					"tag_end": "1970-01-01 00:00:00",
					"tag_start": "1970-01-01 00:00:00",
					"tags": "",
					"tversion": "0",
					"update_time": "2020-03-02 03:03:14",
					"uv": 1702511,
					"vip_end_time": 0,
					"vip_start_time": 0
				},
				{
					"author": "青山依旧在",
					"authorInfoList": [
						{
							"auditStatus": 0,
							"authorIcon": "",
							"authorId": 718710,
							"author_name": "青山依旧在",
							"circleId": "",
							"relation_name": "作者",
							"relation_type": 10,
							"synopsis": "",
							"usr": "i1732616745"
						},
						{
							"auditStatus": 0,
							"authorIcon": "",
							"authorId": 1014541,
							"author_name": "正月初一",
							"circleId": "",
							"relation_name": "",
							"relation_type": 140,
							"synopsis": "",
							"usr": "i2214501420"
						},
						{
							"auditStatus": 0,
							"authorIcon": "",
							"authorId": 1083205,
							"author_name": "萧秋子",
							"circleId": "",
							"relation_name": "",
							"relation_type": 140,
							"synopsis": "",
							"usr": "i2675520716"
						}
					],
					"book_id": "30025850",
					"book_level": "0",
					"book_level_name": "",
					"book_type_id": "208",
					"book_type_name": "都市",
					"book_type_pic_url": "group1/M00/00/0E/wKgKFVPN98iEWZpCAAAAAPdCZ8k338444244.jpg",
					"c_tag": "0",
					"chapter_count": "460",
					"comment_count": "0",
					"common_chapters_url": "group21/M00/1D/11/CmQVFF4W8y2EGajbAAAAAI6U0yY233526539.xml?v=F_ng0Xzf&t=CmQVFF58c-s.",
					"consume_coin_num": "0",
					"copyright_belong": "2644",
					"copyright_belong_name": "看书网（有声）",
					"copyright_type": "10",
					"copyright_type_name": "CP合作",
					"create_time": "2020-01-09 17:32:29",
					"description": "修仙天才重生花花都市，校花排成排，美女扎成堆，师妹俏皮可爱，师姐成熟感性，邻居甜美乖巧，还有个傲娇未婚妻，秦世压力很大呀；以废材之躯，重塑修仙荣耀，左手红颜，右手霸业，一个转身，已然睥睨天下，败尽群雄！",
					"fee_download_num": "0",
					"fee_user_num": "0",
					"file_key": "cnis3ZHsqNmwhPr/TSxO0gQVzf4AfbxIdIojrWqYJeZ7GITo",
					"finish_state": "N",
					"finish_state_name": "连载",
					"free_download_num": "0",
					"free_user_num": "0",
					"has_pic": "Y",
					"isFree": "0",
					"isVip": 0,
					"is_inner": "N",
					"is_preview": "N",
					"is_single": "N",
					"keywords": "修真,重生,装逼,校花,性感,校园,爽文,修仙,灵兽,未婚妻,老师,总裁",
					"l_recommend": "",
					"last_chapter_time": "2020-03-26 17:20:43",
					"name": "豪门赘婿：征服霸道女总裁",
					"ori_pic": "group21/M00/1C/FC/CmQVFF4W3geEGyM1AAAAAFprGv8102644769.jpg?v=BqGrM9Ij&t=CmQVFF4W3gc.",
					"pic_url": "group21/M00/1C/FF/CmQVE14W3geERov-AAAAALSzgMA929480503.jpg?v=1Ex0cbKe&t=CmQVE14W3gc.",
					"player": "正月初一,萧秋子",
					"price_info": {
						"chapter_count": 460,
						"chapter_price": 0.2,
						"discount": 1,
						"fee_chapter_num": 31,
						"fee_type": "1",
						"limit_end": 1585540800,
						"limit_price": 0,
						"limit_start": 1585058400,
						"vip_sub_code": 0,
						"whole_book_price": 0
					},
					"producer": "刘晨",
					"publish_state": "N",
					"pv": "0",
					"quality_128": "Y",
					"quality_32": "Y",
					"read_hardcover_book_id": "0",
					"read_ori_book_id": "0",
					"relation_type": "0",
					"s_recommend": "",
					"schemes": [
						47,
						50,
						49,
						41,
						34,
						35,
						36,
						42,
						43,
						44,
						45,
						51,
						55,
						56,
						58,
						59,
						60,
						61,
						62,
						63,
						65,
						67,
						68,
						69,
						70,
						71,
						73,
						76,
						77,
						78,
						52
					],
					"show_name": "豪门赘婿：征服霸道女总裁",
					"status": "Y",
					"subtitle": "",
					"tag_end": "1970-01-01 00:00:00",
					"tag_start": "1970-01-01 00:00:00",
					"tags": "",
					"tversion": "0",
					"update_time": "2020-03-26 17:20:43",
					"uv": 135112,
					"vip_end_time": 0,
					"vip_start_time": 0
				},
				{
					"author": "罗晓",
					"authorInfoList": [
						{
							"auditStatus": 1,
							"authorIcon": "",
							"authorId": 370181,
							"author_name": "罗晓",
							"circleId": "author_370181",
							"relation_name": "作者",
							"relation_type": 10,
							"synopsis": "罗晓，著名作家，鲁迅文学院培训学员，中文在线金牌作家，原名丁道兵，湖北恩施人，资深盗墓、淘宝小说作家，主要著有《黄金手》《淘宝笔记》《最强相师》《超级妖瞳》《摸金传人》《纪委书记》。",
							"usr": "i1731575486"
						},
						{
							"auditStatus": 1,
							"authorIcon": "group61/M00/9A/95/CmQUOFya582EIi6_AAAAAPPSNaI835425975.jpg?v=MEwXHVQv&t=CmQUOFya580.",
							"authorId": 829894,
							"author_name": "思大咖",
							"circleId": "",
							"relation_name": "",
							"relation_type": 140,
							"synopsis": "思大咖",
							"usr": "i1919192538"
						}
					],
					"book_id": "30017286",
					"book_level": "0",
					"book_level_name": "",
					"book_type_id": "208",
					"book_type_name": "都市",
					"book_type_pic_url": "group1/M00/00/0E/wKgKFVPN98iEWZpCAAAAAPdCZ8k338444244.jpg",
					"c_tag": "0",
					"chapter_count": "547",
					"comment_count": "0",
					"common_chapters_url": "group2/M00/18/E4/wKgKNVx1CYKEP2WpAAAAAMgwX2g734370822.xml?v=OQK5vnbS&t=wKgKNVyaalA.",
					"consume_coin_num": "0",
					"copyright_belong": "1194",
					"copyright_belong_name": "鸿达以太",
					"copyright_type": "10",
					"copyright_type_name": "CP合作",
					"create_time": "2019-02-26 17:40:19",
					"description": "跌入人生谷底的少年许东偶得异能，能看到珍宝所发出来的“宝气”，从此鉴宝寻宝，能人所不能！珠有光，宝有气，人生就是一出美女与财富混杂的戏！",
					"fee_download_num": "0",
					"fee_user_num": "0",
					"file_key": "dAdxzeb3rbCIHqarHpKVcA7qDrN8PRqKyPzerSZlID9E6sRA",
					"finish_state": "Y",
					"finish_state_name": "完结",
					"free_download_num": "0",
					"free_user_num": "0",
					"has_pic": "Y",
					"isFree": "0",
					"is_inner": "N",
					"is_preview": "N",
					"is_single": "N",
					"keywords": "都市,财富,美女",
					"l_recommend": "",
					"last_chapter_time": "2019-03-27 02:07:12",
					"name": "大宝鉴",
					"ori_pic": "group2/M00/18/CA/wKgKNVx056WEDBIOAAAAAA6qTO4102045878.jpg?v=bSTfp9H8&t=wKgKNV2lOr4.",
					"pic_url": "group2/M00/18/CA/wKgKNVx056WEKMXWAAAAAEzrbcs665336135.jpg?v=97TV97DD&t=wKgKNV2lOr4.",
					"player": "思大咖",
					"price_info": {
						"chapter_count": 547,
						"chapter_price": 0.2,
						"discount": 1,
						"fee_chapter_num": 51,
						"fee_type": "1",
						"limit_end": 1585540800,
						"limit_price": 0,
						"limit_start": 1585058400,
						"vip_sub_code": 0,
						"whole_book_price": 0
					},
					"producer": "杨少卿",
					"publish_state": "N",
					"pv": "0",
					"quality_128": "Y",
					"quality_32": "Y",
					"read_hardcover_book_id": "11299667",
					"read_ori_book_id": "10980995",
					"relation_type": "2",
					"s_recommend": "",
					"schemes": [
						47,
						50,
						49,
						41,
						34,
						35,
						36,
						42,
						43,
						44,
						45,
						51,
						55,
						56,
						58,
						59,
						60,
						61,
						62,
						63,
						65,
						67,
						68,
						69,
						70,
						73,
						76,
						77,
						78,
						52
					],
					"show_name": "大宝鉴",
					"status": "Y",
					"subtitle": "",
					"tag_end": "1970-01-01 00:00:00",
					"tag_start": "1970-01-01 00:00:00",
					"tags": "",
					"tversion": "0",
					"update_time": "2020-03-02 03:03:39",
					"uv": 177591
				},
				{
					"author": "桃慕慕",
					"authorInfoList": [
						{
							"auditStatus": 0,
							"authorIcon": "",
							"authorId": 972034,
							"author_name": "桃慕慕",
							"circleId": "",
							"relation_name": "作者",
							"relation_type": 10,
							"synopsis": "",
							"usr": "i1826126631"
						},
						{
							"auditStatus": 1,
							"authorIcon": "",
							"authorId": 830217,
							"author_name": "网易听书",
							"circleId": "",
							"relation_name": "",
							"relation_type": 140,
							"synopsis": "网易听书",
							"usr": "i1919198796"
						}
					],
					"book_id": "30016758",
					"book_level": "0",
					"book_level_name": "",
					"book_type_id": "210",
					"book_type_name": "言情",
					"book_type_pic_url": "group1/M00/00/0E/wKgKFVPN9z2ELh-GAAAAANS2fao722368760.jpg",
					"c_tag": "0",
					"chapter_count": "496",
					"comment_count": "0",
					"common_chapters_url": "group2/M00/06/DF/wKgKNFwcu3CEM-tBAAAAABUQB0I300915016.xml?v=jVBcaGym&t=wKgKNF571sg.",
					"consume_coin_num": "0",
					"copyright_belong": "2436",
					"copyright_belong_name": "网易（有声）",
					"copyright_type": "10",
					"copyright_type_name": "CP合作",
					"create_time": "2018-12-21 18:07:46",
					"description": "“不要叫我后妈，我没你那么大的儿子！”\r\n艾天晴一直以为自己要嫁的人是一个快六十的老头，直到某天晚上那个邪魅冷血的男人将她抵在了门上，从此她的日子就……",
					"fee_download_num": "0",
					"fee_user_num": "0",
					"file_key": "YFUVmiFR3JD2ZUhJndBE3uBnS7NeGnQoQtJeVBhlgdYHGdjR",
					"finish_state": "N",
					"finish_state_name": "连载",
					"free_download_num": "0",
					"free_user_num": "0",
					"has_pic": "Y",
					"isFree": "0",
					"is_inner": "N",
					"is_preview": "N",
					"is_single": "N",
					"keywords": "婚恋,闪婚,总裁豪门,契约情人",
					"l_recommend": "",
					"last_chapter_time": "2020-03-26 06:10:16",
					"name": "宠你一世又何妨",
					"ori_pic": "group2/M00/06/C0/wKgKNVwclxGEOzGrAAAAALYkv1c369613788.jpg?v=M8CQlu2s&t=wKgKNVwclxM.",
					"pic_url": "group2/M00/06/C0/wKgKNVwclxGEY1QqAAAAAKxXXw4860700542.jpg?v=P7vYc-Nr&t=wKgKNVwclxI.",
					"player": "网易听书",
					"price_info": {
						"chapter_count": 496,
						"chapter_price": 0.2,
						"discount": 1,
						"fee_chapter_num": 31,
						"fee_type": "1",
						"limit_end": 1585540800,
						"limit_price": 0,
						"limit_start": 1585058400,
						"vip_sub_code": 0,
						"whole_book_price": 0
					},
					"producer": "韩金珂",
					"publish_state": "N",
					"pv": "0",
					"quality_128": "Y",
					"quality_32": "Y",
					"read_hardcover_book_id": "0",
					"read_ori_book_id": "0",
					"relation_type": "2",
					"s_recommend": "",
					"schemes": [
						47,
						50,
						49,
						41,
						34,
						35,
						36,
						42,
						43,
						44,
						45,
						51,
						55,
						56,
						58,
						59,
						60,
						61,
						62,
						63,
						65,
						67,
						68,
						69,
						70,
						73,
						76,
						77,
						78,
						52
					],
					"show_name": "宠你一世又何妨",
					"status": "Y",
					"subtitle": "",
					"tag_end": "1970-01-01 00:00:00",
					"tag_start": "1970-01-01 00:00:00",
					"tags": "",
					"tversion": "0",
					"update_time": "2020-03-26 06:10:16",
					"uv": 124390
				},
				{
					"author": "六月",
					"authorInfoList": [
						{
							"auditStatus": 1,
							"authorIcon": "",
							"authorId": 742632,
							"author_name": "六月",
							"circleId": "author_742632",
							"relation_name": "作者",
							"relation_type": 10,
							"synopsis": "六月，网络小说作者，著有《将女惊华：将军大人太霸道》《权宠天下》《倾世医妃要休夫》《摄政王的医品狂妃》《权宠悍妻》。",
							"usr": "i707064512"
						},
						{
							"auditStatus": 0,
							"authorIcon": "",
							"authorId": 1015783,
							"author_name": "梁珈源",
							"circleId": "",
							"relation_name": "",
							"relation_type": 140,
							"synopsis": "",
							"usr": "i2268421262"
						},
						{
							"auditStatus": 0,
							"authorIcon": "",
							"authorId": 1015784,
							"author_name": "韩磊",
							"circleId": "",
							"relation_name": "",
							"relation_type": 140,
							"synopsis": "",
							"usr": "i2268421265"
						}
					],
					"book_id": "30024344",
					"book_level": "0",
					"book_level_name": "",
					"book_type_id": "209",
					"book_type_name": "穿越",
					"book_type_pic_url": "group1/M00/00/0E/wKgKFVPN9z2ELh-GAAAAANS2fao722368760.jpg",
					"c_tag": "0",
					"chapter_count": "370",
					"comment_count": "0",
					"common_chapters_url": "group2/M00/53/20/wKgKNV2hj06EQRrEAAAAADnv-Wg034316812.xml?v=i1ioabJ4&t=wKgKNV572_Q.",
					"consume_coin_num": "0",
					"copyright_belong": "3560",
					"copyright_belong_name": "声动懒人（有声）",
					"copyright_type": "10",
					"copyright_type_name": "CP合作",
					"create_time": "2019-10-12 16:31:12",
					"description": "现代特工军医穿越为相府嫡女，受父亲与庶母迫害，嫁与摄政王，种种陷阱，处处陷害，凭着一身的医术，她在府中斗争与深宫之争中游刃有余，诛太子，救梁王，除瘟疫，从一个畏畏缩缩的相府小姐蜕变成可以与他并肩而立的坚毅女子。",
					"fee_download_num": "0",
					"fee_user_num": "0",
					"file_key": "LxAY/mCsFc9xhPdH0N23v5aFL31gSuK2jnp1kWvmIj434hDO",
					"finish_state": "N",
					"finish_state_name": "连载",
					"free_download_num": "0",
					"free_user_num": "0",
					"has_pic": "Y",
					"isFree": "0",
					"isVip": 0,
					"is_inner": "N",
					"is_preview": "N",
					"is_single": "N",
					"keywords": "穿越,宫斗,爱情,王妃,逆袭",
					"l_recommend": "",
					"last_chapter_time": "2020-03-26 06:32:20",
					"name": "摄政王的医品狂妃",
					"ori_pic": "group2/M00/30/C0/CmRenV2hi5KEcijvAAAAABNfMcY022277769.jpg?v=qrC3tLPk&t=CmRenV2hi5I.",
					"pic_url": "group2/M00/53/20/wKgKNV2hi5CEEGxHAAAAABdDZYg047337666.jpg?v=ixG5eTy7&t=wKgKNV2hi5I.",
					"player": "梁珈源,韩磊",
					"price_info": {
						"chapter_count": 370,
						"chapter_price": 0.2,
						"discount": 1,
						"fee_chapter_num": 41,
						"fee_type": "1",
						"limit_end": 1585540800,
						"limit_price": 0,
						"limit_start": 1585058400,
						"vip_sub_code": 0,
						"whole_book_price": 0
					},
					"producer": "刘晨",
					"publish_state": "N",
					"pv": "0",
					"quality_128": "Y",
					"quality_32": "Y",
					"read_hardcover_book_id": "11508903",
					"read_ori_book_id": "11508900",
					"relation_type": "2",
					"s_recommend": "",
					"schemes": [
						47,
						50,
						49,
						41,
						34,
						35,
						36,
						42,
						43,
						44,
						45,
						51,
						55,
						56,
						58,
						59,
						60,
						61,
						62,
						63,
						65,
						67,
						68,
						69,
						70,
						73,
						76,
						77,
						78,
						52
					],
					"show_name": "摄政王的医品狂妃",
					"status": "Y",
					"subtitle": "",
					"tag_end": "1970-01-01 00:00:00",
					"tag_start": "1970-01-01 00:00:00",
					"tags": "",
					"tversion": "0",
					"update_time": "2020-03-26 06:32:20",
					"uv": 119171,
					"vip_end_time": 0,
					"vip_start_time": 0
				}
			],
			"campaignId": 0,
			"campaignType": 0,
			"classifyBookNum": 10,
			"classifyTitle": "7.1有声内容",
			"createDate": "2018-01-30",
			"customize": 0,
			"data_end_time": "",
			"data_etime": "",
			"data_start_time": "",
			"data_stime": "",
			"description": "",
			"description2": "限时免费时间：3月25日到3月30日12:00",
			"downloadChapterNum": 0,
			"downloadType": 1,
			"epubOnekeyDownload": 0,
			"ext": "",
			"extendAttr": "",
			"fixed_bk_count": 0,
			"icon": "",
			"id": 29239,
			"isDisplayTitle": 1,
			"isDrive": 0,
			"isMore": 3,
			"isOnekeyDownload": 0,
			"isRefreshChange": 1,
			"isShare": "0",
			"isTop": 0,
			"laisee": 1,
			"moreTitle": "",
			"moreUrl": "",
			"name": "限时免费",
			"number": 6,
			"page_info": {
				"total_page": 1,
				"total_record": 6
			},
			"readingNum": 0,
			"ruleId": 0,
			"secColor": 1,
			"sectionUv": 0,
			"section_update_time": 1585102457,
			"shareContent": "",
			"sharePic": "",
			"shareTitle": "",
			"sortId": 24353,
			"status": 1,
			"styleId": 1,
			"styleIdNew": 27,
			"styleName": "NEW-STYLE136",
			"styleSwitch": 0,
			"tagList": [],
			"tags": "",
			"themeImgUrl": "",
			"topBookUv": 0,
			"topStableNum": 0,
			"type": "20"
		},
		"code": 0,
		"msg": "Success"
	}`

	var data result
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		return ListenSection{}, err
	}
	fmt.Println("code", data.Code)
	fmt.Println("msg", data.Msg)
	return data.Body, nil
}

func TestStructStr(t *testing.T) {
	rs, err := StructStr()
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(rs.ClassifyTitle)
	t.Log(rs.PageInfo)
}

func TestYunsuan(t *testing.T) {
	a := "123"
	b := "123"
	if strings.Contains(a, b) {
		t.Log("包含")
	} else {
		t.Log("不包含")
	}
}
