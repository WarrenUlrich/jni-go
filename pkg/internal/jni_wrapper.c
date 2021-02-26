#include "jni_wrapper.h"

int jni_destroy_java_vm(JavaVM* vm)
{
    return (*vm)->DestroyJavaVM(vm);
}

int jni_attach_current_thread(JavaVM* vm, void** env_buf, void* args)
{
    return (*vm)->AttachCurrentThread(vm, env_buf, 0);
}

int jni_detach_current_thread(JavaVM* vm)
{
    return (*vm)->DetachCurrentThread(vm);
}

int jni_get_env(JavaVM* vm, void** env_buf, int version)
{
    return (*vm)->GetEnv(vm, env_buf, version);
}

int jni_attach_current_thread_as_daemon(JavaVM* vm, void** env_buf, int version)
{
    return (*vm)->GetEnv(vm, env_buf, version);
}

int jni_get_version(JNIEnv* env)
{
    return (*env)->GetVersion(env);
}

jclass jni_find_class(JNIEnv* env, const char* name)
{
    return (*env)->FindClass(env, name);
}

jmethodID jni_from_reflected_method(JNIEnv* env, jobject method)
{
    return (*env)->FromReflectedMethod(env, method);
}

jfieldID jni_from_reflected_field(JNIEnv* env, jobject field)
{
    return (*env)->FromReflectedField(env, field);
}

jobject jni_to_reflected_method(JNIEnv* env, jclass cls, jmethodID method_id, jboolean is_static)
{
    return (*env)->ToReflectedMethod(env, cls, method_id, is_static);
}

jclass jni_get_superclass(JNIEnv* env, jclass sub_class)
{
    return (*env)->GetSuperclass(env, sub_class);
}

jboolean jni_is_assignable_from(JNIEnv* env, jclass sub, jclass sup)
{
    return (*env)->IsAssignableFrom(env, sub, sup);
}

jobject jni_to_reflected_field(JNIEnv* env, jclass cls, jfieldID field_id, jboolean is_static)
{
    return (*env)->ToReflectedField(env, cls, field_id, is_static);
}

jint jni_throw(JNIEnv* env, jthrowable obj)
{
    return (*env)->Throw(env, obj);
}

jint jni_throw_new(JNIEnv* env, jclass cls, const char* msg)
{
    return (*env)->ThrowNew(env, cls, msg);
}

jthrowable jni_exception_occured(JNIEnv* env)
{
    return (*env)->ExceptionOccurred(env);
}

void jni_exception_describe(JNIEnv* env)
{
    (*env)->ExceptionDescribe(env);
}

void jni_exception_clear(JNIEnv* env)
{
    (*env)->ExceptionClear(env);
}

void jni_fatal_error(JNIEnv* env, const char* msg)
{
    (*env)->FatalError(env, msg);
}

jint jni_push_local_frame(JNIEnv* env, jint capacity)
{
    return (*env)->PushLocalFrame(env, capacity);
}

jobject jni_pop_local_frame(JNIEnv* env, jobject result)
{
    return (*env)->PopLocalFrame(env, result);
}

jobject jni_new_global_ref(JNIEnv* env, jobject lobj)
{
    return (*env)->NewGlobalRef(env, lobj);
}

void jni_delete_global_ref(JNIEnv* env, jobject gref) 
{
    (*env)->DeleteGlobalRef(env, gref);
}

void jni_delete_local_ref(JNIEnv* env, jobject obj)
{
    (*env)->DeleteLocalRef(env, obj);
}

jboolean jni_is_same_object(JNIEnv* env, jobject obj1, jobject obj2)
{
    return (*env)->IsSameObject(env, obj1, obj2);
}

jobject jni_new_local_ref(JNIEnv* env, jobject ref)
{
    return (*env)->NewLocalRef(env, ref);
}

jint jni_ensure_local_capacity(JNIEnv* env, jint capacity)
{
    return (*env)->EnsureLocalCapacity(env, capacity);
}

jobject jni_alloc_object(JNIEnv* env, jclass cls)
{
    return (*env)->AllocObject(env, cls);
}

jobject jni_new_object(JNIEnv* env, jclass cls, jmethodID method_id, const jvalue* args)
{
    return (*env)->NewObjectA(env, cls, method_id, args);
}

jclass jni_get_object_class(JNIEnv* env, jobject obj)
{
    return (*env)->GetObjectClass(env, obj);
}

