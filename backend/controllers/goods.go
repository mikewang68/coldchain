package controllers

import (
	"bytes"
	"coldchain/backend/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"encoding/json"

	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/module/contract"

	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/types/request"

	"context"

	"github.com/ipfs/boxo/files" // 统一使用 boxo/files
	rpc "github.com/ipfs/kubo/client/rpc"

	"github.com/ipfs/go-cid"
)

// 调用合约input数据体
type ContractCall struct {
	Function string `json:"function"`
	Args     string `json:"args"`
}

// 查询合约信息input数据体
type ContractCa struct {
	Function string `json:"function"`
	Args     string `json:"args"`
	Return   string `json:"return"`
}

var MyPrivateKey string = "priSPKkpuYHiwQ886GdRrb9s6TbCmTqYdQdKEYo1X6njuSiMNP"
var SDK_URL string = "http://test.bifcore.bitfactory.cn"
var MyAccountAddress string = "did:bid:efPLdVAy6AN5wVgViFzfeNZ5yauq7hFs"
var ContractAddress = "did:bid:efS882eonf4Ng5eHJmDA2Z61qnJwrBR7"

// 创建中文时间格式模板
const chineseTimeLayout = "2006年1月2日15时04分"

func GenerateCall(s models.Goods, cid string) ContractCall {
	return ContractCall{
		Function: "addItem(uint256,string)",
		Args: fmt.Sprintf("%s,%s",
			s.ID,
			cid,
		),
	}
}

func ContractCalls(senderAddress string, contractAddress string, senderPrivateKey string, input string) {
	// 1. 初始化合约实例（连接区块链节点）
	bs := contract.GetContractInstance(SDK_URL) // SDK_INSTANCE_URL 需替换为实际节点地址

	// 2. 构建合约调用请求参数
	var r request.BIFContractInvokeRequest

	// 4. 填充请求参数
	r.SenderAddress = senderAddress     // 指定交易发送者
	r.PrivateKey = senderPrivateKey     // 私钥用于签名交易
	r.ContractAddress = contractAddress // 要调用的合约地址
	r.BIFAmount = 0                     // 转账金额（单位：链的最小单位，如 1 BIF = 1e8 最小单位）
	r.Input = input
	r.Remarks = "contract invoke" // 交易备注（可读描述）
	r.FeeLimit = 100000000
	res := bs.ContractInvoke(r)
	fmt.Println(res.ErrorCode)

}

func uploadProduct(client *rpc.HttpApi, product models.Goods) (cid.Cid, error) {

	// jsonData, err := json.Marshal(product)
	// if err != nil {
	// 	return cid.Undef, err
	// }

	// 禁用 HTML 转义
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(product); err != nil {
		return cid.Undef, err
	}
	jsonData := buf.Bytes()

	// 打印原始 JSON 和十六进制格式
	fmt.Printf("=== 原始 JSON 数据 ===\n%s\n", string(jsonData))
	fmt.Printf("=== 十六进制格式 ===\n%x\n", jsonData)

	// 直接上传文件（不再包装为目录）
	fileNode := files.NewBytesFile(jsonData)
	ctx := context.Background()
	ipfsPath, err := client.Unixfs().Add(ctx, fileNode)
	if err != nil {
		return cid.Undef, err
	}

	return ipfsPath.RootCid(), nil
}

// // 通过 CID 查询数据
// func queryProduct(client *rpc.HttpApi, cidStr string) (*models.Goods, error) {
// 	c, err := cid.Parse(cidStr)
// 	if err != nil {
// 		return nil, fmt.Errorf("无效的 CID: %v", err)
// 	}

// 	ctx := context.Background()
// 	reader, err := client.Block().Get(ctx, path.FromCid(c))
// 	if err != nil {
// 		return nil, fmt.Errorf("获取数据失败: %v", err)
// 	}

// 	data, err := io.ReadAll(reader)
// 	data = cleanInvalidChars(data)

// 	if err != nil {
// 		return nil, fmt.Errorf("读取数据失败: %v", err)
// 	}

