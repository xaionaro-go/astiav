#include <libavutil/frame.h>
#include <libavutil/motion_vector.h>

AVRegionOfInterest* astiavConvertRegionsOfInterestFrameSideData(AVFrameSideData *sd) {
    return (AVRegionOfInterest*)sd->data;
}

AVMotionVector* astiavConvertMotionVectorsFrameSideData(AVFrameSideData *sd) {
    return (AVMotionVector*)sd->data;
}