
#include <libavcodec/avcodec.h>
#include <libavcodec/packet.h>

void astiav_free_opaque_data(void *opaque, uint8_t *data) {
    (void)opaque;
    av_free(data);
}

AVBufferRef* astiav_new_opaque(const uint8_t *src, int size) {
    uint8_t *p = av_malloc(size);
    if (p == NULL) {
		return NULL;
	}
    memcpy(p, src, size);
    return av_buffer_create(p, size, astiav_free_opaque_data, NULL, 0);
}

int astiav_copy_from_opaque(const AVBufferRef *opaque_ref, uint8_t *dst, int dst_size) {
    int n = opaque_ref->size;
    if (n > dst_size) {
		n = dst_size;
	}
    memcpy(dst, opaque_ref->data, n);
    return n;
}
