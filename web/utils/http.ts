export const baseURL = 'http://127.0.0.1:8080'
import store from '../store/index';
// 添加拦截器
const httpInterceptor = {
  // 拦截前触发
  invoke(options: UniApp.RequestOptions) {
    // 1. 非 http 开头需拼接地址
    if (!options.url.startsWith('http')) {
      options.url = baseURL + options.url
    }
    // 2. 请求超时, 默认 60s
    options.timeout = 10000
    // 3. 添加小程序端请求头标识
    options.header = {
      ...options.header,
      'source-client': 'miniapp',
    }
    // 4. 添加 JWT 到请求头
    const jwt = store.state.jwt;
    if (jwt) {
      options.header.Authorization = `Bearer ${jwt}`;
    }
  },
}
uni.addInterceptor('request', httpInterceptor)
uni.addInterceptor('uploadFile', httpInterceptor)
/**
 * 请求函数
 * @param  UniApp.RequestOptions
 * @returns Promise
 *  1. 返回 Promise 对象
 *  2. 获取数据成功
 *    2.1 提取核心数据 res.data
 *    2.2 添加类型，支持泛型
 *  3. 获取数据失败
 *    3.1 401错误  -> 清理用户信息，跳转到登录页
 *    3.2 其他错误 -> 根据后端错误信息轻提示
 *    3.3 网络错误 -> 提示用户换网络
 */
export type ApiEnvelope<T> = {
  code: number
  msg: string
  data: T
}

function isSuccessEnvelope<T>(body: ApiEnvelope<T> | null | undefined): boolean {
  if (!body || typeof body !== 'object') {
    return false
  }
  if (body.code === 200) {
    return true
  }
  return 'data' in body && body.data !== undefined && body.data !== null
}

// 2.2 添加类型，支持泛型
export const http = <T>(options: UniApp.RequestOptions) => {
  // 1. 返回 Promise 对象
  return new Promise<ApiEnvelope<T>>((resolve, reject) => {
    uni.request({
      ...options,
      // 响应成功
      success(res) {
        const body = res.data as ApiEnvelope<T>
        // 状态码 2xx
        if (res.statusCode >= 200 && res.statusCode < 300) {
          if (isSuccessEnvelope(body)) {
            resolve(body)
            return
          }
          uni.showToast({
            icon: 'none',
            title: body?.msg || '请求错误',
          })
          reject(res)
        } else if (res.statusCode === 401) {
          uni.navigateTo({ url: '/pages/login/login' })
          reject(res)
        } else {
          uni.showToast({
            icon: 'none',
            title: body?.msg || '请求错误',
          })
          reject(res)
        }
      },
      // 响应失败
      fail(err) {
        uni.showToast({
          icon: 'none',
          title: '网络错误，换个网络试试',
        })
        reject(err)
      },
    })
  })
}
