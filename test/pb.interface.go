package chat

import (
	"context"
	"fmt"
	"net"
	"unicode"

	js "github.com/llsix/GBR/Utils/JS"
	grpc "google.golang.org/grpc"
)

type MethodsHandle struct {
	FullMethod string
	in         interface{}
}

type Methods = func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error)
type msg_buff struct {
	Name    string
	MsgType func() interface{}
}

var _sd_ = _Chat_serviceDesc
var grpc_msg = []msg_buff{
	{Name: "newHelloRequest", MsgType: func() interface{} { return new(HelloRequest) }},
}

var grpc_jsm *js.JSM
var grpc_jsfile = "./script/grpc.js"
var grpc_server *grpc.Server
var network, address string
var ctx, cancel = context.WithCancel(context.Background())

func Wait() {
	<-ctx.Done()
}

func init() {
	if grpc_jsm == nil {
		setup()
		grpc_jsm.StartFileChangeMonitor(func(jsm *js.JSM) {
			setup()
			grpc_server.Stop()
			grpc_server = grpc.NewServer()
			grpc_server.RegisterService(Packe_servierDesc(&_sd_), nil)
			if listener, err := net.Listen(network, address); err == nil {
				go func() {
					if err := grpc_server.Serve(listener); err != nil {
						cancel()
					}
				}()
			} else {
				panic(err)
			}

		})
	}
}

func setup() {
	grpc_jsm = js.NewSourceJS("grpc", grpc_jsfile)
	grpc_jsm.GetOtto().Object("grpc={}")
	g, _ := grpc_jsm.GetOtto().Get("grpc")
	for _, tp := range grpc_msg {
		g.Object().Set(tp.Name, func() interface{} { return tp.MsgType() })
	}

	grpc_jsm.RunJSFile(grpc_jsfile)
}

func NewMethodsHandle(FullMethod string, in interface{}, Func func(context.Context) (interface{}, error)) Methods {
	m := MethodsHandle{FullMethod: FullMethod, in: in}
	return func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
		in := m.in
		if err := dec(in); err != nil {
			return nil, err
		}
		if interceptor == nil {
			return Func(ctx)
		}
		info := &grpc.UnaryServerInfo{
			Server:     srv,
			FullMethod: m.FullMethod,
		}
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return Func(ctx)
		}
		return interceptor(ctx, in, info, handler)
	}
}

func RegisterChatServer2(listener net.Listener, s *grpc.Server, srv ChatServer) {
	network = listener.Addr().Network()
	address = listener.Addr().String()
	grpc_server = s
	grpc_server.RegisterService(Packe_servierDesc(&_sd_), nil)
	go func() {
		if err := grpc_server.Serve(listener); err != nil {
			cancel()
		}
	}()
}

func Packe_servierDesc(sd *grpc.ServiceDesc) *grpc.ServiceDesc {
	_sd_ := &grpc.ServiceDesc{
		ServiceName: sd.ServiceName,
		HandlerType: nil,
		Methods:     []grpc.MethodDesc{},
		Metadata:    sd.Metadata,
	}
	if len(_sd_.Methods) > 0 {
		_sd_.Methods = []grpc.MethodDesc{}
	}
	v, _ := grpc_jsm.GetOtto().Get("grpc")
	if !v.IsObject() {
		grpc_jsm.GetOtto().Object("grpc={}")
		v, _ = grpc_jsm.Get("grpc")
	}

	for _, fc := range v.Object().Keys() {
		fv, _ := v.Object().Get(fc)
		if fv.IsObject() && len(fc) > 0 && unicode.IsUpper(rune(fc[0])) {
			In, err := fv.Object().Get("In")
			in, _ := In.Export()
			_sd_.Methods = append(_sd_.Methods, grpc.MethodDesc{
				MethodName: fc,
				Handler: NewMethodsHandle(fmt.Sprintf("/%v/%v", sd.ServiceName, fc), in, func(ctx context.Context) (interface{}, error) {

					g, _ := grpc_jsm.Get("grpc")
					Func, _ := g.Object().Get(fc)
					function, _ := Func.Object().Get("Func")
					if vv, err := function.Call(function, ctx, in); err == nil {

						return vv.Export()
					} else {

					}
					return nil, err
				}),
			})
		}
	}
	return _sd_
}
