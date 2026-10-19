# Day 1 · 骨架：让 Go 服务跑起来并连上 MySQL

- 日期：2026-10-19
- 阶段：建立项目骨架（对应 `ROADMAP.md` 第 1 天）
- 状态：已完成（验收通过）

---

## 一、今天的验收标准

跑通下面这条链路，今天就算过：

1. `go run ./cmd/server` 能启动，启动日志里出现 `mysql connected`
2. 另开一个终端，`curl localhost:8080/health` 返回：

```json
{"status":"ok"}
```

**范围红线**：今天不碰商品表、不写业务逻辑、不装多余的库。只做「服务能起 + 库能连」。

> 实际验收记录（2026-10-19）：`HTTP 200  body={"status":"ok"}`，启动日志含 `mysql connected`。

---

## 二、环境（已就绪，不用再折腾）

| 组件 | 地址 | 备注 |
|---|---|---|
| MySQL | `127.0.0.1:3307` | 容器 `mall-mysql`，库 `mall`，账号密码 `root/root123` |
| Redis | `127.0.0.1:6379` | 容器 `mall-redis`，今天用不到 |

> 端口是 **3307 不是 3306**：3306 被你本机已装的 MySQL80 服务占着，之前就是它把容器卡住了。

---

## 三、要写的代码（你写，我 review）

### 1. `internal/config/config.go`

- 定义 `Config` struct，字段：`DBHost`、`DBPort`、`DBUser`、`DBPassword`、`DBName`、`ServerPort`
- 从环境变量读；读不到就给默认值：`127.0.0.1` / `3307` / `root` / `root123` / `mall` / `8080`
- 一个 `Load()` 返回 `*Config`

### 2. `cmd/server/main.go`

启动顺序固定成这样：

```
读配置 → 连 MySQL → Ping → 起 Gin → 注册 GET /health → Run
```

- **DSN 由你自己拼**（这是今天的考点之一）
- 连库失败或 Ping 失败 → 打印错误 + `os.Exit(1)`，**绝不带着坏连接启动**

---

## 四、今天要能讲清楚（原理）

做完之后我们逐个过一遍，你必须能用自己的话说出来：

1. 一个 HTTP 请求从进来到返回，经过了哪些环节
2. 路由 / Handler / 数据库连接，各自负责什么
3. DSN 是什么？`parseTime=True` 不加会出什么问题
4. `gorm.Open` 返回的 error 是 nil，就等于真的连上 MySQL 了吗
5. 8080 端口被占用时会怎样，你怎么发现

---

## 五、三个坑（先自己想 5–10 分钟，再问我）

- [ ] DSN 长什么样？`parseTime=True` 的作用是什么？
- [ ] `gorm.Open` 返回 nil error ⇒ 真的连上了吗？
- [ ] 端口被占用时会发生什么？怎么定位是谁占的？

---

## 六、复盘（今天结束前自己填）

- 卡住的点：
- 今天真正搞懂的一件事：
- 明天要补的：

---

## 七、提交

- [x] `git commit -m "step1: 服务骨架 + MySQL 连接 + /health 健康检查"`

---

## 八、验收问答（3 题，口述回答后记录在这里）

1. 一个请求打到 `localhost:8080/health`，从进来到返回，依次经过了哪些环节？
2. DSN 里 `parseTime=True` 去掉会怎样？`tcp(%s:%d)` 里的 `tcp` 代表什么？
3. 8080 端口被占用时 `r.Run` 会发生什么？用什么命令能查出是谁占的？

> 答案待你口述补充。
