module Wenmar
  class Vehicle
    attr_reader :attributes

    def initialize(attributes)
      @attributes = attributes
    end

    def id; @attributes["id"]; end
    def make; @attributes["make"]; end
    def model; @attributes["model"]; end
    def year; @attributes["year"]; end
    def vin; @attributes["vin"]; end
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