jboolean jni_is_instance_of(JNIEnv* env, jobject obj, jclass cls)
{
    return (*env)->IsInstanceOf(env, obj, cls);
}

jmethodID jni_get_method_id(JNIEnv* env, jclass cls, const char* name, const char* sig)
{
    return (*env)->GetMethodID(env, cls, name, sig);
}

jobject jni_call_object_method(JNIEnv* env, jobject obj, jmethodID method_id, const jvalue* args)
{
    return (*env)->CallObjectMethodA(env, obj, method_id, args);
}

jboolean jni_call_boolean_method(JNIEnv* env, jobject obj, jmethodID method_id, const jvalue* args)
{
    return (*env)->CallBooleanMethodA(env, obj, method_id, args);
}

jbyte jni_call_byte_method(JNIEnv* env, jobject obj, jmethodID method_id, const jvalue* args)
{
    return (*env)->CallByteMethodA(env, obj, method_id, args);
}

jchar jni_call_char_method(JNIEnv* env, jobject obj, jmethodID method_id, const jvalue* args)
{
    return (*env)->CallCharMethodA(env, obj, method_id, args);
}

jshort jni_call_short_method(JNIEnv* env, jobject obj, jmethodID method_id, const jvalue* args)
{
    return (*env)->CallShortMethodA(env, obj, method_id, args);
}

jint jni_call_int_method(JNIEnv* env, jobject obj, jmethodID method_id, const jvalue* args)
{
    return (*env)->CallIntMethodA(env, obj, method_id, args);
}

jlong jni_call_long_method(JNIEnv* env, jobject obj, jmethodID method_id, const jvalue* args)
{
    return (*env)->CallLongMethodA(env, obj, method_id, args);
}

jfloat jni_call_float_method(JNIEnv* env, jobject obj, jmethodID method_id, const jvalue* args)
{
    return (*env)->CallFloatMethodA(env, obj, method_id, args);
}

jdouble jni_call_double_method(JNIEnv* env, jobject obj, jmethodID method_id, const jvalue* args)
{
    return (*env)->CallDoubleMethodA(env, obj, method_id, args);
}

void jni_call_void_method(JNIEnv* env, jobject obj, jmethodID method_id, const jvalue* args)
{
    (*env)->CallVoidMethodA(env, obj, method_id, args);
}

jobject jni_call_non_virtual_object_method(JNIEnv* env, jobject obj, jclass cls, jmethodID method_id, const jvalue* args)
{
    return (*env)->CallNonvirtualObjectMethodA(env, obj, cls, method_id, args);
}

jboolean jni_call_non_virtual_boolean_method(JNIEnv* env, jobject obj, jclass cls, jmethodID method_id, const jvalue* args)
{
    return (*env)->CallNonvirtualBooleanMethodA(env, obj, cls, method_id, args);
}

jbyte jni_call_non_virtual_byte_method(JNIEnv* env, jobject obj, jclass cls, jmethodID method_id, const jvalue* args)
{
    return (*env)->CallNonvirtualByteMethodA(env, obj, cls, method_id, args);
}

jchar jni_call_non_virtual_char_method(JNIEnv* env, jobject obj, jclass cls, jmethodID method_id, const jvalue* args)
{
    return (*env)->CallNonvirtualCharMethodA(env, obj, cls, method_id, args);
}

jshort jni_call_non_virtual_short_method(JNIEnv* env, jobject obj, jclass cls, jmethodID method_id, const jvalue* args)
{
    return (*env)->CallNonvirtualShortMethodA(env, obj, cls, method_id, args);
}

jint jni_call_non_virtual_int_method(JNIEnv* env, jobject obj, jclass cls, jmethodID method_id, const jvalue* args)
{
    return (*env)->CallNonvirtualIntMethodA(env, obj, cls, method_id, args);
}

jlong jni_call_non_virtual_long_method(JNIEnv* env, jobject obj, jclass cls, jmethodID method_id, const jvalue* args)
{
    return (*env)->CallNonvirtualLongMethodA(env, obj, cls, method_id, args);
}

jfloat jni_call_non_virtual_float_method(JNIEnv* env, jobject obj, jclass cls, jmethodID method_id, const jvalue* args)
{
    return (*env)->CallNonvirtualFloatMethodA(env, obj, cls, method_id, args);
}

