package main

import (
	"crypto/md5"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io"
	"math"
	"strconv"
	"time"
)

/*func main()  {
	for i:=1;i<10;i++{
		for j:=1;j<=i ;j++  {
			fmt.Printf("%d * %d = %d\t",j,i,j*i)
		}
		fmt.Println()
	}
}*/

/*func main()  {
	for num:=100;num<1000 ;num++  {
		var i = num/100
		var j  = num/10%10
		var k  = num%10
		if math.Pow(float64(i),3)+math.Pow(float64(j),3)+math.Pow(float64(k),3) == float64(num) {
			fmt.Println(num)
		}
	}
}*/

/*func main() {
	for cock := 0; cock < 20; cock++ {
		for hen := 0; hen < 33; hen++ {
			for chicken := 0; chicken < 100; chicken += 3 {
				if 5*cock+3*hen+chicken/3 == 100 && chicken+cock+hen == 100 {
					fmt.Printf("公鸡:%d只;母鸡:%d只;小鸡:%d只\n", cock, hen, chicken)
				}
			}
		}
	}
}*/

/*func main() {
	var line int = 10
	for i := 0; i < line; i++ {
		for j:=0;j<line-i-1 ;j++  {
			fmt.Print(" ")
		}
		for k:=0;k<i*2+1 ;k++  {
			fmt.Print("*")
		}
		fmt.Println()
	}
}*/

type WeekDate struct {
	WeekTh    string
	StartTime time.Time
	EndTime   time.Time
}

// 将开始时间和结束时间分割为周为单位
func GroupByWeekDate(startTime, endTime time.Time) []WeekDate {
	weekDate := make([]WeekDate, 0)
	diffDuration := endTime.Sub(startTime)
	days := int(math.Ceil(float64(diffDuration/(time.Hour*24)))) + 1

	currentWeekDate := WeekDate{}
	currentWeekDate.WeekTh = WeekByDate(endTime)
	currentWeekDay := int(endTime.Weekday())
	if currentWeekDay == 0 {
		currentWeekDay = 7
	}
	currentWeekDate.EndTime = endTime.AddDate(0, 0, 7-currentWeekDay)
	currentWeekDate.StartTime = endTime.AddDate(0, 0, -currentWeekDay+1)
	nextWeekEndTime := currentWeekDate.StartTime
	weekDate = append(weekDate, currentWeekDate)

	for i := 0; i < (days-currentWeekDay)/7; i++ {
		weekData := WeekDate{}
		weekData.StartTime = nextWeekEndTime.AddDate(0, 0, -7)
		weekData.EndTime = nextWeekEndTime.AddDate(0, 0, -1)
		weekData.WeekTh = WeekByDate(weekData.StartTime)
		nextWeekEndTime = weekData.StartTime
		weekDate = append(weekDate, weekData)
	}

	if lastDays := (days - currentWeekDay) % 7; lastDays > 0 {
		lastData := WeekDate{}
		lastData.StartTime = nextWeekEndTime.AddDate(0, 0, - 7)
		lastData.EndTime = nextWeekEndTime.AddDate(0, 0, -1)
		lastData.WeekTh = WeekByDate(lastData.StartTime)
		weekDate = append(weekDate, lastData)
	}

	return weekDate
}

// func main() {
//
// 	l, _ := time.LoadLocation("Asia/Shanghai")
// 	startTime, _ := time.ParseInLocation("2006-01-02", "2020-05-20", l)
// 	endTime, _ := time.ParseInLocation("2006-01-02", "2020-06-15", l)
//
// 	datas := GroupByWeekDate(startTime, endTime)
// 	// if err != nil {
// 	//
// 	// }
// 	// datas = reverse(datas)
// 	for _, d := range datas {
// 		fmt.Println(d)
// 	}
//
// }

// 判断时间是当年的第几周
func WeekByDate(t time.Time) string {
	_, week := t.ISOWeek()
	return fmt.Sprintf("%d第%v周", t.Year(), week)
}

type Mame struct {
	Dat  string `json:"dat"`
	Bbbb string `json:"bbbb"`
}
type Mame1 struct {
	Typss string `json:"typss"`
	Mame
}

