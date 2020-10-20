package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

//
///*
//#include <windows.h>
//#include <conio.h>
//// 使用了WinAPI来移动控制台的光标
//void gotoxy(int x,int y)
//{
//    COORD c;
//    c.X=x,c.Y=y;
//    SetConsoleCursorPosition(GetStdHandle(STD_OUTPUT_HANDLE),c);
//}
//// 从键盘获取一次按键，但不显示到控制台
//int direct()
//{
//    return _getch();
//}
//*/
//import "C" // go中可以嵌入C语言的函数
//
//// 表示光标的位置
//type loct struct {
//	i, j int
//}
//
//var (
//	area = [20][20]byte{} // 记录了蛇、食物的信息
//	food bool             // 当前是否有食物
//	lead byte             // 当前蛇头移动方向
//	head loct             // 当前蛇头位置
//	tail loct             // 当前蛇尾位置
//	size int              // 当前蛇身长度
//)
//
//// 随机生成一个位置，来放置食物
//func place() loct {
//	k := rand.Int() % 400
//	return loct{k / 20, k % 20}
//}
//
//// 用来更新控制台的显示，在指定位置写字符，使用错误输出避免缓冲
//func draw(p loct, c byte) {
//	C.gotoxy(C.int(p.i*2+4), C.int(p.j+2))
//	fmt.Fprintf(os.Stderr, "%c", c)
//}
//
//func init() {
//
//	// 初始化蛇的位置和方向、首尾；初始化随机数
//	head, tail = loct{4, 4}, loct{4, 4}
//	lead, size = 'R', 1
//	area[4][4] = 'H'
//	rand.Seed(int64(time.Now().Unix()))
//
//	// 输出初始画面
//	fmt.Fprintln(os.Stderr,
//		`
//  #-----------------------------------------#
//  |                                         |
//  |                                         |
//  |                                         |
//  |                                         |
//  |         *                               |
//  |                                         |
//  |                                         |
//  |                                         |
//  |                                         |
//  |                                         |
//  |                                         |
//  |                                         |
//  |                                         |
//  |                                         |
//  |                                         |
//  |                                         |
//  |                                         |
//  |                                         |
//  |                                         |
//  |                                         |
//  #-----------------------------------------#
//`)
//
//	// 我们使用一个单独的go程来捕捉键盘的动作，因为是单独的，不怕阻塞
//	go func() {
//		for { // 函数只写入lead，外部只读取lead，无需设锁
//			switch byte(C.direct()) {
//			case 72:
//				lead = 'U'
//			case 75:
//				lead = 'L'
//			case 77:
//				lead = 'R'
//			case 80:
//				lead = 'D'
//			case 32:
//				lead = 'P'
//			}
//		}
//	}()
//}
//
//func main() {
//
//	// 主程序
//	for {
//
//		// 程序更新周期，400毫秒
//		time.Sleep(time.Millisecond * 400)
//
//		// 暂停，还是要有滴
//		if lead == 'P' {
//			continue
//		}
//
//		// 放置食物
//		if !food {
//			give := place()
//			if area[give.i][give.j] == 0 { // 食物只能放在空闲位置
//				area[give.i][give.j] = 'F'
//				draw(give, '$') // 绘制食物
//				food = true
//			}
//		}
//
//		// 我们在蛇头位置记录它移动的方向
//		area[head.i][head.j] = lead
//
//		// 根据lead来移动蛇头
//		switch lead {
//		case 'U':
//			head.j--
//		case 'L':
//			head.i--
//		case 'R':
//			head.i++
//		case 'D':
//			head.j++
//		}
//
//		// 判断蛇头是否出界
//		if head.i < 0 || head.i >= 20 || head.j < 0 || head.j >= 20 {
//			C.gotoxy(0, 23) // 让光标移动到画面下方
//			break           // 跳出死循环
//		}
//
//		// 获取蛇头位置的原值，来判断是否撞车，或者吃到食物
//		eat := area[head.i][head.j]
//
//		if eat == 'F' { // 吃到食物
//			food = false
//
//			// 增加蛇的尺寸，并且不移动蛇尾
//			size++
//		} else if eat == 0 { // 普通移动
//
//			draw(tail, ' ') // 擦除蛇尾
//
//			// 注意我们记录了它移动的方向
//			dir := area[tail.i][tail.j]
//
//			// 我们需要擦除蛇尾的记录
//			area[tail.i][tail.j] = 0
//
//			// 移动蛇尾
//			switch dir {
//			case 'U':
//				tail.j--
//			case 'L':
//				tail.i--
//			case 'R':
//				tail.i++
//			case 'D':
//				tail.j++
//			}
//		} else { // 撞车了
//			C.gotoxy(0, 23)
//			break
//		}
//		draw(head, '*') // 绘制蛇头
//	}
//
//	// 收尾了
//	switch {
//	case size < 22:
//		fmt.Fprintf(os.Stderr, "Faild! You've eaten %d $\\n", size-1)
//	case size < 42:
//		fmt.Fprintf(os.Stderr, "Try your best! You've eaten %d $\\n", size-1)
//	default:
//		fmt.Fprintf(os.Stderr, "Congratulations! You've eaten %d $\\n", size-1)
//	}
//}

