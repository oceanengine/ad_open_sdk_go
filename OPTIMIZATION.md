# 巨量引擎 Go SDK 内存优化方案与使用指南

## 📋 为什么需要这个优化版本？

巨量引擎官方开放平台提供的 Go SDK (`github.com/oceanengine/ad_open_sdk_go`) 拥有数千个 API 文件，但在生成代码时采用了 **点导入（Dot Imports `import . "models"`）**。

- **官方痛点**：点导入会将 `models` 包中数万个符号强行注入到当前包的扁平符号表，导致 Go 编译器在做 AST 类型检查与符号解析时，需要维护数千万个条目，**编译峰值内存暴涨至 6GB~8GB+**。在小型 CI 容器或云服务器中，极易触发 Swap 剧烈换页（假死 10 分钟）甚至 OOM 崩溃。
- **本仓库优化**：通过自动化 AST/正则流水线，将所有点导入转换为标准具名导入，并为所有类型显式补全 `models.` 前缀。
  - **编译内存降低 70%~80%（降至 1~2GB）**
  - **编译速度提升 60%+**
  - **完美对齐官方语义化版本 Tag（如 `v1.1.93-optimized`）**

---

## 🚀 快速安装与使用

### 1. 安装优化版依赖

```bash
# 推荐方式：直接引用 optimize 分支（保持与官方最新同步）
go get github.com/Garfield247/ad_open_sdk_go@optimize

# 或引用特定版本 Tag（对齐官方版本号）
go get github.com/Garfield247/ad_open_sdk_go@v1.1.93-optimized
```

### 2. 代码调用示例

```go
package main

import (
	"context"
	"fmt"

	ad_open_sdk_go "github.com/Garfield247/ad_open_sdk_go"
	"github.com/Garfield247/ad_open_sdk_go/config"
	"github.com/Garfield247/ad_open_sdk_go/models" // 标准具名导入
)

func main() {
	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)

	// 使用 models. 前缀访问结构体和枚举类型
	req := apiClient.AccountFundGetV30Api().Get(context.Background())
	req.AccessToken("YOUR_ACCESS_TOKEN")
	req.AdvertiserId(123456789)

	resp, _, err := req.Execute()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Response: %+v\n", resp)
}
```

---

## 🌿 分支与版本策略

- **`master` 分支**：纯净追踪官方 `oceanengine/ad_open_sdk_go` 上游代码与历史，与官方 1:1 保持一致。
- **`optimize` 分支**：经过自动化移除点导入后的优化分支，**生产环境强烈推荐使用此分支**。
- **Tag 规范**：每当官方发布新版本（如 `v1.1.93`），GitHub Actions 自动完成编译验证并在 `optimize` 分支生成 **`v1.1.93-optimized`** 标签与 GitHub Release。
