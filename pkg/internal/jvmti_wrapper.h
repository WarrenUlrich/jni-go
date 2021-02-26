#pragma once
#include <jvmti.h>

// /*   2 : Set Event Notification Mode */
// jvmtiError jvmti_set_event_notification_mode(jvmtiEnv *env,
//                                     jvmtiEventMode mode,
//                                     jvmtiEvent event_type,
//                                     jthread event_thread,
//                                     ...);

// /*   4 : Get All Threads */
// jvmtiError jvmti_get_all_threads(jvmtiEnv *env,
//                          jint *threads_count_ptr,
//                          jthread **threads_ptr);

// /*   5 : Suspend Thread */
// jvmtiError suspend_thread(jvmtiEnv *env,
//                          jthread thread);

// /*   6 : Resume Thread */
// jvmtiError resume_thread(jvmtiEnv *env,
//                         jthread thread);

// /*   7 : Stop Thread */
// jvmtiError stop_thread(jvmtiEnv *env,
//                       jthread thread,
//                       jobject exception);

// /*   8 : Interrupt Thread */
// jvmtiError interrupt_thread(jvmtiEnv *env,
//                            jthread thread);

// /*   9 : Get Thread Info */
// jvmtiError GetThreadInfo(jvmtiEnv *env,
//                          jthread thread,
//                          jvmtiThreadInfo *info_ptr);

// /*   10 : Get Owned Monitor Info */
// jvmtiError GetOwnedMonitorInfo(jvmtiEnv *env,
//                                jthread thread,
//                                jint *owned_monitor_count_ptr,
//                                jobject **owned_monitors_ptr);

// /*   11 : Get Current Contended Monitor */
// jvmtiError GetCurrentContendedMonitor(jvmtiEnv *env,
//                                       jthread thread,
//                                       jobject *monitor_ptr);

// /*   12 : Run Agent Thread */
// jvmtiError RunAgentThread(jvmtiEnv *env,
//                           jthread thread,
//                           jvmtiStartFunction proc,
//                           const void *arg,
//                           jint priority);

// /*   13 : Get Top Thread Groups */
// jvmtiError GetTopThreadGroups(jvmtiEnv *env,
//                               jint *group_count_ptr,
//                               jthreadGroup **groups_ptr);

// /*   14 : Get Thread Group Info */
// jvmtiError GetThreadGroupInfo(jvmtiEnv *env,
//                               jthreadGroup group,
//                               jvmtiThreadGroupInfo *info_ptr);

// /*   15 : Get Thread Group Children */
// jvmtiError GetThreadGroupChildren(jvmtiEnv *env,
//                                   jthreadGroup group,
//                                   jint *thread_count_ptr,
//                                   jthread **threads_ptr,
//                                   jint *group_count_ptr,
//                                   jthreadGroup **groups_ptr);

// /*   16 : Get Frame Count */
// jvmtiError GetFrameCount(jvmtiEnv *env,
//                          jthread thread,
//                          jint *count_ptr);

// /*   17 : Get Thread State */
// jvmtiError GetThreadState(jvmtiEnv *env,
//                           jthread thread,
//                           jint *thread_state_ptr);

// /*   18 : Get Current Thread */
// jvmtiError GetCurrentThread(jvmtiEnv *env,
//                             jthread *thread_ptr);

// /*   19 : Get Frame Location */
// jvmtiError GetFrameLocation(jvmtiEnv *env,
//                             jthread thread,
//                             jint depth,
//                             jmethodID *method_ptr,
//                             jlocation *location_ptr);

// /*   20 : Notify Frame Pop */
// jvmtiError NotifyFramePop(jvmtiEnv *env,
//                           jthread thread,
//                           jint depth);

// /*   21 : Get Local Variable - Object */
// jvmtiError GetLocalObject(jvmtiEnv *env,
//                           jthread thread,
//                           jint depth,
//                           jint slot,
//                           jobject *value_ptr);

