#!/usr/bin/env python
"""Simple web-server that says "Hello World" for every path

helloweb [--logfile <logfile>] <bind-ip> <bind-port> ...
"""

import sys
import time
import argparse
from BaseHTTPServer import BaseHTTPRequestHandler, HTTPServer
from socket import AF_INET6


class WebHello(BaseHTTPRequestHandler):

    def do_GET(self):
        self.send_response(200) # ok
        self.send_header("Content-type", "text/plain")
        self.end_headers()

        print >>self.wfile, \
            "Hello %s at `%s`  ; %s" % (
                ' '.join(self.server.webhello_argv) or 'world',
                self.path, time.asctime())


class HTTPServerV6(HTTPServer):
    address_family = AF_INET6

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('--logfile', dest='logfile')
    parser.add_argument('bind_ip')
    parser.add_argument('bind_port', type=int)
    parser.add_argument('argv_extra', metavar='...', nargs=argparse.REMAINDER)

    args = parser.parse_args()

    # HTTPServer logs to sys.stderr - override it if we have --logfile
    if args.logfile:
        f = open(args.logfile, 'a', buffering=1)
        sys.stderr = f

    print >>sys.stderr, '* %s helloweb starting at %s' % (
        time.asctime(), (args.bind_ip, args.bind_port))

    # TODO autodetect ipv6/ipv4
    httpd = HTTPServerV6( (args.bind_ip, args.bind_port), WebHello)
    httpd.webhello_argv = args.argv_extra
    httpd.serve_forever()

if __name__ == '__main__':
    main()