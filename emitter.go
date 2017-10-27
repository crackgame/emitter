package emitter

import (
	"fmt"
	"reflect"
)

type EmitFunc func(args ...interface{})

type SeqFunc struct {
	seq      int
	listener reflect.Value
}

type SeqFuncMap map[int]*SeqFunc

func (s SeqFuncMap) clear() {
	s = map[int]*SeqFunc{}
}

func (s SeqFuncMap) remove(seq int) {
	delete(s, seq)
}

type EmitOnResult struct {
	event   interface{}
	seqFunc *SeqFunc
	emitter *Emitter
}

func (r *EmitOnResult) Off() {
	if r.emitter != nil && r.event != nil {
		r.emitter.OffFunc(r.event, r.seqFunc)
		r.emitter = nil
	}
}

type Emitter struct {
	funcs map[interface{}]SeqFuncMap
	seq   int
}

func NewEmitter() *Emitter {
	rv := &Emitter{
		funcs: map[interface{}]SeqFuncMap{},
		seq:   0,
	}
	return rv
}

func (e *Emitter) Emit(event interface{}, args ...interface{}) {
	funcMap, ok := e.funcs[event]
	if !ok {
		return
	}

	if len(funcMap) == 0 {
		return
	}

	// 收集函数seq
	var fnSeqs []int
	for _, v := range funcMap {
		fnSeqs = append(fnSeqs, v.seq)
	}

	// 通过seq来遍历处理
	for _, seq := range fnSeqs {
		seqFn, ok := funcMap[seq]
		if !ok {
			fmt.Println("Emitter::emit func is removed in emit!!!!")
			continue
		}

		fn := seqFn.listener

		var values []reflect.Value
		for i := 0; i < len(args); i++ {
			if args[i] == nil {
				values = append(values, reflect.New(fn.Type().In(i)).Elem())
			} else {
				values = append(values, reflect.ValueOf(args[i]))
			}
		}

		// call callback func
		fn.Call(values)
	}
}

func (e *Emitter) On(event interface{}, listener interface{}) *EmitOnResult {
	e.seq++
	seqFunc := &SeqFunc{
		seq:      e.seq,
		listener: reflect.ValueOf(listener),
	}

	if e.funcs[event] == nil {
		e.funcs[event] = SeqFuncMap{}
	}
	e.funcs[event][seqFunc.seq] = seqFunc

	return &EmitOnResult{
		event:   event,
		seqFunc: seqFunc,
		emitter: e,
	}
}

func (e *Emitter) OffFunc(event interface{}, seqFunc *SeqFunc) {
	fnMap := e.funcs[event]
	fnMap.remove(seqFunc.seq)
}

func (e *Emitter) Off(event interface{}) {
	delete(e.funcs, event)
}

func (e *Emitter) OffAll() {
	e.funcs = map[interface{}]SeqFuncMap{}
}
