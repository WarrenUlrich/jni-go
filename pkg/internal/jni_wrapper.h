#pragma once
#include <jni.h>

int jni_destroy_java_vm(JavaVM* vm);
int jni_attach_current_thread(JavaVM* vm, void** env_buf, void* args);
int jni_detach_current_thread(JavaVM* JavaVM);
int jni_get_env(JavaVM* vm, void** env_buf, int version);
int jni_attach_current_thread_as_daemon(JavaVM* vm, void** env_buf, int version);

// =============== JNIEnv functions ===============
int jni_get_version(JNIEnv* env);

jclass jni_define_class(JNIEnv* env, const char* name, jobject loader, const jbyte* buf, jsize len);
jclass jni_find_class(JNIEnv* env, const char* name);

jmethodID jni_from_reflected_method(JNIEnv* env, jobject method);
jfieldID jni_from_reflected_field(JNIEnv* env, jobject field);

jobject jni_to_reflected_method(JNIEnv* env, jclass cls, jmethodID method_id, jboolean is_static);

jclass jni_get_superclass(JNIEnv* env, jclass sub_class);
jboolean jni_is_assignable_from(JNIEnv* env, jclass sub, jclass sup);

jobject jni_to_reflected_field(JNIEnv* env, jclass cls, jfieldID field_id, jboolean is_static);

jint jni_throw(JNIEnv* env, jthrowable obj);
jint jni_throw_new(JNIEnv* env, jclass cls, const char* msg);
jthrowable jni_exception_occured(JNIEnv* env);
void jni_exception_describe(JNIEnv* env);
void jni_exception_clear(JNIEnv* env);
void jni_fatal_error(JNIEnv* env, const char* msg);

jint jni_push_local_frame(JNIEnv* env, jint capacity);
jobject jni_pop_local_frame(JNIEnv* env, jobject result);

jobject jni_new_global_ref(JNIEnv* env, jobject lobj);

void jni_delete_global_ref(JNIEnv* env, jobject gref);
void jni_delete_local_ref(JNIEnv* env, jobject obj);
jboolean jni_is_same_object(JNIEnv* env, jobject obj1, jobject obj2);
jobject jni_new_local_ref(JNIEnv* env, jobject ref);
jint jni_ensure_local_capacity(JNIEnv* env, jint capacity);

jobject jni_alloc_object(JNIEnv* env, jclass cls);
jobject jni_new_object(JNIEnv* env, jclass cls, jmethodID method_id, const jvalue* args);

jclass jni_get_object_class(JNIEnv* env, jobject obj);
jboolean jni_is_instance_of(JNIEnv* env, jobject obj, jclass cls);

jmethodID jni_get_method_id(JNIEnv* env, jclass cls, const char* name, const char* sig);
jobject jni_call_object_method(JNIEnv* env, jobject obj, jmethodID method_id, const jvalue* args);
jboolean jni_call_boolean_method(JNIEnv* env, jobject obj, jmethodID method_id, const jvalue* args);
jbyte jni_call_byte_method(JNIEnv* env, jobject obj, jmethodID method_id, const jvalue* args);
jchar jni_call_char_method(JNIEnv* env, jobject obj, jmethodID method_id, const jvalue* args);
jshort jni_call_short_method(JNIEnv* env, jobject obj, jmethodID method_id, const jvalue* args);
jint jni_call_int_method(JNIEnv* env, jobject obj, jmethodID method_id, const jvalue* args);
jlong jni_call_long_method(JNIEnv* env, jobject obj, jmethodID method_id, const jvalue* args);
jfloat jni_call_float_method(JNIEnv* env, jobject obj, jmethodID method_id, const jvalue* args);
jdouble jni_call_double_method(JNIEnv* env, jobject obj, jmethodID method_id, const jvalue* args);
void jni_call_void_method(JNIEnv* env, jobject obj, jmethodID method_id, const jvalue* args);

