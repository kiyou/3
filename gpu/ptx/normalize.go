package ptx

//This file is auto-generated. Editing is futile.

func init() { Code["normalize"] = NORMALIZE }

const NORMALIZE = `
//
// Generated by NVIDIA NVVM Compiler
// Compiler built on Sat Sep 22 02:35:14 2012 (1348274114)
// Cuda compilation tools, release 5.0, V0.2.1221
//

.version 3.1
.target sm_30
.address_size 64

	.file	1 "/tmp/tmpxft_00001bdb_00000000-9_normalize.cpp3.i"
	.file	2 "/home/arne/src/code.google.com/p/mx3/gpu/ptx/normalize.cu"
	.file	3 "/usr/local/cuda-5.0/nvvm/ci_include.h"

.visible .entry normalize(
	.param .u64 normalize_param_0,
	.param .u64 normalize_param_1,
	.param .u64 normalize_param_2,
	.param .u32 normalize_param_3
)
{
	.reg .pred 	%p<3>;
	.reg .s32 	%r<15>;
	.reg .f32 	%f<15>;
	.reg .s64 	%rd<11>;


	ld.param.u64 	%rd7, [normalize_param_0];
	ld.param.u64 	%rd8, [normalize_param_1];
	ld.param.u64 	%rd9, [normalize_param_2];
	ld.param.u32 	%r2, [normalize_param_3];
	cvta.to.global.u64 	%rd1, %rd9;
	cvta.to.global.u64 	%rd2, %rd8;
	cvta.to.global.u64 	%rd3, %rd7;
	.loc 2 6 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 2 7 1
	setp.ge.s32 	%p1, %r1, %r2;
	@%p1 bra 	BB0_4;

	.loc 2 9 1
	mul.wide.s32 	%rd10, %r1, 4;
	add.s64 	%rd4, %rd3, %rd10;
	add.s64 	%rd5, %rd2, %rd10;
	add.s64 	%rd6, %rd1, %rd10;
	ld.global.f32 	%f1, [%rd4];
	ld.global.f32 	%f2, [%rd5];
	.loc 2 10 1
	mul.f32 	%f8, %f2, %f2;
	fma.rn.f32 	%f9, %f1, %f1, %f8;
	.loc 2 9 1
	ld.global.f32 	%f3, [%rd6];
	.loc 2 10 1
	fma.rn.f32 	%f10, %f3, %f3, %f9;
	.loc 3 991 5
	sqrt.rn.f32 	%f4, %f10;
	mov.f32 	%f14, 0f00000000;
	.loc 2 10 1
	setp.eq.f32 	%p2, %f4, 0f00000000;
	@%p2 bra 	BB0_3;

	rcp.rn.f32 	%f14, %f4;

BB0_3:
	mul.f32 	%f11, %f14, %f1;
	.loc 2 11 1
	st.global.f32 	[%rd4], %f11;
	.loc 2 10 1
	mul.f32 	%f12, %f14, %f2;
	.loc 2 12 1
	st.global.f32 	[%rd5], %f12;
	.loc 2 10 1
	mul.f32 	%f13, %f14, %f3;
	.loc 2 13 1
	st.global.f32 	[%rd6], %f13;

BB0_4:
	.loc 2 15 2
	ret;
}


`