jdouble jni_call_non_virtual_double_method(JNIEnv* env, jobject obj, jclass cls, jmethodID method_id, const jvalue* args)
{
    return (*env)->CallNonvirtualDoubleMethodA(env, obj, cls, method_id, args);
}

void jni_call_non_virtual_void_method(JNIEnv* env, jobject obj, jclass cls, jmethodID method_id, const jvalue* args)
{
    (*env)->CallNonvirtualVoidMethodA(env, obj, cls, method_id, args);
}

jfieldID jni_get_field_id(JNIEnv* env, jclass cls, const char* name, const char* sig)
{
    return (*env)->GetFieldID(env, cls, name, sig);
}

jobject jni_get_object_field(JNIEnv* env, jobject obj, jfieldID field_id)
{
    return (*env)->GetObjectField(env, obj, field_id);
}

jboolean jni_get_boolean_field(JNIEnv* env, jobject obj, jfieldID field_id)
{
    return (*env)->GetBooleanField(env, obj, field_id);
}

jbyte jni_get_byte_field(JNIEnv* env, jobject obj, jfieldID field_id)
{
    return (*env)->GetByteField(env, obj, field_id);
}

jchar jni_get_char_field(JNIEnv* env, jobject obj, jfieldID field_id)
{
    return (*env)->GetCharField(env, obj, field_id);
}

jshort jni_get_short_field(JNIEnv* env, jobject obj, jfieldID field_id)
{
    return (*env)->GetShortField(env, obj, field_id);
}

jint jni_get_int_field(JNIEnv* env, jobject obj, jfieldID field_id)
{
    return (*env)->GetIntField(env, obj, field_id);
}

jlong jni_get_long_field(JNIEnv* env, jobject obj, jfieldID field_id)
{
    return (*env)->GetLongField(env, obj, field_id);
}

jfloat jni_get_float_field(JNIEnv* env, jobject obj, jfieldID field_id)
{
    return (*env)->GetFloatField(env, obj, field_id);
}

jdouble jni_get_double_field(JNIEnv* env, jobject obj, jfieldID field_id)
{
    return (*env)->GetDoubleField(env, obj, field_id);
}

void jni_set_object_field(JNIEnv* env, jobject obj, jfieldID field_id, jobject value)
{
    (*env)->SetObjectField(env, obj, field_id, value);
}

void jni_set_boolean_field(JNIEnv* env, jobject obj, jfieldID field_id, jboolean value)
{
    (*env)->SetBooleanField(env, obj, field_id, value);
}

void jni_set_byte_field(JNIEnv* env, jobject obj, jfieldID field_id, jbyte value)
{
    (*env)->SetByteField(env, obj, field_id, value);
}

void jni_set_char_field(JNIEnv* env, jobject obj, jfieldID field_id, jchar value)
{
    (*env)->SetCharField(env, obj, field_id, value);
}

void jni_set_short_field(JNIEnv* env, jobject obj, jfieldID field_id, jshort value)
{
    (*env)->SetShortField(env, obj, field_id, value);
}

void jni_set_int_field(JNIEnv* env, jobject obj, jfieldID field_id, jint value)
{
    (*env)->SetIntField(env, obj, field_id, value);
}

void jni_set_long_field(JNIEnv* env, jobject obj, jfieldID field_id, jlong value)
{
    (*env)->SetLongField(env, obj, field_id, value);
}

void jni_set_float_field(JNIEnv* env, jobject obj, jfieldID field_id, jfloat value)
{
    (*env)->SetFloatField(env, obj, field_id, value);
}

void jni_set_double_field(JNIEnv* env, jobject obj, jfieldID field_id, jdouble value)
{
    (*env)->SetDoubleField(env, obj, field_id, value);
}

jmethodID jni_get_static_method_id(JNIEnv* env, jclass cls, const char* name, const char* sig)
{
    return (*env)->GetStaticMethodID(env, cls, name, sig);
}

jobject jni_call_static_object_method(JNIEnv* env, jclass cls, jmethodID method_id, const jvalue* args)
{
    return (*env)->CallStaticObjectMethodA(env, cls, method_id, args);
}

jboolean jni_call_static_boolean_method(JNIEnv* env, jclass cls, jmethodID method_id, const jvalue* args)
{
    return (*env)->CallStaticBooleanMethodA(env, cls, method_id, args);
}

