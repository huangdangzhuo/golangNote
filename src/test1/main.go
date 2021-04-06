package main

import "fmt"

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
//
//func main() {
//	var (
//		paykey   = "8ceaa9dba9f8c199df5c5a2f5c169f5c"
//		req_info = "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"
//	)
//	//src := []byte(req_info)
//	key := []byte(paykey)
//
//	//dst , _ := openssl.AesECBDecrypt(src, key, openssl.PKCS7_PADDING)
//	//fmt.Println(string(dst)) // 123456
//	// aes-256-ecb 解密 body
//	// 解密 待开发
//	decryptBody, _ := base64.StdEncoding.DecodeString(req_info)
//
//	decryptBody, _ = openssl.AesECBDecrypt(decryptBody, key, openssl.PKCS5_PADDING)
//	fmt.Println(string(decryptBody))
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

type addOrderHouse struct {
	OrderNo                 string           `json:"order_no"`                                     // 购房订单号
	CustomerID              string           `json:"customer_id" binding:"required"`               // 关联客户的记录ID
	CustomerName            string           `json:"customer_name" binding:"required"`             // 客户姓名
	CustomerType            int              `json:"customer_type" binding:"required"`             // 客户类别(1今玖客户2渠道客户)
	CustomerPhone           string           `json:"customer_phone" binding:"required"`            // 客户电话
	OrderTime               string           `json:"order_time" binding:"required"`                // 成单日期
	Area                    int              `json:"area" binding:"required"`                      // 区域(1惠阳2海南3柬埔寨4泰国5珠海6广州7深圳)
	ProjectName             string           `json:"project_name" binding:"required"`              // 项目名称
	HouseCode               string           `json:"house_code" binding:"required"`                // 购买房号
	DealArea                float64          `json:"deal_area" binding:"required"`                 // 成交面积(平方)
	DealAmount              float64          `json:"deal_amount" binding:"required"`               // 成交金额(元)
	CommissionRate          float64          `json:"commission_rate" binding:"required"`           // 佣金比例
	IsInstallment           int              `json:"is_installment" binding:"required"`            // 是否分期(1是2否)
	FirstPayTime            string           `json:"first_pay_time" binding:"required"`            // 首付时间
	FirstPayCommissionRate  float64          `json:"first_pay_commission_rate"`                    // 首付佣金比例
	FirstPaySettlementTime  string           `json:"first_pay_settlement_time"`                    // 首付结算时间
	LastPayCommissionRate   float64          `json:"last_pay_commission_rate"`                     // 尾款佣金比例
	LastPayCommission       float64          `json:"last_pay_commission"`                          // 尾款佣金(元)
	LastPaySettlementTime   string           `json:"last_pay_settlement_time"`                     // 尾款结算时间
	USDRate                 float64          `json:"usd_rate"`                                     // 美金汇率
	ChannelCommission       float64          `json:"channel_commission" binding:"required"`        // 渠道分销佣金(元)
	PayWay                  int              `json:"pay_way" binding:"required"`                   // 付款方式(1线上2线下)
	MortgageTime            string           `json:"mortgage_time" binding:"required"`             // 办理按揭时间
	SettlementCommissionWay string           `json:"settlement_commission_way" binding:"required"` // 结佣渠道
	DepartmentID            string           `json:"department_id" binding:"required"`             // 部门ID
	FollowUserID            string           `json:"follow_user_id" binding:"required"`            // 跟进人的记录ID
	GroupUser               []string         `json:"group_user"`                                   // 带团人的记录ID列表
	Remark                  string           `json:"remark"`                                       // 备注
	Creator                 string           `json:"creator"`                                      // 创建者
	Images                  OrderHouseImages `json:"images"`                                       // 认购书图片列表
}

type OrderHouseImage struct {
	RecordID  string `json:"record_id"`  // 记录ID
	Name      string `json:"name"`       // 名称
	AliyunURL string `json:"aliyun_url"` // 阿里云地址(不带域名)
	HouseID   string `json:"house_id"`   // 关联购房订单的记录ID
}

// OrderHouseImages 房产购房订单认购书图片管理显示项列表
type OrderHouseImages []*OrderHouseImage

//func main()  {
//	var old OrderHouseImage
//	old.RecordID = "1111"
//
//	newItem := old
//	old.Name ="122354156"
//	newItem.RecordID = "222"
//	fmt.Printf("%+v",newItem)
//	fmt.Printf("%+v",old)
//}

//func main() {
//
//	str := "{\"caseCode\":\"QX000000128707\",\"insureNum\":\"20201030012510\",\"partnerId\":1048483,\"policys\":[{\"applicant\":\"陈嘉瑜\",\"endDate\":\"2021-10-31\",\"insurant\":\"陈嘉瑜\",\"planId\":128707,\"planName\":\"尊享版\",\"policyNum\":\"BD0122012006001200016267\",\"productId\":104151,\"productName\":\"锦一卫·个人意外综合保障计划\",\"startDate\":\"2020-10-31\"}],\"state\":true}"
//	//sdd :="{\"caseCode\":\"QX000000128707\",\"insureNum\":\"20201030012510\",\"partnerId\":1048483,\"policys\":[{\"applicant\":\"陈嘉瑜\",\"endDate\":\"2021-10-31\",\"insurant\":\"陈嘉瑜\",\"planId\":128707,\"planName\":\"尊享版\",\"policyNum\":\"BD0122012006001200016267\",\"productId\":104151,\"productName\":\"锦一卫·个人意外综合保障计划\",\"startDate\":\"2020-10-31\"}],\"state\":true}"
//
//	sign := fmt.Sprintf("%x", md5.Sum([]byte("ONOGI3MDU2MjE4Yjc1048460"+str)))
//	fmt.Println(sign)
//
//}

func main()  {
	//	now := time.Now()
	//
	//	offset := int(time.Monday - now.Weekday())
	//	if offset > 0 {
	//		offset = -6
	//	}
	//
	//	//weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	//fmt.Println( now.Format("01月02日"))

	var a bool
	fmt.Println(a)

}