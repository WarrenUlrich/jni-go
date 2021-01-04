#pragma once
#include <jni.h>

//For ease
#define JENV void*
#define JOBJECT void*
#define JCLASS void*
#define JTHROWABLE void*
#define JMETHODID void*
#define JFIELDID void*

int jni_get_version(JENV env);

JCLASS jni_define_class(JENV env, const char* name, JOBJECT loader, void* buf);

JCLASS jni_find_class(JENV env, const char* name);

JMETHODID jni_from_reflected_method(JENV env, JOBJECT method);

JFIELDID jni_from_reflected_field(JENV env, JOBJECT field);

JCLASS jni_get_super_class(JENV env, JCLASS sub);

unsigned char jni_is_assignable_from(JENV env_ptr, JCLASS sub, JCLASS sup);

JOBJECT jni_to_reflected_field(JENV env, JCLASS cls, JFIELDID field_id, unsigned char is_static);

int jni_throw(JENV env, JTHROWABLE throwable);

int jni_throw_new(JENV env, JCLASS cls, const char* msg);

JTHROWABLE jni_exception_occured(JENV env);

void jni_exception_describe(JENV env);

void jni_exception_clear(JENV env);

void jni_fatal_error(JENV env, const char* msg);

int push_local_frame(JENV env, int capacity);

JOBJECT jni_pop_local_frame(JENV env, JOBJECT result);