//func main() {
//	var (
//		paykey   = "9c2773ce639d2681498c85f8fb8770b6"
//		req_info = "bg/rcry1fKG00mc7XdAIc2Ln0iDnpUIBNhwr0C7cCkVCuCtz6ryousAtY3roecJKJ+Go4Y9TOnCepYT4IMtS+C9DfJg5hC5+xcaqF0lhvRKwUOMxdJ9m0defrTQM40UdVk9Fk4KBd4d50Z1d7DCuudA09cfUbFKODk1ZcDGtJwuRe5UqbDqSWmTuBW01PXILhrbgoYwY+dxkLBmM1WxVyFIggavyhnjq7fCXh3mAOeb9467bRhg8ro6/9HDkY81KvwJJpDTgDvkJ802x/Tzk5fJ+Zsl9aYEN6YZAJdWY6QOmrdo/50iAhqOXTCQirWkmW/dG2zuHexhIymXUog+O82861x+m8G1+QhThSycqMlVRMmh3Gra34iq2vCgNQmpYtZ+eId9u+Bosu6Vpim6eEwBNt39NEwGLnJRojnksydB1iCabd0RxTk/on5zlheBwe6m9efbOP5mVEz1c35hWDlQRrZ1sCNwk1yJa5mEJRw6ahtdnkN11noM0iUsJEk7W3TK30exMoWR8KNrc6w/MdFxNH6TuNlE5K5Wdr/CoU6v5yNOO5uSrCNhRVtvo+wLsvzlbp7K5+XDhN1kWk1Cm800a2XysCkyoNyDGnTfJhJyhQCU8Fe8ltaVawOan58lHhfcZwrfZOxA5XXT3k0KOGywFlBnizR+Yc81wis4RRWGoO0RbUsghDFfwE3F4AD8agfbpTkqYA3CZPQgnfmL/cxb+25T1zTtLOpAnM+gFlQC4NiuIv7y0tiAhJAbhbqmvdRxkkLpltvLSeME6Y9EM6z0A2PS/2cPSoWBR/1XGTZncBYIoGqjuP6btaZgU0SnfALzOkWxKIEbpPTX1+udkpUVmHGRibs2aP/tVjrYbqg/41Rr7vqeymaIX7Cb7zRkMxFSStOTwZDokRVjkhMAQu0Q/Dd6xJ4KwMwEuamg3TS5yWjbPuZrgz9ov9x/awsL+JXYQeSccHl3OwZbwLmVPMYd9qDLc/Y0dHfxvQ2XdlR12M2Zj0AmPvdwr8iACY292lluXWxDcl4gJligbYH0HlPUNeptZgoMNa+EdmIMZIH9H6+rD6WkrGCe3Q7d1nju1Tly0w/hSK2j+SVzSHaZo3yEYoOYrHPz2YJyfa6tXbEBtybSlfhjYOXch5ZRcu+zwmnOYYnP+C5B8EQlffm+O6n3Mp24k49tqnxGOdFuK1fS4NiuIv7y0tiAhJAbhbqmv/7xRGN7cr0CIbp62Fk+aQ76BFBzC+kE+pbd2elrB2/9lBaLvq2j2Lte0KJcW79DlSuQGhBZql0F9eCUXy7EZ0he/LR1wedu5jkqjeSSvJeuQ82fGxdOIM9EGi5UcwlSApdTmOPpV89MDoSZTsrDHdBmmWuSCoI0OmsNZVXP3fIMKV2p12teHq9R3pRKSzHJxygmW/jzEsqHTvo2QVIHzUZOfFOJFYmZkZL6RA7d4sL0hZf8G5ODtw2W1QfHQ6qQO2BA85ES2KYE/ByR1NiVdazl1i8sH4YRh0nL1khCKdiK40q/nMf8yLdH+pwXtTcZoQuVxyRBd4bZDC+XnKd8ME9kgIWN7HJrQg6fChlRmdQuQILQsadkoPgrWAI6yXC8b"
//	)
//	src := []byte(req_info)
//	key := []byte(paykey)
//
//	dst , _ := openssl.AesECBDecrypt(src, key, openssl.PKCS7_PADDING)
//	fmt.Println(string(dst)) // 123456
//}

