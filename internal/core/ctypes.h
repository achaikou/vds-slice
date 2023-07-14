#ifndef VDS_SLICE_CTYPES_H
#define VDS_SLICE_CTYPES_H
#include <stdlib.h>

struct response {
    char*         data;
    unsigned long size;
};
typedef struct response response;

enum axis_name {
    I         = 0,
    J         = 1,
    K         = 2,
    INLINE    = 3,
    CROSSLINE = 4,
    DEPTH     = 5,
    TIME      = 6,
    SAMPLE    = 7,
};

enum coordinate_system {
    INDEX      = 0,
    ANNOTATION = 1,
    CDP        = 2,
};

enum interpolation_method {
    NEAREST,
    LINEAR,
    CUBIC,
    ANGULAR,
    TRIANGULAR
};

enum attribute {
    VALUE,
    MIN,
    MAX,
    MAXABS,
    MEAN,
    MEANABS,
    MEANPOS,
    MEANNEG,
    MEDIAN,
    RMS,
    VAR,
    SD,
    SUMPOS,
    SUMNEG
};

#endif // VDS_SLICE_CTYPES_H