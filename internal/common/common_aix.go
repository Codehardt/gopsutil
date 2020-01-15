//+build aix

package common

import (
	"context"
	"sync/atomic"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

var cachedBootTime uint64

func BootTimeWithContext(ctx context.Context) (uint64, error) {
	t := atomic.LoadUint64(&cachedBootTime)
	if t != 0 {
		return t, nil
	}
	buf, err := unix.SysctlRaw("kern.boottime")
	if err != nil {
		return 0, err
	}

	tv := *(*syscall.Timeval)(unsafe.Pointer((&buf[0])))
	atomic.StoreUint64(&cachedBootTime, uint64(tv.Sec))

	return t, nil
}