// /*   22 : Get Local Variable - Int */
// jvmtiError GetLocalInt(jvmtiEnv *env,
//                        jthread thread,
//                        jint depth,
//                        jint slot,
//                        jint *value_ptr);

// /*   23 : Get Local Variable - Long */
// jvmtiError GetLocalLong(jvmtiEnv *env,
//                         jthread thread,
//                         jint depth,
//                         jint slot,
//                         jlong *value_ptr);

// /*   24 : Get Local Variable - Float */
// jvmtiError GetLocalFloat(jvmtiEnv *env,
//                          jthread thread,
//                          jint depth,
//                          jint slot,
//                          jfloat *value_ptr);

// /*   25 : Get Local Variable - Double */
// jvmtiError GetLocalDouble(jvmtiEnv *env,
//                           jthread thread,
//                           jint depth,
//                           jint slot,
//                           jdouble *value_ptr);

// /*   26 : Set Local Variable - Object */
// jvmtiError SetLocalObject(jvmtiEnv *env,
//                           jthread thread,
//                           jint depth,
//                           jint slot,
//                           jobject value);

// /*   27 : Set Local Variable - Int */
// jvmtiError SetLocalInt(jvmtiEnv *env,
//                        jthread thread,
//                        jint depth,
//                        jint slot,
//                        jint value);

// /*   28 : Set Local Variable - Long */
// jvmtiError SetLocalLong(jvmtiEnv *env,
//                         jthread thread,
//                         jint depth,
//                         jint slot,
//                         jlong value);

// /*   29 : Set Local Variable - Float */
// jvmtiError SetLocalFloat(jvmtiEnv *env,
//                          jthread thread,
//                          jint depth,
//                          jint slot,
//                          jfloat value);

// /*   30 : Set Local Variable - Double */
// jvmtiError SetLocalDouble(jvmtiEnv *env,
//                           jthread thread,
//                           jint depth,
//                           jint slot,
//                           jdouble value);

// /*   31 : Create Raw Monitor */
// jvmtiError CreateRawMonitor(jvmtiEnv *env,
//                             const char *name,
//                             jrawMonitorID *monitor_ptr);

// /*   32 : Destroy Raw Monitor */
// jvmtiError DestroyRawMonitor(jvmtiEnv *env,
//                              jrawMonitorID monitor);

// /*   33 : Raw Monitor Enter */
// jvmtiError RawMonitorEnter(jvmtiEnv *env,
//                            jrawMonitorID monitor);

// /*   34 : Raw Monitor Exit */
// jvmtiError RawMonitorExit(jvmtiEnv *env,
//                           jrawMonitorID monitor);

// /*   35 : Raw Monitor Wait */
// jvmtiError RawMonitorWait(jvmtiEnv *env,
//                           jrawMonitorID monitor,
//                           jlong millis);

// /*   36 : Raw Monitor Notify */
// jvmtiError RawMonitorNotify(jvmtiEnv *env,
//                             jrawMonitorID monitor);

// /*   37 : Raw Monitor Notify All */
// jvmtiError RawMonitorNotifyAll(jvmtiEnv *env,
//                                jrawMonitorID monitor);

// /*   38 : Set Breakpoint */
// jvmtiError SetBreakpoint(jvmtiEnv *env,
//                          jmethodID method,
//                          jlocation location);

// /*   39 : Clear Breakpoint */
// jvmtiError ClearBreakpoint(jvmtiEnv *env,
//                            jmethodID method,
//                            jlocation location);

// /*   41 : Set Field Access Watch */
// jvmtiError SetFieldAccessWatch(jvmtiEnv *env,
//                                jclass klass,
//                                jfieldID field);

// /*   42 : Clear Field Access Watch */
// jvmtiError ClearFieldAccessWatch(jvmtiEnv *env,
//                                  jclass klass,
//                                  jfieldID field);

// /*   43 : Set Field Modification Watch */
// jvmtiError SetFieldModificationWatch(jvmtiEnv *env,
//                                      jclass klass,
//                                      jfieldID field);

