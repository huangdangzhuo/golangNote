package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/forgoer/openssl"
	"math/rand"
	"time"
)

// 未加锁
// func add(count *int, wg *sync.WaitGroup)   {
// 	for i:= 0;i<1000;i++ {
// 		*count +=1
// 	}
// 	wg.Done()
// }
//
// func main()  {
// 	var wg sync.WaitGroup
// 	count := 0
// 	wg.Add(3)
// 	go add(&count,&wg)
// 	go add(&count,&wg)
// 	go add(&count,&wg)
//
// 	wg.Wait()
// 	fmt.Println("count的值为",count)
// }

// 枷锁
// func add(count *int, wg *sync.WaitGroup,lock *sync.Mutex)   {
// 	for i:= 0;i<1000;i++ {
// 		lock.Lock()
// 		*count +=1
// 		lock.Unlock()
// 	}
// 	wg.Done()
// }
//
// func main()  {
// 	var wg sync.WaitGroup
// 	lock := &sync.Mutex{}
// 	count := 0
// 	wg.Add(3)
// 	go add(&count,&wg,lock)
// 	go add(&count,&wg,lock)
// 	go add(&count,&wg,lock)
//
// 	wg.Wait()
// 	fmt.Println("count的值为",count)
// }

// func main()  {
// 	lock := &sync.RWMutex{}
// 	lock.Lock()
// 	for i:=0;i<4;i++ {
// 		go func(i int) {
// 			fmt.Printf("第%d个协程准备开始...\n",i)
// 			lock.RLock()
// 			fmt.Printf("第%d个协程获取读锁,sleep 1s ,释放锁\n",i)
// 			time.Sleep(time.Second)
// 			lock.RUnlock()
// 		}(i)
// 	}
// 	time.Sleep(time.Second*2)
// 	fmt.Println("准备释放写锁，读锁不再阻塞")
// 	// 写锁一释放，读锁就自由了
// 	lock.Unlock()
//
// 	// 由于会等到读锁全部释放，才能获得写锁
// 	// 因为这里一定会在上面 4 个协程全部完成才能往下走
// 	lock.Lock()
// 	fmt.Println("程序退出...")
// 	lock.Unlock()
// }

// IyNotifyResponse i云保-回调通知结构
type IyNotifyResponse struct {
	Data string `json:"data"`
	Sign string `json:"sign"`
}

// IyNotifyReturnResponse i云保-回调通知响应结构
type IyNotifyReturnResponse struct {
	Code int    `json:"code"` // 响应代码，0-成功；1-失败；2-异常
	Msg  string `json:"msg"`  // 消息，如果code不为0，可以输出对应的错误消息
}

// IyOrderNotifyResponse i云保-保单出单通知
type IyOrderNotifyResponse struct {
	Partner           string `json:"partner"`           // 渠道商id，i云服对渠道的唯一标示，如何获取见：开发前必读-基本概念介绍
	ApplicantName     string `json:"applicantName"`     // 投保人姓名
	ApplicantGender   string `json:"applicantGender"`   // 投保人性别，枚举值：M =男、F = 女
	InsureTime        string `json:"insureTime"`        // 投保日期/投保时间，如：2016-12-22 23:31:34
	Amount            int    `json:"amount"`            // 保额，3006000
	ApplicantBirthday string `json:"applicantBirthday"` //
	ApplicantCerType  string `json:"applicantCerType"`  //
	Company           string `json:"company"`           //  产品所属保险公司，如：众安
	EffectiveTime     string `json:"effectiveTime"`     // 保单生效日期，如：2017-02-03 23:31:34
	Expires           string `json:"expires"`           // 保障期，1年 码表见12.4
	ExpiryTime        string `json:"expiryTime"`        // 保单终止日期，如：2018-02-02 23:31:34
	InsurantBirthday  string `json:"insurantBirthday"`  // 被保人生日，如：2018-02-02
	InsurantCerType   string `json:"insurantCerType"`   //
	InsurantGender    string `json:"insurantGender"`    // 被保人性别，枚举值：M = 男，F = 女
	InsurantName      string `json:"insurantName"`      //  被保人姓名，
	InsureData        string `json:"insureData"`        // 承保日期，如：2018-02-02 23:31:34
	Mobile            string `json:"mobile"`            // 投保人手机号，
	OrderNo           string `json:"orderNo"`           // 订单号，如：IYB23852094323423
	PartnerTag        string `json:"partnerTag"`        // 合作方自定义标签，
	PayFreq           string `json:"payFreq"`           // 缴费频次，一次交清, 年交, 月交, 季交,半年交
	Payment           string `json:"payment"`           // 缴费期，1年 码表见12.4
	PolicyNo          string `json:"policyNo"`          //  保单号，IH23852094323423
	PolicyStatus      int    `json:"policyStatus"`      // 保单状态，3-承保
	Price             int    `json:"price"`             // 保费，200
	ProductCode       string `json:"productCode"`       //
	ProductID         string `json:"productId"`         // 保障计划版本id
	ProductName       string `json:"productName"`       // 产品名称/险种名称，如：尊享e生旗舰版
	RelationToStaff   string `json:"relationToStaff"`   //
	Type              int    `json:"type"`              // 产品类型，1长期险 2短期险
	VendorID          int    `json:"vendorId"`          // 公司ID，
	PayCode           string `json:"pay_code"`          //  缴费期code，term_1 码表见12.4
	InsureCode        string `json:"insure_code"`       // 保障期code，to_21码表见12.4
	WareID            string `json:"wareId"`            // 商品ID，一个商品id可能对应多个productid
	AccountID         string `json:"accountId"`         // 帐号id，i云服用户的唯一标示，如何获取见：开发前必读-基本概念介绍
}

