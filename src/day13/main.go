package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"log"
	"sync"
	"time"
)

func test() {
	var mut sync.Mutex
	maxSize := 10
	counter := 0

	// 排水口
	go func() {
		for {
			mut.Lock()
			if counter == maxSize {
				for i := 0; i < maxSize; i++ {
					counter--
					log.Printf(" OUTPUT counter = %d", counter)
				}
			}
			mut.Unlock()
			time.Sleep(1 * time.Second)
		}
	}()

	// 注水口
	for {
		mut.Lock()
		if counter == 0 {
			for i := 0; i < maxSize; i++ {
				counter++
				log.Printf(" INPUT counter = %d", counter)
			}
		}
		mut.Unlock()
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var req_info = "eyJyZXF1ZXN0Ijp7ImVudGVycHJpc2VJbmZvIjp7fSwiYXBwbHlJbmZvIjp7ImhvbGRlckluZm8iOnsiaG9sZGVySWRXZWlnaHQiOiIiLCJob2xkZXJJbmNvbWVUeXBlTmFtZSI6IiIsImhvbGRlckNvbXBhbnlOYW1lIjoiIiwiaG9sZGVySWRIZWlnaHQiOiIiLCJob2xkZXJQcm92aW5jZU5hbWUiOiIiLCJob2xkZXJDaXR5SWQiOiIiLCJob2xkZXJEaXN0cmljdElkIjoiIiwiaG9sZGVyUHJvZmVzc2lvbk5hbWUiOiIiLCJob2xkZXJDb21wYW55QWRkcmVzcyI6IiIsImhvbGRlcklkVmFpbGRFbmREYXRlIjoiIiwiaG9sZGVyUHJvdmluY2VJZCI6IiIsImhvbGRlckdlbmRlciI6IjIiLCJob2xkZXJOYW1lIjoi6LS+5YehIiwiaG9sZGVyRW1haWwiOiIiLCJob2xkZXJDaXR5TmFtZSI6IiIsImhvbGRlckluY29tZVR5cGVJZCI6IiIsImhvbGRlckRpc3RyaWN0TmFtZSI6IiIsImhvbGRlcklkVmFpbGRTdGFydERhdGUiOiIiLCJob2xkZXJBZGRyZXNzIjoiIiwiaG9sZGVyVGF4UmVzaWRlbnRzIjoiIiwiaG9sZGVyUG9zdENvZGUiOiIiLCJob2xkZXJNYXJpdGFsU3RhdHVzIjoiIiwiaG9sZGVyQ2FyZE5vIjoiNjEyNTAxMTk5MTA5MjgwMDI3IiwiaG9sZGVyUHJvZmVzc2lvbklkIjoiIiwiaG9sZGVyQmlydGhkYXkiOiIxOTkxLTA5LTI4IiwiaG9sZGVyTW9iaWxlIjoiMTMyNTkxODAwNjIiLCJob2xkZXJDYXJkVHlwZSI6IjEiLCJob2xkZXJNYXJpdGFsU3RhdHVzTmFtZSI6IiJ9LCJpbnN1cmVkSW5mbyI6eyJpbnN1cmVkTGlzdCI6W3siaW5zdXJlZENpdHlJZCI6IiIsImluc3VyZWRHZW5kZXIiOiIyIiwiaW5zdXJlZElkV2VpZ2h0IjoiIiwiaW5zdXJlZENhcmRObyI6IjYxMjUwMTE5OTEwOTI4MDAyNyIsImluc3VyZWRJZFZhaWxkRW5kRGF0ZSI6IiIsImluc3VyZWRJZEhlaWdodCI6IiIsImluc3VyZWRDb21wYW55TmFtZSI6IiIsImluc3VyZWREaXN0cmljdElkIjoiIiwiaW5zdXJlZEluY29tZVR5cGVJZCI6IiIsImluc3VyZWRQcm92aW5jZUlkIjoiIiwiaW5zdXJlZERpc3RyaWN0TmFtZSI6IiIsImJlbmVmaXRJbmZvIjp7ImlzTGVnYWwiOiJ0cnVlIiwiYmVuZWZpdExpc3QiOltdfSwiaW5zdXJlZFRheFJlc2lkZW50cyI6IiIsImluc3VyZWROYW1lIjoi6LS+5YehIiwiaW5zdXJlZEluY29tZVR5cGVOYW1lIjoiIiwiaW5zdXJlZENpdHlOYW1lIjoiIiwiaW5zdXJlZFBvc3RDb2RlIjoiIiwiaW5zdXJlZE1vYmlsZSI6IjEzMjU5MTgwMDYyIiwiaW5zdXJlZFByb2Zlc3Npb25JZCI6IiIsImluc3VyZWRNYXJpdGFsU3RhdHVzTmFtZSI6IiIsImluc3VyZWRQcm9mZXNzaW9uTmFtZSI6IiIsImluc3VyZWRDb21wYW55QWRkcmVzcyI6IiIsImluc3VyZWRSZWxhdGlvbiI6IjEiLCJpbnN1cmVkQ2FyZFR5cGUiOiIxIiwiaW5zdXJlZEVtYWlsIjoiIiwiaW5zdXJlZEJpcnRoZGF5IjoiMTk5MS0wOS0yOCIsImluc3VyZWRBZGRyZXNzIjoiIiwiaW5zdXJlZElkVmFpbGRTdGFydERhdGUiOiIiLCJpbnN1cmVkTWFyaXRhbFN0YXR1cyI6IiIsImluc3VyZWRQcm92aW5jZU5hbWUiOiIifV0sImlzSG9sZGVyIjoidHJ1ZSJ9fSwicGF5bWVudEluZm8iOnsicGF5UmVuZXdhbCI6eyJwYXllck1vYmlsZSI6IiIsInBheWVyQWNjb3VudCI6IiIsInBheWVyQmFua05hbWUiOiIiLCJwYXllck5hbWUiOiIifSwicGF5Rmlyc3QiOnsicGF5ZXJNb2JpbGUiOiIiLCJwYXllckFjY291bnQiOiIiLCJtZXJjaGFudFBheVRyYWRlSWQiOiIiLCJwYXlPcmRlcklkIjoiMjAxMzU0NzEzNDk0MjgwNTE5NjgwIiwicGF5TW9uZXkiOiI2MDU3MjAiLCJwYXllckJhbmtOYW1lIjoiIiwicGF5ZXJOYW1lIjoiIiwicGF5V2F5IjoiIn19LCJvcmRlciI6eyJpbnN1cmVPcmdOYW1lIjoi5qiq55C05Lq65a+/IiwiY29vcE9yZGVySWQiOiIxMzU0NzEzMTk3MTE3Nzg0MDY0IiwiYW1vdW50IjoiMzgwMDAwIiwic2t1TGlzdCI6W3sic2t1TmFtZSI6IiIsImFtb3VudCI6MzgwMDAwLCJkdXR5TGlzdCI6W10sInNrdUlkIjoiMTI0MzA3OTUyMTU1MjU1MTkzNiJ9XSwicHJvZHVjdElkIjoiMTI0MzA3OTUyMTUzMTU4MDQxNiIsIm9yZGVySWQiOiIxMDEzNTQ3MTM0OTQyMTM0MTA4MTYiLCJlbmREYXRlIjoiMjg4OC0wMS0wMSAwMDowMDowMCIsInBheVRpbWUiOiIyMDIxLTAxLTI4IDE2OjUwOjM1IiwicHJvZHVjdENhdGVnb3J5TmFtZSI6IumHjeeWvumZqSIsImluc1BlcmlvZCI6Ik8iLCJ0b3RhbFByZW1pdW0iOiI2MDU3MjAiLCJhZ2VuY3lJZCI6IjEzMDI4MDgyNDMzNTM1NTkwNDAiLCJpbnN1cmFuY2VUeXBlIjoiMiIsInBvbGljeVN0YXR1cyI6IkNCIiwicHJvZHVjdE5hbWUiOiLmqKrnkLTml6Dlv6fkurrnlJ8yMDIw6YeN5aSn55a+55eF5L+d6ZmpIiwiYmVnaW5EYXRlIjoiMjAyMS0wMS0yOSAwMDowMDowMCIsInBheVR5cGUiOiJZIiwicG9saWN5SWQiOiJQOTcyMTAxMTE0Nzc5IiwiYWNjZXB0ZWRUaW1lIjoiMjAyMS0wMS0yOCAxNjo1MDozNSIsImluc3VyZU9yZ0lkIjoiOTA3NTM4NDg2ODIyMTkxMTY4IiwicGF5UGVyaW9kIjoiWTMwIiwiYml6VGFnIjoiIiwiaW5zU291cmNlIjoibGltYSJ9LCJhZ2VuY3lJbmZvIjp7ImNvbXBhbnlJZCI6IjEyOTEzMjAwNzMxNTk5NTQ0MzIiLCJncm91cElkIjoiMTI5MTMyMDA3MzE3MjUzNzM0NCIsImFnZW5jeUlkIjoiMTMwMjgwODI0MzM1MzU1OTA0MCJ9fSwibWVzc2FnZUhlYWQiOnsicmVxdWVzdFR5cGUiOiJDMDIiLCJhc3luIjoiZmFsc2UiLCJVVUlEIjoiMTM1NDcxMzQ5OTAzMDg1MTU4NCIsInZlcnNpb24iOiIxLjAiLCJjaGFubmVsSWQiOiIxMjkxMzIwMDczMTU5OTU0NDMyIiwic2VuZFRpbWUiOiIyMDIxLTAxLTI4IDE2OjUwOjQyIn19"


	decryptBody, _ := Base64DecodeString(req_info)

	data := []byte(req_info+"LMmd5Test")
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	fmt.Println(md5str1)

	fmt.Println(string(decryptBody))
}

// Base64EncodeString 编码
func Base64EncodeString(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Base64DecodeString 解码
func Base64DecodeString(str string) (string, []byte) {
	resBytes, _ := base64.StdEncoding.DecodeString(str)
	return string(resBytes), resBytes
}

//func main() {
//	date, err := time.ParseInLocation("01-02-06", "01-21-21", time.Local)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(date)
//}
