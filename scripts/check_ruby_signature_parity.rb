#!/usr/bin/env ruby
# frozen_string_literal: true

# CI gate that enforces signature parity between the generated Ruby conformance
# dispatch and the generated Ruby client resources.
#
# The dispatch (conformance/ruby/dispatch.gen.rb) calls client methods passing
# keyword args for query/body params. The client (ruby/lib/wenmar/resources.rb)
# must accept every keyword arg the dispatch passes. This catches generator
# divergence (e.g. query params folded into the dispatch but not the method
# signature) statically, before a conformance fixture exercises the op at
# runtime.
#
# Exit 0 with counts, or exit 1 with a diagnostic.
#
# Usage: ruby scripts/check_ruby_signature_parity.rb

require "json"

ROOT = File.expand_path("..", __dir__)
DISPATCH = File.join(ROOT, "conformance", "ruby", "dispatch.gen.rb")
RESOURCES = File.join(ROOT, "ruby", "lib", "wenmar", "resources.rb")

def fail!(msg)
  warn "ERROR: #{msg}"
  exit 1
end

# Extract { method_name => [kwarg names] } from the dispatch file.
# Each entry looks like:
#   "op" => ->(client, args) { client.method(pos, kw: args["query"]["kw"], ...) }
def dispatch_kwargs(content)
  result = {}
  content.scan(/"([a-z0-9_]+)"\s*=>\s*->\(client, args\)\s*\{\s*client\.([a-z0-9_]+)\((.*?)\)\s*\}/m) do |_op, method, call_args|
    kwargs = call_args.scan(/([a-z0-9_]+):\s*args\[/).flatten
    result[method] = kwargs
  end
  result
end

# Extract { method_name => [kwarg names] } from the resources file.
# Each method looks like:
#   def method(pos, kw: nil, ...)   (optional keyword)
#   def method(pos, kw:, ...)       (required keyword)
def resource_kwargs(content)
  result = {}
  content.scan(/def\s+([a-z0-9_]+)\((.*?)\)/) do |method, sig|
    kwargs = sig.scan(/([a-z0-9_]+):\s*(?:nil|args)?/).flatten
    result[method] = kwargs
  end
  result
end

fail!("dispatch not found at #{DISPATCH}; run make generate first") unless File.exist?(DISPATCH)
fail!("resources not found at #{RESOURCES}; run make generate first") unless File.exist?(RESOURCES)

dispatch = dispatch_kwargs(File.read(DISPATCH))
resources = resource_kwargs(File.read(RESOURCES))

mismatches = []
dispatch.each do |method, kwargs|
  next if kwargs.empty?

  accepted = resources[method]
  if accepted.nil?
    mismatches << "#{method} (dispatch passes kwargs but no such client method)"
    next
  end
  missing = kwargs - accepted
  mismatches << "#{method} passes kwargs not accepted by client: #{missing.join(', ')}" unless missing.empty?
end

unless mismatches.empty?
  fail!("Ruby dispatch/client signature mismatch:\n  #{mismatches.join("\n  ")}")
end

puts "OK: dispatch-methods=#{dispatch.size} resource-methods=#{resources.size} signature parity holds"
