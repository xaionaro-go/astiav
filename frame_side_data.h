#include <libavutil/frame.h>
#include <libavutil/motion_vector.h>

AVRegionOfInterest* astiavConvertRegionsOfInterestFrameSideData(AVFrameSideData *sd);
AVMotionVector* astiavConvertMotionVectorsFrameSideData(AVFrameSideData *sd);