//  SmCustomerInfoExportByTimeResult  sm接口获取指定时间之内发生过更新的所有客户详情-结果
type SmCustomerInfoExportByTimeResponse struct {
	Code    int64                           `json:"code"`
	Data    SmCustomerInfoExportByTimeDatas `json:"data"`
	Message string                          `json:"message"`
}

//  SmCustomerInfoExportByTimeData  sm接口获取指定时间之内发生过更新的所有客户详情-列表
type SmCustomerInfoExportByTimeData struct {
	Age2           interface{} `json:"age2"`
	Birthday       string      `json:"birthday"`
	Phone1         interface{} `json:"phone1"`
	RealName       interface{} `json:"real_name"`
	SalesAccount   string      `json:"salesAccount"`
	SalesWx        string      `json:"salesWx"`
	Sex            interface{} `json:"sex"`
	WechatAlias    interface{} `json:"wechatAlias"`
	WechatComment  string      `json:"wechatComment"`
	WechatNickname string      `json:"wechatNickname"`
	WechatNumber   string      `json:"wechatNumber"`
}

// SmCustomerInfoExportByTimeDatas sm接口获取指定时间之内发生过更新的所有客户详情-列表
type SmCustomerInfoExportByTimeDatas []SmCustomerInfoExportByTimeData

