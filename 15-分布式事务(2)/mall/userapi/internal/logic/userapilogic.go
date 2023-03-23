package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dtm-labs/dtm/client/dtmgrpc"
	"github.com/golang-jwt/jwt/v4"
	"rpc-common/user-score/types/score"
	"rpc-common/user/types/user"
	"strconv"
	"time"
	"userapi/internal/svc"
	"userapi/internal/types"

	_ "github.com/dtm-labs/driver-gozero"
	"github.com/zeromicro/go-zero/core/logx"
)

// dtmServer dtm的conf.yaml中配置的 dtm的注册地址
var dtmServer = "etcd://localhost:2379/dtmservice"

type UserapiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserapiLogic {
	return &UserapiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserapiLogic) Userapi(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}

func (l *UserapiLogic) Register(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	// 一般做超时上下文
	//ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancelFunc()

	// SaveUser和SaveUserScore 纳入到1个事务中
	// dtm client把grpc调用提交给dtm server,由dtm server来完成调用

	// 生成组id
	gid := dtmgrpc.MustGenGid(dtmServer)
	// 使用SAGA模式
	sagaGrpc := dtmgrpc.NewSagaGrpc(dtmServer, gid)

	// 定义grpc调用 让dtm client交给dtm server来完成调用
	userRpcServer, err := l.svcCtx.Config.UserRpc.BuildTarget()
	if err != nil {
		return nil, err
	}

	// 此处后边的形似路由的部分是user_grpc.pb.go中的常量 User_SaveUser_FullMethodName 的值
	// saveUserAction表示要让dtm server调用的 user rpc服务的路由
	saveUserAction := userRpcServer + "/user.User/saveUser"

	// saveUserCallbackAction 表示当事务提交失败 需要回滚时调用的 user rpc服务的路由
	saveUserCallbackAction := userRpcServer + "/user.User/saveUserCallback"

	// 此处的saveReq 表示调用user rpc服务的"/user.User/saveUser"时的参数
	idStr := strconv.FormatInt(req.Id, 10)
	saveReq := &user.UserRequest{
		Id:     idStr,
		Name:   req.Name,
		Gender: req.Gender,
	}

	// 让dtm server 对user rpc服务做调用和回调
	sagaGrpc.Add(saveUserAction, saveUserCallbackAction, saveReq)

	// 让dtm server 对user-score rpc服务做调用和回调
	scoreRpcServer, err := l.svcCtx.Config.UserScoreRpc.BuildTarget()
	if err != nil {
		return nil, err
	}
	saveScoreAction := scoreRpcServer + "/score.UserScore/saveUserScore"
	saveScoreActionCallback := scoreRpcServer + "/score.UserScore/saveUserScoreCallback"
	saveScoreReq := &score.UserScoreRequest{
		// Tips: 此处由于对user rpc服务的调用是由dtm server来完成的 所以在这里是拿不到user rpc服务的返回值的
		// Tips: 也就是说说是拿不到user_id的 所以先随便写一个
		// TODO: 这TM生产咋整 = =
		UserId: req.Id,
		Score:  10,
	}
	sagaGrpc.Add(saveScoreAction, saveScoreActionCallback, saveScoreReq)
	// 等待提交结果 可以认为设置的是同步提交
	sagaGrpc.WaitResult = true
	err = sagaGrpc.Submit()

	if err != nil {
		fmt.Printf("submit transaction error: %#v\n", err.Error())
		return nil, err
	}

	return resp, nil
}

func (l *UserapiLogic) GetUser(t *types.GetUserRequest) (resp *types.GetUserResponse, err error) {
	// 认证通过后 从token中获取userId
	userId := l.ctx.Value("userId")
	logx.Info("获取到的token内容为: %s \n", userId)
	userIdStr := string(userId.(json.Number))
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	// 此处由于发放token时向上下文中存入的userId就是1 故从上下文中拿到的userId值肯定是1
	// 所以必然触发错误
	//if userIdStr == "1" {
	//	return nil, bizError.ParamError
	//}
	userResponse, err := l.svcCtx.UserRpc.GetUser(ctx, &user.IdRequest{Id: userIdStr})
	if err != nil {
		return nil, err
	}
	resp = &types.GetUserResponse{
		Message: "success",
		Data:    userResponse,
	}
	return resp, nil
}

func (l *UserapiLogic) Login(t *types.LoginRequest) (resp *types.LoginResponse, err error) {
	fmt.Printf("正在执行Login方法\n")
	userId := 1
	secret := l.svcCtx.Config.Auth.Secret
	expireIn := l.svcCtx.Config.Auth.ExpireIn
	signAt := time.Now().Unix()
	token, err := genToken(secret, signAt, expireIn, int64(userId))
	if err != nil {
		return nil, err
	}
	resp = &types.LoginResponse{
		Message: "success",
		Data: map[string]string{
			"token": token,
		},
	}
	return resp, nil
}

// genToken 生成token
// secret:密钥
// signAt: 签发时间
// expireIn: token生命周期
// userId: 用户ID
func genToken(secret string, signAt, expireIn, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	// Tips: 以下2个key的值是jwt.MapClaims预制的 必须这么写key值
	claims["exp"] = signAt + expireIn
	claims["iat"] = signAt
	// Tips: 这个key的值可以自定义的 后续验证token通过后 该K-V将会被存入上下文中
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secret))
}
