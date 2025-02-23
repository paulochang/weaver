// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package adservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

var _ codegen.LatestVersion = codegen.Version[[0][17]struct{}](`

ERROR: You generated this file with 'weaver generate' v0.17.0 (codegen
version v0.17.0). The generated code is incompatible with the version of the
github.com/ServiceWeaver/weaver module that you're using. The weaver module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/ServiceWeaver/weaver

We recommend updating the weaver module and the 'weaver generate' command by
running the following.

    go get github.com/ServiceWeaver/weaver@latest
    go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

Then, re-run 'weaver generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/ServiceWeaver/weaver/issues.

`)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "github.com/ServiceWeaver/weaver/examples/onlineboutique/adservice/T",
		Iface: reflect.TypeOf((*T)(nil)).Elem(),
		Impl:  reflect.TypeOf(impl{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return t_local_stub{impl: impl.(T), tracer: tracer, getAdsMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/onlineboutique/adservice/T", Method: "GetAds", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return t_client_stub{stub: stub, getAdsMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/onlineboutique/adservice/T", Method: "GetAds", Remote: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return t_server_stub{impl: impl.(T), addLoad: addLoad}
		},
		RefData: "",
	})
}

// weaver.Instance checks.
var _ weaver.InstanceOf[T] = (*impl)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*impl)(nil)

// Local stub implementations.

type t_local_stub struct {
	impl          T
	tracer        trace.Tracer
	getAdsMetrics *codegen.MethodMetrics
}

// Check that t_local_stub implements the T interface.
var _ T = (*t_local_stub)(nil)

func (s t_local_stub) GetAds(ctx context.Context, a0 []string) (r0 []Ad, err error) {
	// Update metrics.
	begin := s.getAdsMetrics.Begin()
	defer func() { s.getAdsMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "adservice.T.GetAds", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetAds(ctx, a0)
}

// Client stub implementations.

type t_client_stub struct {
	stub          codegen.Stub
	getAdsMetrics *codegen.MethodMetrics
}

// Check that t_client_stub implements the T interface.
var _ T = (*t_client_stub)(nil)

func (s t_client_stub) GetAds(ctx context.Context, a0 []string) (r0 []Ad, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.getAdsMetrics.Begin()
	defer func() { s.getAdsMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "adservice.T.GetAds", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Encode arguments.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_string_4af10117(enc, a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_Ad_86ae3655(dec)
	err = dec.Error()
	return
}

// Server stub implementations.

type t_server_stub struct {
	impl    T
	addLoad func(key uint64, load float64)
}

// Check that t_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*t_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s t_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "GetAds":
		return s.getAds
	default:
		return nil
	}
}

func (s t_server_stub) getAds(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 []string
	a0 = serviceweaver_dec_slice_string_4af10117(dec)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.GetAds(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_Ad_86ae3655(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

// AutoMarshal implementations.

var _ codegen.AutoMarshal = (*Ad)(nil)

type __is_Ad[T ~struct {
	weaver.AutoMarshal
	RedirectURL string
	Text        string
}] struct{}

var _ __is_Ad[Ad]

func (x *Ad) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Ad.WeaverMarshal: nil receiver"))
	}
	enc.String(x.RedirectURL)
	enc.String(x.Text)
}

func (x *Ad) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Ad.WeaverUnmarshal: nil receiver"))
	}
	x.RedirectURL = dec.String()
	x.Text = dec.String()
}

// Encoding/decoding implementations.

func serviceweaver_enc_slice_string_4af10117(enc *codegen.Encoder, arg []string) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		enc.String(arg[i])
	}
}

func serviceweaver_dec_slice_string_4af10117(dec *codegen.Decoder) []string {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = dec.String()
	}
	return res
}

func serviceweaver_enc_slice_Ad_86ae3655(enc *codegen.Encoder, arg []Ad) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_Ad_86ae3655(dec *codegen.Decoder) []Ad {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]Ad, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}