jobject jni_call_non_virtual_object_method(JNIEnv* env, jobject obj, jclass cls, jmethodID method_id, const jvalue* args);
jboolean jni_call_non_virtual_boolean_method(JNIEnv* env, jobject obj, jclass cls, jmethodID method_id, const jvalue* args);
jbyte jni_call_non_virtual_byte_method(JNIEnv* env, jobject obj, jclass cls, jmethodID method_id, const jvalue* args);
jchar jni_call_non_virtual_char_method(JNIEnv* env, jobject obj, jclass cls, jmethodID method_id, const jvalue* args);
jshort jni_call_non_virtual_short_method(JNIEnv* env, jobject obj, jclass cls, jmethodID method_id, const jvalue* args);
jint jni_call_non_virtual_int_method(JNIEnv* env, jobject obj, jclass cls, jmethodID method_id, const jvalue* args);
jlong jni_call_non_virtual_long_method(JNIEnv* env, jobject obj, jclass cls, jmethodID method_id, const jvalue* args);
jfloat jni_call_non_virtual_float_method(JNIEnv* env, jobject obj, jclass cls, jmethodID method_id, const jvalue* args);
jdouble jni_call_non_virtual_double_method(JNIEnv* env, jobject obj, jclass cls, jmethodID method_id, const jvalue* args);
void jni_call_non_virtual_void_method(JNIEnv* env, jobject obj, jclass cls, jmethodID method_id, const jvalue* args);

jfieldID jni_get_field_id(JNIEnv* env, jclass cls, const char* name, const char* sig);
jobject jni_get_object_field(JNIEnv* env, jobject obj, jfieldID field_id);
jboolean jni_get_boolean_field(JNIEnv* env, jobject obj, jfieldID field_id);
jbyte jni_get_byte_field(JNIEnv* env, jobject obj, jfieldID field_id);
jchar jni_get_char_field(JNIEnv* env, jobject obj, jfieldID field_id);
jshort jni_get_short_field(JNIEnv* env, jobject obj, jfieldID field_id);
jint jni_get_int_field(JNIEnv* env, jobject obj, jfieldID field_id);
jlong jni_get_long_field(JNIEnv* env, jobject obj, jfieldID field_id);
jfloat jni_get_float_field(JNIEnv* env, jobject obj, jfieldID field_id);
jdouble jni_get_double_field(JNIEnv* env, jobject obj, jfieldID field_id);

void jni_set_object_field(JNIEnv* env, jobject obj, jfieldID field_id, jobject value);
void jni_set_boolean_field(JNIEnv* env, jobject obj, jfieldID field_id, jboolean value);
void jni_set_byte_field(JNIEnv* env, jobject obj, jfieldID field_id, jbyte value);
void jni_set_char_field(JNIEnv* env, jobject obj, jfieldID field_id, jchar value);
void jni_set_short_field(JNIEnv* env, jobject obj, jfieldID field_id, jshort value);
void jni_set_int_field(JNIEnv* env, jobject obj, jfieldID field_id, jint value);
void jni_set_long_field(JNIEnv* env, jobject obj, jfieldID field_id, jlong value);
void jni_set_float_field(JNIEnv* env, jobject obj, jfieldID field_id, jfloat value);
void jni_set_double_field(JNIEnv* env, jobject obj, jfieldID field_id, jdouble value);

jmethodID jni_get_static_method_id(JNIEnv* env, jclass cls, const char* name, const char* sig);
jobject jni_call_static_object_method(JNIEnv* env, jclass cls, jmethodID method_id, const jvalue* args);
jboolean jni_call_static_boolean_method(JNIEnv* env, jclass cls, jmethodID method_id, const jvalue* args);
jbyte jni_call_static_byte_method(JNIEnv* env, jclass cls, jmethodID method_id, const jvalue* args);
jchar jni_call_static_char_method(JNIEnv* env, jclass cls, jmethodID method_id, const jvalue* args);
jshort jni_call_static_short_method(JNIEnv* env, jclass cls, jmethodID method_id, const jvalue* args);
jint jni_call_static_int_method(JNIEnv* env, jclass cls, jmethodID method_id, const jvalue* args);
jlong jni_call_static_long_method(JNIEnv* env, jclass cls, jmethodID method_id, const jvalue* args);
jfloat jni_call_static_float_method(JNIEnv* env, jclass cls, jmethodID method_id, const jvalue* args);
jdouble jni_call_static_double_method(JNIEnv* env, jclass cls, jmethodID method_id, const jvalue* args);

void jni_call_static_void_method(JNIEnv* env, jclass cls, jmethodID method_id, const jvalue* args);

