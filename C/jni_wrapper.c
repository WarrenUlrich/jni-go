#include <jni.h>
#include "jni_wrapper.h"

int jni_get_version(void* env_ptr)
{
    JNIEnv* env = (JNIEnv*)env_ptr;
    return (*env)->GetVersion(env);
}