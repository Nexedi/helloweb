#!/usr/bin/env ruby
# Simple web-server that says "Hello World" for every path
#
# helloweb.rb [--logfile <logfile>] <bind-ip> <bind-port> ...
require 'webrick'
require 'time'
require 'optparse'
require 'ostruct'

def main
    args = OpenStruct.new
    args.logfile = nil
    opt = OptionParser.new
    opt.banner = "Usage: helloweb.rb [options] bind_ip bind_port ..."
    opt.on('--logfile LOGFILE')       { |o| args.logfile = o }
    opt.parse!

    args.bind_ip = ARGV.delete_at(0)
    args.bind_port = ARGV.delete_at(0)
    args.argv_extra = ARGV
    if args.bind_ip.nil? or args.bind_port.nil?
        puts opt
        exit 1
    end

    args.bind_port = Integer(args.bind_port)

    log = nil
    access_log = nil
    if args.logfile
        log_file = File.open args.logfile, 'a+'
        log_file.sync = true
        log = WEBrick::Log.new log_file
        access_log = [[log_file, WEBrick::AccessLog::COMBINED_LOG_FORMAT]]
    end

    httpd = WEBrick::HTTPServer.new :BindAddress => args.bind_ip,
                :Port => args.bind_port,
                :Logger => log,
                :AccessLog => access_log
    httpd.mount_proc '/' do |req, resp|
        name = args.argv_extra.join(' ')
        name = 'world' if name.empty?
        resp.body = "Hello #{name} at `#{req.path}`  ; #{Time.now.asctime}  (ruby)\n"
    end

    httpd.start
end

main
