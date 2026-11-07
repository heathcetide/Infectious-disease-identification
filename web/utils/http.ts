export const baseURL = 'http://127.0.0.1:8080'
import store from '../store/index'
import { tip } from './tip'

const httpInterceptor = {
  invoke(options: UniApp.RequestOptions) {
    if (!options.url.startsWith('http')) {
      options.url = baseURL + options.url
    }
    if (!options.timeout || options.timeout <= 0) {
      options.timeout = 10000
    }
    options.header = {
      ...options.header,
      'source-client': 'miniapp',
    }
    const jwt = store.state.jwt
    if (jwt) {
      options.header.Authorization = `Bearer ${jwt}`
    }
  },
}
uni.addInterceptor('request', httpInterceptor)
uni.addInterceptor('uploadFile', httpInterceptor)

export type ApiEnvelope<T> = {
  code: number
  msg: string
  message?: string
  data: T
}

function isSuccessEnvelope<T>(body: ApiEnvelope<T> | null | undefined): boolean {
  if (!body || typeof body !== 'object') {
    return false
  }
  if (body.code === 200 || body.code === 1) {
    return true
  }
  return 'data' in body && body.data !== undefined && body.data !== null
}

function errorText(body: ApiEnvelope<unknown> | null | undefined, fallback: string) {
  return (body && (body.msg || body.message)) || fallback
}

/** Quiet request: do not auto toast on business failure (caller handles tip). */
export type HttpOptions = UniApp.RequestOptions & { quiet?: boolean }

export const http = <T>(options: HttpOptions) => {
  const quiet = !!options.quiet
  return new Promise<ApiEnvelope<T>>((resolve, reject) => {
    uni.request({
      ...options,
      success(res) {
        const body = res.data as ApiEnvelope<T>
        if (res.statusCode >= 200 && res.statusCode < 300) {
          if (isSuccessEnvelope(body)) {
            resolve(body)
            return
          }
          if (!quiet) {
            tip.fail(errorText(body, '请求失败'))
          }
          reject(body || res)
          return
        }
        if (res.statusCode === 401) {
          if (!quiet) {
            tip.fail('登录已过期，请重新登录')
          }
          setTimeout(() => {
            uni.reLaunch({ url: '/pages/login/login' })
          }, 400)
          reject(res)
          return
        }
        if (!quiet) {
          tip.fail(errorText(body, `请求错误(${res.statusCode})`))
        }
        reject(body || res)
      },
      fail(err) {
        if (!quiet) {
          tip.fail('网络异常，请检查网络后重试')
        }
        reject(err)
      },
    })
  })
}
