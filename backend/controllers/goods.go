package controllers

import (
	"bytes"
	"coldchain/backend/models"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"encoding/json"

	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/module/contract"

	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/types/request"

	shell "github.com/ipfs/go-ipfs-api"

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

type ContractC struct {
	Function string `json:"function"`
	Args     string `json:"args"`
	Return   string `json:"return"`
}

type ContractResponse struct {
	ErrorCode int    `json:"error_code"`
	ErrorDesc string `json:"error_desc"`
	Result    struct {
		QueryRets []struct {
			Result struct {
				Data string `json:"data"` // 目标字段
			} `json:"result"`
		} `json:"query_rets"`
	} `json:"result"`
}

var MyPrivateKey string = "priSPKkpuYHiwQ886GdRrb9s6TbCmTqYdQdKEYo1X6njuSiMNP"
var SDK_URL string = "http://test.bifcore.bitfactory.cn"
var MyAccountAddress string = "did:bid:efPLdVAy6AN5wVgViFzfeNZ5yauq7hFs"
var ContractAddress = "did:bid:efZu2WS66T4CTp7ZGb5fE925jztapZ4y"

// 全局 IPFS Shell 实例
var ipfsShell *shell.Shell

// 创建中文时间格式模板
const chineseTimeLayout = "2006年1月2日15时04分"

func init() {
	// 初始化全局实例（程序启动时执行一次）
	ipfsShell = shell.NewShell("localhost:5001")
}

func GenerateCall(s models.Goods, cid string) ContractCall {
	return ContractCall{
		Function: "addProduct(string,string)",
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

func QuerycontrctGoodslist(senderAddress string, contractAddress string) []string {
	// 1. 初始化合约实例（连接区块链节点）
	bs := contract.GetContractInstance(SDK_URL) // SDK_INSTANCE_URL 需替换为实际节点地址

	call := ContractC{
		Function: "getAllCids()",
		Return:   "returns(string)",
	}
	// 序列化为JSON字符串
	jsonBytes, err := json.Marshal(call)
	if err != nil {
		panic(err)
	}

	// 2. 构建合约调用请求参数
	var r request.BIFContractCallRequest

	// 4. 填充请求参数
	r.SourceAddress = senderAddress     // 指定交易发送者
	r.ContractAddress = contractAddress // 要调用的合约地址
	r.FeeLimit = 100000000
	r.Input = string(jsonBytes)
	r.Type = 1

	// 5. 调用合约方法
	res := bs.ContractQuery(r) // 发送交易到区块链网络
	if res.ErrorCode != 0 {
		// 错误处理：输出错误详情并终止程序
		log.Fatalf("合约调用失败 (错误码=%d): %s", res.ErrorCode, res.ErrorDesc)
	}

	// // 6. 序列化结果并打印
	dataByte, err := json.Marshal(res)
	if err != nil {
		log.Fatalf("JSON序列化失败: %v", err)
	}

	var response ContractResponse

	// 反序列化到结构体
	if err := json.Unmarshal(dataByte, &response); err != nil {
		log.Fatalf("JSON解析失败: %v", err)
	}

	// 提取data字段
	if len(response.Result.QueryRets) == 0 {
		log.Fatal("无查询结果")
	}
	rawData := response.Result.QueryRets[0].Result.Data
	fmt.Println("原始data字符串:", rawData)
	// 1. 去除方括号
	trimmed := strings.Trim(rawData, "[]")

	// 2. 分割字符串并清理空格
	rawIDs := strings.Split(trimmed, ",")

	// 3. 创建结果切片
	cidList := make([]string, 0, len(rawIDs))

	// 4. 清理每个 CID
	for _, id := range rawIDs {
		cleanID := strings.TrimSpace(id) // 处理可能存在的空格
		if cleanID != "" {
			cidList = append(cidList, cleanID)
		}
	}

	return cidList

}

func uploadProduct(product models.Goods) (cid.Cid, error) {
	// 直接序列化 JSON
	jsonData, err := json.Marshal(product)
	if err != nil {
		return cid.Undef, fmt.Errorf("JSON marshal failed: %w", err)
	}

	// 调试输出（可选）
	fmt.Printf("=== RAW JSON (%d bytes) ===\n%s\n", len(jsonData), jsonData)

	// 上传到 IPFS
	reader := bytes.NewReader(jsonData)
	cidStr, err := ipfsShell.Add(reader)
	if err != nil {
		return cid.Undef, fmt.Errorf("IPFS add failed: %w", err)
	}

	// 解析 CID 字符串
	parsedCID, err := cid.Decode(cidStr)
	if err != nil {
		return cid.Undef, fmt.Errorf("CID decode failed: %w", err)
	}

	return parsedCID, nil
}

func queryProduct(cidStr string) (models.Goods, error) {
	var zeroGoods models.Goods

	reader, err := ipfsShell.Cat(cidStr)
	if err != nil {
		return zeroGoods, fmt.Errorf("IPFS read failed: %w", err)
	}
	defer reader.Close()

	// 使用 io.ReadAll 替代 ioutil.ReadAll
	readData, err := io.ReadAll(reader)
	if err != nil {
		return zeroGoods, fmt.Errorf("data read failed: %w", err)
	}

	fmt.Printf("=== RAW DATA (%d bytes) ===\n%s\n", len(readData), readData)

	var product models.Goods
	if err := json.Unmarshal(readData, &product); err != nil {
		return zeroGoods, fmt.Errorf("JSON unmarshal failed: %w", err)
	}

	return product, nil
}

func QueryProducts(cidStrs []string) ([]models.Goods, error) {
	// 初始化返回结构和错误收集
	results := make([]models.Goods, 0, len(cidStrs))
	var errs []error

	// 顺序处理每个 CID
	for _, cid := range cidStrs {
		product, err := queryProduct(cid)
		if err != nil {
			// 错误包装并收集，包含原始 CID 信息
			errs = append(errs, fmt.Errorf("cid [%s] error: %w", cid, err))
			continue
		}
		results = append(results, product)
	}

	// 处理最终错误返回
	if len(errs) > 0 {
		return results, fmt.Errorf("partial success (%d/%d), errors: %w",
			len(results), len(cidStrs), errors.Join(errs...))
	}
	return results, nil
}

func GetGoodsList(c *gin.Context) {
	var filter models.GoodsFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 打印接收到的过滤条件
	fmt.Printf("Received filter: %+v\n", filter)

	cidlist := QuerycontrctGoodslist(MyAccountAddress, ContractAddress)

	virtualGoodsList, err := QueryProducts(cidlist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query products"})
		return
	}

	// 打印获取到的商品列表
	fmt.Printf("Total goods before filtering: %d\n", len(virtualGoodsList))

	// 过滤数据
	filteredList := make([]models.Goods, 0)
	for _, goods := range virtualGoodsList {

		// 直接比较字符串，不需要类型转换
		if filter.GoodsID != "" && goods.ID != filter.GoodsID {
			continue
		}
		if filter.Type != "" && goods.Type != filter.Type {
			continue
		}
		if filter.Status != "" && goods.Status != filter.Status {
			continue
		}
		filteredList = append(filteredList, goods)
	}

	// 打印过滤后的结果
	fmt.Printf("Total goods after filtering: %d\n", len(filteredList))

	// 返回数据
	c.JSON(http.StatusOK, models.GoodsListResponse{
		Total: int64(len(filteredList)),
		Items: filteredList,
	})
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

	fmt.Println(goods)

	cid, err := uploadProduct(goods)
	if err != nil {
		panic(fmt.Sprintf("上传失败: %v", err))
	}

	fmt.Printf("上传成功，CID: %s\n", cid.String())
	cidStr := cid.String()

	// 上链操作
	contractCall := GenerateCall(goods, cidStr)
	jsonBytes, _ := json.Marshal(contractCall)
	input := string(jsonBytes)
	fmt.Println(input)
	ContractCalls(MyAccountAddress, ContractAddress, MyPrivateKey, input)

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
	id := c.Param("id")

	delect := ContractCall{
		Function: "removeProduct(string)",
		Args:     id,
	}
	jsonBytes, _ := json.Marshal(delect)
	input := string(jsonBytes)
	fmt.Println(input)

	ContractCalls(MyAccountAddress, ContractAddress, MyPrivateKey, input)

	c.JSON(http.StatusOK, gin.H{"message": "goods deleted"})
	//没加错误显示之后加

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
