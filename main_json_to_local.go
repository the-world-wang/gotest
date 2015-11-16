package main

import (
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	//"moemoela.com/netalog"
	"code.google.com/p/go-uuid/uuid"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/context"
	"io/ioutil"
	"os"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"
	"strings"
	"time"
)

const (
	URL         = "http://7xl0tq.com2.z0.glb.qiniucdn.com/"
	FIRST_DATA  = "first_data"
	SECOND_DATA = "second_data"
)

var (
	engine *xorm.Engine
)

type TimeTableEventTable struct {
	UUID             string    `xorm:"pk 'UUID'"`
	NAME             string    `xorm:"'NAME'"`
	TYPE             string    `xorm:"'TYPE'"`
	ADDRESS          string    `xorm:"'ADDRESS'"`
	URL              string    `xorm:"'URL'"`
	SHARE_URL        string    `xorm:"'SHARE_URL'"`
	ICON_ID          string    `xorm:"'ICON_ID'"`
	ICON_NAME        string    `xorm:"'ICON_NAME'"`
	ICON_WIDTH       int64     `xorm:"'ICON_WIDTH'"`
	ICON_HEIGHT      int64     `xorm:"'ICON_HEIGHT'"`
	QR_CODE_ID       string    `xorm:"'QR_CODE_ID'"`
	QR_CODE_NAME     string    `xorm:"'QR_CODE_NAME'"`
	QR_CODE_WIDTH    int64     `xorm:"'QR_CODE_WIDTH'"`
	QR_CODE_HEIGHT   int64     `xorm:"'QR_CODE_HEIGHT'"`
	BEGIN_TIME       time.Time `xorm:"'BEGIN_TIME'"`
	END_TIME         time.Time `xorm:"'END_TIME'"`
	PERIOD_DAYS      int64     `xorm:"'PERIOD_DAYS'"`
	CREATE_USER      string    `xorm:"'CREATE_USER'"`
	CREATE_TIME      time.Time `xorm:"created 'CREATE_TIME'"`
	UPDATE_USER      string    `xorm:"'UPDATE_USER'"`
	UPDATE_TIME      time.Time `xorm:"updated 'UPDATE_TIME'"`
	FROZEN_STATUS    string    `xorm:"'FROZEN_STATUS'"`
	FROZEN_TIME      time.Time `xorm:"'FROZEN_TIME'"`
	VERSION          int64     `xorm:"version 'VERSION'"`
	RECOMMEND_STATUS string    `xorm:"'RECOMMEND_STATUS'"`
	EXTRA_DATA       string    `xorm:"'EXTRA_DATA'"`
}

func (self TimeTableEventTable) TableName() string {
	return "TIME_TABLE_EVENT"
}

type TimeTableEventTypeTable struct {
	UUID             string    `xorm:"pk 'UUID'"`
	EVENT_ID         string    `xorm:"'EVENT_ID'"`
	EVENT_TYPE_ID    string    `xorm:"'EVENT_TYPE_ID'"`
	EVENT_TYPE_NAME  string    `xorm:"'EVENT_TYPE_NAME'"`
	EVENT_TYPE_COLOR string    `xorm:"'EVENT_TYPE_COLOR'"`
	EVENT_TYPE_LEVEL string    `xorm:"'EVENT_TYPE_LEVEL'"`
	CREATE_USER      string    `xorm:"'CREATE_USER'"`
	CREATE_TIME      time.Time `xorm:"created 'CREATE_TIME'"`
}

func (self TimeTableEventTypeTable) TableName() string {
	return "TIME_TABLE_EVENT_TYPE"
}

type HtmlData struct {
	Key   string
	Value interface{}
}

type Style struct {
	HeadBackground string `json:"head_background"`
	FootBackground string `json:"foot_background"`
	Border         string `json:"border"`
	Leaf           string `json:"leaf"`
	Icon           string `json:"icon"`
}

