//go:build ffmpeg7
// +build ffmpeg7

package astiav

//#include <libavfilter/avfilter.h>
import "C"

// https://ffmpeg.org/doxygen/7.0/structAVFilter.html#a04e408702054370fbe35c8d5b49f68cb
func (f *Filter) NbInputs() int {
	return int(f.c.nb_inputs)
}

// https://ffmpeg.org/doxygen/7.0/structAVFilter.html#abb166cb2c9349de54d24aefb879608f4
func (f *Filter) NbOutputs() int {
	return int(f.c.nb_outputs)
}

// https://ffmpeg.org/doxygen/7.0/structAVFilter.html#ad311151fe6e8c87a89f895bef7c8b98b
func (f *Filter) Inputs() (ps []*FilterPad) {
	for idx := 0; idx < f.NbInputs(); idx++ {
		ps = append(ps, newFilterPad(MediaType(C.avfilter_pad_get_type(f.c.inputs, C.int(idx)))))
	}
	return
}

// https://ffmpeg.org/doxygen/7.0/structAVFilter.html#ad0608786fa3e1ca6e4cc4b67039f77d7
func (f *Filter) Outputs() (ps []*FilterPad) {
	for idx := 0; idx < f.NbOutputs(); idx++ {
		ps = append(ps, newFilterPad(MediaType(C.avfilter_pad_get_type(f.c.outputs, C.int(idx)))))
	}
	return
}
