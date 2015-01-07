package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/mumax/3/cuda/cu"
	"sync"
	"unsafe"
)

// CUDA handle for zeromask kernel
var zeromask_code cu.Function

// Stores the arguments for zeromask kernel invocation
type zeromask_args_t struct {
	arg_dst     unsafe.Pointer
	arg_maskLUT unsafe.Pointer
	arg_regions unsafe.Pointer
	arg_N       int
	argptr      [4]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for zeromask kernel invocation
var zeromask_args zeromask_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	zeromask_args.argptr[0] = unsafe.Pointer(&zeromask_args.arg_dst)
	zeromask_args.argptr[1] = unsafe.Pointer(&zeromask_args.arg_maskLUT)
	zeromask_args.argptr[2] = unsafe.Pointer(&zeromask_args.arg_regions)
	zeromask_args.argptr[3] = unsafe.Pointer(&zeromask_args.arg_N)
}

// Wrapper for zeromask CUDA kernel, asynchronous.
func k_zeromask_async(dst unsafe.Pointer, maskLUT unsafe.Pointer, regions unsafe.Pointer, N int, cfg *config) {
	if Synchronous { // debug
		Sync()
	}

	zeromask_args.Lock()
	defer zeromask_args.Unlock()

	if zeromask_code == 0 {
		zeromask_code = fatbinLoad(zeromask_map, "zeromask")
	}

	zeromask_args.arg_dst = dst
	zeromask_args.arg_maskLUT = maskLUT
	zeromask_args.arg_regions = regions
	zeromask_args.arg_N = N

	args := zeromask_args.argptr[:]
	cu.LaunchKernel(zeromask_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
	}
}

// maps compute capability on PTX code for zeromask kernel.
var zeromask_map = map[int]string{0: "",
	20: zeromask_ptx_20,
	30: zeromask_ptx_30,
	35: zeromask_ptx_35}

// zeromask PTX code for various compute capabilities.
const (
	zeromask_ptx_20 = `
.version 4.0
.target sm_20
.address_size 64


.visible .entry zeromask(
	.param .u64 zeromask_param_0,
	.param .u64 zeromask_param_1,
	.param .u64 zeromask_param_2,
	.param .u32 zeromask_param_3
)
{
	.reg .pred 	%p<3>;
	.reg .s32 	%r<10>;
	.reg .f32 	%f<2>;
	.reg .s64 	%rd<14>;


	ld.param.u64 	%rd2, [zeromask_param_0];
	ld.param.u64 	%rd3, [zeromask_param_1];
	ld.param.u64 	%rd4, [zeromask_param_2];
	ld.param.u32 	%r2, [zeromask_param_3];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_3;

	cvta.to.global.u64 	%rd5, %rd3;
	cvta.to.global.u64 	%rd6, %rd4;
	cvt.s64.s32	%rd1, %r1;
	add.s64 	%rd7, %rd6, %rd1;
	ld.global.u8 	%rd8, [%rd7];
	shl.b64 	%rd9, %rd8, 2;
	add.s64 	%rd10, %rd5, %rd9;
	ld.global.f32 	%f1, [%rd10];
	setp.eq.f32	%p2, %f1, 0f00000000;
	@%p2 bra 	BB0_3;

	cvta.to.global.u64 	%rd11, %rd2;
	shl.b64 	%rd12, %rd1, 2;
	add.s64 	%rd13, %rd11, %rd12;
	mov.u32 	%r9, 0;
	st.global.u32 	[%rd13], %r9;

BB0_3:
	ret;
}


`
	zeromask_ptx_30 = `
.version 4.0
.target sm_30
.address_size 64


.visible .entry zeromask(
	.param .u64 zeromask_param_0,
	.param .u64 zeromask_param_1,
	.param .u64 zeromask_param_2,
	.param .u32 zeromask_param_3
)
{
	.reg .pred 	%p<3>;
	.reg .s32 	%r<10>;
	.reg .f32 	%f<2>;
	.reg .s64 	%rd<14>;


	ld.param.u64 	%rd2, [zeromask_param_0];
	ld.param.u64 	%rd3, [zeromask_param_1];
	ld.param.u64 	%rd4, [zeromask_param_2];
	ld.param.u32 	%r2, [zeromask_param_3];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_3;

	cvta.to.global.u64 	%rd5, %rd3;
	cvta.to.global.u64 	%rd6, %rd4;
	cvt.s64.s32	%rd1, %r1;
	add.s64 	%rd7, %rd6, %rd1;
	ld.global.u8 	%rd8, [%rd7];
	shl.b64 	%rd9, %rd8, 2;
	add.s64 	%rd10, %rd5, %rd9;
	ld.global.f32 	%f1, [%rd10];
	setp.eq.f32	%p2, %f1, 0f00000000;
	@%p2 bra 	BB0_3;

	cvta.to.global.u64 	%rd11, %rd2;
	shl.b64 	%rd12, %rd1, 2;
	add.s64 	%rd13, %rd11, %rd12;
	mov.u32 	%r9, 0;
	st.global.u32 	[%rd13], %r9;

BB0_3:
	ret;
}


`
	zeromask_ptx_35 = `
.version 4.0
.target sm_35
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.visible .entry zeromask(
	.param .u64 zeromask_param_0,
	.param .u64 zeromask_param_1,
	.param .u64 zeromask_param_2,
	.param .u32 zeromask_param_3
)
{
	.reg .pred 	%p<3>;
	.reg .s16 	%rs<2>;
	.reg .s32 	%r<10>;
	.reg .f32 	%f<2>;
	.reg .s64 	%rd<15>;


	ld.param.u64 	%rd2, [zeromask_param_0];
	ld.param.u64 	%rd3, [zeromask_param_1];
	ld.param.u64 	%rd4, [zeromask_param_2];
	ld.param.u32 	%r2, [zeromask_param_3];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB2_3;

	cvta.to.global.u64 	%rd5, %rd3;
	cvta.to.global.u64 	%rd6, %rd4;
	cvt.s64.s32	%rd1, %r1;
	add.s64 	%rd7, %rd6, %rd1;
	ld.global.nc.u8 	%rs1, [%rd7];
	cvt.u64.u16	%rd8, %rs1;
	and.b64  	%rd9, %rd8, 255;
	shl.b64 	%rd10, %rd9, 2;
	add.s64 	%rd11, %rd5, %rd10;
	ld.global.nc.f32 	%f1, [%rd11];
	setp.eq.f32	%p2, %f1, 0f00000000;
	@%p2 bra 	BB2_3;

	cvta.to.global.u64 	%rd12, %rd2;
	shl.b64 	%rd13, %rd1, 2;
	add.s64 	%rd14, %rd12, %rd13;
	mov.u32 	%r9, 0;
	st.global.u32 	[%rd14], %r9;

BB2_3:
	ret;
}


`
)
