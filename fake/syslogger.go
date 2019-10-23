// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"sync"

	"github.com/phogolabs/log/handler/syslog"
)

type Syslogger struct {
	AlertStub        func(string) error
	alertMutex       sync.RWMutex
	alertArgsForCall []struct {
		arg1 string
	}
	alertReturns struct {
		result1 error
	}
	alertReturnsOnCall map[int]struct {
		result1 error
	}
	CritStub        func(string) error
	critMutex       sync.RWMutex
	critArgsForCall []struct {
		arg1 string
	}
	critReturns struct {
		result1 error
	}
	critReturnsOnCall map[int]struct {
		result1 error
	}
	DebugStub        func(string) error
	debugMutex       sync.RWMutex
	debugArgsForCall []struct {
		arg1 string
	}
	debugReturns struct {
		result1 error
	}
	debugReturnsOnCall map[int]struct {
		result1 error
	}
	EmergStub        func(string) error
	emergMutex       sync.RWMutex
	emergArgsForCall []struct {
		arg1 string
	}
	emergReturns struct {
		result1 error
	}
	emergReturnsOnCall map[int]struct {
		result1 error
	}
	ErrStub        func(string) error
	errMutex       sync.RWMutex
	errArgsForCall []struct {
		arg1 string
	}
	errReturns struct {
		result1 error
	}
	errReturnsOnCall map[int]struct {
		result1 error
	}
	InfoStub        func(string) error
	infoMutex       sync.RWMutex
	infoArgsForCall []struct {
		arg1 string
	}
	infoReturns struct {
		result1 error
	}
	infoReturnsOnCall map[int]struct {
		result1 error
	}
	NoticeStub        func(string) error
	noticeMutex       sync.RWMutex
	noticeArgsForCall []struct {
		arg1 string
	}
	noticeReturns struct {
		result1 error
	}
	noticeReturnsOnCall map[int]struct {
		result1 error
	}
	WarningStub        func(string) error
	warningMutex       sync.RWMutex
	warningArgsForCall []struct {
		arg1 string
	}
	warningReturns struct {
		result1 error
	}
	warningReturnsOnCall map[int]struct {
		result1 error
	}
	WriteStub        func([]byte) (int, error)
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		arg1 []byte
	}
	writeReturns struct {
		result1 int
		result2 error
	}
	writeReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Syslogger) Alert(arg1 string) error {
	fake.alertMutex.Lock()
	ret, specificReturn := fake.alertReturnsOnCall[len(fake.alertArgsForCall)]
	fake.alertArgsForCall = append(fake.alertArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Alert", []interface{}{arg1})
	fake.alertMutex.Unlock()
	if fake.AlertStub != nil {
		return fake.AlertStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.alertReturns
	return fakeReturns.result1
}

func (fake *Syslogger) AlertCallCount() int {
	fake.alertMutex.RLock()
	defer fake.alertMutex.RUnlock()
	return len(fake.alertArgsForCall)
}

func (fake *Syslogger) AlertCalls(stub func(string) error) {
	fake.alertMutex.Lock()
	defer fake.alertMutex.Unlock()
	fake.AlertStub = stub
}

func (fake *Syslogger) AlertArgsForCall(i int) string {
	fake.alertMutex.RLock()
	defer fake.alertMutex.RUnlock()
	argsForCall := fake.alertArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Syslogger) AlertReturns(result1 error) {
	fake.alertMutex.Lock()
	defer fake.alertMutex.Unlock()
	fake.AlertStub = nil
	fake.alertReturns = struct {
		result1 error
	}{result1}
}

func (fake *Syslogger) AlertReturnsOnCall(i int, result1 error) {
	fake.alertMutex.Lock()
	defer fake.alertMutex.Unlock()
	fake.AlertStub = nil
	if fake.alertReturnsOnCall == nil {
		fake.alertReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.alertReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Syslogger) Crit(arg1 string) error {
	fake.critMutex.Lock()
	ret, specificReturn := fake.critReturnsOnCall[len(fake.critArgsForCall)]
	fake.critArgsForCall = append(fake.critArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Crit", []interface{}{arg1})
	fake.critMutex.Unlock()
	if fake.CritStub != nil {
		return fake.CritStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.critReturns
	return fakeReturns.result1
}

func (fake *Syslogger) CritCallCount() int {
	fake.critMutex.RLock()
	defer fake.critMutex.RUnlock()
	return len(fake.critArgsForCall)
}

func (fake *Syslogger) CritCalls(stub func(string) error) {
	fake.critMutex.Lock()
	defer fake.critMutex.Unlock()
	fake.CritStub = stub
}

func (fake *Syslogger) CritArgsForCall(i int) string {
	fake.critMutex.RLock()
	defer fake.critMutex.RUnlock()
	argsForCall := fake.critArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Syslogger) CritReturns(result1 error) {
	fake.critMutex.Lock()
	defer fake.critMutex.Unlock()
	fake.CritStub = nil
	fake.critReturns = struct {
		result1 error
	}{result1}
}

func (fake *Syslogger) CritReturnsOnCall(i int, result1 error) {
	fake.critMutex.Lock()
	defer fake.critMutex.Unlock()
	fake.CritStub = nil
	if fake.critReturnsOnCall == nil {
		fake.critReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.critReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Syslogger) Debug(arg1 string) error {
	fake.debugMutex.Lock()
	ret, specificReturn := fake.debugReturnsOnCall[len(fake.debugArgsForCall)]
	fake.debugArgsForCall = append(fake.debugArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Debug", []interface{}{arg1})
	fake.debugMutex.Unlock()
	if fake.DebugStub != nil {
		return fake.DebugStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.debugReturns
	return fakeReturns.result1
}

func (fake *Syslogger) DebugCallCount() int {
	fake.debugMutex.RLock()
	defer fake.debugMutex.RUnlock()
	return len(fake.debugArgsForCall)
}

func (fake *Syslogger) DebugCalls(stub func(string) error) {
	fake.debugMutex.Lock()
	defer fake.debugMutex.Unlock()
	fake.DebugStub = stub
}

func (fake *Syslogger) DebugArgsForCall(i int) string {
	fake.debugMutex.RLock()
	defer fake.debugMutex.RUnlock()
	argsForCall := fake.debugArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Syslogger) DebugReturns(result1 error) {
	fake.debugMutex.Lock()
	defer fake.debugMutex.Unlock()
	fake.DebugStub = nil
	fake.debugReturns = struct {
		result1 error
	}{result1}
}

func (fake *Syslogger) DebugReturnsOnCall(i int, result1 error) {
	fake.debugMutex.Lock()
	defer fake.debugMutex.Unlock()
	fake.DebugStub = nil
	if fake.debugReturnsOnCall == nil {
		fake.debugReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.debugReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Syslogger) Emerg(arg1 string) error {
	fake.emergMutex.Lock()
	ret, specificReturn := fake.emergReturnsOnCall[len(fake.emergArgsForCall)]
	fake.emergArgsForCall = append(fake.emergArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Emerg", []interface{}{arg1})
	fake.emergMutex.Unlock()
	if fake.EmergStub != nil {
		return fake.EmergStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.emergReturns
	return fakeReturns.result1
}

func (fake *Syslogger) EmergCallCount() int {
	fake.emergMutex.RLock()
	defer fake.emergMutex.RUnlock()
	return len(fake.emergArgsForCall)
}

func (fake *Syslogger) EmergCalls(stub func(string) error) {
	fake.emergMutex.Lock()
	defer fake.emergMutex.Unlock()
	fake.EmergStub = stub
}

func (fake *Syslogger) EmergArgsForCall(i int) string {
	fake.emergMutex.RLock()
	defer fake.emergMutex.RUnlock()
	argsForCall := fake.emergArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Syslogger) EmergReturns(result1 error) {
	fake.emergMutex.Lock()
	defer fake.emergMutex.Unlock()
	fake.EmergStub = nil
	fake.emergReturns = struct {
		result1 error
	}{result1}
}

func (fake *Syslogger) EmergReturnsOnCall(i int, result1 error) {
	fake.emergMutex.Lock()
	defer fake.emergMutex.Unlock()
	fake.EmergStub = nil
	if fake.emergReturnsOnCall == nil {
		fake.emergReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.emergReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Syslogger) Err(arg1 string) error {
	fake.errMutex.Lock()
	ret, specificReturn := fake.errReturnsOnCall[len(fake.errArgsForCall)]
	fake.errArgsForCall = append(fake.errArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Err", []interface{}{arg1})
	fake.errMutex.Unlock()
	if fake.ErrStub != nil {
		return fake.ErrStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.errReturns
	return fakeReturns.result1
}

func (fake *Syslogger) ErrCallCount() int {
	fake.errMutex.RLock()
	defer fake.errMutex.RUnlock()
	return len(fake.errArgsForCall)
}

func (fake *Syslogger) ErrCalls(stub func(string) error) {
	fake.errMutex.Lock()
	defer fake.errMutex.Unlock()
	fake.ErrStub = stub
}

func (fake *Syslogger) ErrArgsForCall(i int) string {
	fake.errMutex.RLock()
	defer fake.errMutex.RUnlock()
	argsForCall := fake.errArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Syslogger) ErrReturns(result1 error) {
	fake.errMutex.Lock()
	defer fake.errMutex.Unlock()
	fake.ErrStub = nil
	fake.errReturns = struct {
		result1 error
	}{result1}
}

func (fake *Syslogger) ErrReturnsOnCall(i int, result1 error) {
	fake.errMutex.Lock()
	defer fake.errMutex.Unlock()
	fake.ErrStub = nil
	if fake.errReturnsOnCall == nil {
		fake.errReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.errReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Syslogger) Info(arg1 string) error {
	fake.infoMutex.Lock()
	ret, specificReturn := fake.infoReturnsOnCall[len(fake.infoArgsForCall)]
	fake.infoArgsForCall = append(fake.infoArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Info", []interface{}{arg1})
	fake.infoMutex.Unlock()
	if fake.InfoStub != nil {
		return fake.InfoStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.infoReturns
	return fakeReturns.result1
}

func (fake *Syslogger) InfoCallCount() int {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return len(fake.infoArgsForCall)
}

func (fake *Syslogger) InfoCalls(stub func(string) error) {
	fake.infoMutex.Lock()
	defer fake.infoMutex.Unlock()
	fake.InfoStub = stub
}

func (fake *Syslogger) InfoArgsForCall(i int) string {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	argsForCall := fake.infoArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Syslogger) InfoReturns(result1 error) {
	fake.infoMutex.Lock()
	defer fake.infoMutex.Unlock()
	fake.InfoStub = nil
	fake.infoReturns = struct {
		result1 error
	}{result1}
}

func (fake *Syslogger) InfoReturnsOnCall(i int, result1 error) {
	fake.infoMutex.Lock()
	defer fake.infoMutex.Unlock()
	fake.InfoStub = nil
	if fake.infoReturnsOnCall == nil {
		fake.infoReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.infoReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Syslogger) Notice(arg1 string) error {
	fake.noticeMutex.Lock()
	ret, specificReturn := fake.noticeReturnsOnCall[len(fake.noticeArgsForCall)]
	fake.noticeArgsForCall = append(fake.noticeArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Notice", []interface{}{arg1})
	fake.noticeMutex.Unlock()
	if fake.NoticeStub != nil {
		return fake.NoticeStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.noticeReturns
	return fakeReturns.result1
}

func (fake *Syslogger) NoticeCallCount() int {
	fake.noticeMutex.RLock()
	defer fake.noticeMutex.RUnlock()
	return len(fake.noticeArgsForCall)
}

func (fake *Syslogger) NoticeCalls(stub func(string) error) {
	fake.noticeMutex.Lock()
	defer fake.noticeMutex.Unlock()
	fake.NoticeStub = stub
}

func (fake *Syslogger) NoticeArgsForCall(i int) string {
	fake.noticeMutex.RLock()
	defer fake.noticeMutex.RUnlock()
	argsForCall := fake.noticeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Syslogger) NoticeReturns(result1 error) {
	fake.noticeMutex.Lock()
	defer fake.noticeMutex.Unlock()
	fake.NoticeStub = nil
	fake.noticeReturns = struct {
		result1 error
	}{result1}
}

func (fake *Syslogger) NoticeReturnsOnCall(i int, result1 error) {
	fake.noticeMutex.Lock()
	defer fake.noticeMutex.Unlock()
	fake.NoticeStub = nil
	if fake.noticeReturnsOnCall == nil {
		fake.noticeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.noticeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Syslogger) Warning(arg1 string) error {
	fake.warningMutex.Lock()
	ret, specificReturn := fake.warningReturnsOnCall[len(fake.warningArgsForCall)]
	fake.warningArgsForCall = append(fake.warningArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Warning", []interface{}{arg1})
	fake.warningMutex.Unlock()
	if fake.WarningStub != nil {
		return fake.WarningStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.warningReturns
	return fakeReturns.result1
}

func (fake *Syslogger) WarningCallCount() int {
	fake.warningMutex.RLock()
	defer fake.warningMutex.RUnlock()
	return len(fake.warningArgsForCall)
}

func (fake *Syslogger) WarningCalls(stub func(string) error) {
	fake.warningMutex.Lock()
	defer fake.warningMutex.Unlock()
	fake.WarningStub = stub
}

func (fake *Syslogger) WarningArgsForCall(i int) string {
	fake.warningMutex.RLock()
	defer fake.warningMutex.RUnlock()
	argsForCall := fake.warningArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Syslogger) WarningReturns(result1 error) {
	fake.warningMutex.Lock()
	defer fake.warningMutex.Unlock()
	fake.WarningStub = nil
	fake.warningReturns = struct {
		result1 error
	}{result1}
}

func (fake *Syslogger) WarningReturnsOnCall(i int, result1 error) {
	fake.warningMutex.Lock()
	defer fake.warningMutex.Unlock()
	fake.WarningStub = nil
	if fake.warningReturnsOnCall == nil {
		fake.warningReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.warningReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Syslogger) Write(arg1 []byte) (int, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.writeMutex.Lock()
	ret, specificReturn := fake.writeReturnsOnCall[len(fake.writeArgsForCall)]
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	fake.recordInvocation("Write", []interface{}{arg1Copy})
	fake.writeMutex.Unlock()
	if fake.WriteStub != nil {
		return fake.WriteStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.writeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Syslogger) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *Syslogger) WriteCalls(stub func([]byte) (int, error)) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = stub
}

func (fake *Syslogger) WriteArgsForCall(i int) []byte {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	argsForCall := fake.writeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Syslogger) WriteReturns(result1 int, result2 error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *Syslogger) WriteReturnsOnCall(i int, result1 int, result2 error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = nil
	if fake.writeReturnsOnCall == nil {
		fake.writeReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.writeReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *Syslogger) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.alertMutex.RLock()
	defer fake.alertMutex.RUnlock()
	fake.critMutex.RLock()
	defer fake.critMutex.RUnlock()
	fake.debugMutex.RLock()
	defer fake.debugMutex.RUnlock()
	fake.emergMutex.RLock()
	defer fake.emergMutex.RUnlock()
	fake.errMutex.RLock()
	defer fake.errMutex.RUnlock()
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	fake.noticeMutex.RLock()
	defer fake.noticeMutex.RUnlock()
	fake.warningMutex.RLock()
	defer fake.warningMutex.RUnlock()
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Syslogger) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ syslog.Logger = new(Syslogger)
