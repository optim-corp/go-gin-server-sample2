
package openapi

import (
	"math"
	"reflect"
	"sort"
)

type SampleResponseStream []SampleResponse

func SampleResponseStreamOf(arg ...SampleResponse) SampleResponseStream {
	return arg
}
func SampleResponseStreamFrom(arg []SampleResponse) SampleResponseStream {
	return arg
}

func (self *SampleResponseStream) Add(arg SampleResponse) *SampleResponseStream {
	return self.AddAll(arg)

}

func (self *SampleResponseStream) AddAll(arg ...SampleResponse) *SampleResponseStream {
	*self = append(*self, arg...)
	return self
}

func (self *SampleResponseStream) AddSafe(arg *SampleResponse) *SampleResponseStream {
    if arg != nil {
        self.Add(*arg)
    }
	return self

}
func (self *SampleResponseStream) AllMatch(fn func(SampleResponse, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}

func (self *SampleResponseStream) AnyMatch(fn func(SampleResponse, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *SampleResponseStream) Clone() *SampleResponseStream {
	temp := make([]SampleResponse, self.Len())
	copy(temp, *self)
	return (*SampleResponseStream)(&temp)
}

func (self *SampleResponseStream) Copy() *SampleResponseStream {
	return self.Clone()
}

func (self *SampleResponseStream) Concat(arg []SampleResponse) *SampleResponseStream {
	return self.AddAll(arg...)
}

func (self *SampleResponseStream) Contains(arg SampleResponse) bool {
	return self.FindIndex(func(_arg SampleResponse, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}

func (self *SampleResponseStream) Delete(index int) *SampleResponseStream {
	return self.DeleteRange(index, index)
}

func (self *SampleResponseStream) DeleteRange(startIndex, endIndex int) *SampleResponseStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}

func (self *SampleResponseStream) Distinct() *SampleResponseStream {
	m := map[SampleResponse]bool{}
	self.Filter(func(arg SampleResponse, index int) bool {
		_, ok := m[arg]
		if !ok {
			m[arg] = true
		}
		return !ok
	})
	return self
}

func (self *SampleResponseStream) Equals(arr []SampleResponse) bool {
	if (*self == nil) != (arr == nil) || len(*self) != len(arr) {
		return false
	}
	for i := range *self {
		if !reflect.DeepEqual((*self)[i], arr[i]) {
			return false
		}
	}
	return true
}
func (self *SampleResponseStream) Filter(fn func(SampleResponse, int) bool) *SampleResponseStream {
	_array := SampleResponseStreamOf()
	self.ForEach(func(v SampleResponse, i int) {
		if fn(v, i) {
			_array.Add(v)
		}
	})
	*self = _array
	return self
}

func (self *SampleResponseStream) Find(fn func(SampleResponse, int) bool) *SampleResponse {
	i := self.FindIndex(fn)
	if -1 != i {
		return &(*self)[i]
	}
	return nil
}

func (self *SampleResponseStream) FindIndex(fn func(SampleResponse, int) bool) int {
	for i, v := range self.Val() {
		if fn(v, i) {
			return i
		}
	}
	return -1
}

func (self *SampleResponseStream) First() *SampleResponse {
	return self.Get(0)
}

func (self *SampleResponseStream) ForEach(fn func(SampleResponse, int)) *SampleResponseStream {
	for i, v := range self.Val() {
		fn(v, i)
	}
	return self
}
func (self *SampleResponseStream) ForEachRight(fn func(SampleResponse, int)) *SampleResponseStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *SampleResponseStream) GroupBy(fn func(SampleResponse, int) string) map[string][]SampleResponse {
    m := map[string][]SampleResponse{}
    for i, v := range self.Val() {
        key := fn(v, i)
        m[key] = append(m[key], v)
    }
    return m
}
func (self *SampleResponseStream) GroupByValues(fn func(SampleResponse, int) string) [][]SampleResponse {
	tmp := [][]SampleResponse{}
	m := self.GroupBy(fn)
	for _, v := range m {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *SampleResponseStream) IndexOf(arg SampleResponse) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *SampleResponseStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *SampleResponseStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *SampleResponseStream) Last() *SampleResponse {
	return self.Get(self.Len() - 1)
}

func (self *SampleResponseStream) Len() int {
    if self == nil {
		return 0
	}
	return len(*self)
}
func (self *SampleResponseStream) Limit(limit int) *SampleResponseStream {
	self.Slice(0, limit)
	return self
}
func (self *SampleResponseStream) Map(fn func(SampleResponse, int) SampleResponse) *SampleResponseStream {
	return self.ForEach(func(v SampleResponse, i int) { self.Set(i, fn(v, i)) })
}

func (self *SampleResponseStream) MapAny(fn func(SampleResponse, int) interface{}) []interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *SampleResponseStream) Map2Int(fn func(SampleResponse, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *SampleResponseStream) Map2Int32(fn func(SampleResponse, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *SampleResponseStream) Map2Int64(fn func(SampleResponse, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *SampleResponseStream) Map2Float32(fn func(SampleResponse, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *SampleResponseStream) Map2Float64(fn func(SampleResponse, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *SampleResponseStream) Map2Bool(fn func(SampleResponse, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *SampleResponseStream) Map2Bytes(fn func(SampleResponse, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *SampleResponseStream) Map2String(fn func(SampleResponse, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SampleResponseStream) Max(fn func(SampleResponse, int) float64) *SampleResponse {
	f := self.Get(0)
	if f == nil {
		return nil
	}
	m := fn(*f, 0)
	index := 0
	for i := 1; i < self.Len(); i++ {
		v := fn(*self.Get(i), i)
		m = math.Max(m, v)
		if m == v {
			index = i
		}
	}
	return self.Get(index)
}
func (self *SampleResponseStream) Min(fn func(SampleResponse, int) float64) *SampleResponse {
	f := self.Get(0)
	if f == nil {
		return nil
	}
	m := fn(*f, 0)
	index := 0
	for i := 1; i < self.Len(); i++ {
		v := fn(*self.Get(i), i)
		m = math.Min(m, v)
		if m == v {
			index = i
		}
	}
	return self.Get(index)
}

func (self *SampleResponseStream) NoneMatch(fn func(SampleResponse, int) bool) bool {
	return !self.AnyMatch(fn)
}

func (self *SampleResponseStream) Get(index int) *SampleResponse {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
        return &tmp
	}
	return nil
}
func (self *SampleResponseStream) Peek(fn func(*SampleResponse, int)) *SampleResponseStream {
    for i, v := range *self {
        fn(&v, i)
        self.Set(i, v)
    }
    return self
}
func (self *SampleResponseStream) Reduce(fn func(SampleResponse, SampleResponse, int) SampleResponse) *SampleResponseStream {
	return self.ReduceInit(fn, SampleResponse{})
}
func (self *SampleResponseStream) ReduceInit(fn func(SampleResponse, SampleResponse, int) SampleResponse, initialValue SampleResponse) *SampleResponseStream {
	result :=SampleResponseStreamOf()
	self.ForEach(func(v SampleResponse, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}

func (self *SampleResponseStream) ReduceInterface(fn func(interface{}, SampleResponse, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(SampleResponse{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *SampleResponseStream) ReduceString(fn func(string, SampleResponse, int) string) []string {
	result := []string{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn("", v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *SampleResponseStream) ReduceInt(fn func(int, SampleResponse, int) int) []int {
	result := []int{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *SampleResponseStream) ReduceInt32(fn func(int32, SampleResponse, int) int32) []int32 {
	result := []int32{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *SampleResponseStream) ReduceInt64(fn func(int64, SampleResponse, int) int64) []int64 {
	result := []int64{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *SampleResponseStream) ReduceFloat32(fn func(float32, SampleResponse, int) float32) []float32 {
	result := []float32{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0.0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *SampleResponseStream) ReduceFloat64(fn func(float64, SampleResponse, int) float64) []float64 {
	result := []float64{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0.0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *SampleResponseStream) ReduceBool(fn func(bool, SampleResponse, int) bool) []bool {
	result := []bool{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(false, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}

func (self *SampleResponseStream) Reverse() *SampleResponseStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}

func (self *SampleResponseStream) Replace(fn func(SampleResponse, int) SampleResponse) *SampleResponseStream {
	return self.Map(fn)
}

func (self *SampleResponseStream) Set(index int, val SampleResponse) *SampleResponseStream {
    if len(*self) > index {
        (*self)[index] = val
    }
    return self
}

func (self *SampleResponseStream) Skip(skip int) *SampleResponseStream {
	self.Slice(skip, self.Len()-skip)
	return self
}

func (self *SampleResponseStream) SkippingEach(fn func(SampleResponse, int) int) *SampleResponseStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}

func (self *SampleResponseStream) Slice(startIndex, n int) *SampleResponseStream {
    last := startIndex+n
    if len(*self)-1 < startIndex {
        *self = []SampleResponse{}
    } else if len(*self) < last {
        *self = (*self)[startIndex:len(*self)]
    } else {
        *self = (*self)[startIndex:last]
    }
	return self
}

func (self *SampleResponseStream) Sort(fn func(i, j int) bool) *SampleResponseStream {
	sort.Slice(*self, fn)
	return self
}

func (self *SampleResponseStream) SortStable(fn func(i, j int) bool) *SampleResponseStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *SampleResponseStream) ToList() []SampleResponse {
	return self.Val()
}

func (self *SampleResponseStream) Val() []SampleResponse {
	if self == nil {
		return []SampleResponse{}
	}
	return *self.Copy()
}


func (self *SampleResponseStream) While(fn func(SampleResponse, int) bool) *SampleResponseStream {
    for i, v := range self.Val() {
        if !fn(v, i) {
            break
        }
    }
    return self
}
