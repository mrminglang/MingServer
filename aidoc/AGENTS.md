# MingServer 项目概览

## 项目定位
- 基于 TAF/TafGo 的服务端程序（Go 语言）。
- 协议与接口由 `.jce` 文件定义，代码由 TafGo 生成并落在 `taf-protocol/`。

## 核心结构
- 协议定义：`MingHello.jce`、`ESDriver.jce`。
- 生成代码：`taf-protocol/` 下的 `*.jce.go` / `*.taf.go`。
- 接口接入层：`*_imp.go`（本项目为 `MingHello_imp.go`），只做接入与转发。
- 业务逻辑：`logic/` 包中与接口同名的函数实现具体功能（当前集中在 `logic/logic.go`）。
- 启动与初始化：`main.go` + `boot/boot.go`。
- 数据访问：`repositories/`（当前为 `teacher_repository`）。
- HTTP 接入（可选）：`logic/http_proxy/http.go`。

## 启动流程（简）
1. `main.go` 读取 TAF 配置并调用 `boot.Boot`。
2. `boot/boot.go` 完成：日志、配置、数据库、ES、DCache、业务初始化。
3. 注册 servant：`MingHelloImp` -> `MingApp.MingServer.MingHelloObj`。
4. 若配置开启 `isHttpOn=true`，启动 HTTP 代理监听。
5. `taf.Run()` 启动服务循环。

## 配置说明
- 服务配置默认使用 `MingServer.conf`（同名 `.conf` 文件）。
- TAF 服务本地监听与对象适配器配置在 `config/config.conf`。
- 配置读取与数据库连接初始化来自 `gitlab.upchinaproduct.com/upgo/utils/server_utils`：
  - 读取配置：`server_utils/confs`
  - 数据库连接：`server_utils/ormdb`，连接参数在配置文件 `<db>` 节点中
- ES、DCache 与日志/配置读取使用 `gitlab.upchinaproduct.com/upgo/utils/server_utils` 相关包。

## 协议与接口
- `MingHello.jce` 定义了结构体与接口：
  - `getTeacherList`
  - `setStringCache` / `getStringCache`
  - `setESData` / `getESDataById`
- `MingHello_imp.go` 将接口调用转发到 `logic` 包同名方法。

## 业务逻辑概览（logic/logic.go）
- 教师列表：从 `repositories/teacher_repository` 读取并组装响应。
- DCache：通过 `dcache_rpc` 进行字符串缓存读写。
- ES：通过 `es_repository` 写入与按 ID 查询。

## 构建与运行
- 编译脚本：`./rebuild.sh`。
- 每次修改后建议编译验证；需临时覆盖 GoLand 设置的 `GOROOT` 版本路径后再运行。

## 约定与备注
- `aidoc/` 用于存放测试脚本、文档、中间文件（不存在时再创建即可）。
- 若新增接口，请同步更新 `.jce` 并重新生成协议代码。
