# 敏感词过滤

> 使用 golang 高性能语言实现

#### 技术相关
- go >= 1.14
- protobuf

#### 特点
- @todo

---

### 目录结构

  ```
  ├─probotuf                 # proto文件生成管理
  ├─protos                   # proto文件管理
  ├─server                   # 服务代码
  ├─main.go                  # 入口文件
  ├─.gitignore               # git提交忽略配置
  ```
----

### golang客户端使用示例

```go
  import "google.golang.org/grpc"

  func example() {
    conn, err := grpc.Dial("127.0.0.1:8897", grpc.WithInsecure())
    if err != nil {
      log.Fatalf("did not connect: %v", err)
    }
    client = wordFilter.NewWordFilterClient(conn)
  
    defer conn.Close()
    greq := &wordFilter.ContentReq{Content: username}
    validateReq, _ := client.Validate(context.TODO(), greq)
    if !validateReq.Res {
      exception.Logic("昵称含有违规信息" + validateReq.Word)
    }
  }
```


### 参考文献

github.com/importcjj/sensitive
