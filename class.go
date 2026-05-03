package astiav

//#include "class.h"
import "C"
import (
	"fmt"
	"sync"
	"unsafe"
)

// https://ffmpeg.org/doxygen/8.0/structAVClass.html
type Class struct {
	c   *C.AVClass
	ptr unsafe.Pointer
}

func newClassFromC(ptr unsafe.Pointer) *Class {
	if ptr == nil {
		return nil
	}
	c := (**C.AVClass)(ptr)
	if c == nil {
		return nil
	}
	return &Class{
		c:   *c,
		ptr: ptr,
	}
}

// https://ffmpeg.org/doxygen/8.0/structAVClass.html#a5fc161d93a0d65a608819da20b7203ba
//
// Returns ClassCategoryNa for a Class with a nil receiver or nil
// AVClass pointer — defense in depth against malformed Class objects
// reaching user code (e.g. via stale pointers in log callbacks).
func (c *Class) Category() ClassCategory {
	if c == nil || c.c == nil {
		return ClassCategoryNa
	}
	return ClassCategory(C.astiavClassCategory(c.c, c.ptr))
}

// https://ffmpeg.org/doxygen/8.0/structAVClass.html#ad763b2e6a0846234a165e74574a550bd
func (c *Class) ItemName() string {
	if c == nil || c.c == nil {
		return ""
	}
	return C.GoString(C.astiavClassItemName(c.c, c.ptr))
}

// https://ffmpeg.org/doxygen/8.0/structAVClass.html#aa8883e113a3f2965abd008f7667db7eb
func (c *Class) Name() string {
	if c == nil || c.c == nil {
		return ""
	}
	return C.GoString(c.c.class_name)
}

// https://ffmpeg.org/doxygen/8.0/structAVClass.html#a88948c8a7c6515181771615a54a808bf
func (c *Class) Parent() *Class {
	if c == nil || c.c == nil {
		return nil
	}
	return newClassFromC(unsafe.Pointer(C.astiavClassParent(c.c, c.ptr)))
}

func (c *Class) String() string {
	if c == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%s [%s] @ %p", c.ItemName(), c.Name(), c.ptr)
}

type Classer interface {
	Class() *Class
}

var _ Classer = (*UnknownClasser)(nil)

type UnknownClasser struct {
	c   *Class
	ptr unsafe.Pointer
}

// PointerString returns a hex-formatted address of the originating
// FFmpeg context for diagnostic log lines (e.g. "0xc000123456"). It
// returns "<nil>" for a nil receiver or when no pointer is associated.
//
// The raw pointer itself is intentionally not exposed: the underlying
// memory may already be freed (the entire reason this classer is
// "unknown" — see newUnknownClasser), so dereferencing it is a UAF.
// Returning a string instead of unsafe.Pointer makes that misuse
// impossible at the API boundary.
func (c *UnknownClasser) PointerString() string {
	if c == nil || c.ptr == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%p", c.ptr)
}

// newUnknownClasser returns an UnknownClasser identified by ptr.
//
// It deliberately does NOT dereference ptr. FFmpeg may emit log
// callbacks with a ptr to a context that has already been freed by a
// background thread (e.g. android_camera HAL callbacks racing
// avformat_close_input). Reading *(**AVClass)(ptr) on freed memory
// is a use-after-read that yields a garbage *AVClass — any later
// call to a Class method would dereference that garbage and crash
// (observed signature: SIGSEGV inside astiavClassCategory).
//
// Callers that need to format Class info must obtain it via the
// classers pool while the object is registered. An UnknownClasser
// returned here has Class() == nil to signal "no usable class info";
// existing user code (e.g. logrus formatter) already treats nil
// Class as a no-op.
func newUnknownClasser(ptr unsafe.Pointer) *UnknownClasser {
	return &UnknownClasser{ptr: ptr}
}

func (c *UnknownClasser) Class() *Class {
	return c.c
}

var _ Classer = (*ClonedClasser)(nil)

type ClonedClasser struct {
	c *Class
}

func newClonedClasser(c Classer) *ClonedClasser {
	cl := c.Class()
	if cl == nil {
		return nil
	}
	return &ClonedClasser{c: newClassFromC(cl.ptr)}
}

func (c *ClonedClasser) Class() *Class {
	return c.c
}

var classers = newClasserPool()

type classerPool struct {
	m sync.Mutex
	p map[unsafe.Pointer]Classer
}

func newClasserPool() *classerPool {
	return &classerPool{p: make(map[unsafe.Pointer]Classer)}
}

func (p *classerPool) unsafePointer(c Classer) unsafe.Pointer {
	if c == nil {
		return nil
	}
	cl := c.Class()
	if cl == nil {
		return nil
	}
	return cl.ptr
}

func (p *classerPool) set(c Classer) {
	p.m.Lock()
	defer p.m.Unlock()
	if ptr := p.unsafePointer(c); ptr != nil {
		p.p[ptr] = c
	}
}

func (p *classerPool) del(c Classer) {
	p.m.Lock()
	defer p.m.Unlock()
	if ptr := p.unsafePointer(c); ptr != nil {
		delete(p.p, ptr)
	}
}

func (p *classerPool) get(ptr unsafe.Pointer) (Classer, bool) {
	p.m.Lock()
	defer p.m.Unlock()
	c, ok := p.p[ptr]
	return c, ok
}