jbyte jni_call_static_byte_method(JNIEnv* env, jclass cls, jmethodID method_id, const jvalue* args)
{
    return (*env)->CallStaticByteMethodA(env, cls, method_id, args);
}

jchar jni_call_static_char_method(JNIEnv* env, jclass cls, jmethodID method_id, const jvalue* args)
{
    return (*env)->CallStaticCharMethodA(env, cls, method_id, args);
}

jshort jni_call_static_short_method(JNIEnv* env, jclass cls, jmethodID method_id, const jvalue* args)
{
    return (*env)->CallStaticShortMethodA(env, cls, method_id, args);
}

jint jni_call_static_int_method(JNIEnv* env, jclass cls, jmethodID method_id, const jvalue* args)
{
    return (*env)->CallStaticIntMethodA(env, cls, method_id, args);
}

jlong jni_call_static_long_method(JNIEnv* env, jclass cls, jmethodID method_id, const jvalue* args)
{
    return (*env)->CallStaticLongMethodA(env, cls, method_id, args);
}

jfloat jni_call_static_float_method(JNIEnv* env, jclass cls, jmethodID method_id, const jvalue* args)
{
    return (*env)->CallStaticFloatMethodA(env, cls, method_id, args);
}

jdouble jni_call_static_double_method(JNIEnv* env, jclass cls, jmethodID method_id, const jvalue* args)
{
    return (*env)->CallStaticDoubleMethodA(env, cls, method_id, args);
}

void jni_call_static_void_method(JNIEnv* env, jclass cls, jmethodID method_id, const jvalue* args)
{
    (*env)->CallStaticVoidMethodA(env, cls, method_id, args);
}

jfieldID jni_get_static_field_id(JNIEnv* env, jclass cls, const char* name, const char* sig)
{
    return (*env)->GetStaticFieldID(env, cls, name, sig);
}

jobject jni_get_static_object_field(JNIEnv* env, jclass cls, jfieldID field_id)
{
    return (*env)->GetStaticObjectField(env, cls, field_id);
}

jboolean jni_get_static_boolean_field(JNIEnv* env, jclass cls, jfieldID field_id)
{
    return (*env)->GetStaticBooleanField(env, cls, field_id);
}

jbyte jni_get_static_byte_field(JNIEnv* env, jclass cls, jfieldID field_id)
{
    return (*env)->GetStaticByteField(env, cls, field_id);
}

jchar jni_get_static_char_field(JNIEnv* env, jclass cls, jfieldID field_id)
{
    return (*env)->GetStaticCharField(env, cls, field_id);
}

jshort jni_get_static_short_field(JNIEnv* env, jclass cls, jfieldID field_id)
{
    return (*env)->GetStaticShortField(env, cls, field_id);
}

jint jni_get_static_int_field(JNIEnv* env, jclass cls, jfieldID field_id)
{
    return (*env)->GetStaticIntField(env, cls, field_id);
}
jlong jni_get_static_long_field(JNIEnv* env, jclass cls, jfieldID field_id)
{
    return (*env)->GetStaticLongField(env, cls, field_id);
}

jfloat jni_get_static_float_field(JNIEnv* env, jclass cls, jfieldID field_id)
{
    return (*env)->GetStaticFloatField(env, cls, field_id);
}

jdouble jni_get_static_double_field(JNIEnv* env, jclass cls, jfieldID field_id)
{
    return (*env)->GetStaticDoubleField(env, cls, field_id);
}

void jni_set_static_object_field(JNIEnv* env, jclass cls, jfieldID field_id, jobject value)
{
    (*env)->SetStaticObjectField(env, cls, field_id, value);
}

void jni_set_static_boolean_field(JNIEnv* env, jclass cls, jfieldID field_id, jboolean value)
{
    (*env)->SetStaticBooleanField(env, cls, field_id, value);
}

void jni_set_static_byte_field(JNIEnv* env, jclass cls, jfieldID field_id, jbyte value)
{
    (*env)->SetStaticByteField(env, cls, field_id, value);
}

void jni_set_static_char_field(JNIEnv* env, jclass cls, jfieldID field_id, jchar value)
{
    (*env)->SetStaticCharField(env, cls, field_id, value);
}

void jni_set_static_short_field(JNIEnv* env, jclass cls, jfieldID field_id, jshort value)
{
    (*env)->SetStaticShortField(env, cls, field_id, value);
}

