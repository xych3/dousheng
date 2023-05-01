// Code generated by Kitex v0.4.4. DO NOT EDIT.

package videoservice

import (
	"context"
	"fmt"
	video_rpc "gitee.com/derrickball/douyin/app/video/kitex_gen/video_rpc"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return videoServiceServiceInfo
}

var videoServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "VideoService"
	handlerType := (*video_rpc.VideoService)(nil)
	methods := map[string]kitex.MethodInfo{
		"QueryVideo":            kitex.NewMethodInfo(queryVideoHandler, newQueryVideoArgs, newQueryVideoResult, false),
		"CreateVideo":           kitex.NewMethodInfo(createVideoHandler, newCreateVideoArgs, newCreateVideoResult, false),
		"QueryVideoByUser":      kitex.NewMethodInfo(queryVideoByUserHandler, newQueryVideoByUserArgs, newQueryVideoByUserResult, false),
		"QueryVideoByIdBatch":   kitex.NewMethodInfo(queryVideoByIdBatchHandler, newQueryVideoByIdBatchArgs, newQueryVideoByIdBatchResult, false),
		"IncreaseFavoriteCount": kitex.NewMethodInfo(increaseFavoriteCountHandler, newIncreaseFavoriteCountArgs, newIncreaseFavoriteCountResult, false),
		"DecreaseFavoriteCount": kitex.NewMethodInfo(decreaseFavoriteCountHandler, newDecreaseFavoriteCountArgs, newDecreaseFavoriteCountResult, false),
		"IncreaseCommentCount":  kitex.NewMethodInfo(increaseCommentCountHandler, newIncreaseCommentCountArgs, newIncreaseCommentCountResult, false),
		"DecreaseCommentCount":  kitex.NewMethodInfo(decreaseCommentCountHandler, newDecreaseCommentCountArgs, newDecreaseCommentCountResult, false),
		"GetAuthor":             kitex.NewMethodInfo(getAuthorHandler, newGetAuthorArgs, newGetAuthorResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "video",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func queryVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video_rpc.QueryVideoReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video_rpc.VideoService).QueryVideo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *QueryVideoArgs:
		success, err := handler.(video_rpc.VideoService).QueryVideo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*QueryVideoResult)
		realResult.Success = success
	}
	return nil
}
func newQueryVideoArgs() interface{} {
	return &QueryVideoArgs{}
}

func newQueryVideoResult() interface{} {
	return &QueryVideoResult{}
}

type QueryVideoArgs struct {
	Req *video_rpc.QueryVideoReq
}

