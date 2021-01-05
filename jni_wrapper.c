#include "jni_wrapper.h"

int jni_get_created_jvm(JVM* jvm_buf)
{
    return JNI_GetCreatedJavaVMs((JavaVM**)jvm_buf, 1, NULL);
}

int jni_attach_current_thread(JVM jvm, JENV* env_buf)
{
    JavaVM* j = (JavaVM*)jvm;
    return (*j)->AttachCurrentThread(j, env_buf, NULL);
}

int jni_detach_current_thread(JVM jvm)
{
    JavaVM* j = (JavaVM*)jvm;
    return (*j)->DetachCurrentThread(j);
}

int jni_get_env(JVM jvm, JENV* env_buf, int version)
{
    JavaVM* j = (JavaVM*)jvm;
    return (*j)->GetEnv(j, env_buf, version);
}

int jni_attach_current_thread_as_daemon(JVM jvm, JENV* env_buf)
{
    JavaVM* j = (JavaVM*)jvm;
    return (*j)->AttachCurrentThreadAsDaemon(j, env_buf, NULL);
}

int jni_get_version(JENV env)
{
    JNIEnv* e = (JNIEnv*)env;
    return (*e)->GetVersion(e);
}

JCLASS jni_define_class(JENV env, const char* name, JOBJECT loader, void* buf, int len)
{
    JNIEnv* e = (JNIEnv*)env;
    return (*e)->DefineClass(e, name, loader, buf, len);
}

JCLASS jni_find_class(JENV env, const char* name)
{
    JNIEnv* e = (JNIEnv*)env;
    return (*e)->FindClass(e, name);
}

JMETHODID jni_from_reflected_method(JENV env, JOBJECT method)
{
    JNIEnv* e = (JNIEnv*)env;
    return (*e)->FromReflectedMethod(e, method);
}

JFIELDID jni_from_reflected_field(JENV env, JOBJECT field)
{
    JNIEnv* e = (JNIEnv*)env;
    return (*e)->FromReflectedField(e, field);
}

JOBJECT jni_to_reflected_method(JENV env, JCLASS cls, JMETHODID method_id, unsigned char is_static)
{
    JNIEnv* e = (JNIEnv*)env;
    return (*e)->ToReflectedMethod(e, cls, method_id, is_static);
}

JCLASS jni_get_super_class(JENV env, JCLASS sub)
{
    JNIEnv* e = (JNIEnv*)env;
    return (*e)->GetSuperclass(e, sub);
}

unsigned char jni_is_assignable_from(JENV env, JCLASS sub, JCLASS sup)
{
    JNIEnv* e = (JNIEnv*)env;
    return (*e)->IsAssignableFrom(e, sub, sup);
}