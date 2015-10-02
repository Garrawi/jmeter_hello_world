#!/usr/bin/env ruby

require 'sinatra'
require 'json'

def n_samples(size)
  (0..size).reduce({}) do |map, i|
    map.merge("key#{i}" => i)
  end.to_json
end

get '/small' do
  content_type :json

  n_samples(20)
end


get '/large' do
  content_type :json

  n_samples(2000)
end


get '/random' do
  content_type :json

  n_samples(rand(2000))
end
