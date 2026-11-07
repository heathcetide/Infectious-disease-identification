/**
 * Dev-only middleware: browser cannot reach *.amap.com / *.autonavi.com
 * (ERR_CONNECTION_CLOSED). Rewrite client URLs to /_AMapProxy/<proto>/<host>/...
 * and forward from the Vite Node process instead.
 */
import http from 'node:http'
import https from 'node:https'

const ALLOWED_HOST =
  /(?:^|\.)(?:amap\.com|autonavi\.com)$/i

function forward(req, res, targetUrl) {
  let parsed
  try {
    parsed = new URL(targetUrl)
  } catch {
    res.statusCode = 400
    res.end('bad proxy url')
    return
  }
  if (!ALLOWED_HOST.test(parsed.hostname)) {
    res.statusCode = 403
    res.end('host not allowed')
    return
  }

  const lib = parsed.protocol === 'http:' ? http : https
  const headers = { ...req.headers }
  delete headers.host
  headers.host = parsed.host
  // Avoid compressing so Vite/browser get predictable payloads
  delete headers['accept-encoding']

  const upstream = lib.request(
    {
      protocol: parsed.protocol,
      hostname: parsed.hostname,
      port: parsed.port || (parsed.protocol === 'http:' ? 80 : 443),
      path: parsed.pathname + parsed.search,
      method: req.method,
      headers,
      timeout: 30000,
    },
    (upRes) => {
      const outHeaders = { ...upRes.headers }
      // Same-origin proxy — loosen CORS just in case
      outHeaders['access-control-allow-origin'] = '*'
      res.writeHead(upRes.statusCode || 502, outHeaders)
      upRes.pipe(res)
    }
  )
  upstream.on('error', (err) => {
    res.statusCode = 502
    res.end(`amap proxy error: ${err.message}`)
  })
  upstream.on('timeout', () => {
    upstream.destroy()
    res.statusCode = 504
    res.end('amap proxy timeout')
  })
  req.pipe(upstream)
}

export function amapProxyPlugin() {
  return {
    name: 'amap-dev-proxy',
    configureServer(server) {
      server.middlewares.use((req, res, next) => {
        const url = req.url || ''
        if (!url.startsWith('/_AMapProxy/')) return next()

        // /_AMapProxy/https/webapi.amap.com/maps?...
        const rest = url.slice('/_AMapProxy/'.length)
        const slash = rest.indexOf('/')
        if (slash < 0) {
          res.statusCode = 400
          res.end('usage: /_AMapProxy/<proto>/<host>/...')
          return
        }
        const proto = rest.slice(0, slash)
        const afterProto = rest.slice(slash + 1)
        if (proto !== 'http' && proto !== 'https') {
          res.statusCode = 400
          res.end('proto must be http|https')
          return
        }
        const target = `${proto}://${afterProto}`
        forward(req, res, target)
      })
    },
  }
}