jfieldID jni_get_static_field_id(JNIEnv* env, jclass cls, const char* name, const char* sig);
jobject jni_get_static_object_field(JNIEnv* env, jclass cls, jfieldID field_id);
jboolean jni_get_static_boolean_field(JNIEnv* env, jclass cls, jfieldID field_id);
jbyte jni_get_static_byte_field(JNIEnv* env, jclass cls, jfieldID field_id);
jchar jni_get_static_char_field(JNIEnv* env, jclass cls, jfieldID field_id);
jshort jni_get_static_short_field(JNIEnv* env, jclass cls, jfieldID field_id);
jint jni_get_static_int_field(JNIEnv* env, jclass cls, jfieldID field_id);
jlong jni_get_static_long_field(JNIEnv* env, jclass cls, jfieldID field_id);
jfloat jni_get_static_float_field(JNIEnv* env, jclass cls, jfieldID field_id);
jdouble jni_get_static_double_field(JNIEnv* env, jclass cls, jfieldID field_id);

void jni_set_static_object_field(JNIEnv* env, jclass cls, jfieldID field_id, jobject value);
void jni_set_static_boolean_field(JNIEnv* env, jclass cls, jfieldID field_id, jboolean value);
void jni_set_static_byte_field(JNIEnv* env, jclass cls, jfieldID field_id, jbyte value);
void jni_set_static_char_field(JNIEnv* env, jclass cls, jfieldID field_id, jchar value);
void jni_set_static_short_field(JNIEnv* env, jclass cls, jfieldID field_id, jshort value);
void jni_set_static_int_field(JNIEnv* env, jclass cls, jfieldID field_id, jint value);
void jni_set_static_long_field(JNIEnv* env, jclass cls, jfieldID field_id, jlong value);
void jni_set_static_float_field(JNIEnv* env, jclass cls, jfieldID field_id, jfloat value);
void jni_set_static_double_field(JNIEnv* env, jclass cls, jfieldID field_id, jdouble value);

jstring jni_new_string(JNIEnv* env, const jchar* unicode, jsize len);
jsize jni_get_string_length(JNIEnv* env, jstring str);
const jchar* jni_get_string_chars(JNIEnv* env, jstring str, jboolean* is_copy);
void jni_release_string_chars(JNIEnv* env, jstring str, const jchar* chars);

jstring jni_new_string_utf(JNIEnv* env, const char* utf);
jsize jni_get_string_utf_length(JNIEnv* env, jstring str);
const char* jni_get_string_utf_chars(JNIEnv* env, jstring str, jboolean* is_copy);
void jni_release_string_utf_chars(JNIEnv* env, jstring str, const char* chars);

jsize jni_get_array_length(JNIEnv* env, jarray array);

jobjectArray jni_new_object_array(JNIEnv* env, jsize len, jclass cls, jobject init);
jobject jni_get_object_array_element(JNIEnv* env, jobjectArray array, jsize index);
void jni_set_object_array_element(JNIEnv* env, jobjectArray array, jsize index, jobject val);

jbooleanArray jni_new_boolean_array(JNIEnv* env, jsize len);
jbyteArray jni_new_byte_array(JNIEnv* env, jsize len);
jcharArray jni_new_char_array(JNIEnv* env, jsize len);
jshortArray jni_new_short_array(JNIEnv* env, jsize len);
jintArray jni_new_int_array(JNIEnv* env, jsize len);
jlongArray jni_new_long_array(JNIEnv* env, jsize len);
jfloatArray jni_new_float_array(JNIEnv* env, jsize len);
jdoubleArray jni_new_double_array(JNIEnv* env, jsize len);

jboolean* jni_get_boolean_array_elements(JNIEnv* env, jbooleanArray array, jboolean* is_copy);
jbyte* jni_get_byte_array_elements(JNIEnv* env, jbyteArray array, jboolean* is_copy);
jchar* jni_get_char_array_elements(JNIEnv* env, jcharArray array, jboolean* is_copy);
jshort* jni_get_short_array_elements(JNIEnv* env, jshortArray array, jboolean* is_copy);
jint* jni_get_int_array_elements(JNIEnv* env, jintArray array, jboolean* is_copy);
jlong* jni_get_long_array_elements(JNIEnv* env, jlongArray array, jboolean* is_copy);
jfloat* jni_get_float_array_elements(JNIEnv* env, jfloatArray array, jboolean* is_copy);
jdouble* jni_get_double_array_elements(JNIEnv* env, jdoubleArray array, jboolean* is_copy);

