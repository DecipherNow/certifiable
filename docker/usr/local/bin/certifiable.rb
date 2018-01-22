#!/usr/bin/ruby

require 'securerandom'
require 'sinatra'

set :bind, '0.0.0.0'

VAR_DIRECTORY = File.join(File::SEPARATOR, 'usr', 'local', 'var', 'certifiable')
INTERMEDIATE_CERTS_DIRECTORY = File.join(VAR_DIRECTORY, 'intermediate', 'certs')
ROOT_CERTS_DIRECTORY = File.join(VAR_DIRECTORY, 'root', 'certs')
SERVERS_PATH = File.join(VAR_DIRECTORY, 'servers')
USERS_PATH = File.join(VAR_DIRECTORY, 'users')

get '/intermediate' do
    content_type :json
    files(INTERMEDIATE_CERTS_DIRECTORY).to_json
end

get '/intermediate/:file' do
    send_file File.join(INTERMEDIATE_CERTS_DIRECTORY, File.basename(params['file'])),
              :filename => File.basename(params['file']),
              :type => 'application/octet-stream'
end

get '/root' do
    content_type :json
    files(ROOT_CERTS_DIRECTORY).to_json
end

get '/root/:file' do
    send_file File.join(ROOT_CERTS_DIRECTORY, File.basename(params['file'])),
              :filename => File.basename(params['file']),
              :type => 'application/octet-stream'
end

get '/servers' do
    content_type :json
    directories(SERVERS_PATH).to_json
end

get '/servers/:server' do
    content_type :json
    files(File.join(SERVERS_PATH, File.basename(params['server']), 'files')).to_json
end

get '/servers/:server/:file' do
    send_file File.join(SERVERS_PATH, File.basename(params['server']), 'files', File.basename(params['file'])),
              :filename => File.basename(params['file']),
              :type => 'application/octet-stream'
end

post '/servers/:server' do
    `/usr/local/bin/server.sh #{params['server']} "#{params['cn']}" #{SecureRandom.hex(3)}`
end

get '/users' do
    content_type :json
    directories(USERS_PATH).to_json
end

get '/users/:user' do
    content_type :json
    files(File.join(USERS_PATH, File.basename(params['user']), 'files')).to_json
end

get '/users/:user/:file' do
    send_file File.join(USERS_PATH, File.basename(params['user']), 'files', File.basename(params['file'])),
              :filename => File.basename(params['file']),
              :type => 'application/octet-stream'
end

post '/users/:user' do
    `/usr/local/bin/user.sh #{params['user']} "#{params['cn']}" #{SecureRandom.hex(3)}`
end

def files(path)
    Dir.glob(File.join(path, '*')).select { |entry| File.file?(entry) }.map { |entry| File.basename(entry) }
end

def directories(path)
    Dir.glob(File.join(path, '*')).select { |entry| File.directory?(entry) }.map { |entry| File.basename(entry) }
end
