class AuthController < ApplicationController
  # Skip authentication for the auth endpoint itself
  skip_before_action :authenticate_request, only: [:token]

  def token
    # Return the static API token
    render json: {
      token: ApplicationController::STATIC_API_TOKEN,
      token_type: "Bearer",
      expires_in: nil # Static token doesn't expire
    }
  end
end

