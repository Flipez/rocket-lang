#!/usr/bin/env python3
import http.server
h = http.server.SimpleHTTPRequestHandler
h.extensions_map = {'': 'text/html', '.wasm': 'application/wasm', '.js': 'application/javascript'}
http.server.HTTPServer(('127.0.0.1', 2000), h).serve_forever()