func reverse(s []WeekDate) []WeekDate {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// GroupByQuarterDate 将开始时间和结束时间分割为季度为单位
func GroupByQuarterDate(startTime, endTime time.Time) ([]QuarterDate, error) {
	quarterDate := make([]QuarterDate, 0)

	// 获取起始季度时间
	quarter := (int(startTime.Month()) + 2) / 3
	startQuarterTime, _, err := QuarterByDate(quarter, startTime.Year())

	if err != nil {
		return nil, err
	}

	i := 0
	for {
		currentTime := startQuarterTime.AddDate(0, 3*i, 0)
		if !currentTime.Before(endTime) {
			break
		}
		quarter := (int(currentTime.Month()) + 2) / 3
		startQuarter, endQuarter, err := QuarterByDate(quarter, currentTime.Year())

		if err != nil {
			return nil, err
		}
		quarterDate = append(quarterDate, QuarterDate{
			QuarterTh: "第" + strconv.Itoa(quarter) + "季度",
			StartTime: startQuarter,
			EndTime:   endQuarter,
		})
		i++
	}

	return quarterDate, nil
}

// QuarterDate 季度列表
type QuarterDate struct {
	QuarterTh string
	StartTime time.Time
	EndTime   time.Time
}

// QuarterByDate 获取季度对应的月份
func QuarterByDate(quarter int, year int) (startTime time.Time, endTime time.Time, err error) {
	if quarter == 1 {
		startTime, err = dateStringToTime(strconv.Itoa(year)+"-01-01", false)
		if err != nil {
			return time.Time{}, time.Time{}, err
		}
		endTime = startTime.AddDate(0, 3, -1)
	} else if quarter == 2 {
		startTime, err = dateStringToTime(strconv.Itoa(year)+"-04-01", false)
		if err != nil {
			return time.Time{}, time.Time{}, err
		}
		endTime = startTime.AddDate(0, 3, -1)
	} else if quarter == 3 {
		startTime, err = dateStringToTime(strconv.Itoa(year)+"-07-01", false)
		if err != nil {
			return time.Time{}, time.Time{}, err
		}
		endTime = startTime.AddDate(0, 3, -1)
	} else {
		startTime, err = dateStringToTime(strconv.Itoa(year)+"-10-01", false)
		if err != nil {
			return time.Time{}, time.Time{}, err
		}
		endTime = startTime.AddDate(0, 3, -1)
	}

	return startTime, endTime, nil
}

// 将字符串日期转成time的日期
func dateStringToTime(strTime string, isFormat bool) (date time.Time, err error) {
	if isFormat {
		date, err = time.ParseInLocation("2006-01-02 15:04:05", strTime, time.Local)
	} else {
		date, err = time.ParseInLocation("2006-01-02", strTime, time.Local)
	}
	return
}

// KeywordReply 关键字对应的统计量
type KeywordReply struct {
	HjAdd         string `json:"hj_add,omitempty"`         // 虎鲸添加量：日通过“首套自住/环深投资/海南置业/海外投资/广深置业/投资机会”关键字推送微信号并添加的通过的虎鲸客户人数
	SaleRecommend string `json:"sale_recommend,omitempty"` // 销售推荐量：虎鲸添加量中操作对接客户时选择的对接销售(置业顾问)并添加通过微信的客户数(以推送销售名片方式推送给客户并添加成功)
	Precise       string `json:"precise,omitempty"`        // 精准量：虎鲸添加量中在虎鲸系统标记为精准客户人数
	NotReply      string `json:"not_reply,omitempty"`      // 未回复量：虎鲸添加量中在虎鲸系统标记为未回复客户人数
}

// func main() {
// 	// 计算日期相差多少天
// 	wxIDs := make([]string,0)
// 	wxIDs = append(wxIDs,"DHZ337", "SXCF66", "we4064", "DHZ4063", "we4065")
// 	wxIDs = append(wxIDs,"ddd")
// 	fmt.Println(wxIDs)
// }

func SubDays(t1, t2 time.Time) (day int) {
	day = int(t1.Sub(t2).Hours() / 24)
	return
}

type requestBody struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

func main() {
	// timestamp := time.Now().Unix()
	// md5(appid+时间戳+body传参+密钥)
	// sign := crypto.MD5
	appID := "si_xzkj0620"
	secret := "73u4qv497d1tg0jz0w531kmbnpi31hmu"
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)

	body := "{\"startTime\":\"2020-12-08 00:00:00\",\"endTime\":\"2020-12-08 23:59:59\"}"
	fmt.Println(timestamp)
	//fmt.Println(string(bodyStr))
	str := fmt.Sprintf("%v%v%v%v", appID, timestamp, body, secret)

	// 方法一
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) // 将[]byte转成16进制

	fmt.Println(md5str1)

	// 方法二

	w := md5.New()
	_, _ = io.WriteString(w, str)            // 将str写入到w中
	md5str2 := fmt.Sprintf("%x", w.Sum(nil)) // w.Sum(nil)将w的hash转成[]byte格式

	fmt.Println(md5str2)

}


// func main() {
// 	var (
// 		a, b float64
// 		c    int64
// 	)
// 	a = 2.55
// 	b = 0.0
//
// 	c = int64((a + b) * 100.0)
// 	fmt.Printf("第1次 c=%d\n", c) // 第1次 c=254
//
// 	c = int64(a * 100.0)
// 	fmt.Printf("第2次 c=%d\n", c) // 第2次 c=254
//
// 	fmt.Println(a * 100) // 254.99999999999997
// 	// 发现将小数的元乘以100后强制转换为整数分，少了1分
// 	// 解决办发1：
// 	tmpStr1 := fmt.Sprintf("%.2f", a)
// 	tmpStr2 := fmt.Sprintf("%.2f", b)
//
// 	tmpnum1, _ := strconv.ParseInt(strings.Replace(tmpStr1, ".", "", 1), 10, 64)
// 	tmpnum2, _ := strconv.ParseInt(strings.Replace(tmpStr2, ".", "", 1), 10, 64)
// 	c = tmpnum1 + tmpnum2
// 	fmt.Printf("第3次 c=%d\n", c) // 第3次 c=255
// }