void jni_set_static_int_field(JNIEnv* env, jclass cls, jfieldID field_id, jint value)
{
    (*env)->SetStaticIntField(env, cls, field_id, value);
}

void jni_set_static_long_field(JNIEnv* env, jclass cls, jfieldID field_id, jlong value)
{
    (*env)->SetStaticLongField(env, cls, field_id, value);
}

void jni_set_static_float_field(JNIEnv* env, jclass cls, jfieldID field_id, jfloat value)
{
    (*env)->SetStaticFloatField(env, cls, field_id, value);
}

void jni_set_static_double_field(JNIEnv* env, jclass cls, jfieldID field_id, jdouble value)
{
    (*env)->SetStaticDoubleField(env, cls, field_id, value);
}

jstring jni_new_string(JNIEnv* env, const jchar* unicode, jsize len)
{
    return (*env)->NewString(env, unicode, len);
}

jsize jni_get_string_length(JNIEnv* env, jstring str)
{
    return (*env)->GetStringLength(env, str);
}

const jchar* jni_get_string_chars(JNIEnv* env, jstring str, jboolean* is_copy)
{
    return (*env)->GetStringChars(env, str, is_copy);
}

void jni_release_string_chars(JNIEnv* env, jstring str, const jchar* chars)
{
    (*env)->ReleaseStringChars(env, str, chars);
}

jstring jni_new_string_utf(JNIEnv* env, const char* utf)
{
    return (*env)->NewStringUTF(env, utf);
}

jsize jni_get_string_utf_length(JNIEnv* env, jstring str)
{
    return (*env)->GetStringUTFLength(env, str);
}

const char* jni_get_string_utf_chars(JNIEnv* env, jstring str, jboolean* is_copy)
{
    return (*env)->GetStringUTFChars(env, str, is_copy);
}

void jni_release_string_utf_chars(JNIEnv* env, jstring str, const char* chars)
{
    return (*env)->ReleaseStringUTFChars(env, str, chars);
}

jsize jni_get_array_length(JNIEnv* env, jarray array)
{
    return (*env)->GetArrayLength(env, array);
}

jobjectArray jni_new_object_array(JNIEnv* env, jsize len, jclass cls, jobject init)
{
    return (*env)->NewObjectArray(env, len, cls, init);
}

jobject jni_get_object_array_element(JNIEnv* env, jobjectArray array, jsize index)
{
    return (*env)->GetObjectArrayElement(env, array, index);
}

void jni_set_object_array_element(JNIEnv* env, jobjectArray array, jsize index, jobject val)
{
    (*env)->SetObjectArrayElement(env, array, index, val);
}

jbooleanArray jni_new_boolean_array(JNIEnv* env, jsize len)
{
    return (*env)->NewBooleanArray(env, len);
}

jbyteArray jni_new_byte_array(JNIEnv* env, jsize len)
{
    return (*env)->NewByteArray(env, len);
}

jcharArray jni_new_char_array(JNIEnv* env, jsize len)
{
    return (*env)->NewCharArray(env, len);
}

jshortArray jni_new_short_array(JNIEnv* env, jsize len)
{
    return (*env)->NewShortArray(env, len);
}

jintArray jni_new_int_array(JNIEnv* env, jsize len)
{
    return (*env)->NewIntArray(env, len);
}

jlongArray jni_new_long_array(JNIEnv* env, jsize len)
{
    return (*env)->NewLongArray(env, len);
}

jfloatArray jni_new_float_array(JNIEnv* env, jsize len)
{
    return (*env)->NewFloatArray(env, len);
}

jdoubleArray jni_new_double_array(JNIEnv* env, jsize len)
{
    return (*env)->NewDoubleArray(env, len);
}

jboolean* jni_get_boolean_array_elements(JNIEnv* env, jbooleanArray array, jboolean* is_copy)
{
    return (*env)->GetBooleanArrayElements(env, array, is_copy);
}

jbyte* jni_get_byte_array_elements(JNIEnv* env, jbyteArray array, jboolean* is_copy)
{
    return (*env)->GetByteArrayElements(env, array, is_copy);
}

jchar* jni_get_char_array_elements(JNIEnv* env, jcharArray array, jboolean* is_copy)
{
    return (*env)->GetCharArrayElements(env, array, is_copy);
}

