# goctl

## PART1. 生成文件格式

### 1.1 api指令

[goctl api指令说明](https://go-zero.dev/cn/docs/goctl/api)

#### 1.1.1 根据给定的api文件,将生成的代码存放到给定的路径

```
(base) yanglei@192 order % pwd
/Users/yanglei/Desktop/go-zero-study/goZeroStudy/10-goctl/mall/order
(base) yanglei@192 order % goctl api go --api order.api -dir ./gen
Done.
```

- `--api`:指定api文件
- `-dir`:指定生成代码的存放路径

#### 1.1.2 生成代码时指定文件命名风格

```
(base) yanglei@192 order % goctl api go --api order.api -dir ./gen --style goZero
Done.
```

```
(base) yanglei@192 order % goctl api go --api order.api -dir ./gen --style go_zero
Done.
```

- `--style`参数:指定生成代码时的文件命名风格

[生成代码时的文件命名风格文档](https://github.com/zeromicro/go-zero/blob/master/tools/goctl/config/readme.md)

### 1.2 rpc指令

[goctl rpc指令说明](https://go-zero.dev/cn/docs/goctl/zrpc)

#### 1.2.1 生成proto模板文件

```
(base) yanglei@192 order % pwd
/Users/yanglei/Desktop/go-zero-study/goZeroStudy/10-goctl/mall/order
(base) yanglei@192 order % goctl rpc -o=order.proto
```

该命令会在当前位置生成一个proto文件.

#### 1.2.2 根据proto文件生成rpc服务代码

```
goctl rpc protoc order.proto --go_out=./rpcGen --go-grpc_out=./rpcGen --zrpc_out=./zrpcGen
```

这个命令之前写rpc服务时用过.就不再多讲了.注意:生成的`pb.go`和`_grpc.pb.go`文件必须在同一个目录下.也就是说`--go_out`选项和`--go-grpc_out`选项的值必须相同.

### 1.3 model指令

#### 1.3.1 根据SQL语句生成模型代码

不推荐直接使用他生成的CRUD方法,但可以用它生成的结构体.

```
goctl model mysql ddl -src="./*.sql" -dir="./sql/model" -c
```

类型转换规则:

| mysql dataType | golang dataType | golang dataType(if null&&default null) |
| -------------- | --------------- | -------------------------------------- |
| bool           | int64           | sql.NullInt64                          |
| boolean        | int64           | sql.NullInt64                          |
| tinyint        | int64           | sql.NullInt64                          |
| smallint       | int64           | sql.NullInt64                          |
| mediumint      | int64           | sql.NullInt64                          |
| int            | int64           | sql.NullInt64                          |
| integer        | int64           | sql.NullInt64                          |
| bigint         | int64           | sql.NullInt64                          |
| float          | float64         | sql.NullFloat64                        |
| double         | float64         | sql.NullFloat64                        |
| decimal        | float64         | sql.NullFloat64                        |
| date           | time.Time       | sql.NullTime                           |
| datetime       | time.Time       | sql.NullTime                           |
| timestamp      | time.Time       | sql.NullTime                           |
| time           | string          | sql.NullString                         |
| year           | time.Time       | sql.NullInt64                          |
| char           | string          | sql.NullString                         |
| varchar        | string          | sql.NullString                         |
| binary         | string          | sql.NullString                         |
| varbinary      | string          | sql.NullString                         |
| tinytext       | string          | sql.NullString                         |
| text           | string          | sql.NullString                         |
| mediumtext     | string          | sql.NullString                         |
| longtext       | string          | sql.NullString                         |
| enum           | string          | sql.NullString                         |
| set            | string          | sql.NullString                         |
| json           | string          | sql.NullString                         |

### 1.4 生成dockerfile

```
goctl docker -go order.go
```

- `-go`参数:指定入口文件

### 1.5 生成k8s资源清单

```
goctl kube deploy -name redis -namespace adhoc -image redis:6-alpine -o redis.yaml -port 6379
```

- `-name`参数:指定容器的名称
- `-namespace`参数:指定pod在K8S中的命名空间
- `-image`参数:指定使用的镜像名称
- `-o`参数:指定生成的资源清单yaml文件的名称
- `-port`参数:指定容器暴露的端口

实际使用时,也基本上就是生成资源清单之后再改

### 1.6 plugin指令

[plugin指令](https://go-zero.dev/cn/docs/goctl/plugin)

### 1.7 template指令

[template指令](https://go-zero.dev/cn/docs/goctl/template-cmd)

### 1.8 goctl指令大全

[goctl指令大全](https://go-zero.dev/cn/docs/goctl/commands)

### 1.9 api语法

[api语法](https://go-zero.dev/cn/docs/design/grammar)

在不使用goctl工具的前提下,也能把go-zero这个框架用起来.在这个基础上,再用goctl来加快开发速度.因为你首先要掌握这个框架怎么用,再考虑怎么能加速你的生产效率.你写几遍就知道它的代码细节大概怎么回事了.