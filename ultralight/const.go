// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Sun, 02 Oct 2022 19:40:55 PDT.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package ultralight

/*
#cgo CFLAGS: -I../include
#cgo LDFLAGS: -Wl,--allow-multiple-definition -L${SRCDIR}/libs/linux/x64 -Wl,-rpath,/home/co/code/Muon/ultralight/libs/linux/x64 -lUltralightCore -lAppCore -lUltralight -lWebCore
#include "AppCore/CAPI.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// JSCOBJCAPIENABLED as defined in JavaScriptCore/JSBase.h:151
	JSCOBJCAPIENABLED = 0
)

// ULWindowFlags as declared in AppCore/CAPI.h:47
type ULWindowFlags int32

// ULWindowFlags enumeration from AppCore/CAPI.h:47
const (
	KWindowFlagsBorderless  ULWindowFlags = 1
	KWindowFlagsTitled      ULWindowFlags = 2
	KWindowFlagsResizable   ULWindowFlags = 4
	KWindowFlagsMaximizable ULWindowFlags = 8
)

// ULMessageSource as declared in Ultralight/CAPI.h:75
type ULMessageSource int32

// ULMessageSource enumeration from Ultralight/CAPI.h:75
const (
	KMessageSourceXML            ULMessageSource = iota
	KMessageSourceJS             ULMessageSource = 1
	KMessageSourceNetwork        ULMessageSource = 2
	KMessageSourceConsoleAPI     ULMessageSource = 3
	KMessageSourceStorage        ULMessageSource = 4
	KMessageSourceAppCache       ULMessageSource = 5
	KMessageSourceRendering      ULMessageSource = 6
	KMessageSourceCSS            ULMessageSource = 7
	KMessageSourceSecurity       ULMessageSource = 8
	KMessageSourceContentBlocker ULMessageSource = 9
	KMessageSourceOther          ULMessageSource = 10
)

// ULMessageLevel as declared in Ultralight/CAPI.h:83
type ULMessageLevel int32

// ULMessageLevel enumeration from Ultralight/CAPI.h:83
const (
	KMessageLevelLog     ULMessageLevel = 1
	KMessageLevelWarning ULMessageLevel = 2
	KMessageLevelError   ULMessageLevel = 3
	KMessageLevelDebug   ULMessageLevel = 4
	KMessageLevelInfo    ULMessageLevel = 5
)

// ULCursor as declared in Ultralight/CAPI.h:130
type ULCursor int32

// ULCursor enumeration from Ultralight/CAPI.h:130
const (
	KCursorPointer                  ULCursor = iota
	KCursorCross                    ULCursor = 1
	KCursorHand                     ULCursor = 2
	KCursorIBeam                    ULCursor = 3
	KCursorWait                     ULCursor = 4
	KCursorHelp                     ULCursor = 5
	KCursorEastResize               ULCursor = 6
	KCursorNorthResize              ULCursor = 7
	KCursorNorthEastResize          ULCursor = 8
	KCursorNorthWestResize          ULCursor = 9
	KCursorSouthResize              ULCursor = 10
	KCursorSouthEastResize          ULCursor = 11
	KCursorSouthWestResize          ULCursor = 12
	KCursorWestResize               ULCursor = 13
	KCursorNorthSouthResize         ULCursor = 14
	KCursorEastWestResize           ULCursor = 15
	KCursorNorthEastSouthWestResize ULCursor = 16
	KCursorNorthWestSouthEastResize ULCursor = 17
	KCursorColumnResize             ULCursor = 18
	KCursorRowResize                ULCursor = 19
	KCursorMiddlePanning            ULCursor = 20
	KCursorEastPanning              ULCursor = 21
	KCursorNorthPanning             ULCursor = 22
	KCursorNorthEastPanning         ULCursor = 23
	KCursorNorthWestPanning         ULCursor = 24
	KCursorSouthPanning             ULCursor = 25
	KCursorSouthEastPanning         ULCursor = 26
	KCursorSouthWestPanning         ULCursor = 27
	KCursorWestPanning              ULCursor = 28
	KCursorMove                     ULCursor = 29
	KCursorVerticalText             ULCursor = 30
	KCursorCell                     ULCursor = 31
	KCursorContextMenu              ULCursor = 32
	KCursorAlias                    ULCursor = 33
	KCursorProgress                 ULCursor = 34
	KCursorNoDrop                   ULCursor = 35
	KCursorCopy                     ULCursor = 36
	KCursorNone                     ULCursor = 37
	KCursorNotAllowed               ULCursor = 38
	KCursorZoomIn                   ULCursor = 39
	KCursorZoomOut                  ULCursor = 40
	KCursorGrab                     ULCursor = 41
	KCursorGrabbing                 ULCursor = 42
	KCursorCustom                   ULCursor = 43
)

// ULBitmapFormat as declared in Ultralight/CAPI.h:150
type ULBitmapFormat int32

// ULBitmapFormat enumeration from Ultralight/CAPI.h:150
const (
	KBitmapFormatA8UNORM        ULBitmapFormat = iota
	KBitmapFormatBGRA8UNORMSRGB ULBitmapFormat = 1
)

// ULKeyEventType as declared in Ultralight/CAPI.h:180
type ULKeyEventType int32

// ULKeyEventType enumeration from Ultralight/CAPI.h:180
const (
	KKeyEventTypeKeyDown    ULKeyEventType = iota
	KKeyEventTypeKeyUp      ULKeyEventType = 1
	KKeyEventTypeRawKeyDown ULKeyEventType = 2
	KKeyEventTypeChar       ULKeyEventType = 3
)

// ULMouseEventType as declared in Ultralight/CAPI.h:186
type ULMouseEventType int32

// ULMouseEventType enumeration from Ultralight/CAPI.h:186
const (
	KMouseEventTypeMouseMoved ULMouseEventType = iota
	KMouseEventTypeMouseDown  ULMouseEventType = 1
	KMouseEventTypeMouseUp    ULMouseEventType = 2
)

// ULMouseButton as declared in Ultralight/CAPI.h:193
type ULMouseButton int32

// ULMouseButton enumeration from Ultralight/CAPI.h:193
const (
	KMouseButtonNone   ULMouseButton = iota
	KMouseButtonLeft   ULMouseButton = 1
	KMouseButtonMiddle ULMouseButton = 2
	KMouseButtonRight  ULMouseButton = 3
)

// ULScrollEventType as declared in Ultralight/CAPI.h:198
type ULScrollEventType int32

// ULScrollEventType enumeration from Ultralight/CAPI.h:198
const (
	KScrollEventTypeScrollByPixel ULScrollEventType = iota
	KScrollEventTypeScrollByPage  ULScrollEventType = 1
)

// ULFaceWinding as declared in Ultralight/CAPI.h:203
type ULFaceWinding int32

// ULFaceWinding enumeration from Ultralight/CAPI.h:203
const (
	KFaceWindingClockwise       ULFaceWinding = iota
	KFaceWindowCounterClockwise ULFaceWinding = 1
)

// ULFontHinting as declared in Ultralight/CAPI.h:225
type ULFontHinting int32

// ULFontHinting enumeration from Ultralight/CAPI.h:225
const (
	KFontHintingSmooth     ULFontHinting = iota
	KFontHintingNormal     ULFontHinting = 1
	KFontHintingMonochrome ULFontHinting = 2
)

// ULLogLevel as declared in Ultralight/CAPI.h:1477
type ULLogLevel int32

// ULLogLevel enumeration from Ultralight/CAPI.h:1477
const (
	KLogLevelError   ULLogLevel = iota
	KLogLevelWarning ULLogLevel = 1
	KLogLevelInfo    ULLogLevel = 2
)

// ULVertexBufferFormat as declared in Ultralight/CAPI.h:1552
type ULVertexBufferFormat int32

// ULVertexBufferFormat enumeration from Ultralight/CAPI.h:1552
const (
	KVertexBufferFormat2f4ub2f      ULVertexBufferFormat = iota
	KVertexBufferFormat2f4ub2f2f28f ULVertexBufferFormat = 1
)

// ULShaderType as declared in Ultralight/CAPI.h:1585
type ULShaderType int32

// ULShaderType enumeration from Ultralight/CAPI.h:1585
const (
	KShaderTypeFill     ULShaderType = iota
	KShaderTypeFillPath ULShaderType = 1
)

// ULCommandType as declared in Ultralight/CAPI.h:1660
type ULCommandType int32

// ULCommandType enumeration from Ultralight/CAPI.h:1660
const (
	KCommandTypeClearRenderBuffer ULCommandType = iota
	KCommandTypeDrawGeometry      ULCommandType = 1
)

// JSType as declared in JavaScriptCore/JSValueRef.h:55
type JSType int32

// JSType enumeration from JavaScriptCore/JSValueRef.h:55
const (
	KJSTypeUndefined JSType = iota
	KJSTypeNull      JSType = 1
	KJSTypeBoolean   JSType = 2
	KJSTypeNumber    JSType = 3
	KJSTypeString    JSType = 4
	KJSTypeObject    JSType = 5
	KJSTypeSymbol    JSType = 6
)

// JSTypedArrayType as declared in JavaScriptCore/JSValueRef.h:85
type JSTypedArrayType int32

// JSTypedArrayType enumeration from JavaScriptCore/JSValueRef.h:85
const (
	KJSTypedArrayTypeInt8Array         JSTypedArrayType = iota
	KJSTypedArrayTypeInt16Array        JSTypedArrayType = 1
	KJSTypedArrayTypeInt32Array        JSTypedArrayType = 2
	KJSTypedArrayTypeUint8Array        JSTypedArrayType = 3
	KJSTypedArrayTypeUint8ClampedArray JSTypedArrayType = 4
	KJSTypedArrayTypeUint16Array       JSTypedArrayType = 5
	KJSTypedArrayTypeUint32Array       JSTypedArrayType = 6
	KJSTypedArrayTypeFloat32Array      JSTypedArrayType = 7
	KJSTypedArrayTypeFloat64Array      JSTypedArrayType = 8
	KJSTypedArrayTypeArrayBuffer       JSTypedArrayType = 9
	KJSTypedArrayTypeNone              JSTypedArrayType = 10
)

// ULInvalidFileHandle as declared in Ultralight/CAPI.h:1413
const ULInvalidFileHandle ULFileHandle = -1
const (
	// KJSPropertyAttributeNone as declared in JavaScriptCore/JSObjectRef.h:51
	KJSPropertyAttributeNone = iota
	// KJSPropertyAttributeReadOnly as declared in JavaScriptCore/JSObjectRef.h:52
	KJSPropertyAttributeReadOnly = 2
	// KJSPropertyAttributeDontEnum as declared in JavaScriptCore/JSObjectRef.h:53
	KJSPropertyAttributeDontEnum = 4
	// KJSPropertyAttributeDontDelete as declared in JavaScriptCore/JSObjectRef.h:54
	KJSPropertyAttributeDontDelete = 8
)

const (
	// KJSClassAttributeNone as declared in JavaScriptCore/JSObjectRef.h:69
	KJSClassAttributeNone = iota
	// KJSClassAttributeNoAutomaticPrototype as declared in JavaScriptCore/JSObjectRef.h:70
	KJSClassAttributeNoAutomaticPrototype = 2
)