jshort* jni_get_short_array_elements(JNIEnv* env, jshortArray array, jboolean* is_copy)
{
    return (*env)->GetShortArrayElements(env, array, is_copy);
}

jint* jni_get_int_array_elements(JNIEnv* env, jintArray array, jboolean* is_copy)
{
    return (*env)->GetIntArrayElements(env, array, is_copy);
}

jlong* jni_get_long_array_elements(JNIEnv* env, jlongArray array, jboolean* is_copy)
{
    return (*env)->GetLongArrayElements(env, array, is_copy);
}

jfloat* jni_get_float_array_elements(JNIEnv* env, jfloatArray array, jboolean* is_copy)
{
    return (*env)->GetFloatArrayElements(env, array, is_copy);
}

jdouble* jni_get_double_array_elements(JNIEnv* env, jdoubleArray array, jboolean* is_copy)
{
    return (*env)->GetDoubleArrayElements(env, array, is_copy);
}

void jni_release_boolean_array_elements(JNIEnv* env, jbooleanArray array, jboolean* elements, jint mode)
{
    (*env)->ReleaseBooleanArrayElements(env, array, elements, mode);
}

void jni_release_byte_array_elements(JNIEnv* env, jbyteArray array, jbyte* elements, jint mode)
{
    (*env)->ReleaseByteArrayElements(env, array, elements, mode);
}

void jni_release_char_array_elements(JNIEnv* env, jcharArray array, jchar* elements, jint mode)
{
    (*env)->ReleaseCharArrayElements(env, array, elements, mode);
}

void jni_release_short_array_elements(JNIEnv* env, jshortArray array, jshort* elements, jint mode)
{
    (*env)->ReleaseShortArrayElements(env, array, elements, mode);
}

void jni_release_int_array_elements(JNIEnv* env, jintArray array, jint* elements, jint mode)
{
    (*env)->ReleaseIntArrayElements(env, array, elements, mode);
}

void jni_release_long_array_elements(JNIEnv* env, jlongArray array, jlong* elements, jint mode)
{
    (*env)->ReleaseLongArrayElements(env, array, elements, mode);
}

void jni_release_float_array_elements(JNIEnv* env, jfloatArray array, jfloat* elements, jint mode)
{
    (*env)->ReleaseFloatArrayElements(env, array, elements, mode);
}

void jni_release_double_array_elements(JNIEnv* env, jdoubleArray array, jdouble* elements, jint mode)
{
    (*env)->ReleaseDoubleArrayElements(env, array, elements, mode);
}

void jni_get_boolean_array_region(JNIEnv* env, jbooleanArray array, jsize start, jsize len, jboolean* buffer)
{
    (*env)->GetBooleanArrayRegion(env, array, start, len, buffer);
}

void jni_get_byte_array_region(JNIEnv* env, jbyteArray array, jsize start, jsize len, jbyte* buffer)
{
    (*env)->GetByteArrayRegion(env, array, start, len, buffer);
}

void jni_get_char_array_region(JNIEnv* env, jcharArray array, jsize start, jsize len, jchar* buffer)
{
    (*env)->GetCharArrayRegion(env, array, start, len, buffer);
}

void jni_get_short_array_region(JNIEnv* env, jshortArray array, jsize start, jsize len, jshort* buffer)
{
    (*env)->GetShortArrayRegion(env, array, start, len, buffer);
}

void jni_get_int_array_region(JNIEnv* env, jintArray array, jsize start, jsize len, jint* buffer)
{
    (*env)->GetIntArrayRegion(env, array, start, len, buffer);
}

void jni_get_long_array_region(JNIEnv* env, jlongArray array, jsize start, jsize len, jlong* buffer)
{
    (*env)->GetLongArrayRegion(env, array, start, len, buffer);
}

void jni_get_float_array_region(JNIEnv* env, jfloatArray array, jsize start, jsize len, jfloat* buffer)
{
    (*env)->GetFloatArrayRegion(env, array, start, len, buffer);
}

void jni_get_double_array_region(JNIEnv* env, jdoubleArray array, jsize start, jsize len, jdouble* buffer)
{
    (*env)->GetDoubleArrayRegion(env, array, start, len, buffer);
}

void jni_set_boolean_array_region(JNIEnv* env, jbooleanArray array, jsize start, jsize len, const jboolean* buffer)
{
    (*env)->SetBooleanArrayRegion(env, array, start, len, buffer);
}

