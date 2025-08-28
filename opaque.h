
#include <libavcodec/avcodec.h>
#include <libavcodec/packet.h>

void astiav_free_opaque_data(void *opaque, uint8_t *data);
AVBufferRef* astiav_new_opaque(const uint8_t *src, int size);
int astiav_copy_from_opaque(const AVBufferRef *opaque_ref, uint8_t *dst, int dst_size);