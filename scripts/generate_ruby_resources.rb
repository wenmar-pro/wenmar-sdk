#!/usr/bin/env ruby
# frozen_string_literal: true

# Builds ruby/lib/wenmar/resources.rb from spec/operations.json. This file
# contains all public resource methods on Wenmar::Client.
#
# Usage: ruby scripts/generate_ruby_resources.rb [manifest]
#   manifest defaults to spec/operations.json

require "json"

MANIFEST_PATH = ARGV[0] || "spec/operations.json"
MANIFEST = JSON.parse(File.read(MANIFEST_PATH))
OPS = MANIFEST["operations"]

def body_kwargs(op)
  shape = op["requestShape"]
  schema = op["requestSchema"]
  return [] unless schema
  return [] if shape.nil? || shape["wrapper"].nil? && shape["flat"].nil?

  if shape["wrapper"]
    [shape["wrapper"]]
  else
    shape["flat"]
  end
end

def path_params(op)
  op["pathParams"]
end

# Convert a query param name (e.g. "filters[has_open_work_order]") into a valid
# Ruby keyword identifier (e.g. "filters_has_open_work_order").
def ruby_param_name(name)
  name.gsub(/[\[\]]/, "_").gsub(/_+/, "_").sub(/_+\z/, "")
end

def build_method(op)
  m = op["id"]
  path = op["path"]
  path_args = path_params(op)
  # Substitute path params into the path: /customers/{customer_id}/drivers
  substituted = path.gsub(/\{([^}]+)\}/) do
    id = Regexp.last_match(1)
    "\#{#{id}}"
  end

  body_keys = body_kwargs(op)
  method = op["method"]

  # Positional path params first, then keyword body args.
  pos = path_args
  kw = body_keys.map { |k| "#{k}:" }

  # Query params become keyword args for GET operations that have them but
  # no dedicated params object (Ruby passes a params hash to get).
  query = op["queryParams"] || []

  case method
  when "get"
    if op["paginated"]
      # list -> returns paginator. get_all_* only for list_* operations.
      list_method = m
      base = m.start_with?("list_") ? m.sub(/\Alist_/, "") : nil
      if query.empty?
        sig = (pos + kw).join(", ")
        out = <<~RUBY
          # Lists #{m} resources (paginated).
          # @return [Wenmar::Paginator]
          def #{m}(#{sig.empty? ? "" : sig})
            get("#{substituted}")
          end
        RUBY
        if base
          out << <<~RUBY

            # Fetches all #{base}, up to 1000 by default.
            # @return [Array<Hash>]
            def get_all_#{base}(#{sig.empty? ? "" : sig})
              paginator_to_a(#{m}(#{path_args.join(", ")}), 1000)
            end
          RUBY
        end
        out
      else
        # list with query params -> pass a params hash
        query_kw = query.map { |q| "#{ruby_param_name(q["name"])}: nil" }
        sig = (pos + query_kw).join(", ")
        params_body = "params = { #{query.map { |q| "\"#{q["name"]}\" => #{ruby_param_name(q["name"])}" }.join(", ")} }"
        out = <<~RUBY
          # Lists #{m} resources (paginated).
          # @return [Wenmar::Paginator]
          def #{m}(#{sig.empty? ? "" : sig})
            #{params_body}
            get("#{substituted}", params.compact)
          end
        RUBY
        if base
          query_kw_args = query.map { |q| "#{ruby_param_name(q["name"])}: #{ruby_param_name(q["name"])}" }
          all_args = (path_args + query_kw_args).join(", ")
          out << <<~RUBY

            # Fetches all #{base}, up to 1000 by default.
            # @return [Array<Hash>]
            def get_all_#{base}(#{sig.empty? ? "" : sig})
              paginator_to_a(#{m}(#{all_args.empty? ? "" : all_args}), 1000)
            end
          RUBY
        end
        out
      end
    else
      # non-paginated GET
      query_kw = query.map { |q| "#{ruby_param_name(q["name"])}: nil" }
      sig = (pos + query_kw).join(", ")
      get_arg = if query.empty?
        "get(\"#{substituted}\")"
      else
        params_body = "params = { #{query.map { |q| "\"#{q["name"]}\" => #{ruby_param_name(q["name"])}" }.join(", ")} }"
        "#{params_body}\n      get(\"#{substituted}\", params.compact)"
      end
      <<~RUBY
        # Fetches #{m}.
        def #{m}(#{sig.empty? ? "" : sig})
          #{get_arg}
        end
      RUBY
    end
  when "post", "patch"
    verb = method
    query_kw = query.map { |q| "#{q["name"]}: nil" }
    params_line = if query.empty?
      ""
    else
      "params = { #{query.map { |q| "#{q["name"]}: #{q["name"]}" }.join(", ")} }"
    end
    if op["requestContentType"] && op["requestContentType"] != "application/json"
      # Multipart upload (e.g. import validate): accept a file path/IO and
      # send it as multipart/form-data.
      sig = (pos + ["file:"] + query_kw).join(", ")
      call_args = +"\"#{substituted}\", { multipart: { file: file } }"
      call_args << ", params.compact" unless query.empty?
      body_lines = [params_line, "#{verb}(#{call_args})"].reject(&:empty?)
      <<~RUBY
        # Runs #{m} (#{verb.upcase} #{path}) as a multipart upload.
        def #{m}(#{sig.empty? ? "" : sig})
          #{body_lines.join("\n  ")}
        end
      RUBY
    elsif body_keys.empty?
      # no body or empty body
      sig = (pos + query_kw).join(", ")
      call_args = +"\"#{substituted}\""
      call_args << ", params.compact" unless query.empty?
      body_lines = [params_line, "#{verb}(#{call_args})"].reject(&:empty?)
      <<~RUBY
        # Runs #{m} (#{verb.upcase} #{path}).
        def #{m}(#{sig.empty? ? "" : sig})
          #{body_lines.join("\n  ")}
        end
      RUBY
    else
      kw = body_keys.map { |k| "#{k}:" }
      sig = (pos + kw + query_kw).join(", ")
      body = "{ #{body_keys.map { |k| "#{k}: #{k}" }.join(", ")} }"
      call_args = +"\"#{substituted}\", #{body}"
      call_args << ", params.compact" unless query.empty?
      body_lines = [params_line, "#{verb}(#{call_args})"].reject(&:empty?)
      <<~RUBY
        # Runs #{m} (#{verb.upcase} #{path}).
        def #{m}(#{sig.empty? ? "" : sig})
          #{body_lines.join("\n  ")}
        end
      RUBY
    end
  when "delete"
    query_kw = query.map { |q| "#{q["name"]}: nil" }
    sig = (pos + query_kw).join(", ")
    params_line = if query.empty?
      ""
    else
      "params = { #{query.map { |q| "#{q["name"]}: #{q["name"]}" }.join(", ")} }"
    end
    call_args = +"\"#{substituted}\""
    call_args << ", params.compact" unless query.empty?
    body_lines = [params_line, "delete(#{call_args})"].reject(&:empty?)
    <<~RUBY
      # Deletes #{m}.
      def #{m}(#{sig.empty? ? "" : sig})
        #{body_lines.join("\n  ")}
      end
    RUBY
  end
end

content = +<<~'RUBY'
  # frozen_string_literal: true

  # AUTO-GENERATED by scripts/generate_ruby_resources.rb from spec/operations.json. DO NOT EDIT.

  module Wenmar
    class Client
RUBY

OPS.each do |op|
  content << build_method(op).to_s
end

content << <<~'RUBY'
    end
  end
RUBY

File.write("ruby/lib/wenmar/resources.rb", content)
puts "Wrote ruby/lib/wenmar/resources.rb (#{OPS.size} operations)"
