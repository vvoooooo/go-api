package apple

/*
	人生无常  大肠包小肠
*/
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	// AppleSandbox 沙盒模式下校验地址 https://help.apple.com/app-store-connect/#/dev7e89e149d
	AppleSandbox string = "https://sandbox.itunes.apple.com/verifyReceipt"

	// AppleProd 正式上线校验地址
	AppleProd string = "https://buy.itunes.apple.com/verifyReceipt"
)

// requestVerifyApi
/*  app store 校验支付请求
	url：取 AppleSandbox 或 AppleProd
	reqData : app store pay 返回的收据信息
 	文档：https://developer.apple.com/documentation/appstorereceipts/verifyreceipt
*/
func requestVerify(url string, reqData VerifyRequest) (jsonResp *VerifyResponse, err error) {
	buf := &bytes.Buffer{}
	err = json.NewEncoder(buf).Encode(reqData)
	if err != nil {
		return
	}
	client := http.Client{Timeout: time.Second * 10}
	resp, err := client.Post(url, "application/json", buf)
	if err != nil || resp == nil { //报错了 或者没有返回信息 exit
		return
	}
	defer func(Body io.ReadCloser) {
		e := Body.Close()
		if e != nil {
			fmt.Println(e.Error())
			return
		}
	}(resp.Body)
	jsonResp = &VerifyResponse{}
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusRequestTimeout { // 比如超时啊什么的 需要让其通过校验
			jsonResp.Status = 0
		}
		return
	}
	jsonRespByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = json.Unmarshal(jsonRespByte, jsonResp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

// iosVerify 订单正确性验证  Receipt  客户端拿到的收据 发给我校验
func iosVerify(receipt string) (response *Receipt, err error) {
	reqData := VerifyRequest{
		Receipt:  receipt,
		Password: "",
	}
	request := func(url string) (jsonResp *VerifyResponse, isTest bool, err error) {
		jsonResp, err = requestVerify(url, reqData)
		if err != nil {
			return
		}
		if jsonResp.Status == SandBox {
			return nil, true, nil
		}
		return
	}
	//先往正式URL发
	JsonResp, isTest, err := request(AppleProd)
	if err != nil {
		return nil, err
	}
	if isTest {
		// 如果是沙盒的  来沙盒发请求
		JsonResp, isTest, err = request(AppleSandbox)
		if err != nil {
			return nil, err
		}
		if isTest {
			return nil, fmt.Errorf("[appStore PayVerify] testurl return should in testmod")
		}
	}
	if JsonResp.Status != 0 {
		return nil, fmt.Errorf("[appStore PayVerify] JsonResp.Status[%d]!=0", JsonResp.Status)
	}
	return JsonResp.Receipt, nil
}

// Check 外部引用   没有demo测试啊
func Check(receipt string) bool {
	_, err := iosVerify(receipt)
	if err != nil {
		return false
	}
	return true
}
