module Wenmar
  class Customer
    attr_reader :attributes

    def initialize(attributes)
      @attributes = attributes
    end

    def id; @attributes["id"]; end
    def first_name; @attributes["first_name"]; end
    def last_name; @attributes["last_name"]; end
    def full_name; @attributes["full_name"]; end
    def app_url; @attributes["app_url"]; end
    def url; @attributes["url"]; end

    def method_missing(name, *args)
      @attributes[name.to_s]
    end

    def respond_to_missing?(name, include_private = false)
      @attributes.key?(name.to_s) || super
    end
  end
end