void jni_release_boolean_array_elements(JNIEnv* env, jbooleanArray array, jboolean* elements, jint mode);
void jni_release_byte_array_elements(JNIEnv* env, jbyteArray array, jbyte* elements, jint mode);
void jni_release_char_array_elements(JNIEnv* env, jcharArray array, jchar* elements, jint mode);
void jni_release_short_array_elements(JNIEnv* env, jshortArray array, jshort* elements, jint mode);
void jni_release_int_array_elements(JNIEnv* env, jintArray array, jint* elements, jint mode);
void jni_release_long_array_elements(JNIEnv* env, jlongArray array, jlong* elements, jint mode);
void jni_release_float_array_elements(JNIEnv* env, jfloatArray array, jfloat* elements, jint mode);
void jni_release_double_array_elements(JNIEnv* env, jdoubleArray array, jdouble* elements, jint mode);

void jni_get_boolean_array_region(JNIEnv* env, jbooleanArray array, jsize start, jsize len, jboolean* buffer);
void jni_get_byte_array_region(JNIEnv* env, jbyteArray array, jsize start, jsize len, jbyte* buffer);
void jni_get_char_array_region(JNIEnv* env, jcharArray array, jsize start, jsize len, jchar* buffer);
void jni_get_short_array_region(JNIEnv* env, jshortArray array, jsize start, jsize len, jshort* buffer);
void jni_get_int_array_region(JNIEnv* env, jintArray array, jsize start, jsize len, jint* buffer);
void jni_get_long_array_region(JNIEnv* env, jlongArray array, jsize start, jsize len, jlong* buffer);
void jni_get_float_array_region(JNIEnv* env, jfloatArray array, jsize start, jsize len, jfloat* buffer);
void jni_get_double_array_region(JNIEnv* env, jdoubleArray array, jsize start, jsize len, jdouble* buffer);

void jni_set_boolean_array_region(JNIEnv* env, jbooleanArray array, jsize start, jsize len, const jboolean* buffer);
void jni_set_byte_array_region(JNIEnv* env, jbyteArray array, jsize start, jsize len, const jbyte* buffer);
void jni_set_char_array_region(JNIEnv* env, jcharArray array, jsize start, jsize len, const jchar* buffer);
void jni_set_short_array_region(JNIEnv* env, jshortArray array, jsize start, jsize len, const jshort* buffer);
void jni_set_int_array_region(JNIEnv* env, jintArray array, jsize start, jsize len, const jint* buffer);
void jni_set_long_array_region(JNIEnv* env, jlongArray array, jsize start, jsize len, const jlong* buffer);
void jni_set_float_array_region(JNIEnv* env, jfloatArray array, jsize start, jsize len, const jfloat* buffer);
void jni_set_double_array_region(JNIEnv* env, jdoubleArray array, jsize start, jsize len, const jdouble* buffer);

jint jni_register_natives(JNIEnv* env, jclass cls, const JNINativeMethod* methods, jint method_count);
jint jni_unregister_native(JNIEnv* env, jclass cls);

jint jni_monitor_enter(JNIEnv* env, jobject obj);
jint jni_monitor_exit(JNIEnv* env, jobject obj);

jint get_java_vm(JNIEnv* env, JavaVM** vm);

void jni_get_string_region(JNIEnv* env, jstring str, jsize start, jsize len, jchar* buf);
void jni_get_string_utf_region(JNIEnv* env, jstring str, jsize start, jsize len, char* buf);

void* jni_get_primitive_array_critical(JNIEnv* env, jarray array, jboolean* is_copy);
void jni_release_primitive_array_critical(JNIEnv* env, jarray array, void* carray, jint mode);

const jchar* jni_get_string_critical(JNIEnv* env, jstring string, jboolean* is_copy);
void jni_release_string_critical(JNIEnv* env, jstring str, const jchar* cstring);

jweak jni_new_weak_global_ref(JNIEnv* env, jobject obj);
void jni_delete_weak_global_ref(JNIEnv* env, jweak ref);

jboolean jni_exception_check(JNIEnv* env);

jobject jni_new_direct_byte_buffer(JNIEnv* env, void* address, jlong capacity);
void* jni_get_direct_buffer_address(JNIEnv* env, jobject buf);
jlong jni_get_direct_buffer_capacity(JNIEnv* env, jobject buf);

jobjectRefType jni_get_object_ref_type(JNIEnv* env, jobject obj);