// IyOrderStatusNotifyResponse 保单状态变更通知
type IyOrderStatusNotifyResponse struct {
	OperationTime string `json:"operationTime"` // 退保操作时间，如：2016-12-22 23:31:34
	Partner       string `json:"partner"`       // 渠道商id，i云服对渠道的唯一标示，如何获取见：开发前必读-基本概念介绍
	PartnerTag    string `json:"partnerTag"`    // 渠道自定义参数
	PolicyNo      string `json:"policyNo"`      // 保单号，如：IH23852094323423
	PolicyStatus  int    `json:"policyStatus"`  // 保单状态，8=犹退,9=退保
	VendorID      int    `json:"vendorId"`      // 保单所属公司ID，如：4
	SurrenderTime string `json:"surrenderTime"` // 退保时间，如：2016-12-22 00:00:00
	HestiTime     string `json:"hestiTime"`     // 犹豫期截止时间，如：2016-12-22 00:00:00
}

func MD5Hash(b []byte) string {
	h := md5.New()
	h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}
func GetTimeTick64() int64 {
	return time.Now().UnixNano() / 1e6
}

func GetTimeTick32() int32 {
	return int32(time.Now().Unix())
}

func GetFormatTime(time time.Time) string {
	return time.Format("20060102")
}

// 基础做法 日期20191025时间戳1571987125435+3位随机数
func GenerateCode() {
	date := GetFormatTime(time.Now())
	r := rand.Intn(1000)
	code := fmt.Sprintf("%s%d%03d", date, GetTimeTick64(), r)
	fmt.Println(code, " rand ID generate successed!\n")
}

func main() {
	src := []byte("bg/rcry1fKG00mc7XdAIc2Ln0iDnpUIBNhwr0C7cCkVCuCtz6ryousAtY3roecJKJ+Go4Y9TOnCepYT4IMtS+C9DfJg5hC5+xcaqF0lhvRKwUOMxdJ9m0defrTQM40UdVk9Fk4KBd4d50Z1d7DCuudA09cfUbFKODk1ZcDGtJwuRe5UqbDqSWmTuBW01PXILhrbgoYwY+dxkLBmM1WxVyFIggavyhnjq7fCXh3mAOeb9467bRhg8ro6/9HDkY81KvwJJpDTgDvkJ802x/Tzk5fJ+Zsl9aYEN6YZAJdWY6QOmrdo/50iAhqOXTCQirWkmW/dG2zuHexhIymXUog+O82861x+m8G1+QhThSycqMlVRMmh3Gra34iq2vCgNQmpYtZ+eId9u+Bosu6Vpim6eEwBNt39NEwGLnJRojnksydB1iCabd0RxTk/on5zlheBwe6m9efbOP5mVEz1c35hWDlQRrZ1sCNwk1yJa5mEJRw6ahtdnkN11noM0iUsJEk7W3TK30exMoWR8KNrc6w/MdFxNH6TuNlE5K5Wdr/CoU6v5yNOO5uSrCNhRVtvo+wLsvzlbp7K5+XDhN1kWk1Cm800a2XysCkyoNyDGnTfJhJyhQCU8Fe8ltaVawOan58lHhfcZwrfZOxA5XXT3k0KOGywFlBnizR+Yc81wis4RRWGoO0RbUsghDFfwE3F4AD8agfbpTkqYA3CZPQgnfmL/cxb+25T1zTtLOpAnM+gFlQC4NiuIv7y0tiAhJAbhbqmvdRxkkLpltvLSeME6Y9EM6z0A2PS/2cPSoWBR/1XGTZncBYIoGqjuP6btaZgU0SnfALzOkWxKIEbpPTX1+udkpUVmHGRibs2aP/tVjrYbqg/41Rr7vqeymaIX7Cb7zRkMxFSStOTwZDokRVjkhMAQu0Q/Dd6xJ4KwMwEuamg3TS5yWjbPuZrgz9ov9x/awsL+JXYQeSccHl3OwZbwLmVPMYd9qDLc/Y0dHfxvQ2XdlR12M2Zj0AmPvdwr8iACY292lluXWxDcl4gJligbYH0HlPUNeptZgoMNa+EdmIMZIH9H6+rD6WkrGCe3Q7d1nju1Tly0w/hSK2j+SVzSHaZo3yEYoOYrHPz2YJyfa6tXbEBtybSlfhjYOXch5ZRcu+zwmnOYYnP+C5B8EQlffm+O6n3Mp24k49tqnxGOdFuK1fS4NiuIv7y0tiAhJAbhbqmv/7xRGN7cr0CIbp62Fk+aQ76BFBzC+kE+pbd2elrB2/9lBaLvq2j2Lte0KJcW79DlSuQGhBZql0F9eCUXy7EZ0he/LR1wedu5jkqjeSSvJeuQ82fGxdOIM9EGi5UcwlSApdTmOPpV89MDoSZTsrDHdBmmWuSCoI0OmsNZVXP3fIMKV2p12teHq9R3pRKSzHJxygmW/jzEsqHTvo2QVIHzUZOfFOJFYmZkZL6RA7d4sL0hZf8G5ODtw2W1QfHQ6qQO2BA85ES2KYE/ByR1NiVdazl1i8sH4YRh0nL1khCKdiK40q/nMf8yLdH+pwXtTcZoQuVxyRBd4bZDC+XnKd8ME9kgIWN7HJrQg6fChlRmdQuQILQsadkoPgrWAI6yXC8b")
	key := []byte("9c2773ce639d2681498c85f8fb8770b6")
	src, _ = base64.StdEncoding.DecodeString(string(src))

	dst, _ := openssl.AesECBDecrypt(src, key, openssl.PKCS5_PADDING)
	fmt.Println(string(dst)) // 123456
}
