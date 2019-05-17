/* Code generated by cmd/cgo; DO NOT EDIT. */

/* package golib */


#line 1 "cgo-builtin-export-prolog"

#include <stddef.h> /* for ptrdiff_t below */

#ifndef GO_CGO_EXPORT_PROLOGUE_H
#define GO_CGO_EXPORT_PROLOGUE_H

#ifndef GO_CGO_GOSTRING_TYPEDEF
typedef struct { const char *p; ptrdiff_t n; } _GoString_;
#endif

#endif

/* Start of preamble from import "C" comments.  */




/* End of preamble from import "C" comments.  */


/* Start of boilerplate cgo prologue.  */
#line 1 "cgo-gcc-export-header-prolog"

#ifndef GO_CGO_PROLOGUE_H
#define GO_CGO_PROLOGUE_H

typedef signed char GoInt8;
typedef unsigned char GoUint8;
typedef short GoInt16;
typedef unsigned short GoUint16;
typedef int GoInt32;
typedef unsigned int GoUint32;
typedef long long GoInt64;
typedef unsigned long long GoUint64;
typedef GoInt64 GoInt;
typedef GoUint64 GoUint;
typedef __SIZE_TYPE__ GoUintptr;
typedef float GoFloat32;
typedef double GoFloat64;
typedef float _Complex GoComplex64;
typedef double _Complex GoComplex128;

/*
  static assertion to make sure the file is being used on architecture
  at least with matching size of GoInt.
*/
typedef char _check_for_64_bit_pointer_matching_GoInt[sizeof(void*)==64/8 ? 1:-1];

#ifndef GO_CGO_GOSTRING_TYPEDEF
typedef _GoString_ GoString;
#endif
typedef void *GoMap;
typedef void *GoChan;
typedef struct { void *t; void *v; } GoInterface;
typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;

#endif

/* End of boilerplate cgo prologue.  */

#ifdef __cplusplus
extern "C" {
#endif


extern GoInterface FileExist(GoString p0);

/* Return type for IsExist */
struct IsExist_return {
	GoUint8 r0;
	GoInterface r1;
};

extern struct IsExist_return IsExist(GoString p0);

/* Return type for ReadDir */
struct ReadDir_return {
	GoSlice r0;
	GoInterface r1;
};

extern struct ReadDir_return ReadDir(GoString p0);

extern GoInterface CreateFile(GoString p0);

/* Return type for MetaData */
struct MetaData_return {
	GoString r0;
	GoInterface r1;
};

/**
 * @note 获取文件信息
 * @params path,pathType string
 * @param path 文件路径
 * @param pathType 文件空间
 * @return Entry,error
 */

extern struct MetaData_return MetaData(GoString p0, GoString p1);

//params username string,password string

extern GoInterface Login(GoString p0, GoString p1);

//params filePath, path, pathType, neid, from string, overwrite bool

extern GoInterface UploadFile(GoString p0, GoString p1, GoString p2, GoString p3, GoString p4, GoUint8 p5);

#ifdef __cplusplus
}
#endif
