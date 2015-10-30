# minimal gem specification for helloweb.rb
Gem::Specification.new do |gem|
    gem.name            = "helloweb"
    gem.version         = "0.1"
    gem.summary         = "Hello Web in Ruby"

    gem.bindir          = ''
    gem.executables     = 'helloweb.rb'

    # NOTE we can require other gems via
    # gem.add_runtime_dependency ...
end