func main() {
	//http://7xl0tq.com2.z0.glb.qiniucdn.com/
	/*netaLog := netalog.New(netalog.LEVEL_TRACE, os.Stdout)
	netaLog.Info("time table html rebuild")*/
	var err error
	engine, err = xorm.NewEngine("postgres", "host=123.58.136.98 port=12345 dbname=otaku user=otaku password='123456' connect_timeout=60 sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}

	timeEventSlice := make([]TimeTableEventTable, 0, 10000)
	session := engine.NewSession()
	session.Asc(`"UUID"`).Find(&timeEventSlice)

	fmt.Println(len(timeEventSlice))
	for _, timeTable := range timeEventSlice {
		timeTableType := make([]TimeTableEventTypeTable, 0, 2)
		//最后存放的数组，内容以tmlData{key，value}保存
		data := make([]HtmlData, 0, 7)
		session.Where(`"EVENT_ID" = $1`, timeTable.UUID).Find(&timeTableType)
		for _, v := range timeTableType {
			if v.EVENT_TYPE_LEVEL == "NODE" {
				switch v.EVENT_TYPE_NAME {
				case "纪念日":
					style := Style{"#86de6e",
						"#b6d55b",
						"#3fbca4",
						"#4cb1e5",
						URL + "timetable/img/ea7dea29-b015-4a98-9245-51d690c39e38.png"}
					data = append(data, HtmlData{"style", style})
				case "活动":
					style := Style{"#ed5a52",
						"#e679b6",
						"#b23659",
						"#fcc015",
						URL + "timetable/img/bd216a18-ec0b-45f7-aed2-14140b8f1222.png"}
					data = append(data, HtmlData{"style", style})
				case "宅物":
					style := Style{
						"#6bd7fe",
						"#8bdbc5",
						"#78b05b",
						"#3ea8f7",
						URL + "timetable/img/3d9fc13e-0ae0-4785-91b5-ee26b4aacb17.png"}
					data = append(data, HtmlData{"style", style})
				case "新番":
					style := Style{"#f5e674",
						"#f5b842",
						"#e58217",
						"#cf6555",
						URL + "timetable/img/bd86e500-a65d-4416-b1f4-8a848d7d55db.png"}
					data = append(data, HtmlData{"style", style})
				case "游戏":
					style := Style{
						"#5c4f8f",
						"#54829e",
						"#482e72",
						"#3ad8d8",
						URL + "timetable/img/c761ceb1-8bf3-432d-8156-9c17bd30fd44.png"}
					data = append(data, HtmlData{"style", style})
				}

			}
			if v.EVENT_TYPE_LEVEL == "LEAF" {
				data = append(data, HtmlData{"leaf", v.EVENT_TYPE_NAME})
			}
		}

		// 对应HTML的URL
		httpURL := URL + timeTable.URL
		// 活动名称
		eventName := timeTable.NAME

		// 将title放置其中
		data = append(data, HtmlData{"title", eventName})

		// time处理
		timeData := HtmlData{Key: "time"}
		if timeTable.BEGIN_TIME == timeTable.END_TIME {
			timeData.Value = timeTable.BEGIN_TIME.Format("2006年01月")
			data = append(data, timeData)
		} else {
			// beignTime和endTime的比较
			beginTime := timeTable.BEGIN_TIME.Format("2006年01月02日")
			beginTime_month := timeTable.BEGIN_TIME.Format("01月")
			endTime_month := timeTable.END_TIME.Format("01月")
			var endTime string
			if beginTime_month == endTime_month {
				endTime = timeTable.END_TIME.Format("02日")
			} else {
				endTime = timeTable.END_TIME.Format("01年02日")
			}
			timeData.Value = beginTime + "-" + endTime
			data = append(data, timeData)
		}

		// 生成json
		dataStr := GetJson(httpURL, data)
		// file, _ := os.Create("D:/temp/data/" + timeTable.UUID + ".json")
		// file.WriteString(dataStr)
		var key string
		err, key = Upload()
		if err != nil {
			fmt.Println(err)
		}
		timeTable.EXTRA_DATA = key
		session.Cols(`"EXTRA_DATA"`).Update(&timeTable)
	}

}

/* 根据url获取json数据
 *
 */
func GetJson(url string, data []HtmlData) string {

	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err)
	}

	keyTemp := make([]string, 0, 5) //暂时存放取出来的key

	//key
	doc.Find("div div span").Each(func(i int, s *goquery.Selection) {
		key := s.Text()
		keyTemp = append(keyTemp, key)
	})

	// value,以数组形式保存
	doc.Find("div div p").Each(func(i int, s *goquery.Selection) {
		value, _ := s.Html()
		value = strings.Replace(value, "\n", "<br>", -1)
		httpData := HtmlData{keyTemp[i], value}
		data = append(data, httpData)
		data = addExtendData(data, httpData.Key, value)
	})

	// 处理图片
	imgList := make([]string, 0, 1)
	doc.Find("div div img").Each(func(i int, s *goquery.Selection) {
		value, _ := s.Attr("src")
		imgList = append(imgList, value)
	})

	imgData := HtmlData{"img", imgList}
	data = append(data, imgData)

	b, erro := json.Marshal(data)
	if erro != nil {
		fmt.Println(erro)
	}
	return string(b)
}

