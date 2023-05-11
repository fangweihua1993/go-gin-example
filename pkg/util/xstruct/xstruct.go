// @Description struct 转换为map
// @Author shiyibo
// @Copyright 2021 sndks.com. All rights reserved.
// @Datetime 2021/10/25 3:30 下午

package xstruct

import (
	"github.com/mitchellh/mapstructure"
)

var (
	defaultMapStruct = NewDefaultMapStruct()
)

type Option func(m *MapStruct)

type MapStruct struct {
	ErrorUnused      bool
	ZeroFields       bool
	WeaklyTypedInput bool
	Squash           bool
	TagName          string
	Result           interface{}
	DecodeHook       mapstructure.DecodeHookFunc
	Metadata         *mapstructure.Metadata
	MatchName        func(mapKey, fieldName string) bool
}

// WithErrorUnused
// 	@Description
//	@Param errorUnused
// 	@Return Option
func WithErrorUnused(errorUnused bool) Option {
	return func(m *MapStruct) {
		m.ErrorUnused = errorUnused
	}
}

// WithZeroFields
// 	@Description
//	@Param zeroFields
// 	@Return Option
func WithZeroFields(zeroFields bool) Option {
	return func(m *MapStruct) {
		m.ZeroFields = zeroFields
	}
}

// WithWeaklyTypedInput
// 	@Description
//	@Param weaklyTypedInput
// 	@Return Option
func WithWeaklyTypedInput(weaklyTypedInput bool) Option {
	return func(m *MapStruct) {
		m.WeaklyTypedInput = weaklyTypedInput
	}
}

// WithSquash
// 	@Description
//	@Param squash
// 	@Return Option
func WithSquash(squash bool) Option {
	return func(m *MapStruct) {
		m.Squash = squash
	}
}

// WithTagName
// 	@Description
//	@Param tagName
// 	@Return Option
func WithTagName(tagName string) Option {
	return func(m *MapStruct) {
		m.TagName = tagName
	}
}

// NewDefaultMapStruct
// 	@Description
// 	@Return *MapStruct
func NewDefaultMapStruct() *MapStruct {
	return &MapStruct{
		Metadata: nil,
	}
}

// Decode
// 	@Description
//	@Param input
//	@Param output
//	@Param opts decode 可选项
// 	@Return error
func Decode(input interface{}, output interface{}, opts ...Option) error {
	return defaultMapStruct.Decode(input, output, opts...)
}

// Decode
// 	@Description
// 	@Receiver s
//	@Param input
//	@Param output
//	@Param opts
// 	@Return error
func (s *MapStruct) Decode(input interface{}, output interface{}, opts ...Option) error {
	newMapStruct := s.clone()
	for _, f := range opts {
		f(newMapStruct)
	}
	newMapStruct.Result = output
	cfg := &mapstructure.DecoderConfig{
		DecodeHook:       newMapStruct.DecodeHook,
		ErrorUnused:      newMapStruct.ErrorUnused,
		ZeroFields:       newMapStruct.ZeroFields,
		WeaklyTypedInput: newMapStruct.WeaklyTypedInput,
		Squash:           newMapStruct.Squash,
		Metadata:         newMapStruct.Metadata,
		Result:           newMapStruct.Result,
		TagName:          newMapStruct.TagName,
		MatchName:        newMapStruct.MatchName,
	}
	decoder, err := mapstructure.NewDecoder(cfg)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func (s *MapStruct) clone() *MapStruct {
	return &MapStruct{
		DecodeHook:       s.DecodeHook,
		ErrorUnused:      s.ErrorUnused,
		ZeroFields:       s.ZeroFields,
		WeaklyTypedInput: s.WeaklyTypedInput,
		Squash:           s.Squash,
		Metadata:         s.Metadata,
		TagName:          s.TagName,
		MatchName:        s.MatchName,
	}
}
