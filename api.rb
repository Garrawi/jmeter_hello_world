#!/usr/bin/env ruby

require 'sinatra'
require 'json'

configure do
  set :port, 8080
end

def n_samples(size)
  (0..size).reduce({}) do |map, i|
    map.merge("key#{i}" => rand(200))
  end.to_json
end

get '/random' do
  content_type :json
  size = (params['size'] || rand(2000)).to_i

  n_samples(size)
end
