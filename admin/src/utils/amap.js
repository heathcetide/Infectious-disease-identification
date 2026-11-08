const AMAP_KEY =
  (typeof import.meta !== 'undefined' && import.meta.env?.VITE_AMAP_KEY) ||
  '55f746c1718bbd2930ba179a96bdd895'
const AMAP_SECURITY =
  (typeof import.meta !== 'undefined' && import.meta.env?.VITE_AMAP_SECURITY_CODE) ||
  ''

// Prefer same-origin proxy when Vite DEV or when explicitly enabled for
// embedded admin (VITE_AMAP_PROXY=1). Otherwise hit the official CDN.
const useAmapProxy =
  (typeof import.meta !== 'undefined' && !!import.meta.env?.DEV) ||
  (typeof import.meta !== 'undefined' && import.meta.env?.VITE_AMAP_PROXY === '1') ||
  (typeof window !== 'undefined' && window.location?.port === '8080')

/** Same-origin proxy when needed; else official CDN. */
function mapsScriptUrl(cacheBust = 0) {
  const qs = `v=2.0&key=${AMAP_KEY}${cacheBust ? `&_r=${cacheBust}` : ''}`
  if (useAmapProxy) {
    return `/_AMapProxy/https/webapi.amap.com/maps?${qs}`
  }
  return `https://webapi.amap.com/maps?${qs}`
}

let loadingPromise = null
let rewriteInstalled = false

function applySecurityConfig() {
  if (typeof window === 'undefined') return
  // Always set the object so AMap 2.0 does not fall back to brittle Referer checks alone.
  // Prefer VITE_AMAP_SECURITY_CODE from admin/.env.local when available.
  window._AMapSecurityConfig = {
    securityJsCode: AMAP_SECURITY || '',
  }
}

export function rewriteAmapUrl(url) {
  if (typeof url !== 'string' || !url) return url
  if (!useAmapProxy) return url
  if (url.startsWith('/_AMapProxy/')) return url

  // protocol-relative
  if (url.startsWith('//') && /(?:amap\.com|autonavi\.com)/i.test(url)) {
    return `/_AMapProxy/https/${url.slice(2)}`
  }

  const m = url.match(/^(https?):\/\/([^/]+)(\/.*)?$/i)
  if (!m) return url
  const proto = m[1].toLowerCase()
  const host = m[2]
  const path = m[3] || '/'
  if (!/(?:^|\.)(?:amap\.com|autonavi\.com)$/i.test(host.split(':')[0])) {
    return url
  }
  return `/_AMapProxy/${proto}/${host}${path}`
}

/**
 * Browser → *.amap.com / *.autonavi.com often dies with ERR_CONNECTION_CLOSED.
 * In Vite DEV, rewrite script / image / XHR / fetch URLs through /_AMapProxy.
 */