// 	// 调试输出
// 	fmt.Printf("=== 查询到的原始字节 (HEX) ===\n%x\n", data)
// 	fmt.Printf("=== 查询到的字符串 ===\n%s\n", string(data))

// 	var product models.Goods
// 	if err := json.Unmarshal(data, &product); err != nil {
// 		return nil, fmt.Errorf("解析 JSON 失败: %v", err)
// 	}

// 	return &product, nil
// }

// func cleanInvalidChars(data []byte) []byte {
// 	re := regexp.MustCompile(`[^\x20-\x7E]`) // 匹配非可打印 ASCII 字符
// 	return re.ReplaceAllLiteral(data, []byte{})
// }

// GetGoodsList 获取货物列表
func GetGoodsList(c *gin.Context) {
	// var filter models.GoodsFilter
	// if err := c.ShouldBindQuery(&filter); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// // 过滤数据
	// filteredList := make([]models.Goods, 0)
	// for _, goods := range goodsList {
	// 	if filter.GoodsID != "" && goods.ID != filter.GoodsID {
	// 		continue
	// 	}
	// 	if filter.Type != "" && goods.Type != filter.Type {
	// 		continue
	// 	}
	// 	if filter.Status != "" && goods.Status != filter.Status {
	// 		continue
	// 	}
	// 	filteredList = append(filteredList, goods)
	// }

	// c.JSON(http.StatusOK, models.GoodsListResponse{
	// 	Total: int64(len(filteredList)),
	// 	Items: filteredList,
	// })
}

// CreateGoods 创建货物
func CreateGoods(c *gin.Context) {

	var goods models.Goods
	if err := c.ShouldBindJSON(&goods); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 生成批次编号
	goods.ID = "SF" + time.Now().Format("20060102") + uuid.New().String()[:3]
	goods.InTime = time.Now().Format(chineseTimeLayout)
	goods.Status = "instock"

	// // 连接本地 IPFS 节点（默认端口 5001）
	// client, err := rpc.NewLocalApi()
	// if err != nil {
	// 	panic(err)
	// }

	// cid, err := uploadProduct(client, goods)
	// if err != nil {
	// 	panic(fmt.Sprintf("上传失败: %v", err))
	// }
	// fmt.Printf("上传成功，CID: %s\n", cid.String())

	fmt.Println(goods)

	c.JSON(http.StatusOK, goods)
}

// UpdateGoods 更新货物信息
func UpdateGoods(c *gin.Context) {
	// id := c.Param("id")
	// var updatedGoods models.Goods
	// if err := c.ShouldBindJSON(&updatedGoods); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// for i, goods := range goodsList {
	// 	if goods.ID == id {
	// 		updatedGoods.ID = id               // 保持ID不变
	// 		updatedGoods.InTime = goods.InTime // 保持入库时间不变
	// 		goodsList[i] = updatedGoods
	// 		c.JSON(http.StatusOK, updatedGoods)
	// 		return
	// 	}
	// }

	// c.JSON(http.StatusNotFound, gin.H{"error": "goods not found"})
}

// DeleteGoods 删除货物
func DeleteGoods(c *gin.Context) {
	// id := c.Param("id")

	// for i, goods := range goodsList {
	// 	if goods.ID == id {
	// 		goodsList = append(goodsList[:i], goodsList[i+1:]...)
	// 		c.JSON(http.StatusOK, gin.H{"message": "goods deleted"})
	// 		return
	// 	}
	// }

	// c.JSON(http.StatusNotFound, gin.H{"error": "goods not found"})
}

// OutGoods 货物出库
func OutGoods(c *gin.Context) {
	// id := c.Param("id")

	// for i, goods := range goodsList {
	// 	if goods.ID == id {
	// 		if goods.Status != models.StatusInStock {
	// 			c.JSON(http.StatusBadRequest, gin.H{"error": "goods is not in stock"})
	// 			return
	// 		}
	// 		goodsList[i].Status = models.StatusOutStock
	// 		c.JSON(http.StatusOK, goodsList[i])
	// 		return
	// 	}
	// }

	// c.JSON(http.StatusNotFound, gin.H{"error": "goods not found"})
}
