class ApplicationController < ActionController::API
  before_action :authenticate_request

  def authenticate_request
    # your logic here
  end
end
