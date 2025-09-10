/**
 * @author: dn-jinmin/dn-jinmin
 * @doc:
 */

package zlog

import "runtime"

// 获取代码的执行行数
func getCallerFrame(skip int) (frame runtime.Frame, ok bool) {
	pc := make([]uintptr, 1)
	numFrames := runtime.Callers(skip, pc)
	if numFrames < 1 {
		return
	}

	frame, _ = runtime.CallersFrames(pc).Next()
	return frame, frame.PC != 0
}