void jni_set_byte_array_region(JNIEnv* env, jbyteArray array, jsize start, jsize len, const jbyte* buffer)
{
    (*env)->SetByteArrayRegion(env, array, start, len, buffer);
}

void jni_set_char_array_region(JNIEnv* env, jcharArray array, jsize start, jsize len, const jchar* buffer)
{
    (*env)->SetCharArrayRegion(env, array, start, len, buffer);
}

void jni_set_short_array_region(JNIEnv* env, jshortArray array, jsize start, jsize len, const jshort* buffer)
{
    (*env)->SetShortArrayRegion(env, array, start, len, buffer);
}

void jni_set_int_array_region(JNIEnv* env, jintArray array, jsize start, jsize len, const jint* buffer)
{
    (*env)->SetIntArrayRegion(env, array, start, len, buffer);
}

void jni_set_long_array_region(JNIEnv* env, jlongArray array, jsize start, jsize len, const jlong* buffer)
{
    (*env)->SetLongArrayRegion(env, array, start, len, buffer);
}

void jni_set_float_array_region(JNIEnv* env, jfloatArray array, jsize start, jsize len, const jfloat* buffer)
{
    (*env)->SetFloatArrayRegion(env, array, start, len, buffer);
}

void jni_set_double_array_region(JNIEnv* env, jdoubleArray array, jsize start, jsize len, const jdouble* buffer)
{
    (*env)->SetDoubleArrayRegion(env, array, start, len, buffer);
}

jint jni_register_natives(JNIEnv* env, jclass cls, const JNINativeMethod* methods, jint method_count)
{
    return (*env)->RegisterNatives(env, cls, methods, method_count);
}

jint jni_unregister_native(JNIEnv* env, jclass cls)
{
    return (*env)->UnregisterNatives(env, cls);
}

jint jni_monitor_enter(JNIEnv* env, jobject obj)
{
    return (*env)->MonitorEnter(env, obj);
}

jint jni_monitor_exit(JNIEnv* env, jobject obj)
{
    return (*env)->MonitorExit(env, obj);
}

jint get_java_vm(JNIEnv* env, JavaVM** vm)
{
    return (*env)->GetJavaVM(env, vm);
}

void jni_get_string_region(JNIEnv* env, jstring str, jsize start, jsize len, jchar* buf)
{
    (*env)->GetStringRegion(env, str, start, len, buf);
}

void jni_get_string_utf_region(JNIEnv* env, jstring str, jsize start, jsize len, char* buf)
{
    (*env)->GetStringUTFRegion(env, str, start, len, buf);
}

void* jni_get_primitive_array_critical(JNIEnv* env, jarray array, jboolean* is_copy)
{
    return (*env)->GetPrimitiveArrayCritical(env, array, is_copy);
}

void jni_release_primitive_array_critical(JNIEnv* env, jarray array, void* carray, jint mode)
{
    (*env)->ReleasePrimitiveArrayCritical(env, array, carray, mode);
}

const jchar* jni_get_string_critical(JNIEnv* env, jstring string, jboolean* is_copy)
{
    return (*env)->GetStringCritical(env, string, is_copy);
}

void jni_release_string_critical(JNIEnv* env, jstring str, const jchar* cstring)
{
    return (*env)->ReleaseStringCritical(env, str, cstring);
}

jweak jni_new_weak_global_ref(JNIEnv* env, jobject obj)
{
    return (*env)->NewWeakGlobalRef(env, obj);
}

void jni_delete_weak_global_ref(JNIEnv* env, jweak ref)
{
    (*env)->DeleteWeakGlobalRef(env, ref);
}

jboolean jni_exception_check(JNIEnv* env)
{
    return (*env)->ExceptionCheck(env);
}

jobject jni_new_direct_byte_buffer(JNIEnv* env, void* address, jlong capacity)
{
    return (*env)->NewDirectByteBuffer(env, address, capacity);
}

void* jni_get_direct_buffer_address(JNIEnv* env, jobject buf)
{
    return (*env)->GetDirectBufferAddress(env, buf);
}

jlong jni_get_direct_buffer_capacity(JNIEnv* env, jobject buf)
{
    return (*env)->GetDirectBufferCapacity(env, buf);
}

jobjectRefType jni_get_object_ref_type(JNIEnv* env, jobject obj)
{
    return (*env)->GetObjectRefType(env, obj);
}





