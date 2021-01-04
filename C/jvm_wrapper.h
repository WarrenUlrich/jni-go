#pragma once
#include <jni.h>

int jvm_get_created_vm(void** buf);

int jvm_attach_current_thread(void* vm_ptr, void** env);

int jvm_detach_current_thread(void* vm_ptr);

int jvm_get_environment(void* vm_ptr, void** env, int version);

int jvm_attach_current_thread_as_daemon(void* vm_ptr, void** env);