func main() {

	body := []byte("{\"code\":0,\"message\":\"success\",\"data\":[{\"birthday\":\"\",\"salesAccount\":\"XN70014\",\"wechatComment\":\"vip 环深 中珠海😉🎀190505\",\"wechatAlias\":null,\"sex\":null,\"salesWx\":\"wxid_0w2uj547nmyk12\",\"wechatNickname\":\"💫梦宇🍑🌝\",\"real_name\":null,\"wechatNumber\":\"dongmengyu29\",\"age2\":null,\"phone1\":null},{\"birthday\":\"\",\"salesAccount\":\"XN70023\",\"wechatComment\":\"20.10.14💫梦宇🍑🌝\",\"wechatAlias\":null,\"sex\":null,\"salesWx\":\"wxid_0qfsnvf9ow8b22\",\"wechatNickname\":\"💫梦宇🍑🌝\",\"real_name\":null,\"wechatNumber\":\"dongmengyu29\",\"age2\":null,\"phone1\":null},{\"birthday\":\"\",\"salesAccount\":null,\"wechatComment\":\"181219彩云\",\"wechatAlias\":null,\"sex\":null,\"salesWx\":\"wxid_g2paz7qxp8ut12\",\"wechatNickname\":\"彩云\",\"real_name\":null,\"wechatNumber\":\"wxid_jy002xhy79pg22\",\"age2\":null,\"phone1\":null},{\"birthday\":\"\",\"salesAccount\":\"XN70014\",\"wechatComment\":\"19.01.17  低南沙 诊断意向\",\"wechatAlias\":null,\"sex\":null,\"salesWx\":\"wxid_61kijkrnuo7l12\",\"wechatNickname\":\"彩云\",\"real_name\":null,\"wechatNumber\":\"wxid_jy002xhy79pg22\",\"age2\":null,\"phone1\":null},{\"birthday\":\"\",\"salesAccount\":\"XN70023\",\"wechatComment\":\"190717彩云\",\"wechatAlias\":null,\"sex\":null,\"salesWx\":\"wxid_vhbxkxiclg6622\",\"wechatNickname\":\"彩云\",\"real_name\":null,\"wechatNumber\":\"wxid_jy002xhy79pg22\",\"age2\":null,\"phone1\":null},{\"birthday\":\"\",\"salesAccount\":\"XN70024\",\"wechatComment\":\"20.10.14彩云\",\"wechatAlias\":null,\"sex\":null,\"salesWx\":\"wxid_z2a2pyxey7sy22\",\"wechatNickname\":\"彩云\",\"real_name\":null,\"wechatNumber\":\"wxid_jy002xhy79pg22\",\"age2\":null,\"phone1\":null},{\"birthday\":\"\",\"salesAccount\":\"XN70020\",\"wechatComment\":null,\"wechatAlias\":\"A13266136628\",\"sex\":null,\"salesWx\":\"wxid_39o439zpjj2p22\",\"wechatNickname\":\"淡雅如菊\",\"real_name\":null,\"wechatNumber\":\"wxid_oh5873lpc7k622\",\"age2\":null,\"phone1\":null},{\"birthday\":\"\",\"salesAccount\":\"XN70023\",\"wechatComment\":\"高惠阳淡雅如菊\",\"wechatAlias\":\"A13266136628\",\"sex\":null,\"salesWx\":\"wxid_0qfsnvf9ow8b22\",\"wechatNickname\":\"淡雅如菊\",\"real_name\":null,\"wechatNumber\":\"wxid_oh5873lpc7k622\",\"age2\":null,\"phone1\":null},{\"birthday\":\"\",\"salesAccount\":\"XN70015\",\"wechatComment\":\"20.10.14淡雅如菊\",\"wechatAlias\":\"A13266136628\",\"sex\":null,\"salesWx\":\"wxid_swmjvfpdw97p12\",\"wechatNickname\":\"淡雅如菊\",\"real_name\":null,\"wechatNumber\":\"wxid_oh5873lpc7k622\",\"age2\":null,\"phone1\":null},{\"birthday\":\"\",\"salesAccount\":\"XN70024\",\"wechatComment\":null,\"wechatAlias\":null,\"sex\":null,\"salesWx\":\"wxid_z2a2pyxey7sy22\",\"wechatNickname\":\"不叮点\",\"real_name\":null,\"wechatNumber\":\"wangdongnevison\",\"age2\":null,\"phone1\":null},{\"birthday\":\"\",\"salesAccount\":\"XN70020\",\"wechatComment\":\"20.10.02🍃caterpillar🍃\",\"wechatAlias\":\"Caterpillar658865\",\"sex\":null,\"salesWx\":\"wxid_39o439zpjj2p22\",\"wechatNickname\":\"🍃caterpillar🍃\",\"real_name\":null,\"wechatNumber\":\"liuting658865\",\"age2\":null,\"phone1\":null}]}")

	// 解析response
	var item SmCustomerInfoExportByTimeResponse
	if err := json.Unmarshal(body, &item); err != nil {
		fmt.Println(err)
	}

	for _, v := range item.Data {
		tags := []string{"低", "中", "高"}
		areas := []string{"惠阳", "南沙", "清远", "珠海"}
		for _, tag := range tags {
			for _, area := range areas {
				if v.WechatComment != "" && strings.Contains(v.WechatComment, tag+area) {
					fmt.Println(v.WechatComment)
				}
			}
		}
	}

}