// /*   44 : Clear Field Modification Watch */
// jvmtiError ClearFieldModificationWatch(jvmtiEnv *env,
//                                        jclass klass,
//                                        jfieldID field);

// /*   45 : Is Modifiable Class */
// jvmtiError IsModifiableClass(jvmtiEnv *env,
//                              jclass klass,
//                              jboolean *is_modifiable_class_ptr);

// /*   46 : Allocate */
// jvmtiError Allocate(jvmtiEnv *env,
//                     jlong size,
//                     unsigned char **mem_ptr);

/*   47 : Deallocate */
jvmtiError jvmti_deallocate(jvmtiEnv *env,
                      void* mem);

/*   48 : Get Class Signature */
jvmtiError jvmti_get_class_signature(jvmtiEnv *env,
                             jclass klass,
                             char **signature_ptr,
                             char **generic_ptr);

// /*   49 : Get Class Status */
// jvmtiError GetClassStatus(jvmtiEnv *env,
//                           jclass klass,
//                           jint *status_ptr);

// /*   50 : Get Source File Name */
// jvmtiError GetSourceFileName(jvmtiEnv *env,
//                              jclass klass,
//                              char **source_name_ptr);

// /*   51 : Get Class Modifiers */
// jvmtiError GetClassModifiers(jvmtiEnv *env,
//                              jclass klass,
//                              jint *modifiers_ptr);

// /*   52 : Get Class Methods */
// jvmtiError GetClassMethods(jvmtiEnv *env,
//                            jclass klass,
//                            jint *method_count_ptr,
//                            jmethodID **methods_ptr);

// /*   53 : Get Class Fields */
// jvmtiError GetClassFields(jvmtiEnv *env,
//                           jclass klass,
//                           jint *field_count_ptr,
//                           jfieldID **fields_ptr);

// /*   54 : Get Implemented Interfaces */
// jvmtiError GetImplementedInterfaces(jvmtiEnv *env,
//                                     jclass klass,
//                                     jint *interface_count_ptr,
//                                     jclass **interfaces_ptr);

// /*   55 : Is Interface */
// jvmtiError IsInterface(jvmtiEnv *env,
//                        jclass klass,
//                        jboolean *is_interface_ptr);

// /*   56 : Is Array Class */
// jvmtiError IsArrayClass(jvmtiEnv *env,
//                         jclass klass,
//                         jboolean *is_array_class_ptr);

// /*   57 : Get Class Loader */
// jvmtiError GetClassLoader(jvmtiEnv *env,
//                           jclass klass,
//                           jobject *classloader_ptr);

// /*   58 : Get Object Hash Code */
// jvmtiError GetObjectHashCode(jvmtiEnv *env,
//                              jobject object,
//                              jint *hash_code_ptr);

// /*   59 : Get Object Monitor Usage */
// jvmtiError GetObjectMonitorUsage(jvmtiEnv *env,
//                                  jobject object,
//                                  jvmtiMonitorUsage *info_ptr);

// /*   60 : Get Field Name (and Signature) */
// jvmtiError GetFieldName(jvmtiEnv *env,
//                         jclass klass,
//                         jfieldID field,
//                         char **name_ptr,
//                         char **signature_ptr,
//                         char **generic_ptr);

// /*   61 : Get Field Declaring Class */
// jvmtiError GetFieldDeclaringClass(jvmtiEnv *env,
//                                   jclass klass,
//                                   jfieldID field,
//                                   jclass *declaring_class_ptr);

// /*   62 : Get Field Modifiers */
// jvmtiError GetFieldModifiers(jvmtiEnv *env,
//                              jclass klass,
//                              jfieldID field,
//                              jint *modifiers_ptr);

// /*   63 : Is Field Synthetic */
// jvmtiError IsFieldSynthetic(jvmtiEnv *env,
//                             jclass klass,
//                             jfieldID field,
//                             jboolean *is_synthetic_ptr);