export function installAmapUrlRewrite() {
  if (!useAmapProxy || rewriteInstalled || typeof window === 'undefined') return
  rewriteInstalled = true

  const rewrite = rewriteAmapUrl

  const patchSrc = (Ctor) => {
    try {
      const proto = Ctor.prototype
      const desc = Object.getOwnPropertyDescriptor(proto, 'src')
      if (!desc?.set || !desc?.get) return
      Object.defineProperty(proto, 'src', {
        configurable: true,
        enumerable: desc.enumerable,
        get() {
          return desc.get.call(this)
        },
        set(v) {
          desc.set.call(this, rewrite(String(v)))
        },
      })
    } catch (_) {
      /* ignore */
    }
  }
  patchSrc(HTMLScriptElement)
  patchSrc(HTMLImageElement)

  try {
    const rawSetAttribute = Element.prototype.setAttribute
    Element.prototype.setAttribute = function setAttribute(name, value) {
      if (
        (this instanceof HTMLScriptElement || this instanceof HTMLImageElement) &&
        String(name).toLowerCase() === 'src'
      ) {
        return rawSetAttribute.call(this, name, rewrite(String(value)))
      }
      return rawSetAttribute.call(this, name, value)
    }
  } catch (_) {
    /* ignore */
  }

  try {
    const patchAppend = (proto) => {
      const raw = proto.appendChild
      proto.appendChild = function appendChild(node) {
        if (
          (node instanceof HTMLScriptElement || node instanceof HTMLImageElement) &&
          node.getAttribute?.('src')
        ) {
          const cur = node.getAttribute('src')
          const next = rewrite(cur)
          if (next !== cur) node.setAttribute('src', next)
        }
        return raw.call(this, node)
      }
    }
    patchAppend(Node.prototype)
  } catch (_) {
    /* ignore */
  }

  try {
    const XO = window.XMLHttpRequest
    if (XO?.prototype?.open) {
      const rawOpen = XO.prototype.open
      XO.prototype.open = function (method, url, ...rest) {
        return rawOpen.call(this, method, rewrite(String(url)), ...rest)
      }
    }
  } catch (_) {
    /* ignore */
  }

  try {
    const rawFetch = window.fetch?.bind(window)
    if (rawFetch) {
      window.fetch = (input, init) => {
        if (typeof input === 'string') {
          return rawFetch(rewrite(input), init)
        }
        if (input instanceof Request) {
          const next = rewrite(input.url)
          if (next !== input.url) {
            return rawFetch(new Request(next, input), init)
          }
        }
        return rawFetch(input, init)
      }
    }
  } catch (_) {
    /* ignore */
  }
}

function waitForAMap(timeoutMs = 15000) {
  return new Promise((resolve, reject) => {
    if (window.AMap) {
      resolve(window.AMap)
      return
    }
    const started = Date.now()
    const timer = setInterval(() => {
      if (window.AMap) {
        clearInterval(timer)
        resolve(window.AMap)
        return
      }
      if (Date.now() - started > timeoutMs) {
        clearInterval(timer)
        reject(new Error('AMap global not ready'))
      }
    }, 50)
  })
}

function injectScript(attempt) {
  return new Promise((resolve, reject) => {
    applySecurityConfig()
    installAmapUrlRewrite()

    if (window.AMap) {
      resolve(window.AMap)
      return
    }

    document.querySelectorAll('script[data-amap="1"]').forEach((el) => el.remove())

    const callbackName = `__amap_init_${Date.now()}_${attempt}`
    let settled = false
    const finishOk = (AMap) => {
      if (settled) return
      settled = true
      try {
        delete window[callbackName]
      } catch (_) {
        /* ignore */
      }
      resolve(AMap)
    }
    const finishErr = (err) => {
      if (settled) return
      settled = true
      try {
        delete window[callbackName]
      } catch (_) {
        /* ignore */
      }
      reject(err instanceof Error ? err : new Error(String(err)))
    }

    window[callbackName] = () => {
      waitForAMap(8000).then(finishOk).catch(finishErr)
    }

    const script = document.createElement('script')
    script.charset = 'utf-8'
    script.async = true
    script.dataset.amap = '1'
    script.src = `${mapsScriptUrl(attempt)}&callback=${callbackName}`
    script.onerror = () => {
      script.remove()
      finishErr(new Error('AMap script failed to load'))
    }
    script.onload = () => {
      waitForAMap(12000).then(finishOk).catch(() => {
        setTimeout(() => {
          if (!settled) finishErr(new Error('AMap script loaded but AMap missing'))
        }, 3000)
      })
    }
    document.head.appendChild(script)
  })
}

/** Lazily load AMap JS API once (only when a map page needs it). */
export function loadAMap(maxAttempts = 3) {
  if (typeof window !== 'undefined' && window.AMap) {
    return Promise.resolve(window.AMap)
  }
  if (loadingPromise) {
    return loadingPromise
  }
  loadingPromise = (async () => {
    let lastErr
    for (let attempt = 0; attempt < maxAttempts; attempt += 1) {
      try {
        const AMap = await injectScript(attempt)
        return AMap
      } catch (err) {
        lastErr = err
        await new Promise((r) => setTimeout(r, 400 * (attempt + 1)))
      }
    }
    loadingPromise = null
    throw lastErr || new Error('AMap script failed to load')
  })()
  return loadingPromise
}