/*判断Key是否包含CV，作品，事务所等*/
func addExtendData(data []HtmlData, key string, value string) []HtmlData {
	// 纪念
	if strings.Contains(key, "CV") {
		extendData := HtmlData{FIRST_DATA, key + "：" + value}
		data = append(data, extendData)
	} else if key == "声优" {
		extendData := HtmlData{FIRST_DATA, key + "：" + value}
		data = append(data, extendData)
	} else if strings.Contains(key, "作品") {
		extendData := HtmlData{SECOND_DATA, key + "：" + value}
		data = append(data, extendData)
	} else if strings.Contains(key, "事务所") {
		extendData := HtmlData{FIRST_DATA, key + "：" + value}
		data = append(data, extendData)
	} else if strings.Contains(key, "地点") {
		// 活动
		extendData := HtmlData{FIRST_DATA, key + "：" + value}
		data = append(data, extendData)
	} else if strings.Contains(key, "地址") {
		extendData := HtmlData{FIRST_DATA, key + "：" + value}
		data = append(data, extendData)
	} else if strings.Contains(key, "厂商") {
		// 宅物
		extendData := HtmlData{FIRST_DATA, key + "：" + value}
		data = append(data, extendData)
	} else if strings.Contains(key, "售价") {
		extendData := HtmlData{SECOND_DATA, key + "：" + value}
		data = append(data, extendData)
	} else if strings.Contains(key, "更新日") {
		//新番
		extendData := HtmlData{FIRST_DATA, key + "：" + value}
		data = append(data, extendData)
	} else if strings.Contains(key, "新番") {
		index := strings.Index(value, "简介")
		if index < 0 {
			index = 0
		}
		var str string
		if len(value) < 150 {
			str = value[index:len(value)]
		} else {
			str = value[index : index+150]
		}
		//找到了"简介"
		if index > 0 {
			str = strings.Replace(str, "】", "：", 1)
		}
		extendData := HtmlData{SECOND_DATA, str}
		data = append(data, extendData)
	} else if strings.Contains(key, "游戏类型") {
		// 游戏
		extendData := HtmlData{FIRST_DATA, key + "：" + value}
		data = append(data, extendData)
	} else if strings.Contains(key, "语言") {
		extendData := HtmlData{SECOND_DATA, key + "：" + value}
		data = append(data, extendData)
	}
	return data
}

func Upload() (error, string) {
	access_key := "epbKZnxFUtJ9bTWufWtvXkAwtsseutpa8xRpJ3KI"
	secret_key := "TEWIuwOCs9KTeRqrOuQEDmUHDDd6RYBkoZ32m2Is"
	kodo.SetMac(access_key, secret_key)

	zone := 0 // 您空间(Bucket)所在的区域

	client := kodo.New(zone, nil) // 用默认配置创建 Client
	deadline_sec := time.Now().Add(3000 * time.Second).Unix()

	files := ListDir("D:/temp/data", suffix)

	for _, v := range files {
		reg := "timetable/data/"
		filename := uuid.New()
		reg = reg + filename + ".json"
		fmt.Println(reg)
		put_policy := kodo.PutPolicy{
			Scope:   "otaku-resource:" + reg,
			Expires: uint32(deadline_sec),
		}
		utoken := client.MakeUptoken(&put_policy)

		local_image_file := v //本地img文件地址
		uploader := kodocli.NewUploader(zone, nil)
		ret := kodocli.PutRet{}
		ctx := context.Background()
		err := uploader.PutFile(ctx, &ret, utoken, reg, local_image_file, nil)
		//http://7xl0tq.com2.z0.glb.qiniucdn.com
		if err != nil {
			fmt.Println(err)
			return err, nil
		}
		return err, ret.Key
	}
}

//获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)

	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	pathSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+pathSep+fi.Name())
		}
	}
	return files, nil
}