// /*   64 : Get Method Name (and Signature) */
// jvmtiError GetMethodName(jvmtiEnv *env,
//                          jmethodID method,
//                          char **name_ptr,
//                          char **signature_ptr,
//                          char **generic_ptr);

// /*   65 : Get Method Declaring Class */
// jvmtiError GetMethodDeclaringClass(jvmtiEnv *env,
//                                    jmethodID method,
//                                    jclass *declaring_class_ptr);

// /*   66 : Get Method Modifiers */
// jvmtiError GetMethodModifiers(jvmtiEnv *env,
//                               jmethodID method,
//                               jint *modifiers_ptr);

// /*   68 : Get Max Locals */
// jvmtiError GetMaxLocals(jvmtiEnv *env,
//                         jmethodID method,
//                         jint *max_ptr);

// /*   69 : Get Arguments Size */
// jvmtiError GetArgumentsSize(jvmtiEnv *env,
//                             jmethodID method,
//                             jint *size_ptr);

// /*   70 : Get Line Number Table */
// jvmtiError GetLineNumberTable(jvmtiEnv *env,
//                               jmethodID method,
//                               jint *entry_count_ptr,
//                               jvmtiLineNumberEntry **table_ptr);

// /*   71 : Get Method Location */
// jvmtiError GetMethodLocation(jvmtiEnv *env,
//                              jmethodID method,
//                              jlocation *start_location_ptr,
//                              jlocation *end_location_ptr);

// /*   72 : Get Local Variable Table */
// jvmtiError GetLocalVariableTable(jvmtiEnv *env,
//                                  jmethodID method,
//                                  jint *entry_count_ptr,
//                                  jvmtiLocalVariableEntry **table_ptr);

// /*   73 : Set Native Method Prefix */
// jvmtiError SetNativeMethodPrefix(jvmtiEnv *env,
//                                  const char *prefix);

// /*   74 : Set Native Method Prefixes */
// jvmtiError SetNativeMethodPrefixes(jvmtiEnv *env,
//                                    jint prefix_count,
//                                    char **prefixes);

// /*   75 : Get Bytecodes */
// jvmtiError GetBytecodes(jvmtiEnv *env,
//                         jmethodID method,
//                         jint *bytecode_count_ptr,
//                         unsigned char **bytecodes_ptr);

// /*   76 : Is Method Native */
// jvmtiError IsMethodNative(jvmtiEnv *env,
//                           jmethodID method,
//                           jboolean *is_native_ptr);

// /*   77 : Is Method Synthetic */
// jvmtiError IsMethodSynthetic(jvmtiEnv *env,
//                              jmethodID method,
//                              jboolean *is_synthetic_ptr);

/*   78 : Get Loaded Classes */
jvmtiError jvmti_get_loaded_classes(jvmtiEnv *env,
                            jint *class_count_ptr,
                            jclass **classes_ptr);

// /*   79 : Get Classloader Classes */
// jvmtiError GetClassLoaderClasses(jvmtiEnv *env,
//                                  jobject initiating_loader,
//                                  jint *class_count_ptr,
//                                  jclass **classes_ptr);

// /*   80 : Pop Frame */
// jvmtiError PopFrame(jvmtiEnv *env,
//                     jthread thread);

// /*   81 : Force Early Return - Object */
// jvmtiError ForceEarlyReturnObject(jvmtiEnv *env,
//                                   jthread thread,
//                                   jobject value);

// /*   82 : Force Early Return - Int */
// jvmtiError ForceEarlyReturnInt(jvmtiEnv *env,
//                                jthread thread,
//                                jint value);

// /*   83 : Force Early Return - Long */
// jvmtiError ForceEarlyReturnLong(jvmtiEnv *env,
//                                 jthread thread,
//                                 jlong value);

// /*   84 : Force Early Return - Float */
// jvmtiError ForceEarlyReturnFloat(jvmtiEnv *env,
//                                  jthread thread,
//                                  jfloat value);

