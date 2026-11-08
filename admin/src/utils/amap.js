const AMAP_KEY =
  (typeof import.meta !== 'undefined' && import.meta.env?.VITE_AMAP_KEY) ||
  '55f746c1718bbd2930ba179a96bdd895'
const AMAP_SECURITY =
  (typeof import.meta !== 'undefined' && import.meta.env?.VITE_AMAP_SECURITY_CODE) ||
  ''

// Proxy only in Vite DEV (browser often cannot reach amap.com) or when explicitly forced.
// Embedded admin on :8080 should prefer the official CDN — Go proxy can 502/503 under load.
const forceProxy =
  (typeof import.meta !== 'undefined' && import.meta.env?.VITE_AMAP_PROXY === '1') ||
  (typeof import.meta !== 'undefined' && !!import.meta.env?.DEV)

let preferProxy = forceProxy
let rewriteInstalled = false
let loadingPromise = null

function applySecurityConfig() {
  if (typeof window === 'undefined') return
  window._AMapSecurityConfig = {
    securityJsCode: AMAP_SECURITY || '',
  }
}

export function rewriteAmapUrl(url) {
  if (typeof url !== 'string' || !url) return url
  if (!preferProxy) return url
  if (url.startsWith('/_AMapProxy/')) return url

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
 * Rewrite script / image / XHR / fetch through /_AMapProxy when proxy mode is on.
 */
export function installAmapUrlRewrite() {
  if (!preferProxy || rewriteInstalled || typeof window === 'undefined') return
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

function mapsScriptUrl(attempt = 0) {
  const qs = `v=2.0&key=${AMAP_KEY}${attempt ? `&_r=${attempt}` : ''}`
  if (preferProxy) {
    return `/_AMapProxy/https/webapi.amap.com/maps?${qs}`
  }
  return `https://webapi.amap.com/maps?${qs}`
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
    if (preferProxy) installAmapUrlRewrite()

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

/** Lazily load AMap JS API once. Direct CDN first; fall back to same-origin proxy. */
export function loadAMap(maxAttempts = 3) {
  if (typeof window !== 'undefined' && window.AMap) {
    return Promise.resolve(window.AMap)
  }
  if (loadingPromise) {
    return loadingPromise
  }
  loadingPromise = (async () => {
    let lastErr
    // Pass 1: current mode (direct unless Vite DEV / VITE_AMAP_PROXY=1)
    for (let attempt = 0; attempt < maxAttempts; attempt += 1) {
      try {
        return await injectScript(attempt)
      } catch (err) {
        lastErr = err
        await new Promise((r) => setTimeout(r, 300 * (attempt + 1)))
      }
    }
    // Pass 2: if direct failed, try Go /_AMapProxy once
    if (!preferProxy) {
      preferProxy = true
      rewriteInstalled = false
      for (let attempt = 0; attempt < 2; attempt += 1) {
        try {
          return await injectScript(100 + attempt)
        } catch (err) {
          lastErr = err
          await new Promise((r) => setTimeout(r, 400))
        }
      }
    }
    loadingPromise = null
    throw lastErr || new Error('AMap script failed to load')
  })()
  return loadingPromise
}