func (p *QueryVideoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video_rpc.QueryVideoReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *QueryVideoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *QueryVideoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *QueryVideoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in QueryVideoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *QueryVideoArgs) Unmarshal(in []byte) error {
	msg := new(video_rpc.QueryVideoReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var QueryVideoArgs_Req_DEFAULT *video_rpc.QueryVideoReq

func (p *QueryVideoArgs) GetReq() *video_rpc.QueryVideoReq {
	if !p.IsSetReq() {
		return QueryVideoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *QueryVideoArgs) IsSetReq() bool {
	return p.Req != nil
}

type QueryVideoResult struct {
	Success *video_rpc.QueryVideoResp
}

var QueryVideoResult_Success_DEFAULT *video_rpc.QueryVideoResp

func (p *QueryVideoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video_rpc.QueryVideoResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *QueryVideoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *QueryVideoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *QueryVideoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in QueryVideoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *QueryVideoResult) Unmarshal(in []byte) error {
	msg := new(video_rpc.QueryVideoResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *QueryVideoResult) GetSuccess() *video_rpc.QueryVideoResp {
	if !p.IsSetSuccess() {
		return QueryVideoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *QueryVideoResult) SetSuccess(x interface{}) {
	p.Success = x.(*video_rpc.QueryVideoResp)
}

func (p *QueryVideoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func createVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video_rpc.CreateVideoReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video_rpc.VideoService).CreateVideo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CreateVideoArgs:
		success, err := handler.(video_rpc.VideoService).CreateVideo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateVideoResult)
		realResult.Success = success
	}
	return nil
}
func newCreateVideoArgs() interface{} {
	return &CreateVideoArgs{}
}

func newCreateVideoResult() interface{} {
	return &CreateVideoResult{}
}

type CreateVideoArgs struct {
	Req *video_rpc.CreateVideoReq
}

func (p *CreateVideoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video_rpc.CreateVideoReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CreateVideoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CreateVideoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CreateVideoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CreateVideoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CreateVideoArgs) Unmarshal(in []byte) error {
	msg := new(video_rpc.CreateVideoReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateVideoArgs_Req_DEFAULT *video_rpc.CreateVideoReq

func (p *CreateVideoArgs) GetReq() *video_rpc.CreateVideoReq {
	if !p.IsSetReq() {
		return CreateVideoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateVideoArgs) IsSetReq() bool {
	return p.Req != nil
}

type CreateVideoResult struct {
	Success *video_rpc.CreateVideoResp
}

var CreateVideoResult_Success_DEFAULT *video_rpc.CreateVideoResp

func (p *CreateVideoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video_rpc.CreateVideoResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CreateVideoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CreateVideoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CreateVideoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CreateVideoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CreateVideoResult) Unmarshal(in []byte) error {
	msg := new(video_rpc.CreateVideoResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateVideoResult) GetSuccess() *video_rpc.CreateVideoResp {
	if !p.IsSetSuccess() {
		return CreateVideoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateVideoResult) SetSuccess(x interface{}) {
	p.Success = x.(*video_rpc.CreateVideoResp)
}

func (p *CreateVideoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func queryVideoByUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video_rpc.QueryVideoByUserReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video_rpc.VideoService).QueryVideoByUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *QueryVideoByUserArgs:
		success, err := handler.(video_rpc.VideoService).QueryVideoByUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*QueryVideoByUserResult)
		realResult.Success = success
	}
	return nil
}
func newQueryVideoByUserArgs() interface{} {
	return &QueryVideoByUserArgs{}
}

func newQueryVideoByUserResult() interface{} {
	return &QueryVideoByUserResult{}
}

type QueryVideoByUserArgs struct {
	Req *video_rpc.QueryVideoByUserReq
}

func (p *QueryVideoByUserArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video_rpc.QueryVideoByUserReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *QueryVideoByUserArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *QueryVideoByUserArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *QueryVideoByUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in QueryVideoByUserArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *QueryVideoByUserArgs) Unmarshal(in []byte) error {
	msg := new(video_rpc.QueryVideoByUserReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var QueryVideoByUserArgs_Req_DEFAULT *video_rpc.QueryVideoByUserReq

func (p *QueryVideoByUserArgs) GetReq() *video_rpc.QueryVideoByUserReq {
	if !p.IsSetReq() {
		return QueryVideoByUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *QueryVideoByUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type QueryVideoByUserResult struct {
	Success *video_rpc.QueryVideoByUserResp
}

var QueryVideoByUserResult_Success_DEFAULT *video_rpc.QueryVideoByUserResp

func (p *QueryVideoByUserResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video_rpc.QueryVideoByUserResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *QueryVideoByUserResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *QueryVideoByUserResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *QueryVideoByUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in QueryVideoByUserResult")
	}
	return proto.Marshal(p.Success)
}

func (p *QueryVideoByUserResult) Unmarshal(in []byte) error {
	msg := new(video_rpc.QueryVideoByUserResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *QueryVideoByUserResult) GetSuccess() *video_rpc.QueryVideoByUserResp {
	if !p.IsSetSuccess() {
		return QueryVideoByUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *QueryVideoByUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*video_rpc.QueryVideoByUserResp)
}

func (p *QueryVideoByUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func queryVideoByIdBatchHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video_rpc.QueryVideoByVideoIdBatchReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video_rpc.VideoService).QueryVideoByIdBatch(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *QueryVideoByIdBatchArgs:
		success, err := handler.(video_rpc.VideoService).QueryVideoByIdBatch(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*QueryVideoByIdBatchResult)
		realResult.Success = success
	}
	return nil
}
func newQueryVideoByIdBatchArgs() interface{} {
	return &QueryVideoByIdBatchArgs{}
}

func newQueryVideoByIdBatchResult() interface{} {
	return &QueryVideoByIdBatchResult{}
}

type QueryVideoByIdBatchArgs struct {
	Req *video_rpc.QueryVideoByVideoIdBatchReq
}

func (p *QueryVideoByIdBatchArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video_rpc.QueryVideoByVideoIdBatchReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *QueryVideoByIdBatchArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *QueryVideoByIdBatchArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *QueryVideoByIdBatchArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in QueryVideoByIdBatchArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *QueryVideoByIdBatchArgs) Unmarshal(in []byte) error {
	msg := new(video_rpc.QueryVideoByVideoIdBatchReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var QueryVideoByIdBatchArgs_Req_DEFAULT *video_rpc.QueryVideoByVideoIdBatchReq

func (p *QueryVideoByIdBatchArgs) GetReq() *video_rpc.QueryVideoByVideoIdBatchReq {
	if !p.IsSetReq() {
		return QueryVideoByIdBatchArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *QueryVideoByIdBatchArgs) IsSetReq() bool {
	return p.Req != nil
}

type QueryVideoByIdBatchResult struct {
	Success *video_rpc.QueryVideoByVideoIdBatchResp
}

var QueryVideoByIdBatchResult_Success_DEFAULT *video_rpc.QueryVideoByVideoIdBatchResp

func (p *QueryVideoByIdBatchResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video_rpc.QueryVideoByVideoIdBatchResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *QueryVideoByIdBatchResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *QueryVideoByIdBatchResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *QueryVideoByIdBatchResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in QueryVideoByIdBatchResult")
	}
	return proto.Marshal(p.Success)
}

func (p *QueryVideoByIdBatchResult) Unmarshal(in []byte) error {
	msg := new(video_rpc.QueryVideoByVideoIdBatchResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *QueryVideoByIdBatchResult) GetSuccess() *video_rpc.QueryVideoByVideoIdBatchResp {
	if !p.IsSetSuccess() {
		return QueryVideoByIdBatchResult_Success_DEFAULT
	}
	return p.Success
}

func (p *QueryVideoByIdBatchResult) SetSuccess(x interface{}) {
	p.Success = x.(*video_rpc.QueryVideoByVideoIdBatchResp)
}

func (p *QueryVideoByIdBatchResult) IsSetSuccess() bool {
	return p.Success != nil
}

func increaseFavoriteCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video_rpc.IncreaseFavoriteCountReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video_rpc.VideoService).IncreaseFavoriteCount(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *IncreaseFavoriteCountArgs:
		success, err := handler.(video_rpc.VideoService).IncreaseFavoriteCount(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*IncreaseFavoriteCountResult)
		realResult.Success = success
	}
	return nil
}
func newIncreaseFavoriteCountArgs() interface{} {
	return &IncreaseFavoriteCountArgs{}
}

func newIncreaseFavoriteCountResult() interface{} {
	return &IncreaseFavoriteCountResult{}
}

type IncreaseFavoriteCountArgs struct {
	Req *video_rpc.IncreaseFavoriteCountReq
}

func (p *IncreaseFavoriteCountArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video_rpc.IncreaseFavoriteCountReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *IncreaseFavoriteCountArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *IncreaseFavoriteCountArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *IncreaseFavoriteCountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in IncreaseFavoriteCountArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *IncreaseFavoriteCountArgs) Unmarshal(in []byte) error {
	msg := new(video_rpc.IncreaseFavoriteCountReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var IncreaseFavoriteCountArgs_Req_DEFAULT *video_rpc.IncreaseFavoriteCountReq

func (p *IncreaseFavoriteCountArgs) GetReq() *video_rpc.IncreaseFavoriteCountReq {
	if !p.IsSetReq() {
		return IncreaseFavoriteCountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *IncreaseFavoriteCountArgs) IsSetReq() bool {
	return p.Req != nil
}

type IncreaseFavoriteCountResult struct {
	Success *video_rpc.IncreaseFavoriteCountResp
}

var IncreaseFavoriteCountResult_Success_DEFAULT *video_rpc.IncreaseFavoriteCountResp

func (p *IncreaseFavoriteCountResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video_rpc.IncreaseFavoriteCountResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *IncreaseFavoriteCountResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *IncreaseFavoriteCountResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *IncreaseFavoriteCountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in IncreaseFavoriteCountResult")
	}
	return proto.Marshal(p.Success)
}

func (p *IncreaseFavoriteCountResult) Unmarshal(in []byte) error {
	msg := new(video_rpc.IncreaseFavoriteCountResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *IncreaseFavoriteCountResult) GetSuccess() *video_rpc.IncreaseFavoriteCountResp {
	if !p.IsSetSuccess() {
		return IncreaseFavoriteCountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *IncreaseFavoriteCountResult) SetSuccess(x interface{}) {
	p.Success = x.(*video_rpc.IncreaseFavoriteCountResp)
}

func (p *IncreaseFavoriteCountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func decreaseFavoriteCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video_rpc.DecreaseFavoriteCountReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video_rpc.VideoService).DecreaseFavoriteCount(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *DecreaseFavoriteCountArgs:
		success, err := handler.(video_rpc.VideoService).DecreaseFavoriteCount(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DecreaseFavoriteCountResult)
		realResult.Success = success
	}
	return nil
}
func newDecreaseFavoriteCountArgs() interface{} {
	return &DecreaseFavoriteCountArgs{}
}

func newDecreaseFavoriteCountResult() interface{} {
	return &DecreaseFavoriteCountResult{}
}

type DecreaseFavoriteCountArgs struct {
	Req *video_rpc.DecreaseFavoriteCountReq
}

func (p *DecreaseFavoriteCountArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video_rpc.DecreaseFavoriteCountReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DecreaseFavoriteCountArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DecreaseFavoriteCountArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DecreaseFavoriteCountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DecreaseFavoriteCountArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *DecreaseFavoriteCountArgs) Unmarshal(in []byte) error {
	msg := new(video_rpc.DecreaseFavoriteCountReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DecreaseFavoriteCountArgs_Req_DEFAULT *video_rpc.DecreaseFavoriteCountReq

func (p *DecreaseFavoriteCountArgs) GetReq() *video_rpc.DecreaseFavoriteCountReq {
	if !p.IsSetReq() {
		return DecreaseFavoriteCountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DecreaseFavoriteCountArgs) IsSetReq() bool {
	return p.Req != nil
}

type DecreaseFavoriteCountResult struct {
	Success *video_rpc.DecreaseFavoriteCountResp
}

var DecreaseFavoriteCountResult_Success_DEFAULT *video_rpc.DecreaseFavoriteCountResp

func (p *DecreaseFavoriteCountResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video_rpc.DecreaseFavoriteCountResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DecreaseFavoriteCountResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DecreaseFavoriteCountResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DecreaseFavoriteCountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DecreaseFavoriteCountResult")
	}
	return proto.Marshal(p.Success)
}

func (p *DecreaseFavoriteCountResult) Unmarshal(in []byte) error {
	msg := new(video_rpc.DecreaseFavoriteCountResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DecreaseFavoriteCountResult) GetSuccess() *video_rpc.DecreaseFavoriteCountResp {
	if !p.IsSetSuccess() {
		return DecreaseFavoriteCountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DecreaseFavoriteCountResult) SetSuccess(x interface{}) {
	p.Success = x.(*video_rpc.DecreaseFavoriteCountResp)
}

func (p *DecreaseFavoriteCountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func increaseCommentCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video_rpc.IncreaseCommentCountReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video_rpc.VideoService).IncreaseCommentCount(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *IncreaseCommentCountArgs:
		success, err := handler.(video_rpc.VideoService).IncreaseCommentCount(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*IncreaseCommentCountResult)
		realResult.Success = success
	}
	return nil
}
func newIncreaseCommentCountArgs() interface{} {
	return &IncreaseCommentCountArgs{}
}

func newIncreaseCommentCountResult() interface{} {
	return &IncreaseCommentCountResult{}
}

type IncreaseCommentCountArgs struct {
	Req *video_rpc.IncreaseCommentCountReq
}

func (p *IncreaseCommentCountArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video_rpc.IncreaseCommentCountReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *IncreaseCommentCountArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *IncreaseCommentCountArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *IncreaseCommentCountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in IncreaseCommentCountArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *IncreaseCommentCountArgs) Unmarshal(in []byte) error {
	msg := new(video_rpc.IncreaseCommentCountReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var IncreaseCommentCountArgs_Req_DEFAULT *video_rpc.IncreaseCommentCountReq

func (p *IncreaseCommentCountArgs) GetReq() *video_rpc.IncreaseCommentCountReq {
	if !p.IsSetReq() {
		return IncreaseCommentCountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *IncreaseCommentCountArgs) IsSetReq() bool {
	return p.Req != nil
}

type IncreaseCommentCountResult struct {
	Success *video_rpc.IncreaseCommentCountResp
}

var IncreaseCommentCountResult_Success_DEFAULT *video_rpc.IncreaseCommentCountResp

func (p *IncreaseCommentCountResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video_rpc.IncreaseCommentCountResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *IncreaseCommentCountResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *IncreaseCommentCountResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *IncreaseCommentCountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in IncreaseCommentCountResult")
	}
	return proto.Marshal(p.Success)
}

func (p *IncreaseCommentCountResult) Unmarshal(in []byte) error {
	msg := new(video_rpc.IncreaseCommentCountResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *IncreaseCommentCountResult) GetSuccess() *video_rpc.IncreaseCommentCountResp {
	if !p.IsSetSuccess() {
		return IncreaseCommentCountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *IncreaseCommentCountResult) SetSuccess(x interface{}) {
	p.Success = x.(*video_rpc.IncreaseCommentCountResp)
}

func (p *IncreaseCommentCountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func decreaseCommentCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video_rpc.DecreaseCommentCountReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video_rpc.VideoService).DecreaseCommentCount(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *DecreaseCommentCountArgs:
		success, err := handler.(video_rpc.VideoService).DecreaseCommentCount(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DecreaseCommentCountResult)
		realResult.Success = success
	}
	return nil
}
func newDecreaseCommentCountArgs() interface{} {
	return &DecreaseCommentCountArgs{}
}

func newDecreaseCommentCountResult() interface{} {
	return &DecreaseCommentCountResult{}
}

type DecreaseCommentCountArgs struct {
	Req *video_rpc.DecreaseCommentCountReq
}

func (p *DecreaseCommentCountArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video_rpc.DecreaseCommentCountReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DecreaseCommentCountArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DecreaseCommentCountArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DecreaseCommentCountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DecreaseCommentCountArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *DecreaseCommentCountArgs) Unmarshal(in []byte) error {
	msg := new(video_rpc.DecreaseCommentCountReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DecreaseCommentCountArgs_Req_DEFAULT *video_rpc.DecreaseCommentCountReq

func (p *DecreaseCommentCountArgs) GetReq() *video_rpc.DecreaseCommentCountReq {
	if !p.IsSetReq() {
		return DecreaseCommentCountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DecreaseCommentCountArgs) IsSetReq() bool {
	return p.Req != nil
}

type DecreaseCommentCountResult struct {
	Success *video_rpc.DecreaseCommentCountResp
}

var DecreaseCommentCountResult_Success_DEFAULT *video_rpc.DecreaseCommentCountResp

func (p *DecreaseCommentCountResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video_rpc.DecreaseCommentCountResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DecreaseCommentCountResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DecreaseCommentCountResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DecreaseCommentCountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DecreaseCommentCountResult")
	}
	return proto.Marshal(p.Success)
}

func (p *DecreaseCommentCountResult) Unmarshal(in []byte) error {
	msg := new(video_rpc.DecreaseCommentCountResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DecreaseCommentCountResult) GetSuccess() *video_rpc.DecreaseCommentCountResp {
	if !p.IsSetSuccess() {
		return DecreaseCommentCountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DecreaseCommentCountResult) SetSuccess(x interface{}) {
	p.Success = x.(*video_rpc.DecreaseCommentCountResp)
}

func (p *DecreaseCommentCountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getAuthorHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video_rpc.GetAuthorReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video_rpc.VideoService).GetAuthor(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetAuthorArgs:
		success, err := handler.(video_rpc.VideoService).GetAuthor(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetAuthorResult)
		realResult.Success = success
	}
	return nil
}
func newGetAuthorArgs() interface{} {
	return &GetAuthorArgs{}
}

func newGetAuthorResult() interface{} {
	return &GetAuthorResult{}
}

type GetAuthorArgs struct {
	Req *video_rpc.GetAuthorReq
}

func (p *GetAuthorArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video_rpc.GetAuthorReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetAuthorArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetAuthorArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetAuthorArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetAuthorArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetAuthorArgs) Unmarshal(in []byte) error {
	msg := new(video_rpc.GetAuthorReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetAuthorArgs_Req_DEFAULT *video_rpc.GetAuthorReq

func (p *GetAuthorArgs) GetReq() *video_rpc.GetAuthorReq {
	if !p.IsSetReq() {
		return GetAuthorArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetAuthorArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetAuthorResult struct {
	Success *video_rpc.GetAuthorResp
}

var GetAuthorResult_Success_DEFAULT *video_rpc.GetAuthorResp

func (p *GetAuthorResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video_rpc.GetAuthorResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetAuthorResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetAuthorResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetAuthorResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetAuthorResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetAuthorResult) Unmarshal(in []byte) error {
	msg := new(video_rpc.GetAuthorResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetAuthorResult) GetSuccess() *video_rpc.GetAuthorResp {
	if !p.IsSetSuccess() {
		return GetAuthorResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetAuthorResult) SetSuccess(x interface{}) {
	p.Success = x.(*video_rpc.GetAuthorResp)
}

func (p *GetAuthorResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) QueryVideo(ctx context.Context, Req *video_rpc.QueryVideoReq) (r *video_rpc.QueryVideoResp, err error) {
	var _args QueryVideoArgs
	_args.Req = Req
	var _result QueryVideoResult
	if err = p.c.Call(ctx, "QueryVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateVideo(ctx context.Context, Req *video_rpc.CreateVideoReq) (r *video_rpc.CreateVideoResp, err error) {
	var _args CreateVideoArgs
	_args.Req = Req
	var _result CreateVideoResult
	if err = p.c.Call(ctx, "CreateVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryVideoByUser(ctx context.Context, Req *video_rpc.QueryVideoByUserReq) (r *video_rpc.QueryVideoByUserResp, err error) {
	var _args QueryVideoByUserArgs
	_args.Req = Req
	var _result QueryVideoByUserResult
	if err = p.c.Call(ctx, "QueryVideoByUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryVideoByIdBatch(ctx context.Context, Req *video_rpc.QueryVideoByVideoIdBatchReq) (r *video_rpc.QueryVideoByVideoIdBatchResp, err error) {
	var _args QueryVideoByIdBatchArgs
	_args.Req = Req
	var _result QueryVideoByIdBatchResult
	if err = p.c.Call(ctx, "QueryVideoByIdBatch", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) IncreaseFavoriteCount(ctx context.Context, Req *video_rpc.IncreaseFavoriteCountReq) (r *video_rpc.IncreaseFavoriteCountResp, err error) {
	var _args IncreaseFavoriteCountArgs
	_args.Req = Req
	var _result IncreaseFavoriteCountResult
	if err = p.c.Call(ctx, "IncreaseFavoriteCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DecreaseFavoriteCount(ctx context.Context, Req *video_rpc.DecreaseFavoriteCountReq) (r *video_rpc.DecreaseFavoriteCountResp, err error) {
	var _args DecreaseFavoriteCountArgs
	_args.Req = Req
	var _result DecreaseFavoriteCountResult
	if err = p.c.Call(ctx, "DecreaseFavoriteCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) IncreaseCommentCount(ctx context.Context, Req *video_rpc.IncreaseCommentCountReq) (r *video_rpc.IncreaseCommentCountResp, err error) {
	var _args IncreaseCommentCountArgs
	_args.Req = Req
	var _result IncreaseCommentCountResult
	if err = p.c.Call(ctx, "IncreaseCommentCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DecreaseCommentCount(ctx context.Context, Req *video_rpc.DecreaseCommentCountReq) (r *video_rpc.DecreaseCommentCountResp, err error) {
	var _args DecreaseCommentCountArgs
	_args.Req = Req
	var _result DecreaseCommentCountResult
	if err = p.c.Call(ctx, "DecreaseCommentCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetAuthor(ctx context.Context, Req *video_rpc.GetAuthorReq) (r *video_rpc.GetAuthorResp, err error) {
	var _args GetAuthorArgs
	_args.Req = Req
	var _result GetAuthorResult
	if err = p.c.Call(ctx, "GetAuthor", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}