// /*   85 : Force Early Return - Double */
// jvmtiError ForceEarlyReturnDouble(jvmtiEnv *env,
//                                   jthread thread,
//                                   jdouble value);

// /*   86 : Force Early Return - Void */
// jvmtiError ForceEarlyReturnVoid(jvmtiEnv *env,
//                                 jthread thread);

// /*   87 : Redefine Classes */
// jvmtiError RedefineClasses(jvmtiEnv *env,
//                            jint class_count,
//                            const jvmtiClassDefinition *class_definitions);

// /*   88 : Get Version Number */
// jvmtiError GetVersionNumber(jvmtiEnv *env,
//                             jint *version_ptr);

// /*   89 : Get Capabilities */
// jvmtiError GetCapabilities(jvmtiEnv *env,
//                            jvmtiCapabilities *capabilities_ptr);

// /*   90 : Get Source Debug Extension */
// jvmtiError GetSourceDebugExtension(jvmtiEnv *env,
//                                    jclass klass,
//                                    char **source_debug_extension_ptr);

// /*   91 : Is Method Obsolete */
// jvmtiError IsMethodObsolete(jvmtiEnv *env,
//                             jmethodID method,
//                             jboolean *is_obsolete_ptr);

// /*   92 : Suspend Thread List */
// jvmtiError SuspendThreadList(jvmtiEnv *env,
//                              jint request_count,
//                              const jthread *request_list,
//                              jvmtiError *results);

// /*   93 : Resume Thread List */
// jvmtiError ResumeThreadList(jvmtiEnv *env,
//                             jint request_count,
//                             const jthread *request_list,
//                             jvmtiError *results);

// /*   100 : Get All Stack Traces */
// jvmtiError GetAllStackTraces(jvmtiEnv *env,
//                              jint max_frame_count,
//                              jvmtiStackInfo **stack_info_ptr,
//                              jint *thread_count_ptr);

// /*   101 : Get Thread List Stack Traces */
// jvmtiError GetThreadListStackTraces(jvmtiEnv *env,
//                                     jint thread_count,
//                                     const jthread *thread_list,
//                                     jint max_frame_count,
//                                     jvmtiStackInfo **stack_info_ptr);

// /*   102 : Get Thread Local Storage */
// jvmtiError GetThreadLocalStorage(jvmtiEnv *env,
//                                  jthread thread,
//                                  void **data_ptr);

// /*   103 : Set Thread Local Storage */
// jvmtiError SetThreadLocalStorage(jvmtiEnv *env,
//                                  jthread thread,
//                                  const void *data);

// /*   104 : Get Stack Trace */
// jvmtiError GetStackTrace(jvmtiEnv *env,
//                          jthread thread,
//                          jint start_depth,
//                          jint max_frame_count,
//                          jvmtiFrameInfo *frame_buffer,
//                          jint *count_ptr);

// /*   106 : Get Tag */
// jvmtiError GetTag(jvmtiEnv *env,
//                   jobject object,
//                   jlong *tag_ptr);

// /*   107 : Set Tag */
// jvmtiError SetTag(jvmtiEnv *env,
//                   jobject object,
//                   jlong tag);

// /*   108 : Force Garbage Collection */
// jvmtiError ForceGarbageCollection(jvmtiEnv *env);

// /*   109 : Iterate Over Objects Reachable From Object */
// jvmtiError IterateOverObjectsReachableFromObject(jvmtiEnv *env,
//                                                  jobject object,
//                                                  jvmtiObjectReferenceCallback object_reference_callback,
//                                                  const void *user_data);

// /*   110 : Iterate Over Reachable Objects */
// jvmtiError IterateOverReachableObjects(jvmtiEnv *env,
//                                        jvmtiHeapRootCallback heap_root_callback,
//                                        jvmtiStackReferenceCallback stack_ref_callback,
//                                        jvmtiObjectReferenceCallback object_ref_callback,
//                                        const void *user_data);

