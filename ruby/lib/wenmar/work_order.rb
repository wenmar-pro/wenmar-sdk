module Wenmar
  class WorkOrder
    attr_reader :attributes

    def initialize(attributes)
      @attributes = attributes
    end

    def id; @attributes["id"]; end
    def status; @attributes["status"]; end
    def work_order_number; @attributes["work_order_number"]; end
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
