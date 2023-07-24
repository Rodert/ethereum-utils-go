package clienttest_test

import (
	"ethereum-utils-go/client"
	"testing"
)

// 测试连接节点
func TestGetClient(t *testing.T) {
	infuraNode1 := "https://mainnet.infura.io/v3/d6624913c27d4f5a9540f071767b4d49"
	client.GetClient(infuraNode1)
}