// /*   111 : Iterate Over Heap */
// jvmtiError IterateOverHeap(jvmtiEnv *env,
//                            jvmtiHeapObjectFilter object_filter,
//                            jvmtiHeapObjectCallback heap_object_callback,
//                            const void *user_data);

// /*   112 : Iterate Over Instances Of Class */
// jvmtiError IterateOverInstancesOfClass(jvmtiEnv *env,
//                                        jclass klass,
//                                        jvmtiHeapObjectFilter object_filter,
//                                        jvmtiHeapObjectCallback heap_object_callback,
//                                        const void *user_data);

// /*   114 : Get Objects With Tags */
// jvmtiError GetObjectsWithTags(jvmtiEnv *env,
//                               jint tag_count,
//                               const jlong *tags,
//                               jint *count_ptr,
//                               jobject **object_result_ptr,
//                               jlong **tag_result_ptr);

// /*   115 : Follow References */
// jvmtiError FollowReferences(jvmtiEnv *env,
//                             jint heap_filter,
//                             jclass klass,
//                             jobject initial_object,
//                             const jvmtiHeapCallbacks *callbacks,
//                             const void *user_data);

// /*   116 : Iterate Through Heap */
// jvmtiError IterateThroughHeap(jvmtiEnv *env,
//                               jint heap_filter,
//                               jclass klass,
//                               const jvmtiHeapCallbacks *callbacks,
//                               const void *user_data);

// /*   120 : Set JNI Function Table */
// jvmtiError SetJNIFunctionTable(jvmtiEnv *env,
//                                const jniNativeInterface *function_table);

// /*   121 : Get JNI Function Table */
// jvmtiError GetJNIFunctionTable(jvmtiEnv *env,
//                                jniNativeInterface **function_table);

// /*   122 : Set Event Callbacks */
// jvmtiError SetEventCallbacks(jvmtiEnv *env,
//                              const jvmtiEventCallbacks *callbacks,
//                              jint size_of_callbacks);

// /*   123 : Generate Events */
// jvmtiError GenerateEvents(jvmtiEnv *env,
//                           jvmtiEvent event_type);

// /*   124 : Get Extension Functions */
// jvmtiError GetExtensionFunctions(jvmtiEnv *env,
//                                  jint *extension_count_ptr,
//                                  jvmtiExtensionFunctionInfo **extensions);

// /*   125 : Get Extension Events */
// jvmtiError GetExtensionEvents(jvmtiEnv *env,
//                               jint *extension_count_ptr,
//                               jvmtiExtensionEventInfo **extensions);

// /*   126 : Set Extension Event Callback */
// jvmtiError SetExtensionEventCallback(jvmtiEnv *env,
//                                      jint extension_event_index,
//                                      jvmtiExtensionEvent callback);

// /*   127 : Dispose Environment */
// jvmtiError DisposeEnvironment(jvmtiEnv *env);

// /*   128 : Get Error Name */
// jvmtiError GetErrorName(jvmtiEnv *env,
//                         jvmtiError error,
//                         char **name_ptr);

// /*   129 : Get JLocation Format */
// jvmtiError GetJLocationFormat(jvmtiEnv *env,
//                               jvmtiJlocationFormat *format_ptr);

// /*   130 : Get System Properties */
// jvmtiError GetSystemProperties(jvmtiEnv *env,
//                                jint *count_ptr,
//                                char ***property_ptr);

// /*   131 : Get System Property */
// jvmtiError GetSystemProperty(jvmtiEnv *env,
//                              const char *property,
//                              char **value_ptr);

// /*   132 : Set System Property */
// jvmtiError SetSystemProperty(jvmtiEnv *env,
//                              const char *property,
//                              const char *value);

// /*   133 : Get Phase */
// jvmtiError GetPhase(jvmtiEnv *env,
//                     jvmtiPhase *phase_ptr);

// /*   134 : Get Current Thread CPU Timer Information */
// jvmtiError GetCurrentThreadCpuTimerInfo(jvmtiEnv *env,
//                                         jvmtiTimerInfo *info_ptr);

