#!/usr/bin/env python
# Copyright (C) 2015-2021  Nexedi SA and Contributors.
#                          Kirill Smelkov <kirr@nexedi.com>
#
# This program is free software: you can Use, Study, Modify and Redistribute
# it under the terms of the GNU General Public License version 3, or (at your
# option) any later version, as published by the Free Software Foundation.
#
# You can also Link and Combine this program with other software covered by
# the terms of any of the Free Software licenses or any of the Open Source
# Initiative approved licenses and Convey the resulting work. Corresponding
# source of such a combination shall include the source code for all other
# software used.
#
# This program is distributed WITHOUT ANY WARRANTY; without even the implied
# warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
#
# See COPYING file for full licensing terms.
# See https://www.nexedi.com/licensing for rationale and options.
"""Simple web-server that says "Hello World" for every path

helloweb.py [--logfile <logfile>] <bind-ip> <bind-port> ...
"""

from __future__ import print_function

import sys
PY3 = sys.version_info.major >= 3
import time
import argparse
if PY3:
    from http.server import BaseHTTPRequestHandler, HTTPServer
else:
    from BaseHTTPServer import BaseHTTPRequestHandler, HTTPServer

from socket import AF_INET6


class WebHello(BaseHTTPRequestHandler):

    def do_GET(self):
        self.send_response(200) # ok
        self.send_header("Content-type", "text/plain")
        self.end_headers()

        msg = "Hello %s at `%s`  ; %s  (python %s)" % (
                ' '.join(self.server.webhello_argv) or 'world',
                self.path, time.asctime(), sys.version.replace('\n', ' '))
        if PY3:
            msg = msg.encode()
        self.wfile.write(msg)


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

    print('* %s helloweb.py starting at %s' % (
        time.asctime(), (args.bind_ip, args.bind_port)), file=sys.stderr)

    # TODO autodetect ipv6/ipv4
    httpd = HTTPServerV6( (args.bind_ip, args.bind_port), WebHello)
    httpd.webhello_argv = args.argv_extra
    httpd.serve_forever()

if __name__ == '__main__':
    main()
