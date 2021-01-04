#include "jvm_wrapper.h"

int jvm_get_created_vm(void** buf)
{
    return JNI_GetCreatedJavaVMs((JavaVM**)buf, 1, NULL);
}

int jvm_attach_current_thread(void* vm_ptr, void** env)
{
    JavaVM* vm = (JavaVM*)vm_ptr;
    
    return (*vm)->AttachCurrentThread(vm, env, NULL);
}

int jvm_detach_current_thread(void* vm_ptr)
{   
    JavaVM* vm = (JavaVM*)vm_ptr;

    return (*vm)->DetachCurrentThread(vm);
}

int jvm_get_environment(void* vm_ptr, void** env, int version) 
{
    JavaVM* vm = (JavaVM*)vm_ptr;

    return (*vm)->GetEnv(vm, env, version);
}

int jvm_attach_current_thread_as_daemon(void* vm_ptr, void** env)
{
    JavaVM* vm = (JavaVM*)vm_ptr;

    return (*vm)->AttachCurrentThreadAsDaemon(vm, env, NULL);
}