// /*   135 : Get Current Thread CPU Time */
// jvmtiError GetCurrentThreadCpuTime(jvmtiEnv *env,
//                                    jlong *nanos_ptr);

// /*   136 : Get Thread CPU Timer Information */
// jvmtiError GetThreadCpuTimerInfo(jvmtiEnv *env,
//                                  jvmtiTimerInfo *info_ptr);

// /*   137 : Get Thread CPU Time */
// jvmtiError GetThreadCpuTime(jvmtiEnv *env,
//                             jthread thread,
//                             jlong *nanos_ptr);

// /*   138 : Get Timer Information */
// jvmtiError GetTimerInfo(jvmtiEnv *env,
//                         jvmtiTimerInfo *info_ptr);

// /*   139 : Get Time */
// jvmtiError GetTime(jvmtiEnv *env,
//                    jlong *nanos_ptr);

// /*   140 : Get Potential Capabilities */
// jvmtiError GetPotentialCapabilities(jvmtiEnv *env,
//                                     jvmtiCapabilities *capabilities_ptr);

// /*   142 : Add Capabilities */
// jvmtiError AddCapabilities(jvmtiEnv *env,
//                            const jvmtiCapabilities *capabilities_ptr);

// /*   143 : Relinquish Capabilities */
// jvmtiError RelinquishCapabilities(jvmtiEnv *env,
//                                   const jvmtiCapabilities *capabilities_ptr);

// /*   144 : Get Available Processors */
// jvmtiError GetAvailableProcessors(jvmtiEnv *env,
//                                   jint *processor_count_ptr);

// /*   145 : Get Class Version Numbers */
// jvmtiError GetClassVersionNumbers(jvmtiEnv *env,
//                                   jclass klass,
//                                   jint *minor_version_ptr,
//                                   jint *major_version_ptr);

// /*   146 : Get Constant Pool */
// jvmtiError GetConstantPool(jvmtiEnv *env,
//                            jclass klass,
//                            jint *constant_pool_count_ptr,
//                            jint *constant_pool_byte_count_ptr,
//                            unsigned char **constant_pool_bytes_ptr);

// /*   147 : Get Environment Local Storage */
// jvmtiError GetEnvironmentLocalStorage(jvmtiEnv *env,
//                                       void **data_ptr);

// /*   148 : Set Environment Local Storage */
// jvmtiError SetEnvironmentLocalStorage(jvmtiEnv *env,
//                                       const void *data);

// /*   149 : Add To Bootstrap Class Loader Search */
// jvmtiError AddToBootstrapClassLoaderSearch(jvmtiEnv *env,
//                                            const char *segment);

// /*   150 : Set Verbose Flag */
// jvmtiError SetVerboseFlag(jvmtiEnv *env,
//                           jvmtiVerboseFlag flag,
//                           jboolean value);

// /*   151 : Add To System Class Loader Search */
// jvmtiError AddToSystemClassLoaderSearch(jvmtiEnv *env,
//                                         const char *segment);

// /*   152 : Retransform Classes */
// jvmtiError RetransformClasses(jvmtiEnv *env,
//                               jint class_count,
//                               const jclass *classes);

// /*   153 : Get Owned Monitor Stack Depth Info */
// jvmtiError GetOwnedMonitorStackDepthInfo(jvmtiEnv *env,
//                                          jthread thread,
//                                          jint *monitor_info_count_ptr,
//                                          jvmtiMonitorStackDepthInfo **monitor_info_ptr);

// /*   154 : Get Object Size */
// jvmtiError GetObjectSize(jvmtiEnv *env,
//                          jobject object,
//                          jlong *size_ptr);

// /*   155 : Get Local Instance */
// jvmtiError GetLocalInstance(jvmtiEnv *env,
//                             jthread thread,
//                             jint depth,
//                             jobject *value_ptr);

// jvmtiError jvmti_get_loaded_classes(jvmtiEnv *env, jint *class_count_ptr, jclass **classes_buf);